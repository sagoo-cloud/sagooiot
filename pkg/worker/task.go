package worker

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/pkg/errors"

	"go.opentelemetry.io/otel"
	"io"
	"reflect"
	"sagooiot/internal/tasks"
	"strings"
	"sync"
)

type Tasks struct {
	worker  *Worker
	taskJob tasks.TaskJob
}

var (
	instance *Tasks
	once     sync.Once
)

func TasksInstance() *Tasks {
	once.Do(func() {
		instance, _ = NewTasks() // Initialize only once
	})
	return instance
}

// Once 注册一个任务以运行一次。
func (tk Tasks) Once(options ...func(*RunOptions)) error {
	return tk.worker.Once(options...)
}

// Cron 注册一个任务以由cron表达式运行。
func (tk Tasks) Cron(options ...func(*RunOptions)) error {
	return tk.worker.Cron(options...)
}

// Remove 从任务队列中删除一个任务。
func (tk Tasks) Remove(ctx context.Context, uid string) error {
	return tk.worker.Remove(ctx, uid)
}

// GetTaskJobNameList 获取任务可用方法名列表。
func (tk Tasks) GetTaskJobNameList() (res map[string]string) {
	// 获取结构体的类型
	taskJobType := reflect.TypeOf(tk.taskJob)
	res = make(map[string]string)
	// 获取方法的数量
	numMethod := taskJobType.NumMethod()
	// 获取每个方法的信息
	for i := 0; i < numMethod; i++ {
		// 获取第 i 个方法
		method := taskJobType.Method(i)
		// 获取方法的名称
		methodName := method.Name
		res[methodName] = tk.taskJob.GetFuncNameList()[methodName]
	}
	return

}

// CheckFuncName 检查方法名是否存在
func (tk Tasks) CheckFuncName(funcName string) (exists bool) {
	funcList := tk.taskJob.GetFuncNameList()
	_, exists = funcList[funcName]
	return
}

// ParseParameters 解析参数
func (tk Tasks) ParseParameters(parseData string) (params []interface{}, err error) {
	parts := strings.Split(parseData, "|")
	params = make([]interface{}, len(parts))
	for i, part := range parts {
		params[i] = part
	}
	return
}

func NewTasks() (tk *Tasks, err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = errors.Errorf("%v", e)
		}
	}()
	w := New(
		WithHandler(func(ctx context.Context, p Payload) error {
			return process(task{
				ctx:     ctx,
				payload: p,
			})
		}),
	)
	if w.Error != nil {
		err = errors.WithMessage(w.Error, "initialize worker failed")
		return
	}

	tk = &Tasks{
		worker: w,
	}
	g.Log().Debug(context.Background(), "initialize worker success")
	return
}

type task struct {
	ctx     context.Context
	payload Payload
	job     tasks.TaskJob
}

// process 处理任务
func process(t task) (err error) {
	tr := otel.Tracer("task")
	_, span := tr.Start(t.ctx, "Task")
	defer span.End()
	var taskData tasks.TaskJob
	err = json.Unmarshal(t.payload.Payload, &taskData)
	if err != nil {
		return err
	}

	err = CallMethod(&taskData)
	if err != nil {
		fmt.Println("CallMethod err:", err.Error())
		return err
	}
	return
}

// CallMethod 调用任务的实际方法
func CallMethod(task *tasks.TaskJob) (err error) {
	// 判断task.Params是否为nil
	if task.Params == nil {
		task.Params = []interface{}{}
	} else {
		var paramsData []interface{}
		// 判断task.Params内的值是否为基本类型
		for _, param := range task.Params {
			switch reflect.TypeOf(param).Kind() {
			case reflect.Float64:
				// 如果是数字 ，检查是否可以转换为整数
				v := param.(float64)
				if v == float64(int64(v)) {
					// 如果可以转换为整数，则转换为int类型
					paramsData = append(paramsData, int64(v))
				} else {
					// 如果不能转换为整数，则保留为float64类型
					paramsData = append(paramsData, v)
				}
			case reflect.String, reflect.Bool:
				// 对于字符串、布尔和空值，直接添加到Params中
				paramsData = append(paramsData, param)
				// 可以添加更多的类型处理
			default:
				panic("unhandled default case")
			}
		}
		task.Params = paramsData
	}

	// 获取TaskModel的反射值
	taskValue := reflect.ValueOf(task)
	// 准备要调用的方法的参数
	var args []reflect.Value
	for _, param := range task.Params {
		if !reflect.ValueOf(param).IsValid() {
			err = errors.New("invalid parameter")
			return
		}
		args = append(args, reflect.ValueOf(param))
	}
	// 查找并执行方法
	method := taskValue.MethodByName(task.MethodName)
	if method.IsValid() {
		if method.Type().NumIn() == 0 {
			method.Call(nil)
			return
		} else if method.Type().NumIn() != len(args) {
			err = errors.New("incorrect number of parameters")
			return
		}
		g.Log().Debug(context.Background(), "执行的任务：", task.MethodName)
		if len(task.Params) > 0 {
			method.Call(args)
		} else {
			method.Call(nil)
		}
	} else {
		errInfo := fmt.Sprintf("Method not found: %s", task.MethodName)
		err = errors.New(errInfo)
	}
	return
}

// UnmarshalTask 解析TaskJob，使用自定义JSON解析来处理类型问题
func UnmarshalTask(data []byte) (*tasks.TaskJob, error) {
	var task tasks.TaskJob
	dec := json.NewDecoder(bytes.NewReader(data))
	// 使用json.Decoder逐个解析Token，自定义解析逻辑
	for {
		t, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		// 根据Token类型处理
		switch v := t.(type) {
		case float64:
			// 如果是数字 ，检查是否可以转换为整数
			if v == float64(int64(v)) {
				// 如果可以转换为整数，则转换为int类型
				task.Params = append(task.Params, int64(v))
			} else {
				// 如果不能转换为整数，则保留为float64类型
				task.Params = append(task.Params, v)
			}
		case string, bool, nil:
			// 对于字符串、布尔和空值，直接添加到Params中
			task.Params = append(task.Params, v)
			// 可以添加更多的类型处理
		}
	}
	return &task, nil
}
