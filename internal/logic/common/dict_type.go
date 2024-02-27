package common

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/api/v1/common"
	"sagooiot/internal/consts"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/do"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"
	"sagooiot/pkg/cache"
)

type sDictType struct {
}

func DictType() *sDictType {
	return &sDictType{}
}

func init() {
	service.RegisterDictType(DictType())
}

// List 字典类型列表
func (s *sDictType) List(ctx context.Context, input *model.DictTypeDoInput) (total int, out []*model.SysDictTypeInfoOut, err error) {
	m := dao.SysDictType.Ctx(ctx)
	if input.ModuleClassify != "" {
		m = m.Where(dao.SysDictType.Columns().ModuleClassify, input.ModuleClassify)
	}
	if input.Status != "" {
		m = m.Where(dao.SysDictType.Columns().Status, gconv.Int(input.Status))
	}
	if input.DictName != "" {
		m = m.WhereLike(dao.SysDictType.Columns().DictName, "%"+input.DictName+"%")
		m = m.WhereOrLike(dao.SysDictType.Columns().DictType, "%"+input.DictName+"%")
	}
	total, err = m.Count()
	if err != nil {
		return 0, nil, errors.New("获取字典类型失败")
	}
	if input.PageNum == 0 {
		input.PageNum = 1
	}
	if input.PageSize == 0 {
		input.PageSize = consts.PageSize
	}
	err = m.Fields(model.SysDictTypeInfoRes{}).Page(input.PageNum, input.PageSize).
		OrderDesc(dao.SysDictType.Columns().CreatedAt).Scan(&out)
	if err != nil {
		return 0, nil, errors.New("获取字典类型失败")
	}
	return
}

// Add 添加字典类型
func (s *sDictType) Add(ctx context.Context, input *model.AddDictTypeInput, userId int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = s.ExistsDictType(ctx, input.DictType)
		if err != nil {
			return
		}
		_, err = dao.SysDictType.Ctx(ctx).Insert(do.SysDictType{
			DictName:       input.DictName,
			DictType:       input.DictType,
			Status:         input.Status,
			CreatedBy:      userId,
			Remark:         input.Remark,
			ModuleClassify: input.ModuleClassify,
		})
		if err != nil {
			return
		}
		//清除缓存
		_, err = cache.Instance().Remove(ctx, consts.CacheSysDictTag)
	})
	return
}

// Edit 修改字典类型
func (s *sDictType) Edit(ctx context.Context, input *model.EditDictTypeInput, userId int) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = s.ExistsDictType(ctx, input.DictType, input.DictId)
		if err != nil {
			return nil
		}
		dictType := (*entity.SysDictType)(nil)
		err = dao.SysDictType.Ctx(ctx).Fields(dao.SysDictType.Columns().DictType).WherePri(input.DictId).Scan(&dictType)
		if err != nil {
			return errors.New("获取字典类型失败")
		}
		if dictType == nil {
			return errors.New("字典类型不存在")
		}

		//修改字典类型
		_, err = dao.SysDictType.Ctx(ctx).TX(tx).WherePri(input.DictId).Update(do.SysDictType{
			DictName:       input.DictName,
			DictType:       input.DictType,
			Status:         input.Status,
			UpdatedBy:      userId,
			Remark:         input.Remark,
			ModuleClassify: input.ModuleClassify,
		})
		if err != nil {
			return errors.New("修改字典类型失败")
		}
		//修改字典数据
		_, err = dao.SysDictData.Ctx(ctx).TX(tx).Data(do.SysDictData{DictType: input.DictType}).
			Where(dao.SysDictData.Columns().DictType, dictType.DictType).Update()
		if err != nil {
			return errors.New("修改字典数据失败")
		}
		//清除缓存
		_, err = cache.Instance().Remove(ctx, consts.CacheSysDictTag)

		return err
	})
	return
}

func (s *sDictType) Get(ctx context.Context, req *common.DictTypeGetReq) (dictType *model.SysDictTypeOut, err error) {
	err = dao.SysDictType.Ctx(ctx).Where(dao.SysDictType.Columns().DictId, req.DictId).Scan(&dictType)
	if err != nil {
		return nil, errors.New("修改字典数据失败")
	}

	return
}

// ExistsDictType 检查类型是否已经存在
func (s *sDictType) ExistsDictType(ctx context.Context, dictType string, dictId ...int) (err error) {
	m := dao.SysDictType.Ctx(ctx).Fields(dao.SysDictType.Columns().DictId).
		Where(dao.SysDictType.Columns().DictType, dictType)
	if len(dictId) > 0 {
		m = m.Where(dao.SysDictType.Columns().DictId+" !=? ", dictId[0])
	}
	res, err := m.One()
	if err != nil {
		return
	}
	if !res.IsEmpty() {
		return errors.New("字典类型已存在")
	}

	return
}

// Delete 删除字典类型
func (s *sDictType) Delete(ctx context.Context, dictIds []int) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		discs := ([]*entity.SysDictType)(nil)
		err = dao.SysDictType.Ctx(ctx).Fields(dao.SysDictType.Columns().DictType).
			Where(dao.SysDictType.Columns().DictId+" in (?) ", dictIds).Scan(&discs)
		if err != nil {
			return errors.New("删除失败")
		}
		types := garray.NewStrArray()
		for _, dt := range discs {
			types.Append(dt.DictType)
		}
		if types.Len() > 0 {
			_, err = dao.SysDictType.Ctx(ctx).TX(tx).Delete(dao.SysDictType.Columns().DictId+" in (?) ", dictIds)
			if err != nil {
				return errors.New("删除类型失败")
			}
			_, err = dao.SysDictData.Ctx(ctx).TX(tx).Delete(dao.SysDictData.Columns().DictType+" in (?) ", types.Slice())
			if err != nil {
				return errors.New("删除字典数据失败")
			}
		}
		//清除缓存
		_, err = cache.Instance().Remove(ctx, consts.CacheSysDictTag)
		return err
	})
	return
}
