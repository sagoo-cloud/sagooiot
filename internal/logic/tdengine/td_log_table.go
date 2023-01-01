package tdengine

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 设备日志 TDengine 表结构维护
type sTdLogTable struct{}

func init() {
	service.RegisterTdLogTable(tdLogTableNew())
}

func tdLogTableNew() *sTdLogTable {
	return &sTdLogTable{}
}

// 添加超级表
func (s *sTdLogTable) CreateStable(ctx context.Context) (err error) {
	taos, err := service.TdEngine().GetConn(ctx, dbName)
	if err != nil {
		err = gerror.New("获取链接失败")
		return
	}
	defer taos.Close()

	var name string
	err = taos.QueryRow("SELECT stable_name FROM information_schema.ins_stables WHERE stable_name = 'device_log' LIMIT 1").Scan(&name)
	if name != "" {
		return
	}

	sql := "CREATE STABLE device_log (ts TIMESTAMP, type VARCHAR(20), content VARCHAR(1000)) TAGS (device VARCHAR(255))"
	_, err = taos.Exec(sql)

	return
}

// 写入数据
func (s *sTdLogTable) Insert(ctx context.Context, log *model.TdLogAddInput) (err error) {
	taos, err := service.TdEngine().GetConn(ctx, dbName)
	if err != nil {
		err = gerror.New("获取链接失败")
		return
	}
	defer taos.Close()

	sql := "INSERT INTO ? USING device_log TAGS ('?') VALUES ('?', '?', '?')"
	_, err = taos.Exec(sql, "log_"+log.Device, log.Device, log.Ts.String(), log.Type, log.Content)

	return
}

// 清理过期数据
func (s *sTdLogTable) Clear(ctx context.Context) (err error) {
	taos, err := service.TdEngine().GetConn(ctx, dbName)
	if err != nil {
		err = gerror.New("获取链接失败")
		return
	}
	defer taos.Close()

	ts := gtime.Now().Add(-7 * 24 * time.Hour).Format("Y-m-d")

	sql := "DELETE FROM device_log WHERE ts < '" + ts + "'"
	_, err = taos.Exec(sql)

	return
}

// 超级表查询，多条数据
func (s *sTdLogTable) GetAll(ctx context.Context, sql string, args ...any) (list []model.TdLog, err error) {
	taos, err := service.TdEngine().GetConn(ctx, dbName)
	if err != nil {
		err = gerror.New("获取链接失败")
		return
	}
	defer taos.Close()

	rows, err := taos.Query(sql, args...)
	if err != nil {
		g.Log().Error(ctx, err, sql, args)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var log model.TdLog

		err = rows.Scan(&log.Ts, &log.Type, &log.Content, &log.Device)
		if err != nil {
			return nil, err
		}
		log.Ts = service.TdEngine().Time(gvar.New(log.Ts.Format("Y-m-d H:i:s O T"))).GTime()

		list = append(list, log)
	}

	return
}
