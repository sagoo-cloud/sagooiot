package system

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
	"strings"
)

type sSysRole struct {
}

func init() {
	service.RegisterSysRole(sysRoleNew())
}

func sysRoleNew() *sSysRole {
	return &sSysRole{}
}

// GetAll 获取所有的角色
func (s *sSysRole) GetAll(ctx context.Context) (entity []*entity.SysRole, err error) {
	m := dao.SysRole.Ctx(ctx)

	err = m.Where(g.Map{
		dao.SysRole.Columns().Status:    1,
		dao.SysRole.Columns().IsDeleted: 0,
	}).OrderAsc(dao.SysRole.Columns().ListOrder).Scan(&entity)
	return
}

func (s *sSysRole) GetTree(ctx context.Context, name string, status int) (out []*model.RoleTreeOut, err error) {
	var e []*entity.SysRole
	m := dao.SysRole.Ctx(ctx)
	if name != "" {
		m = m.WhereLike(dao.SysRole.Columns().Name, "%"+name+"%")
	}
	if status != -1 {
		m = m.Where(dao.SysRole.Columns().Status, status)
	}
	m = m.Where(dao.SysRole.Columns().IsDeleted, 0)

	err = m.OrderAsc(dao.SysRole.Columns().ListOrder).Scan(&e)

	if len(e) > 0 {
		out, err = GetRoleTree(ctx, e)
		if err != nil {
			return
		}
	}
	return
}

// Add 添加
func (s *sSysRole) Add(ctx context.Context, input *model.AddRoleInput) (err error) {
	var role *entity.SysRole
	//根据名称查看角色是否存在
	err = dao.SysRole.Ctx(ctx).Where(g.Map{
		dao.SysRole.Columns().Name:      input.Name,
		dao.SysRole.Columns().IsDeleted: 0,
	}).Scan(&role)
	if role != nil {
		return gerror.New("角色已存在,无法添加")
	}
	//判断是否有权限删除当前角色
	if input.ParentId != -1 {
		var parentRole *entity.SysRole
		parentRole, err = s.GetInfoById(ctx, uint(input.ParentId))
		if parentRole == nil {
			err = gerror.Newf("无权限选择当前角色")
			return
		}
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)
	role = new(entity.SysRole)
	role.Name = input.Name
	role.DataScope = 1
	role.ParentId = input.ParentId
	role.ListOrder = input.ListOrder
	role.Remark = input.Remark
	role.Status = input.Status
	role.IsDeleted = 0
	role.CreatedBy = uint(loginUserId)
	_, err = dao.SysRole.Ctx(ctx).Data(do.SysRole{
		DeptId:    service.Context().GetUserDeptId(ctx),
		ParentId:  role.ParentId,
		ListOrder: role.ListOrder,
		Name:      role.Name,
		DataScope: role.DataScope,
		Remark:    role.Remark,
		Status:    role.Status,
		IsDeleted: role.IsDeleted,
		CreatedBy: role.CreatedBy,
		CreatedAt: gtime.Now(),
	}).Insert()
	if err != nil {
		return err
	}
	return
}

// Edit 编辑
func (s *sSysRole) Edit(ctx context.Context, input *model.EditRoleInput) (err error) {
	if input.Id == uint(input.ParentId) {
		return gerror.New("父级不能为自己")
	}
	var role *entity.SysRole
	//根据ID查询角色是否存在
	err = dao.SysRole.Ctx(ctx).Where(g.Map{
		dao.SysRole.Columns().Id:        input.Id,
		dao.SysRole.Columns().IsDeleted: 0,
	}).Scan(&role)
	if role == nil {
		return gerror.New("ID错误,无法修改")
	}
	//判断上级角色是否可以选择
	if input.ParentId != -1 {
		var parentRole *entity.SysRole
		parentRole, err = s.GetInfoById(ctx, uint(input.ParentId))
		if parentRole == nil {
			err = gerror.Newf("无权限选择当前角色")
			return
		}
	}

	//查看角色名称是否存在
	var roleByName *entity.SysRole
	err = dao.SysRole.Ctx(ctx).Where(g.Map{
		dao.SysRole.Columns().Name:      input.Name,
		dao.SysRole.Columns().IsDeleted: 0,
	}).Scan(&roleByName)
	if roleByName != nil && roleByName.Id != input.Id {
		return gerror.New("角色已存在,无法修改")
	}
	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)
	role.Name = input.Name
	role.ParentId = input.ParentId
	role.ListOrder = input.ListOrder
	role.Remark = input.Remark
	role.Status = input.Status
	role.UpdatedBy = uint(loginUserId)
	_, err = dao.SysRole.Ctx(ctx).Data(role).Where(dao.SysRole.Columns().Id, input.Id).Update()
	if err != nil {
		return err
	}
	return
}

// GetInfoById 根据ID获取角色信息
func (s *sSysRole) GetInfoById(ctx context.Context, id uint) (entity *entity.SysRole, err error) {
	m := dao.SysRole.Ctx(ctx)

	err = m.Where(g.Map{
		dao.SysRole.Columns().Id: id,
	}).Scan(&entity)
	return
}

// DelInfoById 根据ID删除角色信息
func (s *sSysRole) DelInfoById(ctx context.Context, id uint) (err error) {
	var roleData *entity.SysRole
	err = dao.SysRole.Ctx(ctx).Where(g.Map{
		dao.SysRole.Columns().Id: id,
	}).Scan(&roleData)
	if roleData == nil {
		return gerror.New("ID错误")
	}
	//查询是否有子节点
	num, err := dao.SysRole.Ctx(ctx).Where(g.Map{
		dao.SysRole.Columns().ParentId:  id,
		dao.SysRole.Columns().IsDeleted: 0,
	}).Count()
	if err != nil {
		return err
	}
	if num > 0 {
		return gerror.New("请先删除子节点!")
	}

	loginUserId := service.Context().GetUserId(ctx)
	//开启事务
	err = dao.SysRole.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		_, err = dao.SysRole.Ctx(ctx).Data(g.Map{
			dao.SysRole.Columns().DeletedBy: uint(loginUserId),
			dao.SysRole.Columns().DeletedAt: gtime.Now(),
			dao.SysRole.Columns().IsDeleted: 1,
		}).Where(dao.SysRole.Columns().Id, id).Update()
		//删除角色信息
		_, err = dao.SysRole.Ctx(ctx).Where(dao.SysRole.Columns().Id, id).Delete()
		if err != nil {
			return gerror.New("删除角色失败")
		}
		//删除与角色相绑定的用户关系
		_, err = dao.SysUserRole.Ctx(ctx).Where(dao.SysUserRole.Columns().RoleId, id).Delete()
		if err != nil {
			return gerror.New("解除与用户绑定关系失败")
		}
		//删除与角色相绑定的部门关系
		_, err = dao.SysRoleDept.Ctx(ctx).Where(dao.SysRoleDept.Columns().RoleId, id).Delete()
		if err != nil {
			return gerror.New("解除与部门绑定关系失败")
		}
		//删除角色已授权的菜单，按钮，列表
		_, err = dao.SysAuthorize.Ctx(ctx).Data(g.Map{
			dao.SysAuthorize.Columns().IsDeleted: 1,
			dao.SysAuthorize.Columns().DeletedBy: loginUserId,
			dao.SysAuthorize.Columns().DeletedAt: gtime.Now(),
		}).Where(dao.SysAuthorize.Columns().RoleId, id).Update()
		//删除权限配置
		_, err = dao.SysAuthorize.Ctx(ctx).Where(dao.SysAuthorize.Columns().RoleId, id).Delete()
		return
	})

	return
}

// GetRoleList 获取角色列表
func (s *sSysRole) GetRoleList(ctx context.Context) (list []*model.RoleInfoOut, err error) {
	//从缓存获取
	iList, err := cache.Instance().GetOrSetFuncLock(ctx, consts.CacheSysRole, s.getRoleListFromDb, 0)
	if iList != nil {
		err = gconv.Struct(iList.Val(), &list)
	}
	return
}

// 从数据库获取所有角色
func (s *sSysRole) getRoleListFromDb(ctx context.Context) (value interface{}, err error) {
	var v []*entity.SysRole
	m := dao.SysRole.Ctx(ctx)
	//从数据库获取
	err = m.
		Order(dao.SysRole.Columns().ListOrder + " asc," + dao.SysRole.Columns().Id + " asc").
		Scan(&v)
	if err != nil {
		return nil, errors.New("获取角色数据失败")
	}
	value = v
	return
}

// GetInfoByIds 根据ID数组获取角色信息
func (s *sSysRole) GetInfoByIds(ctx context.Context, id []int) (entity []*entity.SysRole, err error) {
	m := dao.SysRole.Ctx(ctx)
	err = m.WhereIn(dao.SysRole.Columns().Id, id).Where(g.Map{
		dao.SysRole.Columns().Status:    1,
		dao.SysRole.Columns().IsDeleted: 0,
	}).Scan(&entity)
	return
}

// DataScope 角色数据授权
func (s *sSysRole) DataScope(ctx context.Context, id int, dataScope uint, deptIds []int64) (err error) {
	//判断数据范围是否为自定义数据返回
	if dataScope == 2 {
		if deptIds == nil {
			return gerror.New("请选择部门信息")
		}
	}
	//查询角色ID是否存在
	var role *entity.SysRole
	err = dao.SysRole.Ctx(ctx).Where(dao.SysRole.Columns().Id, id).Scan(&role)
	if role == nil {
		return gerror.New("角色不存在,无法授权")
	}
	if role != nil && role.Status == 0 {
		return gerror.New("角色已禁用,无法授权")
	}

	//获取登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	//开启事务
	err = g.Try(ctx, func(ctx context.Context) {
		//修改角色数据范围
		role.DataScope = dataScope
		role.UpdatedBy = uint(loginUserId)
		_, editErr := dao.SysRole.Ctx(ctx).Data(role).Where(dao.SysRole.Columns().Id, id).Update()
		if editErr != nil {
			err = gerror.New("修改角色信息失败")
			return
		}
		//判断数据范围是否为自定义数据范围
		if dataScope == 2 {
			//自定义数据返回
			//根据部门ID数据获取部门信息
			var deptInfo []*entity.SysDept
			err = dao.SysDept.Ctx(ctx).Where(g.Map{
				dao.SysDept.Columns().IsDeleted: 0,
				dao.SysDept.Columns().Status:    1,
			}).WhereIn(dao.SysDept.Columns().DeptId, deptIds).Scan(&deptInfo)
			if deptInfo == nil {
				err = gerror.New("部门ID错误")
			}
			//删除原有绑定关系
			_, delErr := dao.SysRoleDept.Ctx(ctx).Where(dao.SysRoleDept.Columns().RoleId, id).Delete()
			if delErr != nil {
				err = gerror.New("解决绑定关系失败")
				return
			}
			//封装角色与部门绑定管理
			var roleDepts []*entity.SysRoleDept
			for _, dept := range deptInfo {
				var roleDept = new(entity.SysRoleDept)
				roleDept.RoleId = int64(id)
				roleDept.DeptId = dept.DeptId
				roleDepts = append(roleDepts, roleDept)
			}
			//添加绑定关系
			_, addErr := dao.SysRoleDept.Ctx(ctx).Data(roleDepts).Insert()
			if addErr != nil {
				err = gerror.New("绑定关系失败")
				return
			}
		}
		return
	})
	return
}

func (s *sSysRole) GetAuthorizeById(ctx context.Context, id int) (menuIds []string, menuButtonIds []string, menuColumnIds []string, menuApiIds []string, err error) {
	var e *entity.SysRole
	err = dao.SysRole.Ctx(ctx).Where(g.Map{
		dao.SysRole.Columns().Id: id,
	}).Scan(&e)
	if err != nil {
		return
	}
	if e == nil {
		err = gerror.New("ID错误")
		return
	}
	if e.IsDeleted == 1 {
		err = gerror.New("角色已删除,无法查询")
		return
	}
	if e.Status == 0 {
		err = gerror.New("角色已禁用,无法查询")
		return
	}

	//根据角色ID获取权限信息
	authorizeInfo, err := service.SysAuthorize().GetInfoByRoleId(ctx, id)
	if err != nil {
		return
	}
	for _, authorize := range authorizeInfo {
		if strings.EqualFold(authorize.ItemsType, consts.Menu) {
			menuIds = append(menuIds, gconv.String(authorize.ItemsId)+"_"+gconv.String(authorize.IsCheckAll))
		} else if strings.EqualFold(authorize.ItemsType, consts.Button) {
			menuButtonIds = append(menuButtonIds, gconv.String(authorize.ItemsId)+"_"+gconv.String(authorize.IsCheckAll))
		} else if strings.EqualFold(authorize.ItemsType, consts.Column) {
			menuColumnIds = append(menuColumnIds, gconv.String(authorize.ItemsId)+"_"+gconv.String(authorize.IsCheckAll))
		} else if strings.EqualFold(authorize.ItemsType, consts.Api) {
			menuApiIds = append(menuApiIds, gconv.String(authorize.ItemsId)+"_"+gconv.String(authorize.IsCheckAll))
		}
	}
	return
}
