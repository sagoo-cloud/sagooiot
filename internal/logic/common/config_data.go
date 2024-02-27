package common

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/database/gdb"
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
	"sagooiot/pkg/cache"
)

type sConfigData struct {
}

func sysConfigDataNew() *sConfigData {
	return &sConfigData{}
}

func init() {
	service.RegisterConfigData(sysConfigDataNew())
}

// List 系统参数列表
func (s *sConfigData) List(ctx context.Context, input *model.ConfigDoInput) (total int, out []*model.SysConfigOut, err error) {
	m := dao.SysConfig.Ctx(ctx)
	if input != nil {
		if input.ConfigName != "" {
			m = m.WhereLike(dao.SysConfig.Columns().ConfigName, "%"+input.ConfigName+"%")
		}
		if input.ConfigType != "" {
			m = m.Where(dao.SysConfig.Columns().ConfigType, gconv.Int(input.ConfigType))
		}
		if input.ConfigKey != "" {
			m = m.WhereLike(dao.SysConfig.Columns().ConfigKey, "%"+input.ConfigKey+"%")
		}
		if len(input.DateRange) > 0 {
			m = m.WhereBetween(dao.SysConfig.Columns().CreatedAt, input.DateRange[0], input.DateRange[1])
		}
		if input.ModuleClassify != "" {
			m = m.Where(dao.SysConfig.Columns().ModuleClassify, input.ModuleClassify)
		}
	}
	total, err = m.Count()
	if err != nil {
		return 0, nil, errors.New("获取数据失败")
	}
	if input.PageNum == 0 {
		input.PageNum = 1
	}
	if input.PageSize == 0 {
		input.PageSize = consts.PageSize
	}
	err = m.Page(input.PageNum, input.PageSize).Order("config_id desc").Scan(&out)
	if err != nil {
		return 0, nil, errors.New("获取数据失败")
	}
	return
}

func (s *sConfigData) Add(ctx context.Context, input *model.AddConfigInput, userId int) (err error) {
	err = s.CheckConfigKeyUnique(ctx, input.ConfigKey)
	if err != nil {
		return
	}
	data := &do.SysConfig{
		ConfigName:     input.ConfigName,
		ConfigKey:      input.ConfigKey,
		ConfigValue:    input.ConfigValue,
		ConfigType:     input.ConfigType,
		CreatedBy:      userId,
		Remark:         input.Remark,
		ModuleClassify: input.ModuleClassify,
		Status:         1,
		IsDeleted:      0,
	}
	_, err = dao.SysConfig.Ctx(ctx).Insert(data)
	if err != nil {
		return errors.New("添加系统参数失败")
	}

	//添加到缓存
	err = cache.Instance().Set(ctx, consts.SystemConfigPrefix+input.ConfigKey, data, 0)
	if err != nil {
		return
	}

	return
}

// CheckConfigKeyUnique 验证参数键名是否存在
func (s *sConfigData) CheckConfigKeyUnique(ctx context.Context, configKey string, configId ...int) (err error) {
	data := (*entity.SysConfig)(nil)
	m := dao.SysConfig.Ctx(ctx).Fields(dao.SysConfig.Columns().ConfigId).Where(dao.SysConfig.Columns().ConfigKey, configKey)
	if len(configId) > 0 {
		m = m.Where(dao.SysConfig.Columns().ConfigId+" != ?", configId[0])
	}
	err = m.Scan(&data)
	if err != nil {
		return
	}
	if data != nil {
		return errors.New("参数键名重复")
	}

	return
}

// Get 获取系统参数
func (s *sConfigData) Get(ctx context.Context, id int) (out *model.SysConfigOut, err error) {
	err = dao.SysConfig.Ctx(ctx).WherePri(id).Scan(&out)
	if err != nil {
		return nil, errors.New("获取系统参数失败")
	}
	return
}

// Edit 修改系统参数
func (s *sConfigData) Edit(ctx context.Context, input *model.EditConfigInput, userId int) (err error) {
	err = s.CheckConfigKeyUnique(ctx, input.ConfigKey, input.ConfigId)
	if err != nil {
		return errors.New("参数键名重复")
	}
	data := &do.SysConfig{
		ConfigName:     input.ConfigName,
		ConfigKey:      input.ConfigKey,
		ConfigValue:    input.ConfigValue,
		ConfigType:     input.ConfigType,
		UpdatedBy:      userId,
		Remark:         input.Remark,
		ModuleClassify: input.ModuleClassify,
		Status:         1,
		IsDeleted:      0,
	}
	_, err = dao.SysConfig.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     "ConfigDataByKey",
		Force:    false,
	}).WherePri(input.ConfigId).Update(data)
	if err != nil {
		return errors.New("修改系统参数失败")
	}
	//更新缓存
	_, _, err = cache.Instance().Update(ctx, consts.SystemConfigPrefix+input.ConfigKey, data)

	return
}

// Delete 删除系统参数 //TODO 转为KEY处理
func (s *sConfigData) Delete(ctx context.Context, ids []int) (err error) {
	_, err = dao.SysConfig.Ctx(ctx).Delete(dao.SysConfig.Columns().ConfigId+" in (?)", ids)
	if err != nil {
		return errors.New("删除失败")
	}
	//清除缓存
	_, err = cache.Instance().Remove(ctx, consts.SystemConfigPrefix)

	return
}

// GetConfigByKey 通过key获取参数（从缓存获取）
func (s *sConfigData) GetConfigByKey(ctx context.Context, key string) (config *entity.SysConfig, err error) {
	if key == "" {
		err = gerror.New("参数key不能为空")
		return
	}
	cf, err := cache.Instance().Get(ctx, consts.SystemConfigPrefix+key)
	if cf != nil && !cf.IsEmpty() {
		err = gconv.Struct(cf.Val(), &config)
		return
	} else {
		config, err = s.GetByKey(ctx, key)
		if err != nil {
			return
		}
		if config != nil {
			err = cache.Instance().Set(ctx, consts.SystemConfigPrefix+key, config, 0)
			if err != nil {
				return
			}
		}
	}
	return
}

// GetConfigByKeys 通过key数组获取参数（从缓存获取）
func (s *sConfigData) GetConfigByKeys(ctx context.Context, keys []string) (out []*entity.SysConfig, err error) {

	for _, key := range keys {
		var config *entity.SysConfig
		config, err = s.GetConfigByKey(ctx, key)
		if err != nil {
			return
		}
		out = append(out, config)
	}
	return
}

// GetByKey 通过key获取参数（从数据库获取）
func (s *sConfigData) GetByKey(ctx context.Context, key string) (config *entity.SysConfig, err error) {
	err = dao.SysConfig.Ctx(ctx).Where(dao.SysConfig.Columns().ConfigKey, key).Scan(&config)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("获取配置失败")
	}
	return
}

// GetByKeys 通过keys获取参数（从数据库获取）
func (s *sConfigData) GetByKeys(ctx context.Context, keys []string) (config []*entity.SysConfig, err error) {
	err = dao.SysConfig.Ctx(ctx).WhereIn(dao.SysConfig.Columns().ConfigKey, keys).Scan(&config)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("获取配置失败")
	}
	return
}

func (s *sConfigData) GetSysConfigSetting(ctx context.Context, types int) (out []*entity.SysConfig, err error) {
	var keys []string
	if types == 0 {
		keys = []string{
			consts.SysSystemName,
			consts.SysSystemCopyright,
			consts.SysSystemLogo,
			consts.SysSystemLoginPic,
			consts.SysSystemLogoMini,
		}
	} else if types == 1 {
		keys = []string{
			consts.SysColumnSwitch,
			consts.SysButtonSwitch,
			consts.SysApiSwitch,
			consts.SysIsSingleLogin,
			consts.SysTokenExpiryDate,
			consts.SysPasswordChangePeriod,
			consts.SysPasswordErrorNum,
			consts.SysAgainLoginDate,
			consts.SysPasswordMinimumLength,
			consts.SysRequireComplexity,
			consts.SysRequireDigit,
			consts.SysRequireLowercaseLetter,
			consts.SysRequireUppercaseLetter,
			consts.SysIsSecurityControlEnabled,
			consts.SysChangePasswordForFirstLogin,
			consts.SysPasswordChangePeriodSwitch,
			consts.SysIsRsaEnabled,
		}
	}
	if len(keys) == 0 {
		err = gerror.New("类型选择错误")
		return
	}
	out, err = s.GetByKeys(ctx, keys)
	if err != nil {
		return
	}
	return
}

// EditSysConfigSetting 修改系统配置设置
func (s *sConfigData) EditSysConfigSetting(ctx context.Context, inputs []*model.EditConfigInput) (err error) {
	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)
	for _, input := range inputs {
		err = s.CheckConfigKeyUnique(ctx, input.ConfigKey, input.ConfigId)
		if err != nil {
			return errors.New("参数键名重复")
		}
		_, err = dao.SysConfig.Ctx(ctx).Where(dao.SysConfig.Columns().ConfigKey, input.ConfigKey).Update(do.SysConfig{
			ConfigValue: input.ConfigValue,
			UpdatedBy:   loginUserId,
			UpdatedAt:   gtime.Now(),
		})
		if err != nil {
			return errors.New("修改系统基础配置失败")
		}
		//清除缓存
		_, err = cache.Instance().Remove(ctx, consts.SystemConfigPrefix+input.ConfigKey)
	}
	return
}

// GetLoadCache 获取本地缓存配置
func (s *sConfigData) GetLoadCache(ctx context.Context) (conf *model.CacheConfig, err error) {
	err = g.Cfg().MustGet(ctx, "cache").Scan(&conf)
	return
}
