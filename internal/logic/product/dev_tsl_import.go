package product

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"io"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"
	"sagooiot/pkg/response"
)

type sDevTSLImport struct{}

func init() {
	service.RegisterDevTSLImport(devTSLImportNew())
}

func devTSLImportNew() *sDevTSLImport {
	return &sDevTSLImport{}
}

// Export 导出物模型
func (s *sDevTSLImport) Export(ctx context.Context, key string) (err error) {
	var product *entity.DevProduct
	err = dao.DevProduct.Ctx(ctx).WithAll().Where(dao.DevProduct.Columns().Key, key).Scan(&product)
	jsonData := gjson.New(product.Metadata).MustToJson()
	reader := bytes.NewReader(jsonData)
	var request = g.RequestFromCtx(ctx)
	response.ToJsonFIle(request, reader, "TSL-"+key+"-")

	return
}

// Import 导入物模型
func (s *sDevTSLImport) Import(ctx context.Context, key string, file *ghttp.UploadFile) (err error) {
	jsonData, err := file.Open()
	if err != nil {
		return err
	}
	data, err := io.ReadAll(jsonData)
	if len(data) < 0 || err != nil {
		return err
	}
	var p *entity.DevProduct
	err = dao.DevProduct.Ctx(ctx).Where(dao.DevProduct.Columns().Key, key).Scan(&p)
	if p == nil || err != nil {
		return gerror.New("产品不存在")
	}

	var tsl *model.TSL
	err = json.Unmarshal(data, &tsl)
	if err != nil {
		return
	}
	tsl.Key = p.Key
	tsl.Name = p.Name

	_, err = dao.DevProduct.Ctx(ctx).
		Data(dao.DevProduct.Columns().Metadata, gconv.String(tsl)).
		Where(dao.DevProduct.Columns().Key, key).
		Update()
	return
}
