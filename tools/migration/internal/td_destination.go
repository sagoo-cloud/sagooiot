package internal

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"strconv"
	"strings"
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

const pageSize = 3000

func (s *TdDestination) InsertData() (err error) {
	tdSrc := NewTdSource()
	tables, err := tdSrc.Tables()
	if err != nil || tables.Len() == 0 {
		return
	}

	for _, rs := range tables {
		tb := rs["table_name"].String()
		count, _ := tdSrc.Count(tb)
		if count == 0 {
			continue
		}
		println("insert table : " + tb + ", data count: " + strconv.Itoa(count))

		for i := 1; i <= int(math.Ceil(float64(count)/pageSize)); i++ {
			data, err := tdSrc.Data(tb, i, pageSize)
			if err != nil || data.Len() == 0 {
				continue
			}

			var (
				fed []string
				vle []string
			)
			fed = data[0].GMap().Keys()

			for _, row := range data {
				var vs []string
				for _, k := range fed {
					vs = append(vs, "'"+row[k].String()+"'")
				}
				vle = append(vle, "("+strings.Join(vs, ",")+")")
			}

			sqlStr := fmt.Sprintf("insert into %s (%s) values %s", tb, strings.Join(fed, ","), strings.Join(vle, ","))
			if _, err = GetConn(context.Background(), "tdengineDest").Exec(sqlStr); err != nil {
				return err
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
