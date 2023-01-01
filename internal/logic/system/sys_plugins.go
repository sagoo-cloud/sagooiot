package system

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/service"
)

type sSysPlugins struct{}

func sysPluginsNew() *sSysPlugins {
	return &sSysPlugins{}
}
func init() {
	service.RegisterSysPlugins(sysPluginsNew())
}

// GetSysPluginsList 获取列表数据
func (s *sSysPlugins) GetSysPluginsList(ctx context.Context, in *model.GetSysPluginsListInput) (total, page int, list []*model.SysPluginsOutput, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysPlugins.Ctx(ctx)

		if in.KeyWord != "" {
			m = m.WhereLike(dao.SysPlugins.Columns().Name, "%"+in.KeyWord+"%")
			m = m.WhereLike(dao.SysPlugins.Columns().Title, "%"+in.KeyWord+"%")
			m = m.WhereLike(dao.SysPlugins.Columns().Intro, "%"+in.KeyWord+"%")
		}

		total, err = m.Count()
		if err != nil {
			err = gerror.New("获取总行数失败")
			return
		}
		page = in.PageNum
		if in.PageSize == 0 {
			in.PageSize = consts.PageSize
		}
		err = m.Page(page, in.PageSize).Order("start_time desc").Scan(&list)
		if err != nil {
			err = gerror.New("获取数据失败")
		}
	})
	return
}

// GetSysPluginsById 获取指定ID数据
func (s *sSysPlugins) GetSysPluginsById(ctx context.Context, id int) (out *model.SysPluginsOutput, err error) {
	err = dao.SysPlugins.Ctx(ctx).Where(dao.SysPlugins.Columns().Id, id).Scan(&out)
	return
}

// AddSysPlugins 添加数据
func (s *sSysPlugins) AddSysPlugins(ctx context.Context, in model.SysPluginsAddInput) (err error) {
	_, err = dao.SysPlugins.Ctx(ctx).Insert(in)
	return
}

// EditSysPlugins 修改数据
func (s *sSysPlugins) EditSysPlugins(ctx context.Context, in model.SysPluginsEditInput) (err error) {
	_, err = dao.SysPlugins.Ctx(ctx).FieldsEx(dao.SysPlugins.Columns().Id).Where(dao.SysPlugins.Columns().Id, in.Id).Update(in)
	return
}

// DeleteSysPlugins 删除数据
func (s *sSysPlugins) DeleteSysPlugins(ctx context.Context, Ids []int) (err error) {
	_, err = dao.SysPlugins.Ctx(ctx).Delete(dao.SysPlugins.Columns().Id+" in (?)", Ids)
	return
}

//SaveSysPlugins 存入插件数据，跟据插件类型与名称，数据中只保存一份
func (s *sSysPlugins) SaveSysPlugins(ctx context.Context, in model.SysPluginsAddInput) (err error) {
	var req = g.Map{
		dao.SysPlugins.Columns().Types: in.Types,
		dao.SysPlugins.Columns().Name:  in.Name,
	}
	res, err := dao.SysPlugins.Ctx(ctx).Where(req).One()
	if res != nil {
		_, err = dao.SysPlugins.Ctx(ctx).Data(in).Where(req).Update()

	} else {
		_, err = dao.SysPlugins.Ctx(ctx).Insert(in)
	}
	return
}

func (s *sSysPlugins) EditStatus(ctx context.Context, id int, status int) (err error) {
	//todo 进行插件的启停
	switch status {
	case 0:

	case 1:
	}
	_, err = dao.SysPlugins.Ctx(ctx).Data("status", status).Where(dao.SysPlugins.Columns().Id, id).Update()
	return
}
