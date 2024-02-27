package system

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/internal/consts"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/do"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"
	"strings"
)

type sSysCertificate struct{}

func sSysCertificateNew() *sSysCertificate {
	return &sSysCertificate{}
}
func init() {
	service.RegisterSysCertificate(sSysCertificateNew())
}

// GetList 获取列表数据
func (s *sSysCertificate) GetList(ctx context.Context, input *model.SysCertificateListInput) (total, page int, out []*model.SysCertificateListOut, err error) {
	m := dao.SysCertificate.Ctx(ctx)

	if input.Name != "" {
		m = m.WhereLike(dao.SysCertificate.Columns().Name, "%"+input.Name+"%")
	}
	if input.Status != -1 {
		m = m.Where(dao.SysCertificate.Columns().Status, input.Status)
	}

	total, err = m.Count()
	if err != nil {
		err = gerror.New("获取总行数失败")
		return
	}
	page = input.PageNum
	if input.PageSize == 0 {
		input.PageSize = consts.PageSize
	}
	err = m.Page(page, input.PageSize).OrderDesc(dao.SysCertificate.Columns().CreatedAt).Scan(&out)
	if err != nil {
		err = gerror.New("获取数据失败")
	}
	return
}

// GetInfoById 获取指定ID数据
func (s *sSysCertificate) GetInfoById(ctx context.Context, id int) (out *model.SysCertificateListOut, err error) {
	err = dao.SysCertificate.Ctx(ctx).Where(dao.SysCertificate.Columns().Id, id).Scan(&out)
	return
}

// Add 添加数据
func (s *sSysCertificate) Add(ctx context.Context, input *model.AddSysCertificateListInput) (err error) {
	if strings.TrimSpace(input.Name) == "" {
		err = gerror.New("名称不能为空")
		return
	}
	num, _ := dao.SysCertificate.Ctx(ctx).Where(g.Map{
		dao.SysCertificate.Columns().Name:      input.Name,
		dao.SysCertificate.Columns().IsDeleted: 0,
	}).Count()
	if num > 0 {
		err = gerror.New("证书已存在,无法重复添加")
	}
	var sysCertificate *entity.SysCertificate
	if err = gconv.Scan(input, &sysCertificate); err != nil {
		return
	}
	loginUserId := service.Context().GetUserId(ctx)
	/*sysCertificate.Status = 0
	sysCertificate.IsDeleted = 0
	sysCertificate.CreatedAt = gtime.Now()
	sysCertificate.CreatedBy = uint(loginUserId)*/
	_, err = dao.SysCertificate.Ctx(ctx).Data(do.SysCertificate{
		DeptId:            service.Context().GetUserDeptId(ctx),
		Name:              sysCertificate.Name,
		Standard:          sysCertificate.Standard,
		FileContent:       sysCertificate.FileContent,
		PublicKeyContent:  sysCertificate.PublicKeyContent,
		PrivateKeyContent: sysCertificate.PrivateKeyContent,
		Description:       sysCertificate.Description,
		Status:            0,
		IsDeleted:         0,
		CreatedBy:         uint(loginUserId),
		CreatedAt:         gtime.Now(),
	}).Insert()
	return
}

// Edit 修改数据
func (s *sSysCertificate) Edit(ctx context.Context, input *model.EditSysCertificateListInput) (err error) {
	if strings.TrimSpace(input.Name) == "" {
		err = gerror.New("名称不能为空")
		return
	}
	var sysCertificate *entity.SysCertificate
	err = dao.SysCertificate.Ctx(ctx).Where(g.Map{
		dao.SysCertificate.Columns().Id: input.Id,
	}).Scan(&sysCertificate)

	if sysCertificate == nil {
		err = gerror.New("ID错误")
	}

	if sysCertificate.IsDeleted == 1 {
		err = gerror.New("已删除,无法更新")
	}
	sysCertificate.Name = input.Name
	sysCertificate.Standard = input.Standard
	sysCertificate.FileContent = input.FileContent
	sysCertificate.PublicKeyContent = input.PublicKeyContent
	sysCertificate.PrivateKeyContent = input.PrivateKeyContent
	sysCertificate.Description = input.Description
	loginUserId := service.Context().GetUserId(ctx)
	sysCertificate.UpdatedBy = loginUserId
	sysCertificate.UpdatedAt = gtime.Now()
	_, err = dao.SysCertificate.Ctx(ctx).Data(sysCertificate).Where(dao.SysCertificate.Columns().Id, sysCertificate.Id).Update()
	return
}

// Delete 删除数据
func (s *sSysCertificate) Delete(ctx context.Context, id int) (err error) {
	var sysCertificate *entity.SysCertificate
	err = dao.SysCertificate.Ctx(ctx).Where(g.Map{
		dao.SysCertificate.Columns().Id: id,
	}).Scan(&sysCertificate)

	if sysCertificate == nil {
		err = gerror.New("ID错误")
	}
	if sysCertificate.IsDeleted == 1 {
		err = gerror.New("无法重复删除")
	}

	loginUserId := service.Context().GetUserId(ctx)
	sysCertificate.IsDeleted = 1
	sysCertificate.DeletedBy = loginUserId
	sysCertificate.DeletedAt = gtime.Now()
	_, err = dao.SysCertificate.Ctx(ctx).Data(sysCertificate).Where(dao.SysCertificate.Columns().Id, id).Update()
	return
}

// EditStatus 更新状态
func (s *sSysCertificate) EditStatus(ctx context.Context, id int, status int) (err error) {
	var sysCertificate *entity.SysCertificate
	err = dao.SysCertificate.Ctx(ctx).Where(g.Map{
		dao.SysCertificate.Columns().Id: id,
	}).Scan(&sysCertificate)

	if sysCertificate == nil {
		err = gerror.New("ID错误")
	}
	if sysCertificate.Status == status {
		err = gerror.New("已更新,无法重复更新")
	}

	loginUserId := service.Context().GetUserId(ctx)
	sysCertificate.Status = status
	sysCertificate.UpdatedBy = loginUserId
	sysCertificate.UpdatedAt = gtime.Now()
	_, err = dao.SysCertificate.Ctx(ctx).Data(sysCertificate).Where(dao.SysCertificate.Columns().Id, id).Update()
	return
}

// GetAll 获取所有证书
func (s *sSysCertificate) GetAll(ctx context.Context) (out []*entity.SysCertificate, err error) {
	m := dao.SysCertificate.Ctx(ctx)

	err = m.Where(g.Map{
		dao.SysCertificate.Columns().Status:    1,
		dao.SysCertificate.Columns().IsDeleted: 0,
	}).Scan(&out)
	return
}
