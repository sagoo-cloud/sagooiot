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
	Info map[string]*FInfo
}

type FInfo struct {
	FuncKey  string
	Request  interface{}
	Response chan interface{}
}

var AsyncMapInfo = &AsyncMap{Info: make(map[string]*FInfo)}

func SyncRequest(ctx context.Context, id, funcKey string, params interface{}, timeout int) (interface{}, error) {
	if timeout == 0 {
		timeout = 45
	}

	responseChan := make(chan interface{}, 1)
	AsyncMapInfo.Lock()
	AsyncMapInfo.Info[id] = &FInfo{
		FuncKey:  funcKey,
		Request:  params,
		Response: responseChan,
	}
	AsyncMapInfo.Unlock()

	defer func() {
		AsyncMapInfo.Lock()
		delete(AsyncMapInfo.Info, id)
		AsyncMapInfo.Unlock()
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
	AsyncMapInfo.RLock()
	defer AsyncMapInfo.RUnlock()
	if info, ok := AsyncMapInfo.Info[id]; !ok {
		return "", nil, nil, errors.New("cannot get call info by id " + id)
	} else {
		return info.FuncKey, info.Request, info.Response, nil
	}
}
