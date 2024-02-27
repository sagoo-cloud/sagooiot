package initialize

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
)

func Handles(mux *asynq.ServeMux) {
	//mux.HandleFunc(system.CleanJwt, system.HandleCleanJwtTask)
	//
	//mux.HandleFunc(system.CleanLog, system.HandleCleanOperationLogTask)
	mux.HandleFunc("", HandleCollectTapDataTask)
	mux.HandleFunc("", HandleActiveDataTask)
}

func HandleCollectTapDataTask(context.Context, *asynq.Task) (err error) {

	fmt.Println("=======HandleCollectTapDataTask")
	return
}
func HandleActiveDataTask(context.Context, *asynq.Task) (err error) {

	fmt.Println("=======HandleActiveDataTask")
	return

}
