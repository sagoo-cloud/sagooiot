package product

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/api/v1/product"
	"sagooiot/internal/service"
)

var TSLImport = cTSLImport{}

type cTSLImport struct{}

// ExportTSL 导出物模型
func (c *cTSLImport) ExportTSL(ctx context.Context, req *product.ExportTSLReq) (res *product.ExportTSLRes, err error) {
	g.Log().Debug(ctx, "====导出======", req)
	err = service.DevTSLImport().Export(ctx, req.ProductKey)
	return
}

// ImportTSL 导出物模型
func (c *cTSLImport) ImportTSL(ctx context.Context, req *product.ImportTSLReq) (res *product.ImportTSLRes, err error) {
	g.Log().Debug(ctx, "====导入======", req)
	if req.File == nil {
		err = gerror.New("上传文件必须")
		return
	}
	err = service.DevTSLImport().Import(ctx, req.ProductKey, req.File)
	return
}
