package datahub

import (
	"context"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func TestCreateTable(t *testing.T) {
	table, err := createTable(context.TODO(), 1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(table)
}
