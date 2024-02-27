package analysis

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/internal/consts"
	"sagooiot/internal/service"
	"sagooiot/pkg/general"
	"sagooiot/pkg/tsd"
	"time"
)

type sAnalysisDeviceDataTsd struct {
}

func init() {
	service.RegisterAnalysisDeviceDataTsd(analysisDeviceDataTsdNew())
}

func analysisDeviceDataTsdNew() *sAnalysisDeviceDataTsd {
	return &sAnalysisDeviceDataTsd{}
}

func (s *sAnalysisDeviceDataTsd) GetDeviceData(ctx context.Context, reqData general.SelectReq) (rs []interface{}, err error) {
	// 创建数据库连接
	db := tsd.GetDB()
	defer db.Close()

	if reqData.Param["deviceKey"] == nil {
		err = fmt.Errorf("deviceKey is nil")
		return
	}

	sqlStr := fmt.Sprintf("select * from device_%s", reqData.Param["deviceKey"])
	rows, err := db.Query(sqlStr)
	if err != nil {
		return
	}
	defer rows.Close()

	columns, _ := rows.Columns()

	for rows.Next() {
		values := make([]any, len(columns))
		for i := range values {
			values[i] = new(any)
		}

		err = rows.Scan(values...)
		if err != nil {
			return
		}

		m := make(gdb.Record, len(columns))
		for i, c := range columns {
			// 去除前缀
			if c[:2] == consts.TdPropertyPrefix {
				c = c[2:]
			}
			m[c] = toTime(gvar.New(values[i]))
		}
		rs = append(rs, m)
	}
	return
}

// Time REST连接时区处理
func toTime(v *g.Var) (rs *g.Var) {
	driver := g.Cfg().MustGet(context.TODO(), "tdengine.type")
	if driver.String() == "taosRestful" {
		if t, err := time.Parse("2006-01-02 15:04:05 +0000 UTC", v.String()); err == nil {
			rs = gvar.New(t.Local().Format("2006-01-02 15:04:05"))
			return
		}
	}

	rs = v
	return
}
