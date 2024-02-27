package baseLogic

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type AsyncMap struct {
	sync.RWMutex
	info map[string]*FInfo
}

type FInfo struct {
	FuncKey  string
	Request  interface{}
	Response chan interface{}
}

var asyncMapInfo = &AsyncMap{info: make(map[string]*FInfo)}

func SyncRequest(ctx context.Context, id, funcKey string, params interface{}, timeout int) (interface{}, error) {
	if timeout == 0 {
		timeout = 45
	}

	responseChan := make(chan interface{})
	asyncMapInfo.Lock()
	asyncMapInfo.info[id] = &FInfo{
		FuncKey:  funcKey,
		Request:  params,
		Response: responseChan,
	}
	asyncMapInfo.Unlock()

	defer func() {
		asyncMapInfo.Lock()
		delete(asyncMapInfo.info, id)
		asyncMapInfo.Unlock()
		close(responseChan)
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-time.After(time.Second * time.Duration(timeout)):
		return nil, fmt.Errorf("invoke service %s timed out", funcKey)
	case response := <-responseChan:
		return response, nil
	}
}

func GetCallInfoById(ctx context.Context, id string) (funcKey string, params interface{}, response chan interface{}, err error) {
	asyncMapInfo.RLock()
	defer asyncMapInfo.RUnlock()
	if info, ok := asyncMapInfo.info[id]; !ok {
		return "", nil, nil, errors.New("cannot get call info by id " + id)
	} else {
		return info.FuncKey, info.Request, info.Response, nil
	}
}
