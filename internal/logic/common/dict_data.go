package common

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/internal/consts"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/do"
	"sagooiot/internal/service"
	"sagooiot/pkg/cache"
)

type sDictData struct {
}

func DictData() *sDictData {
	return &sDictData{}
}

func init() {
	service.RegisterDictData(DictData())
}

// GetDictWithDataByType 通过字典键类型获取选项
func (s *sDictData) GetDictWithDataByType(ctx context.Context, input *model.GetDictInput) (dict *model.GetDictOut, err error) {
	cacheKey := consts.CacheSysDict + "_" + input.DictType
	//从缓存获取
	iDict, err := cache.Instance().GetOrSetFuncLock(ctx, cacheKey, func(ctx context.Context) (value interface{}, err error) {
		//从数据库获取
		dict = &model.GetDictOut{}
		//获取类型数据
		err = dao.SysDictType.Ctx(ctx).Where(dao.SysDictType.Columns().DictType, input.DictType).
			Where(dao.SysDictType.Columns().Status, 1).Fields(model.DictTypeOut{}).Scan(&dict.Data)
		if err != nil {
			return nil, errors.New("获取字典类型失败")
		}
		err = dao.SysDictData.Ctx(ctx).Fields(model.DictDataOut{}).
			Where(dao.SysDictData.Columns().DictType, input.DictType).
			Order(dao.SysDictData.Columns().DictSort + " asc," +
				dao.SysDictData.Columns().DictCode + " asc").
			Scan(&dict.Values)
		if err != nil {
			return nil, errors.New("获取字典数据失败")
		}

		value = dict
		return
	}, 0)
	if iDict != nil {
		err = gconv.Struct(iDict, &dict)
		if err != nil {
			return
		}
	}
	//设置给定的默认值
	for _, v := range dict.Values {
		if input.DefaultValue != "" {
			if gstr.Equal(input.DefaultValue, v.DictValue) {
				v.IsDefault = 1
			} else {
				v.IsDefault = 0
			}
		}
	}
	return
}

// List 获取字典数据
func (s *sDictData) List(ctx context.Context, input *model.SysDictSearchInput) (total int, out []*model.SysDictDataOut, err error) {
	m := dao.SysDictData.Ctx(ctx)
	if input != nil {
		if input.DictLabel != "" {
			m = m.Where(dao.SysDictData.Columns().DictLabel+" like ?", "%"+input.DictLabel+"%")
		}
		if input.Status != "-1" {
			m = m.Where(dao.SysDictData.Columns().Status+" = ", gconv.Int(input.Status))
		}
		if input.DictType != "" {
			m = m.Where(dao.SysDictData.Columns().DictType+" = ?", input.DictType)
		}
		total, err = m.Count()
		if err != nil {
			return 0, nil, errors.New("获取字典数据失败")
		}
		if input.PageNum == 0 {
			input.PageNum = 1
		}
	}
	if input.PageSize == 0 {
		input.PageSize = consts.PageSize
	}
	err = m.Page(input.PageNum, input.PageSize).Order(dao.SysDictData.Columns().DictSort + " asc," +
		dao.SysDictData.Columns().DictCode + " asc").Scan(&out)
	if err != nil {
		return 0, nil, errors.New("获取字典数据失败")
	}
	return
}

func (s *sDictData) Add(ctx context.Context, input *model.AddDictDataInput, userId int) (err error) {
	_, err = dao.SysDictData.Ctx(ctx).Insert(do.SysDictData{
		DictSort:  input.DictSort,
		DictLabel: input.DictLabel,
		DictValue: input.DictValue,
		DictType:  input.DictType,
		CssClass:  input.CssClass,
		ListClass: input.ListClass,
		IsDefault: input.IsDefault,
		Status:    input.Status,
		CreatedBy: userId,
		Remark:    input.Remark,
	})
	if err != nil {
		return errors.New("添加字典数据失败")
	}
	//清除缓存
	_, err = cache.Instance().Remove(ctx, consts.CacheSysDictTag)

	return
}

// Get 获取字典数据
func (s *sDictData) Get(ctx context.Context, dictCode uint) (out *model.SysDictDataOut, err error) {
	err = dao.SysDictData.Ctx(ctx).WherePri(dictCode).Scan(&out)
	if err != nil {
		return nil, errors.New("获取字典数据失败")
	}

	return
}

// Edit 修改字典数据
func (s *sDictData) Edit(ctx context.Context, input *model.EditDictDataInput, userId int) (err error) {
	_, err = dao.SysDictData.Ctx(ctx).WherePri(input.DictCode).Update(do.SysDictData{
		DictSort:  input.DictSort,
		DictLabel: input.DictLabel,
		DictValue: input.DictValue,
		DictType:  input.DictType,
		CssClass:  input.CssClass,
		ListClass: input.ListClass,
		IsDefault: input.IsDefault,
		Status:    input.Status,
		UpdatedBy: userId,
		Remark:    input.Remark,
	})
	if err != nil {
		return errors.New("修改字典数据失败")
	}
	//清除缓存
	_, err = cache.Instance().Remove(ctx, consts.CacheSysDictTag)

	return
}

// Delete 删除字典数据
func (s *sDictData) Delete(ctx context.Context, ids []int) (err error) {
	_, err = dao.SysDictData.Ctx(ctx).Where(dao.SysDictData.Columns().DictCode+" in(?)", ids).Delete()
	if err != nil {
		return errors.New("删除字典数据失败")
	}
	//清除缓存
	_, err = cache.Instance().Remove(ctx, consts.CacheSysDictTag)

	return
}

// GetDictDataByType 通过字典键类型获取选项
func (s *sDictData) GetDictDataByType(ctx context.Context, dictType string) (dict *model.GetDictOut, err error) {
	cacheKey := consts.CacheSysDict + "_" + dictType
	//从缓存获取
	iDict, err := cache.Instance().GetOrSetFuncLock(ctx, cacheKey, func(ctx context.Context) (value interface{}, err error) {
		err = g.Try(ctx, func(ctx context.Context) {
			//从数据库获取
			dict = &model.GetDictOut{}
			//获取类型数据
			err = dao.SysDictType.Ctx(ctx).Where(dao.SysDictType.Columns().DictType, dictType).
				Where(dao.SysDictType.Columns().Status, 1).Fields(model.DictTypeOut{}).Scan(&dict.Data)
			if err != nil {
				return
			}
			err = dao.SysDictData.Ctx(ctx).Fields(model.DictDataOut{}).
				Where(dao.SysDictData.Columns().DictType, dictType).
				Order(dao.SysDictData.Columns().DictSort + " asc," +
					dao.SysDictData.Columns().DictCode + " asc").
				Scan(&dict.Values)
			if err != nil {
				return
			}
		})
		value = dict
		return
	}, 0)
	if iDict != nil {
		err = gconv.Struct(iDict, &dict)
		if err != nil {
			return
		}
	}
	return
}
