package alarm

import (
	"context"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func TestAddLog(t *testing.T) {
	log := &model.AlarmLogAddInput{
		Type:       1,
		RuleId:     1,
		RuleName:   "告警规则1--",
		Level:      4,
		Data:       `{"ts":"2022-11-06 11:00:00","pr1":92,"pr2":98,"pr3":89}`,
		ProductKey: "aoxiang",
		DeviceKey:  "device1",
	}
	_, err := service.AlarmLog().Add(context.TODO(), log)
	if err != nil {
		t.Fatal(err)
	}
}
