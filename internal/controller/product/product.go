package product

import (
	"context"
	"sagooiot/api/v1/product"
	"sagooiot/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

var Product = cProduct{}

type cProduct struct{}

//func (c *cProduct) Get(ctx context.Context, req *product.GetProductReq) (res *product.GetProductRes, err error) {
//	p, err := service.DevProduct().Get(ctx, req.Key)
//	// 获取产品的设备数量
//	totals, err := service.DevDevice().TotalByProductId(ctx, []uint{p.Id})
//	if err != nil {
//		return
//	}
//	p.DeviceTotal = totals[p.Id]
//	res = &product.GetProductRes{
//		Data: p,
//	}
//	return
//}

func (c *cProduct) Detail(ctx context.Context, req *product.DetailProductReq) (res *product.DetailProductRes, err error) {
	p, err := service.DevProduct().Detail(ctx, req.ProductKey)
	// 获取产品的设备数量
	totals, err := service.DevDevice().TotalByProductKey(ctx, []string{p.Key})
	if err != nil {
		return
	}
	p.DeviceTotal = totals[p.Key]
	res = &product.DetailProductRes{
		Data: p,
	}
	return
}

func (c *cProduct) ListForPage(ctx context.Context, req *product.ListForPageReq) (res *product.ListForPageRes, err error) {
	out, err := service.DevProduct().ListForPage(ctx, &req.ListForPageInput)
	res = &product.ListForPageRes{
		ListForPageOutput: *out,
	}
	return
}

func (c *cProduct) List(ctx context.Context, req *product.ListReq) (res *product.ListRes, err error) {
	list, err := service.DevProduct().List(ctx)
	res = &product.ListRes{
		Product: list,
	}
	return
}

func (c *cProduct) Add(ctx context.Context, req *product.AddProductReq) (res *product.AddProductRes, err error) {
	err = service.DevProduct().Add(ctx, req.AddProductInput)
	return
}

func (c *cProduct) Edit(ctx context.Context, req *product.EditProductReq) (res *product.EditProductRes, err error) {
	err = service.DevProduct().Edit(ctx, req.EditProductInput)
	return
}

func (c *cProduct) UpdateExtend(ctx context.Context, req *product.UpdateExtendReq) (res *product.UpdateExtendRes, err error) {
	err = service.DevProduct().UpdateExtend(ctx, req.ExtendInput)
	return
}

func (c *cProduct) Del(ctx context.Context, req *product.DelProductReq) (res *product.DelProductRes, err error) {
	err = service.DevProduct().Del(ctx, req.Keys)
	return
}

func (c *cProduct) Deploy(ctx context.Context, req *product.DeployProductReq) (res *product.DeployProductRes, err error) {
	err = service.DevProduct().Deploy(ctx, req.ProductKey)
	return
}

func (c *cProduct) Undeploy(ctx context.Context, req *product.UndeployProductReq) (res *product.UndeployProductRes, err error) {
	err = service.DevProduct().Undeploy(ctx, req.ProductKey)
	return
}

func (c *cProduct) UploadIcon(ctx context.Context, req *product.UploadIconReq) (res *product.UploadIconRes, err error) {
	if req.Icon == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请选择上传的图片")
	}

	uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()

	filename, err := req.Icon.Save(uploadPath+"/product_icon", true)
	if err != nil {
		err = gerror.New("图片上传失败")
		return
	}

	res = &product.UploadIconRes{
		IconPath: uploadPath + "/product_icon/" + filename,
	}

	return
}

func (c *cProduct) ListForSub(ctx context.Context, req *product.ListForSubProductReq) (res *product.ListForSubProductRes, err error) {
	list, err := service.DevProduct().ListForSub(ctx)
	res = &product.ListForSubProductRes{
		Product: list,
	}
	return
}

// UpdateScriptInfo 脚本更新
func (c *cProduct) UpdateScriptInfo(ctx context.Context, req *product.UpdateScriptInfoReq) (res *product.UpdateScriptInfoRes, err error) {
	err = service.DevProduct().UpdateScriptInfo(ctx, req.ScriptInfoInput)
	return
}

// ConnectIntro 获取设备接入信息
func (c *cProduct) ConnectIntro(ctx context.Context, req *product.ConnectIntroReq) (res *product.ConnectIntroRes, err error) {
	data, err := service.DevProduct().ConnectIntro(ctx, req.ProductKey)
	res = &product.ConnectIntroRes{
		Data: data,
	}
	return
}
