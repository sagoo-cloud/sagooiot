package system

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gyaml"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/service"
)

type sSystemPluginsConfig struct{}

func sSystemPluginsConfigNew() *sSystemPluginsConfig {
	return &sSystemPluginsConfig{}
}
func init() {
	service.RegisterSystemPluginsConfig(sSystemPluginsConfigNew())
}

//GetPluginsConfigList 获取列表数据
func (s *sSystemPluginsConfig) GetPluginsConfigList(ctx context.Context, in *model.GetPluginsConfigListInput) (total, page int, list []*model.PluginsConfigOutput, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysPluginsConfig.Ctx(ctx)
		total, err = m.Count()
		if err != nil {
			err = gerror.New("获取总行数失败")
			return
		}
		page = in.PageNum
		if in.PageSize == 0 {
			in.PageSize = consts.PageSize
		}
		err = m.Page(page, in.PageSize).Order("created_at desc").Scan(&list)
		if err != nil {
			err = gerror.New("获取数据失败")
		}
	})
	return
}

//GetPluginsConfigById 获取指定ID数据
func (s *sSystemPluginsConfig) GetPluginsConfigById(ctx context.Context, id int) (out *model.PluginsConfigOutput, err error) {
	err = dao.SysPluginsConfig.Ctx(ctx).Where(dao.SysPluginsConfig.Columns().Id, id).Scan(&out)
	return
}

//GetPluginsConfigByName 获取指定ID数据
func (s *sSystemPluginsConfig) GetPluginsConfigByName(ctx context.Context, types, name string) (out *model.PluginsConfigOutput, err error) {
	var reqData = g.Map{
		dao.SysPluginsConfig.Columns().Type: types,
		dao.SysPluginsConfig.Columns().Name: name,
	}
	err = dao.SysPluginsConfig.Ctx(ctx).Where(reqData).Scan(&out)
	return
}

//AddPluginsConfig 添加数据
func (s *sSystemPluginsConfig) AddPluginsConfig(ctx context.Context, in model.PluginsConfigAddInput) (err error) {
	_, err = dao.SysPluginsConfig.Ctx(ctx).Insert(in)
	return
}

//EditPluginsConfig 修改数据
func (s *sSystemPluginsConfig) EditPluginsConfig(ctx context.Context, in model.PluginsConfigEditInput) (err error) {
	_, err = dao.SysPluginsConfig.Ctx(ctx).FieldsEx(dao.SysPluginsConfig.Columns().Id).Where(dao.SysPluginsConfig.Columns().Id, in.Id).Update(in)
	return
}

//SavePluginsConfig 更新数据，有数据就修改，没有数据就添加
func (s *sSystemPluginsConfig) SavePluginsConfig(ctx context.Context, in model.PluginsConfigAddInput) (err error) {
	var reqData = g.Map{
		dao.SysPluginsConfig.Columns().Id: in.Type,
		dao.SysPluginsConfig.Columns().Id: in.Name,
	}
	_, err = dao.SysPluginsConfig.Ctx(ctx).Where(reqData).Save(in)
	if err != nil {
		return
	}
	err = s.updateCache(ctx, in.Type, in.Name, in.Value)

	return
}

//DeletePluginsConfig 删除数据
func (s *sSystemPluginsConfig) DeletePluginsConfig(ctx context.Context, Ids []int) (err error) {
	_, err = dao.SysPluginsConfig.Ctx(ctx).Delete(dao.SysPluginsConfig.Columns().Id+" in (?)", Ids)
	return
}

//UpdateAllPluginsConfigCache 将插件数据更新到缓存
func (s *sSystemPluginsConfig) UpdateAllPluginsConfigCache() (err error) {

	var dataList []*model.PluginsConfigOutput
	err = dao.SysPluginsConfig.Ctx(context.TODO()).Scan(&dataList)
	if err != nil {
		return
	}

	var (
		ctx = gctx.New()
	)
	for _, datum := range dataList {
		err = s.updateCache(ctx, datum.Type, datum.Name, datum.Value)
	}
	return
}

func (s *sSystemPluginsConfig) updateCache(ctx context.Context, pluginsType, name, value string) (err error) {

	//缓存的key的规则是："plugins+插件类型+插件名称
	key := "plugins" + pluginsType + name
	_, err = g.Redis().Do(ctx, "SET", key, value)

	return
}

//GetPluginsConfigData 获取列表数据
func (s *sSystemPluginsConfig) GetPluginsConfigData(pluginType, pluginName string) (res map[interface{}]interface{}, err error) {
	key := "plugins" + pluginType + pluginName
	fmt.Println(key)
	pcgData, err := g.Redis().Do(context.TODO(), "GET", key)

	err = gyaml.DecodeTo([]byte(pcgData.String()), &res)

	return
}
