package system

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"sagooiot/internal/consts"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/do"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"
	"sagooiot/pkg/cache"
	"sagooiot/pkg/utility/utils"
	"strconv"
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
		//判断获取是否启动安全控制和允许再次登录时间
		configKeys := []string{consts.SysIsSecurityControlEnabled, consts.SysAgainLoginDate}
		var configDatas []*entity.SysConfig
		configDatas, err = service.ConfigData().GetByKeys(ctx, configKeys)
		if err != nil {
			return
		}
		isSecurityControlEnabled := "0" //是否启动安全控制
		againLoginDate := 1             //允许再次登录时间
		for _, configData := range configDatas {
			if strings.EqualFold(configData.ConfigKey, consts.SysIsSecurityControlEnabled) {
				isSecurityControlEnabled = configData.ConfigValue
			}
			if strings.EqualFold(configData.ConfigKey, consts.SysAgainLoginDate) {
				againLoginDate, err = strconv.Atoi(configData.ConfigValue)
				if err != nil {
					err = gerror.New("允许再次登录时间配置错误")
					return nil, err
				}
			}
		}
		if strings.EqualFold(isSecurityControlEnabled, "1") {
			tmpData, _ := cache.Instance().Get(ctx, consts.CacheSysErrorPrefix+"_"+gconv.String(userName))
			var num = 1
			if tmpData.Val() != nil {
				tempValue, _ := strconv.Atoi(tmpData.Val().(string))
				num = tempValue + 1
			}
			//存入缓存
			err := cache.Instance().Set(ctx, consts.CacheSysErrorPrefix+"_"+gconv.String(userName), num, time.Duration(againLoginDate*60)*time.Second)
			if err != nil {
				return nil, err
			}
		}
		err = gerror.New("密码错误")
		return
	}
	return
}

// UpdateLoginInfo 更新用户登录信息
func (s *sSysUser) UpdateLoginInfo(ctx context.Context, id uint64, ip string) (err error) {
	_, err = dao.SysUser.Ctx(ctx).WherePri(id).Update(g.Map{
		dao.SysUser.Columns().LastLoginIp:   ip,
		dao.SysUser.Columns().LastLoginTime: gtime.Now(),
	})
	if err != nil {
		return errors.New("更新用户登录信息失败")
	}

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
		m = m.WhereIn(dao.SysUser.Columns().DeptId, deptIds)
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
		input.PageSize = consts.PageSize
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

	//判断是否有权限选择当前部门
	_, err = service.SysDept().Detail(ctx, int64(input.DeptId))
	if err != nil {
		return
	}

	//开启事务管理
	err = dao.SysUser.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
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
		//添加用户信息
		result, err := dao.SysUser.Ctx(ctx).Data(do.SysUser{
			UserName:     sysUser.UserName,
			UserTypes:    sysUser.UserTypes,
			Mobile:       sysUser.Mobile,
			UserNickname: sysUser.UserNickname,
			Birthday:     sysUser.Birthday,
			UserPassword: sysUser.UserPassword,
			UserSalt:     sysUser.UserSalt,
			UserEmail:    sysUser.UserEmail,
			Sex:          sysUser.Sex,
			Avatar:       sysUser.Avatar,
			DeptId:       sysUser.DeptId,
			Remark:       sysUser.Remark,
			IsAdmin:      sysUser.IsAdmin,
			Address:      sysUser.Address,
			Describe:     sysUser.Describe,
			Status:       sysUser.Status,
			IsDeleted:    sysUser.IsDeleted,
			CreatedBy:    uint(loginUserId),
			CreatedAt:    gtime.Now(),
		}).Insert()
		if err != nil {
			return gerror.New("添加用户失败")
		}

		//获取主键ID
		lastInsertId, err := service.Sequences().GetSequences(ctx, result, dao.SysUser.Table(), dao.SysUser.Columns().Id)
		if err != nil {
			return
		}

		//绑定岗位
		err = service.SysUserPost().BindUserAndPost(ctx, int(lastInsertId), input.PostIds)

		//绑定角色
		err = service.SysUserRole().BindUserAndRole(ctx, int(lastInsertId), input.RoleIds)

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
	err = dao.SysUser.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
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

		//绑定岗位
		err = service.SysUserPost().BindUserAndPost(ctx, int(input.Id), input.PostIds)

		//绑定角色
		err = service.SysUserRole().BindUserAndRole(ctx, int(input.Id), input.RoleIds)
		return err
	})
	return
}

// GetUserById 根据ID获取用户信息
func (s *sSysUser) GetUserById(ctx context.Context, id uint) (out *model.UserInfoOut, err error) {
	var e *entity.SysUser
	m := dao.SysUser.Ctx(ctx)

	err = m.Where(g.Map{
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
		dao.SysUser.Columns().DeletedAt: gtime.Now(),
	}).Where(g.Map{
		dao.SysUser.Columns().Id: id,
	}).Update()
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

	//判断是否启用了安全控制和启用了RSA
	configKeys := []string{consts.SysIsSecurityControlEnabled, consts.SysIsRsaEnabled}
	configDatas, err := service.ConfigData().GetByKeys(ctx, configKeys)
	if err != nil {
		return
	}
	var isSecurityControlEnabled = "0" //是否启用安装控制
	var isRsaEnbled = "0"              //是否启用RSA
	for _, configData := range configDatas {
		if strings.EqualFold(configData.ConfigKey, consts.SysIsSecurityControlEnabled) {
			isSecurityControlEnabled = configData.ConfigValue
		}
		if strings.EqualFold(configData.ConfigKey, consts.SysIsRsaEnabled) {
			isRsaEnbled = configData.ConfigValue
		}
	}

	if strings.EqualFold(isSecurityControlEnabled, "1") {
		if strings.EqualFold(isRsaEnbled, "1") {
			//对用户密码进行解密
			userPassword, err = utils.Decrypt(consts.RsaPrivateKeyFile, userPassword, consts.RsaOAEP)
			if err != nil {
				return
			}
		}

		//校验密码
		err = s.CheckPassword(ctx, userPassword)
		if err != nil {
			return
		}

		//获取用户修改密码的历史记录
		var history []*entity.SysUserPasswordHistory
		if err = dao.SysUserPasswordHistory.Ctx(ctx).Where(dao.SysUserPasswordHistory.Columns().UserId, id).OrderDesc(dao.SysUserPasswordHistory.Columns().CreatedAt).Scan(&history); err != nil {
			return
		}

		if history != nil {
			//判断与最近三次是否一样
			var num int
			if len(history) > 3 {
				num = 3
			} else {
				num = len(history)
			}
			var isExit = false
			for i := 0; i < num; i++ {
				if strings.EqualFold(history[i].AfterPassword, utils.EncryptPassword(userPassword, sysUser.UserSalt)) {
					isExit = true
					break
				}
			}
			if isExit {
				err = gerror.New("密码与前三次重复,请重新输入!")
				return
			}
		}
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

	beforePassword := sysUser.UserPassword
	afterPassword := utils.EncryptPassword(userPassword, sysUser.UserSalt)
	sysUser.UserPassword = afterPassword
	sysUser.UpdatedBy = uint(loginUserId)

	//开启事务管理
	err = dao.SysUser.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		_, err = dao.SysUser.Ctx(ctx).Data(sysUser).Where(g.Map{
			dao.SysUser.Columns().Id: id,
		}).Update()

		//添加
		_, err = dao.SysUserPasswordHistory.Ctx(ctx).Data(&do.SysUserPasswordHistory{
			UserId:         sysUser.Id,
			BeforePassword: beforePassword,
			AfterPassword:  afterPassword,
			ChangeTime:     gtime.Now(),
			CreatedAt:      gtime.Now(),
			CreatedBy:      loginUserId,
		}).Insert()

		return
	})

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
		deptAll, err := sysDeptNew().GetFromCache(ctx)
		if err != nil {
			return
		}
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
	m := dao.SysUser.Ctx(ctx)

	err = m.Where(g.Map{
		dao.SysUser.Columns().IsDeleted: 0,
	}).WhereIn(dao.SysUser.Columns().Id, id).Scan(&data)
	return
}

// GetAll 获取所有用户信息
func (s *sSysUser) GetAll(ctx context.Context) (data []*entity.SysUser, err error) {
	m := dao.SysUser.Ctx(ctx)

	err = m.Where(g.Map{
		dao.SysUser.Columns().IsDeleted: 0,
	}).Scan(&data)
	return
}

func (s *sSysUser) CurrentUser(ctx context.Context) (userInfoOut *model.UserInfoOut, menuTreeOut []*model.UserMenuTreeOut, err error) {
	//获取当前登录用户信息
	loginUserId := service.Context().GetUserId(ctx)
	if loginUserId == 0 {
		err = gerror.New("无登录用户信息,请先登录!")
		return
	}
	tmpUserAuthorize, err := cache.Instance().Get(ctx, consts.CacheUserAuthorize+"_"+gconv.String(loginUserId))
	tmpUserInfo, err := cache.Instance().Get(ctx, consts.CacheUserInfo+"_"+gconv.String(loginUserId))
	if tmpUserAuthorize.Val() != nil && tmpUserInfo.Val() != nil {
		if err = json.Unmarshal([]byte(tmpUserAuthorize.Val().(string)), &menuTreeOut); err != nil {
			return
		}
		err = json.Unmarshal([]byte(tmpUserInfo.Val().(string)), &userInfoOut)
		return
	}

	//获取当前登录用户信息
	userInfo, err := service.SysUser().GetUserById(ctx, uint(loginUserId))
	if err = gconv.Scan(userInfo, &userInfoOut); err != nil {
		return
	}
	if userInfo != nil {
		err := cache.Instance().Set(ctx, consts.CacheUserInfo+"_"+gconv.String(loginUserId), userInfo, time.Hour)
		if err != nil {
			return nil, nil, err
		}
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
			err := cache.Instance().Set(ctx, consts.CacheUserAuthorize+"_"+gconv.String(loginUserId), menuTreeOut, 0)
			if err != nil {
				return nil, nil, err
			}
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

		isSecurityControlEnabled := 0 //是否启用安全控制
		sysColumnSwitch := 0          //列表开关
		sysButtonSwitch := 0          //按钮开关
		sysApiSwitch := 0             //api开关
		//判断是否启用了安全控制、列表开关、按钮开关、api开关
		configKeys := []string{consts.SysIsSecurityControlEnabled, consts.SysColumnSwitch, consts.SysButtonSwitch, consts.SysApiSwitch}
		var configDatas []*entity.SysConfig
		configDatas, err = service.ConfigData().GetByKeys(ctx, configKeys)
		if err != nil {
			return
		}
		for _, configData := range configDatas {
			if strings.EqualFold(configData.ConfigKey, consts.SysIsSecurityControlEnabled) {
				isSecurityControlEnabled = gconv.Int(configData.ConfigValue)
			}
			if strings.EqualFold(configData.ConfigKey, consts.SysColumnSwitch) {
				sysColumnSwitch = gconv.Int(configData.ConfigValue)
			}
			if strings.EqualFold(configData.ConfigKey, consts.SysButtonSwitch) {
				sysButtonSwitch = gconv.Int(configData.ConfigValue)
			}
			if strings.EqualFold(configData.ConfigKey, consts.SysApiSwitch) {
				sysApiSwitch = gconv.Int(configData.ConfigValue)
			}
		}
		//菜单Ids
		var menuIds []int
		var menuButtonIds []int
		//列表Ids
		var menuColumnIds []int
		//API Ids
		var menuApiIds []int

		for _, authorize := range authorizeInfo {
			if strings.EqualFold(authorize.ItemsType, consts.Menu) {
				menuIds = append(menuIds, authorize.ItemsId)
			} else if strings.EqualFold(authorize.ItemsType, consts.Button) && sysButtonSwitch == 1 {
				menuButtonIds = append(menuButtonIds, authorize.ItemsId)
			} else if strings.EqualFold(authorize.ItemsType, consts.Column) && sysColumnSwitch == 1 {
				menuColumnIds = append(menuColumnIds, authorize.ItemsId)
			} else if strings.EqualFold(authorize.ItemsType, consts.Api) && sysApiSwitch == 1 {
				menuApiIds = append(menuApiIds, authorize.ItemsId)
			}
		}
		//判断按钮、列表、API开关状态，如果关闭则获取菜单对应的所有信息
		//判断按钮开关
		if isSecurityControlEnabled == 0 {
			if sysButtonSwitch == 0 {
				//获取所有按钮
				var menuButtons []*entity.SysMenuButton
				menuButtons, err = service.SysMenuButton().GetInfoByMenuIds(ctx, menuIds)
				if err != nil {
					return
				}
				if len(menuButtons) > 0 {
					for _, menuButton := range menuButtons {
						menuButtonIds = append(menuButtonIds, int(menuButton.Id))
					}
				}
			}

			//判断列表开关
			if sysColumnSwitch == 0 {
				//获取所有列表
				var menuColumns []*entity.SysMenuColumn
				menuColumns, err = service.SysMenuColumn().GetInfoByMenuIds(ctx, menuIds)
				if err != nil {
					return
				}
				if len(menuColumns) > 0 {
					for _, menuColumn := range menuColumns {
						menuColumnIds = append(menuColumnIds, int(menuColumn.Id))
					}
				}
			}
			//判断API开关
			if sysApiSwitch == 0 {
				//获取所有API
				var menuApis []*entity.SysMenuApi
				menuApis, err = service.SysMenuApi().GetInfoByMenuIds(ctx, menuIds)
				if err != nil {
					return
				}
				if len(menuApis) > 0 {
					for _, menuApi := range menuApis {
						menuApiIds = append(menuApiIds, int(menuApi.Id))
					}
				}
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
			err := cache.Instance().Set(ctx, consts.CacheUserAuthorize+"_"+gconv.String(loginUserId), menuTreeOut, 0)
			if err != nil {
				return nil, nil, err
			}
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

	//判断是否为当前用户
	if id != uint(loginUserId) {
		err = gerror.New("无法修改其他用户头像")
		return
	}

	sysUser.UpdatedBy = uint(loginUserId)
	sysUser.UpdatedAt = gtime.Now()
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

// CheckPassword 校验用户密码
func (s *sSysUser) CheckPassword(ctx context.Context, userPassword string) (err error) {
	keys := []string{consts.SysPasswordMinimumLength, consts.SysRequireComplexity, consts.SysRequireDigit, consts.SysRequireLowercaseLetter, consts.SysRequireUppercaseLetter}
	var configData []*entity.SysConfig
	configData, err = service.ConfigData().GetByKeys(ctx, keys)
	if err != nil {
		return
	}
	if configData == nil || len(configData) == 0 {
		err = gerror.New(g.I18n().T(ctx, "{#sysUserPwCheckConfig}"))
	}
	var minimumLength int
	var complexity int
	var digit int
	var lowercaseLetter int
	var uppercaseLetter int
	for _, data := range configData {
		if strings.EqualFold(data.ConfigKey, consts.SysPasswordMinimumLength) {
			minimumLength, _ = strconv.Atoi(data.ConfigValue)
		}
		if strings.EqualFold(data.ConfigKey, consts.SysRequireComplexity) {
			complexity, _ = strconv.Atoi(data.ConfigValue)
		}
		if strings.EqualFold(data.ConfigKey, consts.SysRequireDigit) {
			digit, _ = strconv.Atoi(data.ConfigValue)
		}
		if strings.EqualFold(data.ConfigKey, consts.SysRequireLowercaseLetter) {
			lowercaseLetter, _ = strconv.Atoi(data.ConfigValue)
		}
		if strings.EqualFold(data.ConfigKey, consts.SysRequireUppercaseLetter) {
			uppercaseLetter, _ = strconv.Atoi(data.ConfigValue)
		}
	}

	var flag bool
	flag, err = utils.ValidatePassword(userPassword, minimumLength, complexity, digit, lowercaseLetter, uppercaseLetter)
	if err != nil {
		return
	}
	if !flag {
		err = gerror.New(g.I18n().T(ctx, "{#sysUserPwCheckError}"))
		return
	}
	return
}

// EditPassword 修改密码
func (s *sSysUser) EditPassword(ctx context.Context, userName string, oldUserPassword string, userPassword string) (err error) {
	//判断是否启用了安全控制和启用了RSA
	configKeys := []string{consts.SysIsSecurityControlEnabled, consts.SysIsRsaEnabled}
	configDatas, err := service.ConfigData().GetByKeys(ctx, configKeys)
	if err != nil {
		return
	}
	var isSecurityControlEnabled = "0"
	var isRSAEnbled = "0"
	for _, configData := range configDatas {
		if strings.EqualFold(configData.ConfigKey, consts.SysIsSecurityControlEnabled) {
			isSecurityControlEnabled = configData.ConfigValue
		}
		if strings.EqualFold(configData.ConfigKey, consts.SysIsRsaEnabled) {
			isRSAEnbled = configData.ConfigValue
		}
	}
	if strings.EqualFold(isSecurityControlEnabled, "1") && strings.EqualFold(isRSAEnbled, "1") {
		//对账号进行解密
		oldUserPassword, err = utils.Decrypt(consts.RsaPrivateKeyFile, oldUserPassword, consts.RsaOAEP)
		if err != nil {
			return
		}
	}

	//获取用户信息
	userInfo, err := service.SysUser().GetUserByUsername(ctx, userName)
	if err != nil {
		return
	}
	//判断旧密码是否一致
	if !strings.EqualFold(userInfo.UserPassword, utils.EncryptPassword(oldUserPassword, userInfo.UserSalt)) {
		err = gerror.New("原密码输入错误, 请重新输入!")
		return
	}
	//重置密码
	err = s.ResetPassword(ctx, uint(userInfo.Id), userPassword)
	if err != nil {
		return
	}
	return
}
