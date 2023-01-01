package datahub

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func TestUpdateDataRecord(t *testing.T) {
	err := service.DataTemplateRecord().UpdateData(context.TODO(), 10)
	if err != nil {
		t.Fatal(err)
	}
}
