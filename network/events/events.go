package events

import (
	"reflect"
	"sync"
)

type Handler func(args ...interface{})

// EventInterface Events接口
type EventInterface interface {
	Emit(event string, data ...interface{})
	On(event string, fn interface{})
	Once(event string, fn interface{})
	Off(event string, fn interface{})
}

type EventEmitter struct {
	events sync.Map
}

type subscriber struct {
	callback reflect.Value
	once     bool
}

// Emit 发送消息
func (e *EventEmitter) Emit(event string, data ...interface{}) {
	sub, ok1 := e.events.Load(event)
	subAll, ok2 := e.events.Load("*")

	if !ok1 && !ok2 {
		return
	}

	//整理参数
	args := make([]reflect.Value, 0)
	args = append(args, reflect.ValueOf(event))
	for _, v := range data {
		args = append(args, reflect.ValueOf(v))
	}

	if ok1 {
		subscribers := sub.(*sync.Map)
		//args := make([]reflect.Value, 1+len(data))
		subscribers.Range(func(key, value interface{}) bool {
			handler := value.(*subscriber)
			if args != nil && len(args) > 2 {
				handler.callback.Call(args[1:])
			}
			//处理仅订阅一次
			if handler.once {
				subscribers.Delete(key)
			}
			return true
		})
	}

	//处理全局订阅*
	if ok2 {
		subscribers := subAll.(*sync.Map)
		subscribers.Range(func(key, value interface{}) bool {
			handler := value.(*subscriber)
			handler.callback.Call(args)
			//处理仅订阅一次
			if handler.once {
				subscribers.Delete(key)
			}
			return true
		})
	}
}

// On 监听
func (e *EventEmitter) On(event string, fn interface{}) {
	callback := reflect.ValueOf(fn)
	val, ok := e.events.Load(event)
	if !ok {
		val = new(sync.Map)
		e.events.Store(event, val)
	}
	subscribers := val.(*sync.Map)
	subscribers.Store(callback.Pointer(), &subscriber{
		callback: callback,
		once:     false,
	})
}

// Once 监听一次
func (e *EventEmitter) Once(event string, fn interface{}) {
	callback := reflect.ValueOf(fn)
	val, ok := e.events.Load(event)
	if !ok {
		val = new(sync.Map)
		e.events.Store(event, val)
	}
	subscribers := val.(*sync.Map)
	subscribers.Store(callback.Pointer(), &subscriber{
		callback: callback,
		once:     true,
	})
}

// Off 取消监听
func (e *EventEmitter) Off(event string, fn interface{}) {
	callback := reflect.ValueOf(fn)
	val, ok := e.events.Load(event)
	if !ok {
		return
	}
	subscribers := val.(*sync.Map)
	subscribers.Delete(callback.Pointer())
}
