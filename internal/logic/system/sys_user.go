package system

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/logic/common"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"github.com/sagoo-cloud/sagooiot/utility/liberr"
	"github.com/sagoo-cloud/sagooiot/utility/utils"
	"strings"
	"time"
)

type sSysUser struct {
}

func init() {
	service.RegisterSysUser(sysUserNew())
}

func sysUserNew() *sSysUser {
	return &sSysUser{}
}

// GetUserByUsername 通过用户名获取用户信息
func (s *sSysUser) GetUserByUsername(ctx context.Context, userName string) (data *entity.SysUser, err error) {
	var user *entity.SysUser
	err = dao.SysUser.Ctx(ctx).Fields(user).Where(dao.SysUser.Columns().UserName, userName).Scan(&user)
	if user == nil {
		return nil, gerror.New("账号不存在,请重新输入!")
	}
	if user.IsDeleted == 1 {
		return nil, gerror.New("该账号已删除")
	}
	if user.Status == 0 {
		return nil, gerror.New("该账号已禁用")
	}
	if user.Status == 2 {
		return nil, gerror.New("该账号未验证")
	}
	return user, nil
}

// GetAdminUserByUsernamePassword 根据用户名和密码获取用户信息
func (s *sSysUser) GetAdminUserByUsernamePassword(ctx context.Context, userName string, password string) (user *entity.SysUser, err error) {
	//判断账号是否存在
	user, err = s.GetUserByUsername(ctx, userName)
	if err != nil {
		return nil, err
	}
	//验证密码是否正确
	if utils.EncryptPassword(password, user.UserSalt) != user.UserPassword {
		err = gerror.New("密码错误")
		return
	}
	return user, nil
}

// UpdateLoginInfo 更新用户登录信息
func (s *sSysUser) UpdateLoginInfo(ctx context.Context, id uint64, ip string) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysUser.Ctx(ctx).WherePri(id).Update(g.Map{
			dao.SysUser.Columns().LastLoginIp:   ip,
			dao.SysUser.Columns().LastLoginTime: gtime.Now(),
		})
		liberr.ErrIsNil(ctx, err, "更新用户登录信息失败")
	})
	return
}

// UserList 用户列表
func (s *sSysUser) UserList(ctx context.Context, input *model.UserListDoInput) (total int, out []*model.UserListOut, err error) {
	m := dao.SysUser.Ctx(ctx)
	if input.KeyWords != "" {
		keyWords := "%" + input.KeyWords + "%"
		m = m.Where("user_name like ? or  user_nickname like ?", keyWords, keyWords)
	}
	if input.DeptId != 0 {
		//m = m.Where(dao.SysUser.Columns().DeptId, req.DeptId)
		deptIds, _ := s.getSearchDeptIds(ctx, gconv.Int64(input.DeptId))
		m = m.Where("dept_id in (?)", deptIds)
	}
	if input.UserName != "" {
		m = m.Where(dao.SysUser.Columns().UserName, input.UserName)
	}
	if input.Status != -1 {
		m = m.Where(dao.SysUser.Columns().Status, input.Status)
	}
	if input.Mobile != "" {
		m = m.WhereLike(dao.SysUser.Columns().Mobile, "%"+input.Mobile+"%")
	}
	if len(input.DateRange) > 0 {
		m = m.Where("created_at >=? AND created_at <?", input.DateRange[0], gtime.NewFromStrFormat(input.DateRange[1], "Y-m-d").AddDate(0, 0, 1))
	}
	m = m.Where(dao.SysUser.Columns().IsDeleted, 0)
	//获取总数
	total, err = m.Count()
	if err != nil {
		err = gerror.New("获取数据失败")
		return
	}
	if input.PageNum == 0 {
		input.PageNum = 1
	}
	if input.PageSize == 0 {
		input.PageSize = consts.DefaultPageSize
	}
	//获取用户信息
	err = m.Page(input.PageNum, input.PageSize).OrderDesc(dao.SysUser.Columns().CreatedAt).Scan(&out)
	if err != nil {
		err = gerror.New("获取用户列表失败")
	}

	deptIds := g.Slice{}
	for _, v := range out {
		deptIds = append(deptIds, v.DeptId)
	}
	deptData, _ := GetDeptNameDict(ctx, deptIds)
	for _, v := range out {
		err = gconv.Scan(deptData[v.DeptId], &v.Dept)
		v.RolesNames = getUserRoleName(ctx, gconv.Int(v.Id))
	}
	return
}

func getUserRoleName(ctx context.Context, userId int) (rolesNames string) {

	rolesData, _ := gcache.Get(ctx, "RoleListAtName"+gconv.String(userId))
	if rolesData != nil {
		rolesNames = rolesData.String()
		return
	}

	var roleIds []int
	//查询用户角色
	userRoleInfo, _ := service.SysUserRole().GetInfoByUserId(ctx, userId)
	if userRoleInfo != nil {
		for _, userRole := range userRoleInfo {
			roleIds = append(roleIds, userRole.RoleId)
		}
	}
	if roleIds != nil {
		//获取所有的角色信息
		roleInfo, _ := service.SysRole().GetInfoByIds(ctx, roleIds)
		var roleNames []string
		if roleInfo != nil {
			for _, role := range roleInfo {
				roleNames = append(roleNames, role.Name)
			}
			rolesNames = strings.Join(roleNames, ",")

			//放入缓存
			effectiveTime := time.Minute * 5
			_, _ = gcache.SetIfNotExist(ctx, "RoleListAtName"+gconv.String(userId), rolesNames, effectiveTime)

			return

		}
	}

	return ""
}

// Add 添加
func (s *sSysUser) Add(ctx context.Context, input *model.AddUserInput) (err error) {
	//判断账户是否存在
	num, _ := dao.SysUser.Ctx(ctx).Where(g.Map{
		dao.SysUser.Columns().UserName:  input.UserName,
		dao.SysUser.Columns().IsDeleted: 0,
	}).Count()
	if num > 0 {
		return gerror.New("账户已存在")
	}
	num, _ = dao.SysUser.Ctx(ctx).Where(g.Map{
		dao.SysUser.Columns().Mobile:    input.Mobile,
		dao.SysUser.Columns().IsDeleted: 0,
	}).Count()
	if num > 0 {
		return gerror.New("手机号已存在")
	}
	//开启事务管理
	err = dao.SysUser.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) (err error) {
		//添加用户
		var sysUser *entity.SysUser
		if err = gconv.Scan(input, &sysUser); err != nil {
			return err
		}
		//创建用户密码
		sysUser.UserSalt = grand.S(10)
		sysUser.UserPassword = utils.EncryptPassword(input.UserPassword, sysUser.UserSalt)
		//初始化为未删除
		sysUser.IsDeleted = 0
		//获取当前登录用户ID
		loginUserId := service.Context().GetUserId(ctx)
		sysUser.CreateBy = uint(loginUserId)
		//添加用户信息
		lastInsertId, err := dao.SysUser.Ctx(ctx).Data(sysUser).InsertAndGetId()
		if err != nil {
			return gerror.New("添加用户失败")
		}
		err = BindUserAndPost(ctx, int(lastInsertId), input.PostIds)

		return err
	})
	return
}

func (s *sSysUser) Edit(ctx context.Context, input *model.EditUserInput) (err error) {
	//根据ID获取用户信息
	var sysUser *entity.SysUser
	err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().Id, input.Id).Scan(&sysUser)
	if sysUser == nil {
		return gerror.New("Id错误")
	}
	if sysUser.IsDeleted == 1 {
		return gerror.New("该用户已删除")
	}

	//判断账户是否存在
	var sysUserByUserName *entity.SysUser
	err = dao.SysUser.Ctx(ctx).Where(g.Map{
		dao.SysUser.Columns().UserName:  input.UserName,
		dao.SysUser.Columns().IsDeleted: 0,
	}).Scan(&sysUserByUserName)
	if sysUserByUserName != nil && sysUserByUserName.Id != input.Id {
		return gerror.New("账户已存在")
	}
	//判断手机号是否存在
	var sysUserByMobile *entity.SysUser
	err = dao.SysUser.Ctx(ctx).Where(g.Map{
		dao.SysUser.Columns().Mobile:    input.Mobile,
		dao.SysUser.Columns().IsDeleted: 0,
	}).Scan(&sysUserByMobile)
	if sysUserByMobile != nil && sysUserByMobile.Id != input.Id {
		return gerror.New("手机号已存在")
	}
	//开启事务管理
	err = dao.SysUser.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) (err error) {
		//编辑用户
		sysUser.UserNickname = input.UserNickname
		sysUser.DeptId = input.DeptId
		sysUser.Mobile = input.Mobile
		sysUser.Status = input.Status
		sysUser.Address = input.Address
		sysUser.Avatar = input.Avatar
		sysUser.Birthday = gtime.NewFromStrFormat(input.Birthday, "Y-m-d")
		sysUser.Describe = input.Describe
		sysUser.UserTypes = input.UserTypes
		sysUser.UserEmail = input.UserEmail
		sysUser.Sex = input.Sex
		sysUser.IsAdmin = input.IsAdmin
		//获取当前登录用户ID
		loginUserId := service.Context().GetUserId(ctx)
		sysUser.UpdatedBy = uint(loginUserId)
		//编辑用户信息
		_, err = dao.SysUser.Ctx(ctx).Data(sysUser).Where(dao.SysUser.Columns().Id, input.Id).Update()
		if err != nil {
			return gerror.New("编辑用户失败")
		}
		//删除原有用户与岗位绑定管理
		_, err = dao.SysUserPost.Ctx(ctx).Where(dao.SysUserPost.Columns().UserId, input.Id).Delete()
		if err != nil {
			return gerror.New("删除用户与岗位绑定关系失败")
		}

		err = BindUserAndPost(ctx, int(input.Id), input.PostIds)

		//删除原有用户与角色绑定管理
		_, err = dao.SysUserRole.Ctx(ctx).Where(dao.SysUserRole.Columns().UserId, input.Id).Delete()
		if err != nil {
			return gerror.New("删除用户与角色绑定关系失败")
		}

		err = BindUserAndRole(ctx, int(input.Id), input.RoleIds)
		return err
	})
	return
}

// BindUserAndPost 添加用户与岗位绑定关系
func BindUserAndPost(ctx context.Context, userId int, postIds []int) (err error) {
	if len(postIds) > 0 {
		var sysUserPosts []*entity.SysUserPost
		//查询用户与岗位是否存在
		for _, postId := range postIds {
			var sysUserPost *entity.SysUserPost
			err = dao.SysUserPost.Ctx(ctx).Where(g.Map{
				dao.SysUserPost.Columns().UserId: userId,
				dao.SysUserPost.Columns().PostId: postId,
			}).Scan(&sysUserPost)

			if sysUserPost == nil {
				//添加用户与岗位绑定管理
				sysUserPost = new(entity.SysUserPost)
				sysUserPost.UserId = userId
				sysUserPost.PostId = postId
				sysUserPosts = append(sysUserPosts, sysUserPost)
			}
		}
		_, err = dao.SysUserPost.Ctx(ctx).Data(sysUserPosts).Insert()
		if err != nil {
			return gerror.New("绑定岗位失败")
		}
	}
	return
}

// BindUserAndRole 添加用户与角色绑定关系
func BindUserAndRole(ctx context.Context, userId int, roleIds []int) (err error) {
	if len(roleIds) > 0 {
		var sysUserRoles []*entity.SysUserRole
		//查询用户与角色是否存在
		for _, roleId := range roleIds {
			var sysUserRole *entity.SysUserRole
			err = dao.SysUserRole.Ctx(ctx).Where(g.Map{
				dao.SysUserRole.Columns().UserId: userId,
				dao.SysUserRole.Columns().RoleId: roleId,
			}).Scan(&sysUserRole)

			if sysUserRole == nil {
				//添加用户与角色绑定管理
				sysUserRole = new(entity.SysUserRole)
				sysUserRole.UserId = userId
				sysUserRole.RoleId = roleId
				sysUserRoles = append(sysUserRoles, sysUserRole)
			}
		}
		_, err = dao.SysUserRole.Ctx(ctx).Data(sysUserRoles).Insert()
		if err != nil {
			return gerror.New("绑定角色失败")
		}
	}
	return
}

// GetUserById 根据ID获取用户信息
func (s *sSysUser) GetUserById(ctx context.Context, id uint) (out *model.UserInfoOut, err error) {
	var e *entity.SysUser
	err = dao.SysUser.Ctx(ctx).Where(g.Map{
		dao.SysUser.Columns().Id:        id,
		dao.SysUser.Columns().IsDeleted: 0,
	}).Scan(&e)
	if err != nil {
		return
	}
	if e != nil {
		if err = gconv.Scan(e, &out); err != nil {
			return
		}

		//获取用户角色ID
		userRoleInfo, userRoleErr := service.SysUserRole().GetInfoByUserId(ctx, int(e.Id))
		if userRoleErr != nil {
			return nil, userRoleErr
		}
		if userRoleInfo != nil {
			var roleIds []int
			for _, userRole := range userRoleInfo {
				roleIds = append(roleIds, userRole.RoleId)
			}
			out.RoleIds = roleIds
		}
		//获取用户岗位ID
		userPostInfo, userPostErr := service.SysUserPost().GetInfoByUserId(ctx, int(e.Id))
		if userPostErr != nil {
			return nil, userPostErr
		}
		if userPostInfo != nil {
			var postIds []int
			for _, userPost := range userPostInfo {
				postIds = append(postIds, userPost.PostId)
			}
			out.PostIds = postIds
		}
	}

	return
}

// DelInfoById 根据ID删除信息
func (s *sSysUser) DelInfoById(ctx context.Context, id uint) (err error) {
	var sysUser *entity.SysUser
	err = dao.SysUser.Ctx(ctx).Where(g.Map{
		dao.SysUser.Columns().Id: id,
	}).Scan(&sysUser)
	if sysUser == nil {
		err = gerror.New("ID错误")
		return
	}
	if sysUser.IsDeleted == 1 {
		return gerror.New("用户已删除,无须重复删除")
	}
	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)
	_, err = dao.SysUser.Ctx(ctx).Data(g.Map{
		dao.SysUser.Columns().IsDeleted: 1,
		dao.SysUser.Columns().DeletedBy: loginUserId,
	}).Where(g.Map{
		dao.SysUser.Columns().Id: id,
	}).Update()
	//删除用户
	_, err = dao.SysUser.Ctx(ctx).Where(g.Map{
		dao.SysUser.Columns().Id: id,
	}).Delete()
	if err != nil {
		return gerror.New("删除用户失败")
	}
	return
}

// ResetPassword 重置密码
func (s *sSysUser) ResetPassword(ctx context.Context, id uint, userPassword string) (err error) {
	var sysUser *entity.SysUser
	err = dao.SysUser.Ctx(ctx).Where(g.Map{
		dao.SysUser.Columns().Id: id,
	}).Scan(&sysUser)
	if sysUser == nil {
		err = gerror.New("ID错误")
		return
	}
	if sysUser.Status == 0 {
		return gerror.New("用户已禁用,无法重置密码")
	}
	if sysUser.IsDeleted == 1 {
		return gerror.New("用户已删除,无法重置密码")
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)
	//判断当前登录用户是否为超级管理员角色
	var isSuperAdmin = false
	var sysUserRoles []*entity.SysUserRole
	err = dao.SysUserRole.Ctx(ctx).Where(dao.SysUserRole.Columns().UserId, loginUserId).Scan(&sysUserRoles)
	for _, sysUserRole := range sysUserRoles {
		if sysUserRole.RoleId == 1 {
			isSuperAdmin = true
			break
		}
	}
	if !isSuperAdmin {
		return gerror.New("无重置密码权限")
	}

	sysUser.UserSalt = grand.S(10)
	sysUser.UserPassword = utils.EncryptPassword(userPassword, sysUser.UserSalt)
	sysUser.UpdatedBy = uint(loginUserId)
	_, err = dao.SysUser.Ctx(ctx).Data(sysUser).Where(g.Map{
		dao.SysUser.Columns().Id: id,
	}).Update()
	if err != nil {
		return gerror.New("重置密码失败")
	}
	return
}

// EditUserStatus 修改用户状态
func (s *sSysUser) EditUserStatus(ctx context.Context, id uint, status uint) (err error) {
	var sysUser *entity.SysUser
	err = dao.SysUser.Ctx(ctx).Where(g.Map{
		dao.SysUser.Columns().Id: id,
	}).Scan(&sysUser)
	if sysUser == nil {
		err = gerror.New("ID错误")
		return
	}
	if sysUser.Status == status {
		return gerror.New("无须重复修改状态")
	}
	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)
	sysUser.Status = status
	sysUser.UpdatedBy = uint(loginUserId)
	_, err = dao.SysUser.Ctx(ctx).Data(sysUser).Where(g.Map{
		dao.SysUser.Columns().Id: id,
	}).Update()
	if err != nil {
		return gerror.New("修改状态失败")
	}
	return
}

// 获取搜索的部门ID数组
func (s *sSysUser) getSearchDeptIds(ctx context.Context, deptId int64) (deptIds []int64, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		deptAll, e := sysDeptNew().GetFromCache(ctx)
		liberr.ErrIsNil(ctx, e)
		deptWithChildren := sysDeptNew().FindSonByParentId(deptAll, gconv.Int64(deptId))
		deptIds = make([]int64, len(deptWithChildren))
		for k, v := range deptWithChildren {
			deptIds[k] = v.DeptId
		}
		deptIds = append(deptIds, deptId)
	})
	return
}

func GetDeptNameDict(ctx context.Context, ids g.Slice) (dict map[int64]model.DetailDeptRes, err error) {
	var depts []model.DetailDeptRes
	dict = make(map[int64]model.DetailDeptRes)
	err = dao.SysDept.Ctx(ctx).Where(dao.SysDept.Columns().DeptId, ids).Scan(&depts)
	for _, d := range depts {
		dict[d.DeptId] = d
	}
	return
}

// GetUserByIds 根据ID数据获取用户信息
func (s *sSysUser) GetUserByIds(ctx context.Context, id []int) (data []*entity.SysUser, err error) {
	err = dao.SysUser.Ctx(ctx).Where(g.Map{
		dao.SysUser.Columns().IsDeleted: 0,
	}).WhereIn(dao.SysUser.Columns().Id, id).Scan(&data)
	return
}

// GetAll 获取所有用户信息
func (s *sSysUser) GetAll(ctx context.Context) (data []*entity.SysUser, err error) {
	err = dao.SysUser.Ctx(ctx).Where(g.Map{
		dao.SysUser.Columns().IsDeleted: 0,
	}).Scan(&data)
	return
}

func (s *sSysUser) CurrentUser(ctx context.Context) (userInfoOut *model.UserInfoOut, menuTreeOut []*model.UserMenuTreeOut, err error) {
	cache := common.Cache()

	//获取当前登录用户信息
	loginUserId := service.Context().GetUserId(ctx)
	if loginUserId == 0 {
		err = gerror.New("无登录用户信息,请先登录!")
		return
	}
	tmpUserAuthorize := cache.Get(ctx, consts.CacheUserAuthorize+"_"+gconv.String(loginUserId))
	tmpUserInfo := cache.Get(ctx, consts.CacheUserInfo+"_"+gconv.String(loginUserId))
	if tmpUserAuthorize.Val() != nil && tmpUserInfo.Val() != nil {
		json.Unmarshal([]byte(tmpUserAuthorize.Val().(string)), &menuTreeOut)
		json.Unmarshal([]byte(tmpUserInfo.Val().(string)), &userInfoOut)
		return
	}

	//获取当前登录用户信息
	userInfo, err := service.SysUser().GetUserById(ctx, uint(loginUserId))
	if err = gconv.Scan(userInfo, &userInfoOut); err != nil {
		return
	}
	if userInfo != nil {
		cache.Set(ctx, consts.CacheUserInfo+"_"+gconv.String(loginUserId), userInfo, time.Hour)
	}

	//根据当前登录用户ID查询用户角色信息
	userRoleInfo, err := service.SysUserRole().GetInfoByUserId(ctx, loginUserId)
	if userRoleInfo == nil {
		err = gerror.New("用户无权限,禁止访问")
		return
	}
	var isSuperAdmin = false
	var roleIds []int
	//获取角色ID
	for _, role := range userRoleInfo {
		if role.RoleId == 1 {
			isSuperAdmin = true
		}
		roleIds = append(roleIds, role.RoleId)
	}
	if isSuperAdmin {
		//获取所有的菜单
		//根据菜单ID数组获取菜单列表信息
		userMenuTreeOut, menuTreeError := GetMenuInfo(ctx, nil)
		if menuTreeError != nil {
			err = menuTreeError
			return
		}
		var menuIds []int
		//获取菜单绑定的按钮信息
		for _, userMenu := range userMenuTreeOut {
			menuIds = append(menuIds, int(userMenu.Id))
		}
		//获取所有的按钮
		userMenuTreeOut, userItemsTypeTreeErr := GetUserItemsTypeTreeOut(ctx, menuIds, consts.Button, userMenuTreeOut)
		if userItemsTypeTreeErr != nil {
			err = userItemsTypeTreeErr
			return
		}
		//获取所有的列表
		userMenuTreeOut, userItemsTypeTreeErr = GetUserItemsTypeTreeOut(ctx, menuIds, consts.Column, userMenuTreeOut)
		if userItemsTypeTreeErr != nil {
			err = userItemsTypeTreeErr
			return
		}
		//获取所有的接口
		userMenuTreeOut, userItemsTypeTreeErr = GetUserItemsTypeTreeOut(ctx, menuIds, consts.Api, userMenuTreeOut)
		if userItemsTypeTreeErr != nil {
			err = userItemsTypeTreeErr
			return
		}

		//对菜单进行树状重组
		menuTreeOut = GetUserMenuTree(userMenuTreeOut)

		if menuTreeOut != nil {
			cache.Set(ctx, consts.CacheUserAuthorize+"_"+gconv.String(loginUserId), menuTreeOut, time.Hour)
		}
		return
	} else {
		//获取缓存配置信息
		var authorizeInfo []*entity.SysAuthorize
		authorizeInfo, err = service.SysAuthorize().GetInfoByRoleIds(ctx, roleIds)
		if err != nil {
			return
		}
		if authorizeInfo == nil {
			err = gerror.New("无权限配置,请联系管理员")
			return
		}
		//菜单Ids
		var menuIds []int
		//按钮Ids
		var menuButtonIds []int
		//列表Ids
		var menuColumnIds []int
		//API Ids
		var menuApiIds []int
		for _, authorize := range authorizeInfo {
			if strings.EqualFold(authorize.ItemsType, consts.Menu) {
				menuIds = append(menuIds, authorize.ItemsId)
			} else if strings.EqualFold(authorize.ItemsType, consts.Button) {
				menuButtonIds = append(menuButtonIds, authorize.ItemsId)
			} else if strings.EqualFold(authorize.ItemsType, consts.Column) {
				menuColumnIds = append(menuColumnIds, authorize.ItemsId)
			} else if strings.EqualFold(authorize.ItemsType, consts.Api) {
				menuApiIds = append(menuApiIds, authorize.ItemsId)
			}
		}
		//获取所有菜单信息
		var menuInfo []*entity.SysMenu
		menuInfo, err = service.SysMenu().GetInfoByMenuIds(ctx, menuIds)
		if err != nil {
			return
		}
		if menuInfo == nil {
			err = gerror.New("未配置菜单, 请联系系统管理员")
			return
		}
		//根据按钮ID数组获取按钮信息
		menuButtonInfo, menuButtonErr := service.SysMenuButton().GetInfoByButtonIds(ctx, menuButtonIds)
		if menuButtonErr != nil {
			return
		}
		//根据列表字段ID数据获取字段信息
		menuColumnInfo, menuColumnErr := service.SysMenuColumn().GetInfoByColumnIds(ctx, menuColumnIds)
		if menuColumnErr != nil {
			return
		}
		//根据菜单接口ID数组获取接口信息
		menuApiInfo, menuApiErr := service.SysMenuApi().GetInfoByIds(ctx, menuApiIds)
		if menuApiErr != nil {
			return
		}
		/*if menuApiInfo == nil {
			err = gerror.New("未配置相关接口访问权限, 请联系系统管理员")
			return
		}*/
		if err != nil {
			return
		}

		var userMenuTreeOut []*model.UserMenuTreeOut

		for _, menu := range menuInfo {
			var userMenuTree *model.UserMenuTreeOut
			if err = gconv.Scan(menu, &userMenuTree); err != nil {
				return
			}

			//获取菜单按钮权限
			if menuButtonInfo != nil {
				menuButtonTreeData, _ := GetUserMenuButton(int(menu.Id), menuButtonInfo)
				userMenuTree.Button = append(userMenuTree.Button, menuButtonTreeData...)
			}

			//获取列表权限
			if menuColumnInfo != nil {
				menuColumnTreeData, _ := GetUserMenuColumn(int(menu.Id), menuColumnInfo)
				userMenuTree.Column = append(userMenuTree.Column, menuColumnTreeData...)
			}

			//获取接口权限
			if menuApiIds != nil {
				//获取相关接口ID
				var apiIds []int
				for _, menuApi := range menuApiInfo {
					if menuApi.MenuId == int(menu.Id) {
						apiIds = append(apiIds, menuApi.ApiId)
					}
				}
				//获取相关接口信息
				apiInfo, _ := service.SysApi().GetInfoByIds(ctx, apiIds)
				if apiInfo != nil {
					var apiInfoOut []*model.UserApiOut
					if err = gconv.Scan(apiInfo, &apiInfoOut); err != nil {
						return
					}
					userMenuTree.Api = append(userMenuTree.Api, apiInfoOut...)
				}
			}
			userMenuTreeOut = append(userMenuTreeOut, userMenuTree)
		}

		//对菜单进行树状重组
		menuTreeOut = GetUserMenuTree(userMenuTreeOut)
		if menuTreeOut != nil {
			cache.Set(ctx, consts.CacheUserAuthorize+"_"+gconv.String(loginUserId), menuTreeOut, time.Hour)
		}
		return
	}
}

// EditUserAvatar 修改用户头像
func (s *sSysUser) EditUserAvatar(ctx context.Context, id uint, avatar string) (err error) {
	var sysUser *entity.SysUser
	err = dao.SysUser.Ctx(ctx).Where(g.Map{
		dao.SysUser.Columns().Id: id,
	}).Scan(&sysUser)
	if sysUser == nil {
		err = gerror.New("ID错误")
		return
	}
	sysUser.Avatar = avatar
	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)
	sysUser.UpdatedBy = uint(loginUserId)
	_, err = dao.SysUser.Ctx(ctx).Data(sysUser).Where(g.Map{
		dao.SysUser.Columns().Id: id,
	}).Update()
	if err != nil {
		return gerror.New("修改头像失败")
	}
	return
}

// EditUserInfo 修改用户个人资料
func (s *sSysUser) EditUserInfo(ctx context.Context, input *model.EditUserInfoInput) (err error) {
	var sysUser *entity.SysUser
	err = dao.SysUser.Ctx(ctx).Where(g.Map{
		dao.SysUser.Columns().Id: input.Id,
	}).Scan(&sysUser)
	if sysUser == nil {
		err = gerror.New("ID错误")
		return
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	//判断是否为当前用户
	if input.Id != uint64(loginUserId) {
		err = gerror.New("无法修改其他用户资料")
		return
	}
	sysUser.Mobile = input.Mobile
	sysUser.UserNickname = input.UserNickname
	sysUser.Birthday = gtime.NewFromStrFormat(input.Birthday, "Y-m-d")
	if input.UserPassword != "" {
		sysUser.UserPassword = utils.EncryptPassword(input.UserPassword, sysUser.UserSalt)
	}
	sysUser.UserEmail = input.UserEmail
	sysUser.Sex = input.Sex
	sysUser.Avatar = input.Avatar
	sysUser.Address = input.Address
	sysUser.Describe = input.Describe
	//获取当前登录用户ID
	sysUser.UpdatedBy = uint(loginUserId)
	sysUser.UpdatedAt = gtime.Now()
	_, err = dao.SysUser.Ctx(ctx).Data(sysUser).Where(g.Map{
		dao.SysUser.Columns().Id: input.Id,
	}).Update()
	if err != nil {
		return gerror.New("修改个人资料失败")
	}
	return
}
