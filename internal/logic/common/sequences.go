package common

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"sagooiot/internal/consts"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
	"strings"
)

type sSequences struct {
}

func Sequences() *sSequences {
	return &sSequences{}
}

func init() {
	service.RegisterSequences(Sequences())
}

// GetSequences 获取主键ID
func (s *sSequences) GetSequences(ctx context.Context, result sql.Result, tableName string, primaryKey string) (lastInsertId int64, err error) {
	//获取数据源类型
	//TODO 多数据源情况下需对此进行更改优化
	databaseType := gdb.GetConfig(consts.DataBaseGroup)[0].Type
	if strings.EqualFold(databaseType, consts.DatabaseTypeMysql) {
		//获取自增主键
		lastInsertId, _ = result.LastInsertId()
	} else if strings.EqualFold(databaseType, consts.DatabaseTypePostgresql) {
		//获取当前表序列
		var pgSequenceOut *model.PgSequenceOut
		pgSequenceOut, err = service.PgSequences().GetPgSequences(ctx, tableName, primaryKey)
		if err != nil {
			return
		}
		if pgSequenceOut != nil {
			lastInsertId = pgSequenceOut.LastVale
		}
	} else {
		err = gerror.New("暂不支持" + databaseType + "的获取方式")
	}
	return
}
