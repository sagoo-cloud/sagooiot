package internal

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type TdSource struct {
	db *sql.DB
}

func NewTdSource() *TdSource {
	db := GetConn(context.Background(), "tdengineSrc")
	return &TdSource{db: db}
}

func (s *TdSource) ShowCreateStable() (stables, createSql []string, err error) {
	list, err := s.Stables()
	if err != nil || list.Len() == 0 {
		return
	}
	if list == nil {
		gerror.New("数据为空")
		return
	}

	for _, v := range list {
		stable := v["stable_name"].String()
		sqlStr := "show create stable " + stable
		rs, err := s.GetAll(sqlStr)
		if err != nil || rs.Len() == 0 {
			continue
		}
		stables = append(stables, stable)
		createSql = append(createSql, rs[0]["Create Table"].String())
	}
	return
}

func (s *TdSource) ShowCreateTable() (tables, createSql []string, err error) {
	list, err := s.Tables()
	if err != nil || list.Len() == 0 {
		return
	}
	if list == nil {
		return
	}

	for _, v := range list {
		table := v["table_name"].String()
		sqlStr := "show create table " + table
		rs, err := s.GetAll(sqlStr)
		if err != nil || rs.Len() == 0 {
			continue
		}
		tables = append(tables, table)
		createSql = append(createSql, rs[0]["Create Table"].String())
	}
	return
}

func (s *TdSource) Stables() (rs gdb.Result, err error) {
	sqlStr := "show stables"
	rs, err = s.GetAll(sqlStr)
	return
}

func (s *TdSource) Tables() (rs gdb.Result, err error) {
	sqlStr := "show tables"
	rs, err = s.GetAll(sqlStr)
	return
}

func (s *TdSource) Count(table string) (total int, err error) {
	sqlStr := "select count(*) as num from " + table
	rs, err := s.GetAll(sqlStr)
	if err != nil || rs.Len() == 0 {
		return
	}
	total = rs[0]["num"].Int()
	return
}

func (s *TdSource) Data(table string, pageNum, pageSize int) (rs gdb.Result, err error) {
	sqlStr := fmt.Sprintf("select * from %s limit %d, %d", table, (pageNum-1)*pageSize, pageSize)
	rs, err = s.GetAll(sqlStr)
	return
}

// 超级表查询，多条数据
func (s *TdSource) GetAll(sql string, args ...any) (rs gdb.Result, err error) {
	rows, err := s.db.Query(sql, args...)
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
			return nil, err
		}

		m := make(gdb.Record, len(columns))
		for i, c := range columns {
			m[c] = s.Time(gvar.New(values[i]))
		}
		rs = append(rs, m)
	}
	return
}

// REST连接时区处理
func (s *TdSource) Time(v *g.Var) (rs *g.Var) {
	if t, err := time.Parse("2006-01-02 15:04:05 +0000 UTC", v.String()); err == nil {
		rs = gvar.New(t.Local().Format("2006-01-02 15:04:05"))
	} else {
		rs = v
	}
	return
}
