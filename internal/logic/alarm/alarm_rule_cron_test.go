package alarm

import (
	"context"
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	ar := alarmRuleNew()
	if err := ar.start(context.TODO(), 18); err != nil {
		t.Fatal(err)
	}
	time.Sleep(30 * time.Second)
}

func TestStop(t *testing.T) {
	ar := alarmRuleNew()
	if err := ar.stop(context.TODO(), 18); err != nil {
		t.Fatal(err)
	}
	time.Sleep(30 * time.Second)
}
