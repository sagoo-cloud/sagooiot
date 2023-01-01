package product

import (
	"context"
	_ "github.com/sagoo-cloud/sagooiot/internal/logic/tdengine"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func TestTotal(t *testing.T) {
	out, err := service.DevDevice().Total(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(out)
}

func TestTotalForMonths(t *testing.T) {
	out, err := service.DevDevice().TotalForMonths(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(out)
}

func TestAlarmTotalForMonths(t *testing.T) {
	out, err := service.DevDevice().AlarmTotalForMonths(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(out)
}
