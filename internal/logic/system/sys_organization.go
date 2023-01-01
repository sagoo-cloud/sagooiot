package system

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"strconv"
	"strings"
	"time"
)

type sSysOrganization struct {
}

func SysOrganizationNew() *sSysOrganization {
	return &sSysOrganization{}
}

func init() {
	service.RegisterSysOrganization(SysOrganizationNew())
}

// GetList 获取组织数据
func (s *sSysOrganization) GetTree(ctx context.Context, name string, status int) (data []*model.OrganizationOut, err error) {
	orgainzationInfo, err := s.GetData(ctx, name, status)
	var parentNodeOut []*model.OrganizationOut
	if orgainzationInfo != nil {
		//获取所有的根节点
		for _, v := range orgainzationInfo {
			var parentNode *model.OrganizationOut
			if v.ParentId == -1 {
				if err = gconv.Scan(v, &parentNode); err != nil {
					return
				}
				parentNodeOut = append(parentNodeOut, parentNode)
			}
		}
	}
	data = OrganizationTree(parentNodeOut, orgainzationInfo)
	return
}

// OrganizationTree 生成树结构
func OrganizationTree(parentNodeOut []*model.OrganizationOut, data []*model.OrganizationOut) (dataTree []*model.OrganizationOut) {
	//循环所有一级菜单
	for k, v := range parentNodeOut {
		//查询所有该菜单下的所有子菜单
		for _, j := range data {
			var node *model.OrganizationOut
			if j.ParentId == v.Id {
				if err := gconv.Scan(j, &node); err != nil {
					return
				}
				parentNodeOut[k].Children = append(parentNodeOut[k].Children, node)
			}
		}
		OrganizationTree(v.Children, data)
	}
	return parentNodeOut
}

// GetData 执行获取数据操作
func (s *sSysOrganization) GetData(ctx context.Context, name string, status int) (data []*model.OrganizationOut, err error) {
	m := dao.SysOrganization.Ctx(ctx)
	if status != -1 {
		m = m.Where(dao.SysOrganization.Columns().Status, status)
	}
	//模糊查询组织名称
	if name != "" {
		m = m.WhereLike(dao.SysOrganization.Columns().Name, "%"+name+"%")
	}
	err = m.Where(dao.SysOrganization.Columns().IsDeleted, 0).
		OrderDesc(dao.SysOrganization.Columns().OrderNum).
		Scan(&data)
	if err != nil {
		return
	}
	return
}

// Add 添加
func (s *sSysOrganization) Add(ctx context.Context, input *model.AddOrganizationInput) (err error) {
	//根据名称查看组织是否存在
	organization := checkOrganizationName(ctx, input.Name, 0)
	if organization != nil {
		return gerror.New("组织已存在,无法添加")
	}
	organization = new(entity.SysOrganization)
	organization.Number = "org_" + strconv.FormatInt(time.Now().Unix(), 10)
	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)
	organization = new(entity.SysOrganization)
	if err := gconv.Scan(input, &organization); err != nil {
		return err
	}
	organization.IsDeleted = 0
	organization.CreatedBy = uint(loginUserId)
	//开启事务管理
	err = dao.SysOrganization.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) (err error) {
		lastId, err1 := dao.SysOrganization.Ctx(ctx).Data(organization).InsertAndGetId()
		if err1 != nil {
			return err1
		}
		err = setOrganizationAncestors(ctx, input.ParentId, lastId)
		if err != nil {
			return err
		}
		return err
	})
	return
}

// Edit 修改组织
func (s *sSysOrganization) Edit(ctx context.Context, input *model.EditOrganizationInput) (err error) {
	var organization1, organization2 *entity.SysOrganization
	//根据ID查看组织是否存在
	organization1 = checkOrganizationId(ctx, input.Id, organization1)
	organization := organization1.ParentId
	organizationAnces := organization1.Ancestors
	if organization1 == nil {
		return gerror.New("组织不存在")
	}
	organization2 = checkOrganizationName(ctx, input.Name, input.Id)
	if organization2 != nil {
		return gerror.New("相同组织已存在,无法修改")
	}
	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)
	if err := gconv.Scan(input, &organization1); err != nil {
		return err
	}
	organization1.UpdatedBy = loginUserId
	//开启事务管理
	err = dao.SysOrganization.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) (err error) {
		_, err = dao.SysOrganization.Ctx(ctx).Data(organization1).
			Where(dao.SysOrganization.Columns().Id, input.Id).Update()
		if err != nil {
			return gerror.New("修改失败")
		}
		//修改祖籍字段
		if organization != input.ParentId {
			err := setAncestors(ctx, input.ParentId, input.Id)
			if err != nil {
				return gerror.New("祖籍修改失败")
			}
			lId := strconv.FormatInt(input.Id, 10)
			value, _ := dao.SysOrganization.Ctx(ctx).
				Fields(dao.SysOrganization.Columns().Ancestors).
				WhereLike(dao.SysOrganization.Columns().Ancestors, "%"+lId+"%").Array()
			if input.ParentId == -1 {
				for _, v := range value {
					newAncestors := strings.Replace(v.String(), organizationAnces, lId, -1)
					//修改相关祖籍字段
					_, err := dao.SysOrganization.Ctx(ctx).
						Data(dao.SysOrganization.Columns().Ancestors, newAncestors).
						Where(dao.SysOrganization.Columns().Ancestors, v.String()).Update()
					if err != nil {
						return gerror.New("关联祖籍修改失败")
					}
				}
			} else {
				//查询现有的进行拼接
				ancestors, _ := dao.SysOrganization.Ctx(ctx).
					Fields(dao.SysOrganization.Columns().Ancestors).
					Where(dao.SysOrganization.Columns().Id, input.Id).Value()
				for _, v := range value {
					newAncestors := strings.Replace(ancestors.String(), lId, "", -1)
					newAncestor := newAncestors + v.String()
					//修改相关祖籍字段
					_, err := dao.SysOrganization.Ctx(ctx).
						Data(dao.SysOrganization.Columns().Ancestors, newAncestor).
						Where(dao.SysOrganization.Columns().Ancestors, v.String()).
						WhereNot(dao.SysOrganization.Columns().Id, input.Id).
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

// Detail 组织详情
func (s *sSysOrganization) Detail(ctx context.Context, id int64) (entity *entity.SysOrganization, err error) {
	_ = dao.SysOrganization.Ctx(ctx).Where(g.Map{
		dao.SysOrganization.Columns().Id: id,
	}).Scan(&entity)
	if entity == nil {
		return nil, gerror.New("ID错误")
	}
	return
}

// Del 根据ID删除组织信息
func (s *sSysOrganization) Del(ctx context.Context, id int64) (err error) {
	var organization *entity.SysOrganization
	_ = dao.SysOrganization.Ctx(ctx).Where(g.Map{
		dao.SysOrganization.Columns().Id: id,
	}).Scan(&organization)
	if organization == nil {
		return gerror.New("ID错误")
	}
	//查询是否有子节点
	num, err := dao.SysOrganization.Ctx(ctx).Where(g.Map{
		dao.SysOrganization.Columns().ParentId:  id,
		dao.SysOrganization.Columns().IsDeleted: 0,
	}).Count()
	if err != nil {
		return err
	}
	if num > 0 {
		return gerror.New("请先删除子节点!")
	}
	loginUserId := service.Context().GetUserId(ctx)
	//更新组织信息
	_, err = dao.SysOrganization.Ctx(ctx).
		Data(g.Map{
			dao.SysOrganization.Columns().DeletedBy: uint(loginUserId),
			dao.SysOrganization.Columns().IsDeleted: 1,
		}).
		Where(dao.SysOrganization.Columns().Id, id).
		Update()
	//删除组织信息
	_, err = dao.SysOrganization.Ctx(ctx).Where(dao.SysOrganization.Columns().Id, id).Delete()
	return
}

// GetAll 获取全部组织数据
func (s *sSysOrganization) GetAll(ctx context.Context) (data []*entity.SysOrganization, err error) {
	err = dao.SysOrganization.Ctx(ctx).Where(g.Map{
		dao.SysOrganization.Columns().Status:    1,
		dao.SysOrganization.Columns().IsDeleted: 0,
	}).Scan(&data)
	return
}

// 修改祖籍字段
func setOrganizationAncestors(ctx context.Context, ParentId int64, lastId int64) (err error) {
	lId := strconv.FormatInt(lastId, 10)
	if ParentId == -1 { //根级别,修改祖籍为自己
		_, err := dao.SysOrganization.Ctx(ctx).
			Data(dao.SysOrganization.Columns().Ancestors, lId).
			Where(dao.SysOrganization.Columns().Id, lastId).
			Update()
		if err != nil {
			return gerror.New("祖籍修改失败")
		}
	} else {
		var oldorganization *entity.SysOrganization
		_ = dao.SysOrganization.Ctx(ctx).
			Where(dao.SysOrganization.Columns().Id, ParentId).
			Scan(&oldorganization)
		_, err := dao.SysOrganization.Ctx(ctx).
			Data(dao.SysOrganization.Columns().Ancestors, oldorganization.Ancestors+","+lId).
			Where(dao.SysOrganization.Columns().Id, lastId).
			Update()
		if err != nil {
			return gerror.New("祖籍修改失败")
		}
	}
	return
}

// Count 获取组织数量
func (s *sSysOrganization) Count(ctx context.Context) (count int, err error) {
	count, _ = dao.SysOrganization.Ctx(ctx).Where(g.Map{
		dao.SysOrganization.Columns().IsDeleted: 0,
	}).Count()
	return
}

// 检查相同组织名称的数据是否存在
func checkOrganizationName(ctx context.Context, name string, tag int64) (organization *entity.SysOrganization) {
	m := dao.SysOrganization.Ctx(ctx)
	if tag > 0 {
		m = m.WhereNot(dao.SysOrganization.Columns().Id, tag)
	}
	_ = m.Where(g.Map{
		dao.SysOrganization.Columns().Name:      name,
		dao.SysOrganization.Columns().IsDeleted: 0,
	}).Scan(&organization)
	return organization
}

// 检查指定ID的数据是否存在
func checkOrganizationId(ctx context.Context, id int64, organization *entity.SysOrganization) *entity.SysOrganization {
	_ = dao.SysOrganization.Ctx(ctx).Where(g.Map{
		dao.SysOrganization.Columns().Id:        id,
		dao.SysOrganization.Columns().IsDeleted: 0,
	}).Scan(&organization)
	return organization
}
