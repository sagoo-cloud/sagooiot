package tdengine

import (
	"context"
	"reflect"
	"sagooiot/pkg/tsd"
	"testing"

	"github.com/gogf/gf/v2/database/gdb"

	_ "github.com/taosdata/driver-go/v3/taosRestful"
	_ "github.com/taosdata/driver-go/v3/taosWS"
)

func Test_sTdEngine_GetOne(t *testing.T) {
	type args struct {
		ctx  context.Context
		sql  string
		args []any
	}
	tests := []struct {
		name    string
		args    args
		wantRs  gdb.Record
		wantErr bool
	}{
		{
			name: "测试统计设备日志总数",
			args: args{
				ctx:  context.Background(),
				sql:  "select count(*) as num from ?",
				args: []any{"device_log"},
			},
		},
	}

	tsdDb := tsd.DB()
	defer tsdDb.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRs, err := tsdDb.GetTableDataOne(tt.args.ctx, tt.args.sql, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRs, tt.wantRs) {
				t.Errorf("GetOne() gotRs = %v, want %v", gotRs, tt.wantRs)
			}
		})
	}
}
