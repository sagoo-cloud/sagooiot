package datahub

import (
	"context"
	_ "github.com/sagoo-cloud/sagooiot/internal/logic/tdengine"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/taosdata/driver-go/v3/taosRestful"
)

func TestAllSource(t *testing.T) {
	out, err := service.DataSource().AllSource(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(out)
}

func TestGetAllDataS(t *testing.T) {
	in := &model.SourceDataAllInput{
		SourceId: 45,
		Param: map[string]interface{}{
			"pr1=?": "aaa",
		},
	}
	out, err := service.DataSource().GetAllData(context.TODO(), in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(out)
}

func TestUpdateDataS(t *testing.T) {
	err := service.DataSource().UpdateData(context.TODO(), 84)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetDbData(t *testing.T) {
	out, err := service.DataSource().GetDbData(context.TODO(), 77)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(out)
}
