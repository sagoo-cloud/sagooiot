package worker

import "fmt"

var (
	ErrUuidNil                       = fmt.Errorf("uuid is empty")
	ErrRedisNil                      = fmt.Errorf("redis is empty")
	ErrRedisInvalid                  = fmt.Errorf("redis is invalid")
	ErrExprInvalid                   = fmt.Errorf("expr is invalid")
	ErrSaveCron                      = fmt.Errorf("save cron failed")
	ErrHttpCallbackInvalidStatusCode = fmt.Errorf("http callback invalid status code")
)
