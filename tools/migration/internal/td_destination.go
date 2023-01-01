package internal

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/container/gvar"
)

type TdDestination struct {
	db *sql.DB
}

func NewTdDestination() *TdDestination {
	db := GetConn(context.Background(), "tdengineDest")
	return &TdDestination{db: db}
}

func (s *TdDestination) CreateStables() (err error) {
	tdSrc := NewTdSource()
	stables, createSql, err := tdSrc.ShowCreateStable()
	if err != nil || len(stables) == 0 || len(createSql) == 0 {
		return
	}

	for i, v := range stables {
		if err = s.DropStable(v); err == nil {
			_, err = s.db.Exec(createSql[i])
		}
	}
	return
}

func (s *TdDestination) CreateTables() (err error) {
	tdSrc := NewTdSource()
	tables, createSql, err := tdSrc.ShowCreateTable()
	if err != nil {
		return
	}
	if err != nil || len(tables) == 0 || len(createSql) == 0 {
		return
	}

	for i, v := range tables {
		if err = s.DropTable(v); err == nil {
			_, err = s.db.Exec(createSql[i])
		}
	}
	return
}

func (s *TdDestination) InsertData() (err error) {
	tdSrc := NewTdSource()
	tables, err := tdSrc.Tables()
	if err != nil || tables.Len() == 0 {
		return
	}

	for _, rs := range tables {
		tb := rs["table_name"].String()
		count, _ := tdSrc.Count(tb)
		if count > 0 {
			data, err := tdSrc.Data(tb, 1, 1000)
			if err != nil || data.Len() == 0 {
				continue
			}

			for _, row := range data {
				var (
					field []string
					value []string
				)
				for k, v := range row {
					field = append(field, k)
					value = append(value, "'"+gvar.New(v).String()+"'")
				}
				sqlStr := fmt.Sprintf("insert into %s (%s) values (%s)", tb, strings.Join(field, ","), strings.Join(value, ","))
				if _, err = s.db.Exec(sqlStr); err != nil {
					return err
				}
			}
		}
	}

	return
}

func (s *TdDestination) DropStable(stable string) (err error) {
	sqlStr := "drop stable if exists " + stable
	_, err = s.db.Exec(sqlStr)
	return
}

func (s *TdDestination) DropTable(table string) (err error) {
	sqlStr := "drop table if exists " + table
	_, err = s.db.Exec(sqlStr)
	return
}
