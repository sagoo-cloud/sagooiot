package common

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
)

type sPgSequences struct {
}

func PgSequences() *sPgSequences {
	return &sPgSequences{}
}

func init() {
	service.RegisterPgSequences(PgSequences())
}

// GetPgSequences 获取PG指定表序列信息
func (s *sPgSequences) GetPgSequences(ctx context.Context, tableName string, primaryKey string) (out *model.PgSequenceOut, err error) {
	if tableName == "" {
		err = gerror.New("表名不能为空")
		return
	}
	if primaryKey == "" {
		err = gerror.New("主键名不能为空")
		return
	}
	// 创建数据库连接
	db := g.DB()
	result, err := db.Query(ctx, "SELECT schemaname As schemaName, sequencename AS seqUesCeName, sequenceowner AS seqUesCeOwner,data_type AS dataType,start_value AS startVale,min_value AS minValue,max_value AS maxValue,increment_by AS incrementBy,cycle AS cycle,cache_size AS cacheSize,last_value AS lastVale FROM pg_sequences WHERE sequencename = '"+tableName+"_"+primaryKey+"_seq'")
	if err != nil {
		return
	}
	if err = gconv.Scan(result[0], &out); err != nil {
		return
	}
	return
}
