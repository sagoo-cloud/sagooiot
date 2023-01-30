package common

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/do"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"github.com/sagoo-cloud/sagooiot/utility/liberr"
	"time"
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
			m = m.Where("config_name like ?", "%"+input.ConfigName+"%")
		}
		if input.ConfigType != "" {
			m = m.Where("config_type = ", gconv.Int(input.ConfigType))
		}
		if input.ConfigKey != "" {
			m = m.Where("config_key like ?", "%"+input.ConfigKey+"%")
		}
		if len(input.DateRange) > 0 {
			m = m.Where("created_at >= ? AND created_at<=?", input.DateRange[0], input.DateRange[1])
		}
	}
	total, err = m.Count()
	liberr.ErrIsNil(ctx, err, "获取数据失败")
	if input.PageNum == 0 {
		input.PageNum = 1
	}
	if input.PageSize == 0 {
		input.PageSize = consts.PageSize
	}
	err = m.Page(input.PageNum, input.PageSize).Order("config_id desc").Scan(&out)
	liberr.ErrIsNil(ctx, err, "获取数据失败")
	return
}

func (s *sConfigData) Add(ctx context.Context, input *model.AddConfigInput, userId int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = s.CheckConfigKeyUnique(ctx, input.ConfigKey)
		liberr.ErrIsNil(ctx, err)
		_, err = dao.SysConfig.Ctx(ctx).Insert(do.SysConfig{
			ConfigName:  input.ConfigName,
			ConfigKey:   input.ConfigKey,
			ConfigValue: input.ConfigValue,
			ConfigType:  input.ConfigType,
			CreateBy:    userId,
			Remark:      input.Remark,
		})
		liberr.ErrIsNil(ctx, err, "添加系统参数失败")
		//清除缓存
		Cache().RemoveByTag(ctx, consts.CacheSysConfigTag)
	})
	return
}

// CheckConfigKeyUnique 验证参数键名是否存在
func (s *sConfigData) CheckConfigKeyUnique(ctx context.Context, configKey string, configId ...int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		data := (*entity.SysConfig)(nil)
		m := dao.SysConfig.Ctx(ctx).Fields(dao.SysConfig.Columns().ConfigId).Where(dao.SysConfig.Columns().ConfigKey, configKey)
		if len(configId) > 0 {
			m = m.Where(dao.SysConfig.Columns().ConfigId+" != ?", configId[0])
		}
		err = m.Scan(&data)
		liberr.ErrIsNil(ctx, err, "校验失败")
		if data != nil {
			liberr.ErrIsNil(ctx, errors.New("参数键名重复"))
		}
	})
	return
}

// Get 获取系统参数
func (s *sConfigData) Get(ctx context.Context, id int) (out *model.SysConfigOut, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysConfig.Ctx(ctx).WherePri(id).Scan(&out)
		liberr.ErrIsNil(ctx, err, "获取系统参数失败")
	})
	return
}

// Edit 修改系统参数
func (s *sConfigData) Edit(ctx context.Context, input *model.EditConfigInput, userId int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = s.CheckConfigKeyUnique(ctx, input.ConfigKey, input.ConfigId)
		liberr.ErrIsNil(ctx, err)
		_, err = dao.SysConfig.Ctx(ctx).WherePri(input.ConfigId).Update(do.SysConfig{
			ConfigName:  input.ConfigName,
			ConfigKey:   input.ConfigKey,
			ConfigValue: input.ConfigValue,
			ConfigType:  input.ConfigType,
			UpdateBy:    userId,
			Remark:      input.Remark,
		})
		liberr.ErrIsNil(ctx, err, "修改系统参数失败")
		//清除缓存
		Cache().RemoveByTag(ctx, consts.CacheSysConfigTag)
	})
	return
}

// Delete 删除系统参数
func (s *sConfigData) Delete(ctx context.Context, ids []int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysConfig.Ctx(ctx).Delete(dao.SysConfig.Columns().ConfigId+" in (?)", ids)
		liberr.ErrIsNil(ctx, err, "删除失败")
		//清除缓存
		Cache().RemoveByTag(ctx, consts.CacheSysConfigTag)
	})
	return
}

// GetConfigByKey 通过key获取参数（从缓存获取）
func (s *sConfigData) GetConfigByKey(ctx context.Context, key string) (config *entity.SysConfig, err error) {
	if key == "" {
		err = gerror.New("参数key不能为空")
		return
	}
	cache := Cache()
	cf := cache.Get(ctx, consts.CacheSysConfigTag+key)
	if cf != nil && !cf.IsEmpty() {
		err = gconv.Struct(cf, &config)
		return
	}
	config, err = s.GetByKey(ctx, key)
	if err != nil {
		return
	}
	if config != nil {
		//配置数据缓存1分钟
		cache.Set(ctx, consts.CacheSysConfigTag+key, config, time.Minute*1, consts.CacheSysConfigTag)
	}
	return
}

// GetByKey 通过key获取参数（从数据库获取）
func (s *sConfigData) GetByKey(ctx context.Context, key string) (config *entity.SysConfig, err error) {
	err = dao.SysConfig.Ctx(ctx).Where("config_key", key).Scan(&config)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("获取配置失败")
	}
	return
}
