package baseLogic

import (
	"context"
	"reflect"
	"testing"
	"time"
)

// TestSyncRequest 测试 SyncRequest 函数.
func TestSyncRequest(t *testing.T) {
	ctx := context.Background()
	id := "testID"
	funcKey := "testFunc"
	params := "testParams"
	timeout := 10

	// 启动一个 goroutine 模拟异步响应.
	go func() {
		time.Sleep(2 * time.Second) // 模拟处理时间.
		asyncMapInfo.Lock()
		if info, ok := asyncMapInfo.info[id]; ok {
			info.Response <- "testResponse"
		}
		asyncMapInfo.Unlock()
	}()

	got, err := SyncRequest(ctx, id, funcKey, params, timeout)
	if err != nil {
		t.Errorf("SyncRequest() error = %v", err)
		return
	}
	if got != "testResponse" {
		t.Errorf("SyncRequest() got = %v, want %v", got, "testResponse")
	}
}

// TestGetCallInfoById 测试 GetCallInfoById 函数.
func TestGetCallInfoById(t *testing.T) {
	ctx := context.Background()
	id := "testID"
	funcKey := "testFunc"
	params := "testParams"

	asyncMapInfo.Lock()
	asyncMapInfo.info[id] = &FInfo{
		FuncKey:  funcKey,
		Request:  params,
		Response: make(chan interface{}),
	}
	asyncMapInfo.Unlock()

	gotFuncKey, gotParams, _, err := GetCallInfoById(ctx, id)
	if err != nil {
		t.Errorf("GetCallInfoById() error = %v", err)
		return
	}
	if gotFuncKey != funcKey {
		t.Errorf("GetCallInfoById() gotFuncKey = %v, want %v", gotFuncKey, funcKey)
	}
	if !reflect.DeepEqual(gotParams, params) {
		t.Errorf("GetCallInfoById() gotParams = %v, want %v", gotParams, params)
	}
}
