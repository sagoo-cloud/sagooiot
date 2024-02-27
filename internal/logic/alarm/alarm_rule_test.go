package alarm

import (
	"context"
	"sagooiot/internal/service"
	"testing"
)

func TestCacheAllAlarmRule(t *testing.T) {
	err := service.AlarmRule().CacheAllAlarmRule(context.Background())
	if err != nil {
		t.Fatal(err)
	}

}
