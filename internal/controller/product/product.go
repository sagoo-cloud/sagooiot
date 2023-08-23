package product

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/api/v1/product"
	"github.com/sagoo-cloud/sagooiot/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

var Product = cProduct{}

type cProduct struct{}

func (c *cProduct) Get(ctx context.Context, req *product.GetProductReq) (res *product.GetProductRes, err error) {
	p, err := service.DevProduct().Get(ctx, req.Key)
	res = &product.GetProductRes{
		Data: p,
	}
	return
}

func (c *cProduct) Detail(ctx context.Context, req *product.DetailProductReq) (res *product.DetailProductRes, err error) {
	p, err := service.DevProduct().Detail(ctx, req.Id)
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
	list, err := service.DevProduct().List(ctx, req.ListProductInput)
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

func (c *cProduct) Del(ctx context.Context, req *product.DelProductReq) (res *product.DelProductRes, err error) {
	err = service.DevProduct().Del(ctx, req.Ids)
	return
}

func (c *cProduct) Deploy(ctx context.Context, req *product.DeployProductReq) (res *product.DeployProductRes, err error) {
	err = service.DevProduct().Deploy(ctx, req.Id)
	return
}

func (c *cProduct) Undeploy(ctx context.Context, req *product.UndeployProductReq) (res *product.UndeployProductRes, err error) {
	err = service.DevProduct().Undeploy(ctx, req.Id)
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
