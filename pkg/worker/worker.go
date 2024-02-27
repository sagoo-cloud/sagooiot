package worker

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/golang-module/carbon/v2"
	"github.com/gorhill/cronexpr"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"io"
	"net/http"
	"sagooiot/pkg/utility/nx"
	"strings"
	"time"
)

type Worker struct {
	ops       Options
	redis     redis.UniversalClient
	redisOpt  asynq.RedisConnOpt
	lock      *nx.Nx
	client    *asynq.Client
	inspector *asynq.Inspector
	Error     error
}

type periodTask struct {
	Expr      string `json:"expr"` // cron expr github.com/robfig/cron/v3
	Group     string `json:"group"`
	Uid       string `json:"uid"`
	Payload   []byte `json:"payload"`
	Next      int64  `json:"next"`      // next schedule unix timestamp
	Processed int64  `json:"processed"` // run times
	MaxRetry  int    `json:"maxRetry"`
	Timeout   int    `json:"timeout"`
}

func (p *periodTask) String() (str string) {
	bs, _ := json.Marshal(p)
	str = string(bs)
	return
}

func (p *periodTask) FromString(str string) {
	err := json.Unmarshal([]byte(str), p)
	if err != nil {
		return
	}
	return
}

type periodTaskHandler struct {
	tk Worker
}

type Payload struct {
	Group   string `json:"group"`
	Uid     string `json:"uid"`
	Payload []byte `json:"payload"`
}

func (p Payload) String() (str string) {
	bs, _ := json.Marshal(p)
	str = string(bs)
	return
}

func (p periodTaskHandler) ProcessTask(ctx context.Context, t *asynq.Task) (err error) {
	uid := guid.S()
	group := strings.TrimSuffix(strings.TrimSuffix(t.Type(), ".once"), ".cron")
	payload := Payload{
		Group:   group,
		Uid:     t.ResultWriter().TaskID(),
		Payload: t.Payload(),
	}
	defer func() {
		if err != nil {
			glog.Debugf(ctx, "run task failed. uuid: %s task: %s Error:%s", uid, payload, err)
		}
	}()
	if p.tk.ops.handler != nil {
		err = p.tk.ops.handler(ctx, payload)
	} else if p.tk.ops.handlerNeedWorker != nil {
		err = p.tk.ops.handlerNeedWorker(p.tk, ctx, payload)
	} else if p.tk.ops.callback != "" {
		err = p.httpCallback(ctx, payload)
	} else {
		glog.Debugf(ctx, "no task handler. uuid: %s task: %s Error:%s", uid, payload, err)
	}
	// save processed count
	p.tk.processed(ctx, payload.Uid)
	return
}

func (p periodTaskHandler) httpCallback(ctx context.Context, payload Payload) (err error) {
	client := &http.Client{}
	body := payload.String()
	var r *http.Request
	r, _ = http.NewRequestWithContext(ctx, http.MethodPost, p.tk.ops.callback, bytes.NewReader([]byte(body)))
	r.Header.Add("Content-Type", "application/json")
	var res *http.Response
	res, err = client.Do(r)
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			glog.Error(ctx, err)
		}
	}(res.Body)
	if res.StatusCode != http.StatusOK {
		err = ErrHttpCallbackInvalidStatusCode
	}
	return
}

// New 创建一个新的任务处理器
func New(options ...func(*Options)) (tk *Worker) {
	ops := getOptionsOrSetDefault(nil)
	for _, f := range options {
		f(ops)
	}
	tk = &Worker{}
	if ops.redisUri == "" {
		tk.Error = errors.Unwrap(ErrRedisNil)
		return
	}
	rs, err := asynq.ParseRedisURI(ops.redisUri)
	if err != nil {
		tk.Error = errors.Unwrap(ErrRedisInvalid)
		return
	}
	// 将组前缀添加到拆分的差异组
	ops.redisPeriodKey = strings.Join([]string{"ops.group", ops.redisPeriodKey}, ".")
	rd := rs.MakeRedisClient().(redis.UniversalClient)
	client := asynq.NewClient(rs)
	inspector := asynq.NewInspector(rs)
	// initialize redis lock
	nxLock := nx.New(
		nx.WithRedis(rd),
		nx.WithExpire(10),
		nx.WithKey(strings.Join([]string{ops.redisPeriodKey, "lock"}, ".")),
	)
	// initialize server
	srv := asynq.NewServer(
		rs,
		asynq.Config{
			Concurrency: 10, // 最大同时执行任务数
			Queues: map[string]int{
				ops.group: 10,
			},
			LogLevel: 4,
		},
	)
	go func() {
		var h periodTaskHandler
		h.tk = *tk
		if e := srv.Run(h); e != nil {
			glog.Error(context.Background(), "run task handler failed")
		}
	}()
	tk.ops = *ops
	tk.redis = rd
	tk.redisOpt = rs
	tk.lock = nxLock
	tk.client = client
	tk.inspector = inspector
	// initialize scanner
	go func() {
		for {
			time.Sleep(time.Second)
			tk.scan()
		}
	}()
	if tk.ops.clearArchived > 0 {
		// initialize clear archived
		go func() {
			for {
				time.Sleep(time.Duration(tk.ops.clearArchived) * time.Second)
				tk.clearArchived()
			}
		}()
	}
	return
}

func (wk Worker) Once(options ...func(*RunOptions)) (err error) {
	ops := getRunOptionsOrSetDefault(nil)
	for _, f := range options {
		f(ops)
	}
	if ops.uid == "" {
		err = errors.Unwrap(ErrUuidNil)
		return
	}
	t := asynq.NewTask(strings.Join([]string{ops.group, "once"}, "."), ops.payload, asynq.TaskID(ops.uid))
	taskOpts := []asynq.Option{
		asynq.Queue(wk.ops.group),
		asynq.MaxRetry(wk.ops.maxRetry),
		asynq.Timeout(time.Duration(ops.timeout) * time.Second),
	}
	if ops.maxRetry > 0 {
		taskOpts = append(taskOpts, asynq.MaxRetry(ops.maxRetry))
	}
	if ops.retention > 0 {
		taskOpts = append(taskOpts, asynq.Retention(time.Duration(ops.retention)*time.Second))
	} else {
		taskOpts = append(taskOpts, asynq.Retention(time.Duration(wk.ops.retention)*time.Second))
	}
	if ops.in != nil {
		taskOpts = append(taskOpts, asynq.ProcessIn(*ops.in))
	} else if ops.at != nil {
		taskOpts = append(taskOpts, asynq.ProcessAt(*ops.at))
	} else if ops.now {
		taskOpts = append(taskOpts, asynq.ProcessIn(time.Millisecond))
	}
	_, err = wk.client.Enqueue(t, taskOpts...)
	if ops.replace && errors.Is(err, asynq.ErrTaskIDConflict) {
		// remove old one if replace = true
		ctx := wk.getDefaultTimeoutCtx()
		if ops.ctx != nil {
			ctx = ops.ctx
		}
		err = wk.Remove(ctx, ops.uid)
		if err != nil {
			return
		}
		_, err = wk.client.Enqueue(t, taskOpts...)
	}
	return
}

// Cron 设置周期性任务
func (wk Worker) Cron(options ...func(*RunOptions)) (err error) {
	ops := getRunOptionsOrSetDefault(nil)
	for _, f := range options {
		f(ops)
	}
	if ops.uid == "" {
		err = errors.Unwrap(ErrUuidNil)
		return
	}
	var next int64
	next, err = getNext(ops.expr, 0)
	if err != nil {
		err = errors.Unwrap(ErrExprInvalid)
		return
	}
	t := periodTask{
		Expr:     ops.expr,
		Group:    strings.Join([]string{ops.group, "cron"}, "."),
		Uid:      ops.uid,
		Payload:  ops.payload,
		Next:     next,
		MaxRetry: ops.maxRetry,
		Timeout:  ops.timeout,
	}
	ctx := wk.getDefaultTimeoutCtx()
	res, err := wk.redis.HGet(ctx, wk.ops.redisPeriodKey, ops.uid).Result()
	if err == nil {
		var oldT periodTask
		err = json.Unmarshal([]byte(res), &oldT)
		if err != nil {
			return err
		}
		if oldT.Expr != t.Expr {
			// 删除旧任务
			err = wk.Remove(ctx, t.Uid)
			if err != nil {
				return err
			}
		}
	}
	_, err = wk.redis.HSet(ctx, wk.ops.redisPeriodKey, ops.uid, t.String()).Result()
	if err != nil {
		err = errors.Unwrap(ErrSaveCron)
		return
	}
	return
}

// Remove 移除任务
func (wk Worker) Remove(ctx context.Context, uid string) (err error) {
	err = wk.lock.MustLock(ctx)
	if err != nil {
		return
	}
	defer wk.lock.Unlock(ctx)
	wk.redis.HDel(ctx, wk.ops.redisPeriodKey, uid)

	err = wk.inspector.DeleteTask(wk.ops.group, uid)
	return
}

func (wk Worker) processed(ctx context.Context, uid string) {
	err := wk.lock.MustLock(ctx)
	if err != nil {
		return
	}
	defer wk.lock.Unlock(ctx)
	t, e := wk.redis.HGet(ctx, wk.ops.redisPeriodKey, uid).Result()
	if e == nil || !errors.Is(e, redis.Nil) {
		var item periodTask
		item.FromString(t)
		item.Processed++
		wk.redis.HSet(ctx, wk.ops.redisPeriodKey, uid, item.String())
	}
	return
}

// scan 扫描并处理任务队列
func (wk Worker) scan() {
	ctx := wk.getDefaultTimeoutCtx()
	ok := wk.lock.Lock()
	if !ok {
		return
	}
	defer wk.lock.Unlock()
	m, _ := wk.redis.HGetAll(ctx, wk.ops.redisPeriodKey).Result()
	p := wk.redis.Pipeline()
	ops := wk.ops
	if ops.group == "task" {
		for _, v := range m {
			var item periodTask
			item.FromString(v)
			next, _ := getNext(item.Expr, item.Next)
			t := asynq.NewTask(item.Group, item.Payload, asynq.TaskID(item.Uid))
			taskOpts := []asynq.Option{
				asynq.Queue(ops.group),
				asynq.MaxRetry(ops.maxRetry),
				asynq.Timeout(time.Duration(item.Timeout) * time.Second),
			}
			if item.MaxRetry > 0 {
				taskOpts = append(taskOpts, asynq.MaxRetry(item.MaxRetry))
			}
			diff := next - item.Next
			if diff > 10 {
				retention := diff / 3
				if diff > 600 {
					// 最大保留时间为 10 分钟
					retention = 600
				}
				// 设置保留时间以避免在短时间内重复任务
				taskOpts = append(taskOpts, asynq.Retention(time.Duration(retention)*time.Second))
			}
			taskOpts = append(taskOpts, asynq.ProcessAt(time.Unix(item.Next, 0)))
			_, err := wk.client.Enqueue(t, taskOpts...)
			// 入队成功，更新下一个
			if err == nil {
				item.Next = next
				p.HSet(ctx, wk.ops.redisPeriodKey, item.Uid, item.String())
			}
		}
		// 批量保存到缓存
		_, err := p.Exec(ctx)
		if err != nil {
			return
		}
	}

	return
}

// clearArchived 清除已归档的任务
func (wk Worker) clearArchived() {
	list, err := wk.inspector.ListArchivedTasks(wk.ops.group, asynq.Page(1), asynq.PageSize(100))
	if err != nil {
		return
	}
	ctx := wk.getDefaultTimeoutCtx()
	for _, item := range list {
		last := carbon.CreateFromStdTime(item.LastFailedAt)
		if !last.IsZero() && item.Retried < item.MaxRetry {
			continue
		}
		uid := item.ID
		var flag bool
		if strings.HasSuffix(item.Type, ".cron") {
			// cron task
			t, e := wk.redis.HGet(ctx, wk.ops.redisPeriodKey, uid).Result()
			if e == nil || !errors.Is(e, redis.Nil) {
				var task periodTask
				task.FromString(t)
				next, _ := getNext(task.Expr, task.Next)
				diff := next - task.Next
				if diff <= 60 {
					if carbon.Now().Gt(last.AddMinutes(5)) {
						flag = true
					}
				} else if diff <= 600 {
					if carbon.Now().Gt(last.AddMinutes(30)) {
						flag = true
					}
				} else if diff <= 3600 {
					if carbon.Now().Gt(last.AddHours(2)) {
						flag = true
					}
				} else {
					if carbon.Now().Gt(last.AddHours(5)) {
						flag = true
					}
				}
			}
		} else {
			// once task, has failed for more than 5 minutes
			if carbon.Now().Gt(last.AddMinutes(5)) {
				flag = true
			}
		}
		if flag {
			err := wk.inspector.DeleteTask(wk.ops.group, uid)
			if err != nil {
				return
			}
		}
	}
}

// getDefaultTimeoutCtx 获取带有默认超时的上下文
func (wk Worker) getDefaultTimeoutCtx() context.Context {
	c, _ := context.WithTimeout(context.Background(), time.Duration(wk.ops.timeout)*time.Second)
	return c
}

// getNext 计算下一次执行时间
func getNext(expr string, timestamp int64) (next int64, err error) {
	var e *cronexpr.Expression
	e, err = cronexpr.Parse(expr)
	if err != nil {
		return
	}
	t := carbon.Now().ToStdTime()
	if timestamp > 0 {
		t = carbon.CreateFromTimestamp(timestamp).ToStdTime()
	}
	next = e.Next(t).Unix()
	return
}
