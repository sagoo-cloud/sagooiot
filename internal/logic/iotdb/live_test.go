package iotdb

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	_ "sagooiot/internal/logic/tdengine"
	"sagooiot/internal/model"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"
	"sagooiot/pkg/cache"
	"sagooiot/pkg/dcache"
	"sagooiot/pkg/tsd"
	"sagooiot/pkg/tsd/comm"
)

func TestLiveTSLWriteAndPlatformRead(t *testing.T) {
	p := os.Getenv("SAGOO_IOTDB_TEST_PORT")
	if p == "" {
		t.Skip("SAGOO_IOTDB_TEST_PORT is unset")
	}
	port, err := strconv.Atoi(p)
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()
	root := fmt.Sprintf("root.sagoo_tsl_test_%d", time.Now().UnixNano())
	config, err := gcfg.NewAdapterContent(fmt.Sprintf(`{"tsd":{"database":"IoTDB","iotdb":{"host":"127.0.0.1","port":%d,"database":%q}},"database":{"default":{"link":"mysql:unused:unused@tcp(127.0.0.1:1)/unused"}},"cache":{"adapter":"memory","fileDir":""}}`, port, root))
	if err != nil {
		t.Fatal(err)
	}
	previousConfig, previousEngine := g.Cfg().GetAdapter(), service.TdEngine()
	previousTSLTable, previousLogTable := service.TSLTable(), service.TdLogTable()
	defer g.Cfg().SetAdapter(previousConfig)
	defer service.RegisterTdEngine(previousEngine)
	defer service.RegisterTSLTable(previousTSLTable)
	defer service.RegisterTdLogTable(previousLogTable)
	g.Cfg().SetAdapter(config)
	registerServices()
	db, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	impl := service.TSLTable()
	if err := impl.CreateDatabase(ctx); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if _, err := db.ExecContext(ctx, "DELETE DATABASE "+root); err != nil {
			t.Error(err)
		}
	}()
	cache.SetAdapter(ctx)
	modelTSL := &model.TSL{Key: "product-1", Properties: []model.TSLProperty{{Key: "Temp", ValueType: model.TSLValueType{Type: "double"}}}}
	if err := impl.CreateStable(ctx, modelTSL); err != nil {
		t.Fatal(err)
	}
	if ok, err := impl.CheckStable(ctx, comm.ProductTableName(modelTSL.Key)); err != nil || !ok {
		t.Fatalf("schema check: %t (%v)", ok, err)
	}
	if err := impl.CreateTable(ctx, modelTSL.Key, "child-7"); err != nil {
		t.Fatal(err)
	}
	if ok, err := impl.CheckTable(ctx, comm.DeviceTableName("child-7")); err != nil || !ok {
		t.Fatalf("device check: %t (%v)", ok, err)
	}
	metadata := &model.DeviceOutput{DevDevice: &entity.DevDevice{Key: "child-7"}, Product: &entity.DevProduct{Key: modelTSL.Key}, TSL: modelTSL}
	if err := dcache.SetDeviceDetailInfo("child-7", metadata); err != nil {
		t.Fatal(err)
	}
	reported := time.Now().Truncate(time.Second)
	if err := impl.Insert(ctx, "gateway", model.ReportPropertyData{"Temp": {Value: 18.75, CreateTime: reported.Unix()}}, "child-7"); err != nil {
		t.Fatal(err)
	}
	platform := tsd.GetDB()
	defer platform.Close()
	row, err := platform.GetTableDataOne(ctx, "SELECT p_temp, p_temp_time FROM device_child_7 WHERE p_temp_time >= ? AND p_temp_time < ?", reported.UnixMilli(), reported.Add(time.Millisecond).UnixMilli())
	if err != nil || row["temp"].Float64() != 18.75 || row["temp_time"].Int64() != reported.UnixMilli() {
		t.Fatalf("TSL report via platform DB: %v (%v)", row, err)
	}
}
