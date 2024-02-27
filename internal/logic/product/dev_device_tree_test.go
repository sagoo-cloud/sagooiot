package product

import (
	"context"
	"testing"

	"github.com/gogf/gf/v2/frame/g"
)

func TestDevDeviceTreeNew(t *testing.T) {
	treeObj := devDeviceTreeNew()
	list, err := treeObj.List(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(list)
}
