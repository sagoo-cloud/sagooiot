package tdengine

import (
	"context"
	"encoding/json"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gtime"
)

func TestInsertTSL(t *testing.T) {
	deviceKey := "k213213"
	data := map[string]any{
		"ts":          gtime.Now(),
		"property_99": 2,
		"property_98": 2,
		"property_97": 2,
		"property_96": 2,
		"property_95": 2,
	}
	err := service.TSLTable().Insert(context.TODO(), deviceKey, data)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateStable(t *testing.T) {
	metadata := `{"key":"product_cc","name":"产品_1","properties":[{"key":"property_1","name":"属性_1","accessMode":1,"valueType":{"type":"string","maxLength":0},"desc":"描述edit"}],"functions":[{"key":"function_3","name":"功能_3","inputs":[{"key":"input_1","name":"参数_1","valueType":{"type":"string","maxLength":22},"desc":"参数描述"}],"output":{"type":"string","maxLength":22},"desc":"描述编辑"}],"events":[{"key":"function_1","name":"事件_1","level":0,"valueType":{"type":"string","maxLength":22},"desc":"描述"}],"tags":[]}`

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
	err := service.TSLTable().AddDatabaseField(context.TODO(), "product_cc", "test_add", "int", 0)
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
