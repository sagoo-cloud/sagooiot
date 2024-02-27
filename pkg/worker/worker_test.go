package worker

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestWorker(t *testing.T) {
	// cron worker
	wk1 := New(
		WithRedisUri("redis://127.0.0.1:6379/0"),
		WithHandler(handleProcess),
	)
	// example 2: run at every 8 minute
	err := wk1.Cron(
		WithRunUuid("order2"),
		WithRunGroup("task2"),
		WithRunExpr("0/1 * * * ?"),
	)

	fmt.Println(err)

	time.Sleep(time.Hour)
}
func handleProcess(ctx context.Context, p Payload) (err error) {
	fmt.Println(ctx, "=====", p.Uid, p.Group)
	switch p.Group {
	case "task1":

		fmt.Println("=====", p.Uid, p.Group)
	case "task2":
		fmt.Println(p.Uid, p.Group)

	}
	return
}
