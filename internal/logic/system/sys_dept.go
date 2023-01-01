package system

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/logic/common"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"github.com/sagoo-cloud/sagooiot/utility/liberr"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysDept struct {
}

func sysDeptNew() *sSysDept {
	return &sSysDept{}
}

func init() {
	service.RegisterSysDept(sysDeptNew())
}

// GetTree 获取全部部门数据
func (s *sSysDept) GetTree(ctx context.Context, deptName string, status int) (out []*model.DeptOut, err error) {
	dept, err := s.GetData(ctx, deptName, status)
	var parentNodeOut []*model.DeptOut
	if dept != nil {
		//获取所有的根节点
		for _, v := range dept {
			var parentNode *model.DeptOut
			if v.ParentId == -1 {
				if err = gconv.Scan(v, &parentNode); err != nil {
					return
				}
				parentNodeOut = append(parentNodeOut, parentNode)
			}
		}
	}
	out = deptTree(parentNodeOut, dept)
	return
}

// Trees Tree 生成树结构
func deptTree(parentNodeOut []*model.DeptOut, data []*model.DeptOut) (dataTree []*model.DeptOut) {
	//循环所有一级菜单
	for k, v := range parentNodeOut {
		//查询所有该菜单下的所有子菜单
		for _, j := range data {
			var node *model.DeptOut
			if j.ParentId == v.DeptId {
				if err := gconv.Scan(j, &node); err != nil {
					return
				}
				parentNodeOut[k].Children = append(parentNodeOut[k].Children, node)
			}
		}
		deptTree(v.Children, data)
	}
	return parentNodeOut
}

// GetData 执行获取数据操作
func (s *sSysDept) GetData(ctx context.Context, deptName string, status int) (data []*model.DeptOut, err error) {
	m := dao.SysDept.Ctx(ctx)
	if status != -1 {
		m = m.Where(dao.SysDept.Columns().Status, status)
	}
	//模糊查询部门名称
	if deptName != "" {
		m = m.WhereLike(dao.SysDept.Columns().DeptName, "%"+deptName+"%")
	}
	err = m.Where(dao.SysDept.Columns().IsDeleted, 0).
		OrderDesc(dao.SysDept.Columns().OrderNum).
		Scan(&data)
	if err != nil {
		return
	}
	return
}

// Add 添加
func (s *sSysDept) Add(ctx context.Context, input *model.AddDeptInput) (err error) {
	var dept *entity.SysDept
	//根据名称查看部门是否存在
	dept = checkDeptName(ctx, input.DeptName, dept, 0)
	if dept != nil {
		return gerror.New("部门已存在,无法添加")
	}
	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)
	dept = new(entity.SysDept)
	if err := gconv.Scan(input, &dept); err != nil {
		return err
	}
	dept.IsDeleted = 0
	dept.CreatedBy = uint(loginUserId)
	//开启事务管理
	err = dao.SysDept.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) (err error) {
		lastId, err1 := dao.SysDept.Ctx(ctx).Data(dept).InsertAndGetId()
		if err1 != nil {
			return err1
		}
		err = setAncestors(ctx, input.ParentId, lastId)
		if err != nil {
			return err
		}
		return err
	})
	return
}

// Edit 修改部门
func (s *sSysDept) Edit(ctx context.Context, input *model.EditDeptInput) (err error) {
	var dept1, dept2 *entity.SysDept
	//根据ID查看部门是否存在
	dept1 = checkDeptId(ctx, input.DeptId, dept1)
	dept := dept1.ParentId
	deptAnces := dept1.Ancestors
	if dept1 == nil {
		return gerror.New("部门不存在")
	}
	dept2 = checkDeptName(ctx, input.DeptName, dept2, input.DeptId)
	if dept2 != nil {
		return gerror.New("相同部门已存在,无法修改")
	}
	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)
	if err := gconv.Scan(input, &dept1); err != nil {
		return err
	}
	dept1.UpdatedBy = loginUserId
	//开启事务管理
	err = dao.SysDept.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) (err error) {
		_, err = dao.SysDept.Ctx(ctx).Data(dept1).
			Where(dao.SysDept.Columns().DeptId, input.DeptId).Update()
		if err != nil {
			return gerror.New("修改失败")
		}
		//修改祖籍字段
		if dept != input.ParentId {
			err := setAncestors(ctx, input.ParentId, input.DeptId)
			if err != nil {
				return gerror.New("祖籍修改失败")
			}
			lId := strconv.FormatInt(input.DeptId, 10)
			value, _ := dao.SysDept.Ctx(ctx).
				Fields(dao.SysDept.Columns().Ancestors).
				WhereLike(dao.SysDept.Columns().Ancestors, "%"+lId+"%").Array()
			if input.ParentId == -1 {
				for _, v := range value {
					newAncestors := strings.Replace(v.String(), deptAnces, lId, -1)
					//修改相关祖籍字段
					_, err := dao.SysDept.Ctx(ctx).
						Data(dao.SysDept.Columns().Ancestors, newAncestors).
						Where(dao.SysDept.Columns().Ancestors, v.String()).Update()
					if err != nil {
						return gerror.New("关联祖籍修改失败")
					}
				}
			} else {
				//查询现有的进行拼接
				ancestors, _ := dao.SysDept.Ctx(ctx).
					Fields(dao.SysDept.Columns().Ancestors).
					Where(dao.SysDept.Columns().DeptId, input.DeptId).Value()
				for _, v := range value {
					newAncestors := strings.Replace(ancestors.String(), lId, "", -1)
					newAncestor := newAncestors + v.String()
					//修改相关祖籍字段
					_, err := dao.SysDept.Ctx(ctx).
						Data(dao.SysDept.Columns().Ancestors, newAncestor).
						Where(dao.SysDept.Columns().Ancestors, v.String()).
						WhereNot(dao.SysDept.Columns().DeptId, input.DeptId).
						Update()
					if err != nil {
						return gerror.New("关联祖籍修改失败")
					}
				}
			}
		}
		return nil
	})
	return
}

// Detail 部门详情
func (s *sSysDept) Detail(ctx context.Context, deptId int64) (entity *entity.SysDept, err error) {
	_ = dao.SysDept.Ctx(ctx).Where(g.Map{
		dao.SysDept.Columns().DeptId: deptId,
	}).Scan(&entity)
	return
}

// Del 根据ID删除部门信息
func (s *sSysDept) Del(ctx context.Context, deptId int64) (err error) {
	var dept *entity.SysDept
	_ = dao.SysDept.Ctx(ctx).Where(g.Map{
		dao.SysDept.Columns().DeptId: deptId,
	}).Scan(&dept)
	if dept == nil {
		return gerror.New("ID错误")
	}
	//查询是否有子节点
	num, err := dao.SysDept.Ctx(ctx).Where(g.Map{
		dao.SysDept.Columns().ParentId:  deptId,
		dao.SysDept.Columns().IsDeleted: 0,
	}).Count()
	if err != nil {
		return err
	}
	if num > 0 {
		return gerror.New("请先删除子节点!")
	}
	loginUserId := service.Context().GetUserId(ctx)
	//更新部门信息
	_, err = dao.SysDept.Ctx(ctx).
		Data(g.Map{
			dao.SysDept.Columns().DeletedBy: uint(loginUserId),
			dao.SysDept.Columns().IsDeleted: 1,
		}).
		Where(dao.SysDept.Columns().DeptId, deptId).
		Update()
	//删除部门信息
	_, err = dao.SysDept.Ctx(ctx).
		Where(dao.SysDept.Columns().DeptId, deptId).
		Delete()
	return
}

// 修改祖籍字段
func setAncestors(ctx context.Context, ParentId int64, lastId int64) (err error) {
	lId := strconv.FormatInt(lastId, 10)
	if ParentId == -1 { //根级别,修改祖籍为自己
		_, err := dao.SysDept.Ctx(ctx).
			Data(dao.SysDept.Columns().Ancestors, lId).
			Where(dao.SysDept.Columns().DeptId, lastId).
			Update()
		if err != nil {
			return gerror.New("祖籍修改失败")
		}
	} else {
		var oldDept *entity.SysDept
		_ = dao.SysDept.Ctx(ctx).
			Where(dao.SysDept.Columns().DeptId, ParentId).
			Scan(&oldDept)
		_, err := dao.SysDept.Ctx(ctx).
			Data(dao.SysDept.Columns().Ancestors, oldDept.Ancestors+","+lId).
			Where(dao.SysDept.Columns().DeptId, lastId).
			Update()
		if err != nil {
			return gerror.New("祖籍修改失败")
		}
	}
	return
}

// 检查相同部门名称的数据是否存在
func checkDeptName(ctx context.Context, deptName string, dept *entity.SysDept, tag int64) *entity.SysDept {
	m := dao.SysDept.Ctx(ctx)
	if tag > 0 {
		m = m.WhereNot(dao.SysDept.Columns().DeptId, tag)
	}
	_ = m.Where(g.Map{
		dao.SysDept.Columns().DeptName:  deptName,
		dao.SysDept.Columns().IsDeleted: 0,
	}).Scan(&dept)
	return dept
}

// 检查指定ID的数据是否存在
func checkDeptId(ctx context.Context, DeptId int64, dept *entity.SysDept) *entity.SysDept {
	_ = dao.SysDept.Ctx(ctx).Where(g.Map{
		dao.SysDept.Columns().DeptId:    DeptId,
		dao.SysDept.Columns().IsDeleted: 0,
	}).Scan(&dept)
	return dept
}

// GetAll 获取全部部门数据
func (s *sSysDept) GetAll(ctx context.Context) (data []*entity.SysDept, err error) {
	_ = dao.SysDept.Ctx(ctx).Where(g.Map{
		dao.SysDept.Columns().Status:    1,
		dao.SysDept.Columns().IsDeleted: 0,
	}).Scan(&data)
	return
}

func (s *sSysDept) GetFromCache(ctx context.Context) (list []*entity.SysDept, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		cache := common.Cache()
		//从缓存获取
		iList := cache.GetOrSetFuncLock(ctx, consts.CacheSysDept, func(ctx context.Context) (value interface{}, err error) {
			err = dao.SysDept.Ctx(ctx).Scan(&list)
			liberr.ErrIsNil(ctx, err, "获取部门列表失败")
			value = list
			return
		}, 0, consts.CacheSysAuthTag)
		if iList != nil {
			err = gconv.Struct(iList, &list)
			liberr.ErrIsNil(ctx, err)
		}
	})
	return
}
func (s *sSysDept) FindSonByParentId(deptList []*entity.SysDept, deptId int64) []*entity.SysDept {
	children := make([]*entity.SysDept, 0, len(deptList))
	for _, v := range deptList {
		if v.ParentId == deptId {
			children = append(children, v)
			fChildren := s.FindSonByParentId(deptList, v.DeptId)
			children = append(children, fChildren...)
		}
	}
	return children
}
