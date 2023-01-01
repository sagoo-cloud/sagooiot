package notifier

import (
	"testing"
	"time"
)

func TestNewTimer(t *testing.T) {
	notifier := NewNotifier(3 * time.Second)
	notifier.SetCallbacks(
		func(state State) {
			t.Log("告警了")
		},
		func(state State) {
			t.Log("提醒了")
		},
		func(state State) {
			t.Log("修复了")
		})
	i := 0
	for {
		i++
		t.Logf("第 %d 秒", i)
		if i == 6 || i == 12 {
			notifier.Trigger(true)
		}
		if i == 9 || i == 15 {
			notifier.Trigger(false)
		}

		time.Sleep(1 * time.Second)
		if i > 20 {
			break
		}
	}
}
