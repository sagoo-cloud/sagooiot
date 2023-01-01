package common

import (
	"context"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/sagoo-cloud/sagooiot/api/v1/common"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/do"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"github.com/sagoo-cloud/sagooiot/utility/liberr"
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
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysDictType.Ctx(ctx)
		if input.DictName != "" {
			m = m.Where(dao.SysDictType.Columns().DictName+" like ?", "%"+input.DictName+"%")
		}
		if input.DictType != "" {
			m = m.Where(dao.SysDictType.Columns().DictType+" like ?", "%"+input.DictType+"%")
		}
		if input.Status != "" {
			m = m.Where(dao.SysDictType.Columns().Status+" = ", gconv.Int(input.Status))
		}
		total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取字典类型失败")
		if input.PageNum == 0 {
			input.PageNum = 1
		}
		if input.PageSize == 0 {
			input.PageSize = consts.PageSize
		}
		err = m.Fields(model.SysDictTypeInfoRes{}).Page(input.PageNum, input.PageSize).
			Order(dao.SysDictType.Columns().DictId + " asc").Scan(&out)
		liberr.ErrIsNil(ctx, err, "获取字典类型失败")
	})
	return
}

// Add 添加字典类型
func (s *sDictType) Add(ctx context.Context, input *model.AddDictTypeInput, userId int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = s.ExistsDictType(ctx, input.DictType)
		liberr.ErrIsNil(ctx, err)
		_, err = dao.SysDictType.Ctx(ctx).Insert(do.SysDictType{
			DictName: input.DictName,
			DictType: input.DictType,
			Status:   input.Status,
			CreateBy: userId,
			Remark:   input.Remark,
		})
		liberr.ErrIsNil(ctx, err, "添加字典类型失败")
		//清除缓存
		Cache().RemoveByTag(ctx, consts.CacheSysDictTag)
	})
	return
}

// Edit 修改字典类型
func (s *sDictType) Edit(ctx context.Context, input *model.EditDictTypeInput, userId int) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			err = s.ExistsDictType(ctx, input.DictType, input.DictId)
			liberr.ErrIsNil(ctx, err)
			dictType := (*entity.SysDictType)(nil)
			e := dao.SysDictType.Ctx(ctx).Fields(dao.SysDictType.Columns().DictType).WherePri(input.DictId).Scan(&dictType)
			liberr.ErrIsNil(ctx, e, "获取字典类型失败")
			liberr.ValueIsNil(dictType, "字典类型不存在")
			//修改字典类型
			_, e = dao.SysDictType.Ctx(ctx).TX(tx).WherePri(input.DictId).Update(do.SysDictType{
				DictName: input.DictName,
				DictType: input.DictType,
				Status:   input.Status,
				UpdateBy: userId,
				Remark:   input.Remark,
			})
			liberr.ErrIsNil(ctx, e, "修改字典类型失败")
			//修改字典数据
			_, e = dao.SysDictData.Ctx(ctx).TX(tx).Data(do.SysDictData{DictType: input.DictType}).
				Where(dao.SysDictData.Columns().DictType, dictType.DictType).Update()
			liberr.ErrIsNil(ctx, e, "修改字典数据失败")
			//清除缓存
			Cache().RemoveByTag(ctx, consts.CacheSysDictTag)
		})
		return err
	})
	return
}

func (s *sDictType) Get(ctx context.Context, req *common.DictTypeGetReq) (dictType *model.SysDictTypeOut, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysDictType.Ctx(ctx).Where(dao.SysDictType.Columns().DictId, req.DictId).Scan(&dictType)
		liberr.ErrIsNil(ctx, err, "获取字典类型失败")
	})
	return
}

// ExistsDictType 检查类型是否已经存在
func (s *sDictType) ExistsDictType(ctx context.Context, dictType string, dictId ...int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysDictType.Ctx(ctx).Fields(dao.SysDictType.Columns().DictId).
			Where(dao.SysDictType.Columns().DictType, dictType)
		if len(dictId) > 0 {
			m = m.Where(dao.SysDictType.Columns().DictId+" !=? ", dictId[0])
		}
		res, e := m.One()
		liberr.ErrIsNil(ctx, e, "sql err")
		if !res.IsEmpty() {
			liberr.ErrIsNil(ctx, gerror.New("字典类型已存在"))
		}
	})
	return
}

// Delete 删除字典类型
func (s *sDictType) Delete(ctx context.Context, dictIds []int) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			discs := ([]*entity.SysDictType)(nil)
			err = dao.SysDictType.Ctx(ctx).Fields(dao.SysDictType.Columns().DictType).
				Where(dao.SysDictType.Columns().DictId+" in (?) ", dictIds).Scan(&discs)
			liberr.ErrIsNil(ctx, err, "删除失败")
			types := garray.NewStrArray()
			for _, dt := range discs {
				types.Append(dt.DictType)
			}
			if types.Len() > 0 {
				_, err = dao.SysDictType.Ctx(ctx).TX(tx).Delete(dao.SysDictType.Columns().DictId+" in (?) ", dictIds)
				liberr.ErrIsNil(ctx, err, "删除类型失败")
				_, err = dao.SysDictData.Ctx(ctx).TX(tx).Delete(dao.SysDictData.Columns().DictType+" in (?) ", types.Slice())
				liberr.ErrIsNil(ctx, err, "删除字典数据失败")
			}
			//清除缓存
			Cache().RemoveByTag(ctx, consts.CacheSysDictTag)
		})
		return err
	})
	return
}
