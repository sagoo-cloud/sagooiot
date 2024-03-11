package tdengine

import (
	"context"
	"database/sql"
	"sagooiot/internal/consts"
	"sagooiot/internal/service"
	"sync"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type sTdEngine struct{}

var _instance *sTdEngine
var once sync.Once

// 使用单例模式创建sTdEngine实例
func tdEngineNew() *sTdEngine {
	once.Do(func() {
		_instance = &sTdEngine{}
	})
	return _instance
}

var dbName = consts.TdEngineDbName

// 初始化函数
func init() {
	service.RegisterTdEngine(tdEngineNew())
	// 简化配置获取过程
	dbName = g.Cfg().MustGet(context.Background(), consts.TdEngineDbNameKey, consts.TdEngineDbName).String()
}

type connections struct {
	tdEngineType  string
	tdDsn         string
	ConnectionMap map[string]*sql.DB
	sync.RWMutex
}

var connectionMap *connections

// GetConn 获取数据库连接
func (s *sTdEngine) GetConn(ctx context.Context, dbName string) (*sql.DB, error) {
	if connectionMap == nil {
		// 一次性初始化connectionMap
		connectionMap = initConnectionMap(ctx)
		if connectionMap == nil {
			return nil, gerror.New("TDengine配置初始化失败")
		}
	}

	connectionMap.RLock()
	connection, exists := connectionMap.ConnectionMap[dbName]
	connectionMap.RUnlock()

	if !exists {
		return createNewConnection(ctx, dbName)
	}

	return connection, nil
}

// 初始化连接映射
func initConnectionMap(ctx context.Context) *connections {
	driver := g.Cfg().MustGet(ctx, consts.TdEngineTypeKey)
	dsn := g.Cfg().MustGet(ctx, consts.TdEngineDsnKey)
	return &connections{
		tdEngineType:  driver.String(),
		tdDsn:         dsn.String(),
		ConnectionMap: make(map[string]*sql.DB),
	}
}

// 创建新连接
func createNewConnection(ctx context.Context, dbName string) (*sql.DB, error) {
	connectionMap.Lock()
	defer connectionMap.Unlock()

	// 再次检查以防止竞态条件
	if conn, exists := connectionMap.ConnectionMap[dbName]; exists {
		return conn, nil
	}

	connection, err := sql.Open(connectionMap.tdEngineType, connectionMap.tdDsn+dbName)
	if err != nil {
		return nil, gerror.Wrap(err, "TDengine连接失败")
	}

	// 配置连接池设置
	// 配置连接池设置
	connection.SetMaxIdleConns(g.Cfg().MustGet(ctx, consts.TdEngineMaxIdleConnsKey).Int())
	connection.SetMaxOpenConns(g.Cfg().MustGet(ctx, consts.TdEngineMaxOpenConnsKey).Int())

	connectionMap.ConnectionMap[dbName] = connection
	return connection, nil
}

// Time REST连接时区处理
func (s *sTdEngine) Time(v *g.Var) (rs *g.Var) {
	driver := g.Cfg().MustGet(context.TODO(), consts.TdEngineTypeKey)
	if driver.String() == "taosRestful" {
		if t, err := time.Parse("2006-01-02 15:04:05 +0000 UTC", v.String()); err == nil {
			rs = gvar.New(t.Local().Format("2006-01-02 15:04:05"))
			return
		}
	}

	rs = v
	return
}

func Close() {
	if connectionMap == nil {
		return
	}
	connectionMap.Lock()
	for _, node := range connectionMap.ConnectionMap {
		if err := node.Close(); err != nil {
			return
		}
	}
	connectionMap.Unlock()
}

// ClearLogByDays 删除指定天数的设备日志数据
func (s *sTdEngine) ClearLogByDays(ctx context.Context, days int) (err error) {
	db, err := s.GetConn(ctx, dbName)
	if err != nil {
		return
	}

	//
	brforeTime := gtime.Now().AddDate(0, 0, -days).Format("Y-m-d H:i:s.u")
	sqlStr := "delete from device_log  where ts < ?"
	rows, err := db.Query(sqlStr, brforeTime)
	if err != nil {
		g.Log().Error(ctx, err, sqlStr)
		return
	}
	defer rows.Close()
	return
}
