package datahub

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func TestAllTemplate(t *testing.T) {
	out, err := service.DataTemplate().AllTemplate(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(out)
}

func TestGetAllData(t *testing.T) {
	in := &model.TemplateDataAllInput{
		Id: 4,
		Param: map[string]interface{}{
			"city like ?": "%北京%",
		},
	}
	out, err := service.DataTemplate().GetAllData(context.TODO(), in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(out)
}

func TestGetLastData(t *testing.T) {
	in := &model.TemplateDataLastInput{
		Id: 18,
	}
	out, err := service.DataTemplate().GetLastData(context.TODO(), in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(out)
}

func TestUpdateDataTpl(t *testing.T) {
	err := service.DataTemplate().UpdateData(context.TODO(), 15)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetAllBySql(t *testing.T) {
	sql := "select * from data_template_20 order by created_at desc limit 2"
	out, err := service.DataTemplate().GetDataBySql(context.TODO(), sql)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(out)
}

func TestGetAllByTableName(t *testing.T) {
	tableName := "data_template_20"
	out, err := service.DataTemplate().GetDataByTableName(context.TODO(), tableName)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(out)
}

func TestCheckRelation(t *testing.T) {
	out, err := service.DataTemplate().CheckRelation(context.TODO(), 31)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(out)
}
