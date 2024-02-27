package tdengine

import (
	"context"
	"encoding/json"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gtime"
)

func TestInsertTSL(t *testing.T) {
	deviceKey := "k213213"
	data := model.ReportPropertyData{
		"property_99": {2, gtime.Now().Unix()},
		"property_98": {2, gtime.Now().Unix()},
	}
	err := service.TSLTable().Insert(context.TODO(), deviceKey, data)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateStable(t *testing.T) {
	metadata := `{"key":"product_tsl_adjust","name":"物模型调整","properties":[{"key":"property_1","name":"属性_1","accessMode":1,"valueType":{"type":"string","maxLength":0},"desc":"描述edit"}],"functions":[],"events":[],"tags":[]}`

	var tsl *model.TSL
	err := json.Unmarshal([]byte(metadata), &tsl)
	if err != nil {
		t.Fatal(err)
	}

	err = service.TSLTable().CreateStable(context.TODO(), tsl)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDropStable(t *testing.T) {
	err := service.TSLTable().DropStable(context.TODO(), "product_cc")
	if err != nil {
		t.Fatal(err)
	}
}

func TestAddDatabaseField(t *testing.T) {
	err := service.TSLTable().AddDatabaseField(context.TODO(), "product_tsl_adjust", "test_add", "int", 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDelDatabaseField(t *testing.T) {
	err := service.TSLTable().DelDatabaseField(context.TODO(), "product_tsl_adjust", "test_add")
	if err != nil {
		t.Fatal(err)
	}
}

func TestAddTag(t *testing.T) {
	err := service.TSLTable().AddTag(context.TODO(), "product_cc", "test_tag_add", "string", 10)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDelTag(t *testing.T) {
	err := service.TSLTable().DelTag(context.TODO(), "product_cc", "test_tag_add")
	if err != nil {
		t.Fatal(err)
	}
}
