package datahub

import (
	"context"
	_ "github.com/sagoo-cloud/sagooiot/internal/logic/tdengine"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/taosdata/driver-go/v3/taosRestful"
)

func TestGetForTpl(t *testing.T) {
	out, err := service.DataSourceRecord().GetForTpl(context.TODO(), 43, 9)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(out)
}
