// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"net/url"
	"sagooiot/api/v1/system"
	"sagooiot/internal/model"
	"sagooiot/internal/model/entity"
	"sagooiot/pkg/gftoken"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	ISysLoginLog interface {
		Invoke(ctx context.Context, data *model.LoginLogParams)
		// Add 记录登录日志
		Add(ctx context.Context, params *model.LoginLogParams)
		// GetList 获取登录日志数据列表
		GetList(ctx context.Context, req *model.SysLoginLogInput) (total, page int, list []*model.SysLoginLogOut, err error)
		// Detail 登录日志详情
		Detail(ctx context.Context, infoId int) (entity *entity.SysLoginLog, err error)
		// Del 根据ID删除登录日志
		Del(ctx context.Context, infoIds []int) (err error)
		// Export 导出登录日志列表
		Export(ctx context.Context, req *model.SysLoginLogInput) (err error)
	}
	ISysOperLog interface {
		// GetList 获取操作日志数据列表
		GetList(ctx context.Context, input *model.SysOperLogDoInput) (total int, out []*model.SysOperLogOut, err error)
		Invoke(ctx context.Context, userId int, url *url.URL, param g.Map, method string, clientIp string, res map[string]interface{}, err error)
		// Add 添加操作日志
		Add(ctx context.Context, userId int, url *url.URL, param g.Map, method string, clientIp string, res map[string]interface{}, erro error) (err error)
		AnalysisLog(ctx context.Context) (data entity.SysOperLog)
		// RealWrite 真实写入
		RealWrite(ctx context.Context, log entity.SysOperLog) (err error)
		// Detail 操作日志详情
		Detail(ctx context.Context, operId int) (entity *entity.SysOperLog, err error)
		// Del 根据ID删除操作日志
		Del(ctx context.Context, operIds []int) (err error)
		ClearOperationLogByDays(ctx context.Context, days int) (err error)
	}
	ISysToken interface {
		GenerateToken(ctx context.Context, key string, data interface{}) (keys string, err error)
		ParseToken(r *ghttp.Request) (*gftoken.CustomClaims, error)
	}
	ISysApi interface {
		// GetInfoByIds 根据接口APIID数组获取接口信息
		GetInfoByIds(ctx context.Context, ids []int) (data []*entity.SysApi, err error)
		// GetApiByMenuId 根据ApiID获取接口信息
		GetApiByMenuId(ctx context.Context, apiId int) (data []*entity.SysApi, err error)
		// GetInfoById 根据ID获取API
		GetInfoById(ctx context.Context, id int) (entity *entity.SysApi, err error)
		// GetApiAll 获取所有接口
		GetApiAll(ctx context.Context, method string) (data []*entity.SysApi, err error)
		// GetApiTree 获取Api数结构数据
		GetApiTree(ctx context.Context, name string, address string, status int, types int) (out []*model.SysApiTreeOut, err error)
		// Add 添加Api列表
		Add(ctx context.Context, input *model.AddApiInput) (err error)
		// Detail Api列表详情
		Detail(ctx context.Context, id int) (out *model.SysApiOut, err error)
		AddMenuApi(ctx context.Context, addPageSource string, apiIds []int, menuIds []int) (err error)
		// Edit 修改Api列表
		Edit(ctx context.Context, input *model.EditApiInput) (err error)
		// Del 根据ID删除Api列表信息
		Del(ctx context.Context, Id int) (err error)
		// EditStatus 修改状态
		EditStatus(ctx context.Context, id int, status int) (err error)
		// GetInfoByAddress 根据Address获取API
		GetInfoByAddress(ctx context.Context, address string) (entity *entity.SysApi, err error)
		// GetInfoByNameAndTypes 根据名字和类型获取API
		GetInfoByNameAndTypes(ctx context.Context, name string, types int) (entity *entity.SysApi, err error)
		// ImportApiFile 导入API文件
		ImportApiFile(ctx context.Context) (err error)
	}
	ISysJob interface {
		// JobList 获取任务列表
		JobList(ctx context.Context, input *model.GetJobListInput) (total int, out []*model.SysJobOut, err error)
		// GetJobs 获取已开启执行的任务
		GetJobs(ctx context.Context) (jobs []*model.SysJobOut, err error)
		// GetJobFuns 获取任务可用方法列表
		GetJobFuns(ctx context.Context) (jobsList []*model.SysJobFunListOut, err error)
		AddJob(ctx context.Context, input *model.SysJobAddInput) (err error)
		GetJobInfoById(ctx context.Context, id int) (job *model.SysJobOut, err error)
		EditJob(ctx context.Context, input *model.SysJobEditInput) error
		// JobStart 启动任务
		JobStart(ctx context.Context, job *model.SysJobOut) error
		// JobStartMult 批量启动任务
		JobStartMult(ctx context.Context, jobsList []*model.SysJobOut) error
		// JobStop 停止任务
		JobStop(ctx context.Context, job *model.SysJobOut) (err error)
		// JobRun 执行任务
		JobRun(ctx context.Context, job *model.SysJobOut) (err error)
		// DeleteJobByIds 删除任务
		DeleteJobByIds(ctx context.Context, ids []int) (err error)
		WithValue(ctx context.Context, value string) context.Context
		Value(ctx context.Context) uint64
	}
	ISysMenuApi interface {
		// MenuApiList 根据菜单ID获取API列表
		MenuApiList(ctx context.Context, menuId int) (out []*model.SysApiAllOut, err error)
		// GetInfoByIds 根据IDS数组获取菜单信息
		GetInfoByIds(ctx context.Context, ids []int) (data []*entity.SysMenuApi, err error)
		// GetInfoByMenuIds 根据菜单ID数组获取菜单信息
		GetInfoByMenuIds(ctx context.Context, menuIds []int) (data []*entity.SysMenuApi, err error)
		// GetInfoByApiId 根据接口ID数组获取菜单信息
		GetInfoByApiId(ctx context.Context, apiId int) (data []*entity.SysMenuApi, err error)
		// GetAll 获取所有信息
		GetAll(ctx context.Context) (data []*entity.SysMenuApi, err error)
		// GetInfoByMenuId 根据菜单ID获取菜单信息
		GetInfoByMenuId(ctx context.Context, menuId int) (data []*entity.SysMenuApi, err error)
	}
	ISysMenuColumn interface {
		// GetList 获取全部菜单列表数据
		GetList(ctx context.Context, input *model.MenuColumnDoInput) (data []*model.UserMenuColumnOut, err error)
		// GetData 执行获取数据操作
		GetData(ctx context.Context, input *model.MenuColumnDoInput) (data []model.UserMenuColumnOut, err error)
		// Add 添加菜单列表
		Add(ctx context.Context, input *model.AddMenuColumnInput) (err error)
		// Detail 菜单列表详情
		Detail(ctx context.Context, Id int64) (entity *entity.SysMenuColumn, err error)
		// Edit 修改菜单列表
		Edit(ctx context.Context, input *model.EditMenuColumnInput) (err error)
		// Del 根据ID删除菜单列表信息
		Del(ctx context.Context, Id int64) (err error)
		// EditStatus 修改状态
		EditStatus(ctx context.Context, id int, menuId int, status int) (err error)
		// GetInfoByColumnIds 根据列表ID数组获取菜单信息
		GetInfoByColumnIds(ctx context.Context, ids []int) (data []*entity.SysMenuColumn, err error)
		// GetInfoByMenuIds 根据菜单ID数组获取菜单信息
		GetInfoByMenuIds(ctx context.Context, menuIds []int) (data []*entity.SysMenuColumn, err error)
		// GetInfoByMenuId 根据菜单ID获取菜单信息
		GetInfoByMenuId(ctx context.Context, menuId int) (data []*entity.SysMenuColumn, err error)
		// GetAll 获取所有的列表信息
		GetAll(ctx context.Context) (data []*entity.SysMenuColumn, err error)
	}
	ISysOrganization interface {
		// GetTree 获取组织数据
		GetTree(ctx context.Context, name string, status int) (data []*model.OrganizationOut, err error)
		// GetData 执行获取数据操作
		GetData(ctx context.Context, name string, status int) (data []*model.OrganizationOut, err error)
		// Add 添加
		Add(ctx context.Context, input *model.AddOrganizationInput) (err error)
		// Edit 修改组织
		Edit(ctx context.Context, input *model.EditOrganizationInput) (err error)
		// Detail 组织详情
		Detail(ctx context.Context, id int64) (entity *entity.SysOrganization, err error)
		// Del 根据ID删除组织信息
		Del(ctx context.Context, id int64) (err error)
		// GetAll 获取全部组织数据
		GetAll(ctx context.Context) (data []*entity.SysOrganization, err error)
		// Count 获取组织数量
		Count(ctx context.Context) (count int, err error)
	}
	ISystemPluginsConfig interface {
		// GetPluginsConfigList 获取列表数据
		GetPluginsConfigList(ctx context.Context, in *model.GetPluginsConfigListInput) (total, page int, list []*model.PluginsConfigOutput, err error)
		// GetPluginsConfigById 获取指定ID数据
		GetPluginsConfigById(ctx context.Context, id int) (out *model.PluginsConfigOutput, err error)
		// GetPluginsConfigByName 获取指定ID数据
		GetPluginsConfigByName(ctx context.Context, types, name string) (out *model.PluginsConfigOutput, err error)
		// AddPluginsConfig 添加数据
		AddPluginsConfig(ctx context.Context, in model.PluginsConfigAddInput) (err error)
		// EditPluginsConfig 修改数据
		EditPluginsConfig(ctx context.Context, in model.PluginsConfigEditInput) (err error)
		// SavePluginsConfig 更新数据，有数据就修改，没有数据就添加
		SavePluginsConfig(ctx context.Context, in model.PluginsConfigAddInput) (err error)
		// DeletePluginsConfig 删除数据
		DeletePluginsConfig(ctx context.Context, Ids []int) (err error)
		// UpdateAllPluginsConfigCache 将插件数据更新到缓存
		UpdateAllPluginsConfigCache(ctx context.Context) (err error)
		// GetPluginsConfigData 获取列表数据
		GetPluginsConfigData(pluginType, pluginName string) (res map[interface{}]interface{}, err error)
	}
	ISysPost interface {
		// GetTree 获取全部岗位数据
		GetTree(ctx context.Context, postName string, postCode string, status int) (data []*model.PostOut, err error)
		// Add 添加岗位
		Add(ctx context.Context, input *model.AddPostInput) (err error)
		// Edit 修改岗位
		Edit(ctx context.Context, input *model.EditPostInput) (err error)
		// Detail 岗位详情
		Detail(ctx context.Context, postId int64) (entity *entity.SysPost, err error)
		// GetData 执行获取数据操作
		GetData(ctx context.Context, postName string, postCode string, status int) (data []*model.PostOut, err error)
		// Del 根据ID删除岗位信息
		Del(ctx context.Context, postId int64) (err error)
		// GetUsedPost 获取正常状态的岗位
		GetUsedPost(ctx context.Context) (list []*model.DetailPostOut, err error)
	}
	ISysRole interface {
		// GetAll 获取所有的角色
		GetAll(ctx context.Context) (entity []*entity.SysRole, err error)
		GetTree(ctx context.Context, name string, status int) (out []*model.RoleTreeOut, err error)
		// Add 添加
		Add(ctx context.Context, input *model.AddRoleInput) (err error)
		// Edit 编辑
		Edit(ctx context.Context, input *model.EditRoleInput) (err error)
		// GetInfoById 根据ID获取角色信息
		GetInfoById(ctx context.Context, id uint) (entity *entity.SysRole, err error)
		// DelInfoById 根据ID删除角色信息
		DelInfoById(ctx context.Context, id uint) (err error)
		// GetRoleList 获取角色列表
		GetRoleList(ctx context.Context) (list []*model.RoleInfoOut, err error)
		// GetInfoByIds 根据ID数组获取角色信息
		GetInfoByIds(ctx context.Context, id []int) (entity []*entity.SysRole, err error)
		// DataScope 角色数据授权
		DataScope(ctx context.Context, id int, dataScope uint, deptIds []int64) (err error)
		GetAuthorizeById(ctx context.Context, id int) (menuIds []string, menuButtonIds []string, menuColumnIds []string, menuApiIds []string, err error)
	}
	ILogin interface {
		// Login 登录
		Login(ctx context.Context, verifyKey string, captcha string, userName string, password string) (loginUserOut *model.LoginUserOut, token string, isChangePassword int, err error)
		// CheckPwdErrorNum 验证密码错误次数
		CheckPwdErrorNum(ctx context.Context, userName string) (err error)
		IsChangePwd(ctx context.Context, userName string) (isChangePwd int)
		// GenUserToken 生成用户TOKEN
		GenUserToken(ctx context.Context, isSecurityControlEnabled string, ip string, userAgent string, userInfo *entity.SysUser, logMoudel string) (loginUserOut *model.LoginUserOut, token string, err error)
		LoginOut(ctx context.Context) (err error)
	}
	ISysUserOnline interface {
		Invoke(ctx context.Context, data *entity.SysUserOnline)
		// Add 记录用户在线
		Add(ctx context.Context, data *entity.SysUserOnline)
		// DelByToken 根据token删除信息
		DelByToken(ctx context.Context, token string) (err error)
		// GetInfoByToken 根据token获取
		GetInfoByToken(ctx context.Context, token string) (data *entity.SysUserOnline, err error)
		// DelByIds 根据IDS删除信息
		DelByIds(ctx context.Context, ids []int) (err error)
		GetAll(ctx context.Context) (data []*entity.SysUserOnline, err error)
		// UserOnlineList 在线用户列表
		UserOnlineList(ctx context.Context, input *model.UserOnlineDoListInput) (total int, out []*model.UserOnlineListOut, err error)
		UserOnlineStrongBack(ctx context.Context, id int) (err error)
	}
	ISysAuthorize interface {
		AuthorizeQuery(ctx context.Context, itemsType string, menuIds []int) (out []*model.AuthorizeQueryTreeOut, err error)
		// GetInfoByRoleId 根据角色ID获取权限信息
		GetInfoByRoleId(ctx context.Context, roleId int) (data []*entity.SysAuthorize, err error)
		// GetInfoByRoleIds 根据角色ID数组获取权限信息
		GetInfoByRoleIds(ctx context.Context, roleIds []int) (data []*entity.SysAuthorize, err error)
		// GetInfoByRoleIdsAndItemsType 根据角色ID和项目类型获取权限信息
		GetInfoByRoleIdsAndItemsType(ctx context.Context, roleIds []int, itemsType string) (data []*entity.SysAuthorize, err error)
		DelByRoleId(ctx context.Context, roleId int) (err error)
		Add(ctx context.Context, authorize []*entity.SysAuthorize) (err error)
		AddAuthorize(ctx context.Context, roleId int, menuIds []string, buttonIds []string, columnIds []string, apiIds []string) (err error)
		IsAllowAuthorize(ctx context.Context, roleId int) (isAllow bool, err error)
		// InitAuthorize 初始化系统权限
		InitAuthorize(ctx context.Context) (err error)
	}
	ISysCertificate interface {
		// GetList 获取列表数据
		GetList(ctx context.Context, input *model.SysCertificateListInput) (total, page int, out []*model.SysCertificateListOut, err error)
		// GetInfoById 获取指定ID数据
		GetInfoById(ctx context.Context, id int) (out *model.SysCertificateListOut, err error)
		// Add 添加数据
		Add(ctx context.Context, input *model.AddSysCertificateListInput) (err error)
		// Edit 修改数据
		Edit(ctx context.Context, input *model.EditSysCertificateListInput) (err error)
		// Delete 删除数据
		Delete(ctx context.Context, id int) (err error)
		// EditStatus 更新状态
		EditStatus(ctx context.Context, id int, status int) (err error)
		// GetAll 获取所有证书
		GetAll(ctx context.Context) (out []*entity.SysCertificate, err error)
	}
	ISysMenu interface {
		// GetAll 获取全部菜单数据
		GetAll(ctx context.Context) (data []*entity.SysMenu, err error)
		// GetTree 获取菜单数据
		GetTree(ctx context.Context, title string, status int) (data []*model.SysMenuOut, err error)
		// Add 添加菜单
		Add(ctx context.Context, input *model.AddMenuInput) (err error)
		// Detail 菜单详情
		Detail(ctx context.Context, menuId int64) (entity *entity.SysMenu, err error)
		// Edit 修改菜单
		Edit(ctx context.Context, input *model.EditMenuInput) (err error)
		// Del 根据ID删除菜单信息
		Del(ctx context.Context, menuId int64) (err error)
		// GetData 执行获取数据操作
		GetData(ctx context.Context, title string, status int) (data []*model.SysMenuOut, err error)
		// GetInfoByMenuIds 根据菜单ID数组获取菜单信息
		GetInfoByMenuIds(ctx context.Context, menuIds []int) (data []*entity.SysMenu, err error)
		GetInfoById(ctx context.Context, id int) (data *entity.SysMenu, err error)
	}
	ISysMenuButton interface {
		// GetList 获取全部菜单按钮数据
		GetList(ctx context.Context, status int, name string, menuId int) (data []*model.UserMenuButtonOut, err error)
		// GetData 执行获取数据操作
		GetData(ctx context.Context, status int, name string, menuId int) (data []model.UserMenuButtonOut, err error)
		// Add 添加菜单按钮
		Add(ctx context.Context, input *model.AddMenuButtonInput) (err error)
		// Detail 菜单按钮详情
		Detail(ctx context.Context, Id int64) (entity *entity.SysMenuButton, err error)
		// Edit 修改菜单按钮
		Edit(ctx context.Context, input *model.EditMenuButtonInput) (err error)
		// Del 根据ID删除菜单按钮信息
		Del(ctx context.Context, id int64) (err error)
		// GetInfoByButtonIds 根据按钮ID数组获取菜单按钮信息
		GetInfoByButtonIds(ctx context.Context, ids []int) (data []*entity.SysMenuButton, err error)
		// GetInfoByMenuIds 根据菜单ID数组获取菜单按钮信息
		GetInfoByMenuIds(ctx context.Context, menuIds []int) (data []*entity.SysMenuButton, err error)
		// GetInfoByMenuId 根据菜单ID数组获取菜单按钮信息
		GetInfoByMenuId(ctx context.Context, menuId int) (data []*entity.SysMenuButton, err error)
		// GetAll 获取所有的按钮信息
		GetAll(ctx context.Context) (data []*entity.SysMenuButton, err error)
		// EditStatus 修改状态
		EditStatus(ctx context.Context, id int, menuId int, status int) (err error)
	}
	ISysRoleDept interface {
		// GetInfoByRoleId 根据角色ID获取信息
		GetInfoByRoleId(ctx context.Context, roleId int) (data []*entity.SysRoleDept, err error)
	}
	ISysUser interface {
		// GetUserByUsername 通过用户名获取用户信息
		GetUserByUsername(ctx context.Context, userName string) (data *entity.SysUser, err error)
		// GetAdminUserByUsernamePassword 根据用户名和密码获取用户信息
		GetAdminUserByUsernamePassword(ctx context.Context, userName string, password string) (user *entity.SysUser, err error)
		// UpdateLoginInfo 更新用户登录信息
		UpdateLoginInfo(ctx context.Context, id uint64, ip string) (err error)
		// UserList 用户列表
		UserList(ctx context.Context, input *model.UserListDoInput) (total int, out []*model.UserListOut, err error)
		// Add 添加
		Add(ctx context.Context, input *model.AddUserInput) (err error)
		Edit(ctx context.Context, input *model.EditUserInput) (err error)
		// GetUserById 根据ID获取用户信息
		GetUserById(ctx context.Context, id uint) (out *model.UserInfoOut, err error)
		// DelInfoById 根据ID删除信息
		DelInfoById(ctx context.Context, id uint) (err error)
		// ResetPassword 重置密码
		ResetPassword(ctx context.Context, id uint, userPassword string) (err error)
		// EditUserStatus 修改用户状态
		EditUserStatus(ctx context.Context, id uint, status uint) (err error)
		// GetUserByIds 根据ID数据获取用户信息
		GetUserByIds(ctx context.Context, id []int) (data []*entity.SysUser, err error)
		// GetAll 获取所有用户信息
		GetAll(ctx context.Context) (data []*entity.SysUser, err error)
		CurrentUser(ctx context.Context) (userInfoOut *model.UserInfoOut, menuTreeOut []*model.UserMenuTreeOut, err error)
		// EditUserAvatar 修改用户头像
		EditUserAvatar(ctx context.Context, id uint, avatar string) (err error)
		// EditUserInfo 修改用户个人资料
		EditUserInfo(ctx context.Context, input *model.EditUserInfoInput) (err error)
		// CheckPassword 校验用户密码
		CheckPassword(ctx context.Context, userPassword string) (err error)
		// EditPassword 修改密码
		EditPassword(ctx context.Context, userName string, oldUserPassword string, userPassword string) (err error)
	}
	ISysUserRole interface {
		// GetInfoByUserId 根据用户ID获取信息
		GetInfoByUserId(ctx context.Context, userId int) (data []*entity.SysUserRole, err error)
		// BindUserAndRole 添加用户与角色绑定关系
		BindUserAndRole(ctx context.Context, userId int, roleIds []int) (err error)
	}
	ICaptcha interface {
		// GetVerifyImgString 获取字母数字混合验证码
		GetVerifyImgString(ctx context.Context) (idKeyC string, base64stringC string, err error)
		// VerifyString 验证输入的验证码是否正确
		VerifyString(id, answer string) bool
	}
	ISysMessage interface {
		// GetList 获取列表数据
		GetList(ctx context.Context, input *model.MessageListDoInput) (total int, out []*model.MessageListOut, err error)
		// Add 新增
		Add(ctx context.Context, messageInfo *model.AddMessageInput) (err error)
		// GetUnReadMessageAll 获取所有未读消息
		GetUnReadMessageAll(ctx context.Context, input *model.MessageListDoInput) (total int, out []*model.MessageListOut, err error)
		// GetUnReadMessageCount 获取所有未读消息数量
		GetUnReadMessageCount(ctx context.Context) (out int, err error)
		// DelMessage 删除消息
		DelMessage(ctx context.Context, ids []int) (err error)
		// ClearMessage 一键清空消息
		ClearMessage(ctx context.Context) (err error)
		// ReadMessage 阅读消息
		ReadMessage(ctx context.Context, id int) (err error)
		// ReadMessageAll 全部阅读消息
		ReadMessageAll(ctx context.Context) (err error)
		// GetUnReadMessageLast 获取用户最后一条未读消息
		GetUnReadMessageLast(ctx context.Context, userId int) (out []*model.MessageListOut, err error)
	}
	ISysNotifications interface {
		// GetSysNotificationsList 获取列表数据
		GetSysNotificationsList(ctx context.Context, input *model.GetNotificationsListInput) (total, page int, list []*model.NotificationsOut, err error)
		// GetSysNotificationsById 获取指定ID数据
		GetSysNotificationsById(ctx context.Context, id int) (out *model.NotificationsRes, err error)
		// AddSysNotifications 添加数据
		AddSysNotifications(ctx context.Context, in model.NotificationsAddInput) (err error)
		// EditSysNotifications 修改数据
		EditSysNotifications(ctx context.Context, in model.NotificationsEditInput) (err error)
		// DeleteSysNotifications 删除数据
		DeleteSysNotifications(ctx context.Context, in *system.DeleteNotificationsReq) (err error)
	}
	ISysPlugins interface {
		// GetSysPluginsList 获取列表数据
		GetSysPluginsList(ctx context.Context, in *model.GetSysPluginsListInput) (total, page int, list []*model.GetSysPluginsListOut, err error)
		// GetSysPluginsById 获取指定ID数据
		GetSysPluginsById(ctx context.Context, id int) (out *entity.SysPlugins, err error)
		// GetSysPluginsByName 根据名称获取插件数据
		GetSysPluginsByName(ctx context.Context, name string) (out *entity.SysPlugins, err error)
		// GetSysPluginsByTitle 根据TITLE获取插件数据
		GetSysPluginsByTitle(ctx context.Context, title string) (out *entity.SysPlugins, err error)
		// AddSysPlugins 添加数据
		AddSysPlugins(ctx context.Context, file *ghttp.UploadFile) (err error)
		// EditSysPlugins 修改数据
		EditSysPlugins(ctx context.Context, input *model.SysPluginsEditInput) (err error)
		// DeleteSysPlugins 删除数据
		DeleteSysPlugins(ctx context.Context, ids []int) (err error)
		// SaveSysPlugins 存入插件数据，跟据插件类型与名称，数据中只保存一份
		SaveSysPlugins(ctx context.Context, in model.SysPluginsAddInput) (err error)
		EditStatus(ctx context.Context, id int, status int) (err error)
		// GetSysPluginsTypesAll 获取所有插件的通信方式类型
		GetSysPluginsTypesAll(ctx context.Context, types string) (out []*model.SysPluginsInfoOut, err error)
	}
	ISysUserPost interface {
		// GetInfoByUserId 根据用户ID获取信息
		GetInfoByUserId(ctx context.Context, userId int) (data []*entity.SysUserPost, err error)
		// BindUserAndPost 添加用户与岗位绑定关系
		BindUserAndPost(ctx context.Context, userId int, postIds []int) (err error)
	}
	ISysDept interface {
		// GetTree 获取全部部门数据
		GetTree(ctx context.Context, deptName string, status int) (out []*model.DeptOut, err error)
		// GetData 执行获取数据操作
		GetData(ctx context.Context, deptName string, status int) (data []*model.DeptOut, err error)
		// Add 添加
		Add(ctx context.Context, input *model.AddDeptInput) (err error)
		// Edit 修改部门
		Edit(ctx context.Context, input *model.EditDeptInput) (err error)
		// Detail 部门详情
		Detail(ctx context.Context, deptId int64) (entity *entity.SysDept, err error)
		// Del 根据ID删除部门信息
		Del(ctx context.Context, deptId int64) (err error)
		// GetAll 获取全部部门数据
		GetAll(ctx context.Context) (data []*entity.SysDept, err error)
		GetFromCache(ctx context.Context) (list []*entity.SysDept, err error)
		FindSonByParentId(deptList []*entity.SysDept, deptId int64) []*entity.SysDept
		// GetDeptInfosByParentId 根据父ID获取子部门信息
		GetDeptInfosByParentId(ctx context.Context, parentId int) (data []*entity.SysDept, err error)
	}
)

var (
	localSysUserRole         ISysUserRole
	localCaptcha             ICaptcha
	localSysAuthorize        ISysAuthorize
	localSysCertificate      ISysCertificate
	localSysMenu             ISysMenu
	localSysMenuButton       ISysMenuButton
	localSysRoleDept         ISysRoleDept
	localSysUser             ISysUser
	localSysDept             ISysDept
	localSysMessage          ISysMessage
	localSysNotifications    ISysNotifications
	localSysPlugins          ISysPlugins
	localSysUserPost         ISysUserPost
	localSysApi              ISysApi
	localSysLoginLog         ISysLoginLog
	localSysOperLog          ISysOperLog
	localSysToken            ISysToken
	localSysRole             ISysRole
	localLogin               ILogin
	localSysJob              ISysJob
	localSysMenuApi          ISysMenuApi
	localSysMenuColumn       ISysMenuColumn
	localSysOrganization     ISysOrganization
	localSystemPluginsConfig ISystemPluginsConfig
	localSysPost             ISysPost
	localSysUserOnline       ISysUserOnline
)

func SysApi() ISysApi {
	if localSysApi == nil {
		panic("implement not found for interface ISysApi, forgot register?")
	}
	return localSysApi
}

func RegisterSysApi(i ISysApi) {
	localSysApi = i
}

func SysLoginLog() ISysLoginLog {
	if localSysLoginLog == nil {
		panic("implement not found for interface ISysLoginLog, forgot register?")
	}
	return localSysLoginLog
}

func RegisterSysLoginLog(i ISysLoginLog) {
	localSysLoginLog = i
}

func SysOperLog() ISysOperLog {
	if localSysOperLog == nil {
		panic("implement not found for interface ISysOperLog, forgot register?")
	}
	return localSysOperLog
}

func RegisterSysOperLog(i ISysOperLog) {
	localSysOperLog = i
}

func SysToken() ISysToken {
	if localSysToken == nil {
		panic("implement not found for interface ISysToken, forgot register?")
	}
	return localSysToken
}

func RegisterSysToken(i ISysToken) {
	localSysToken = i
}

func Login() ILogin {
	if localLogin == nil {
		panic("implement not found for interface ILogin, forgot register?")
	}
	return localLogin
}

func RegisterLogin(i ILogin) {
	localLogin = i
}

func SysJob() ISysJob {
	if localSysJob == nil {
		panic("implement not found for interface ISysJob, forgot register?")
	}
	return localSysJob
}

func RegisterSysJob(i ISysJob) {
	localSysJob = i
}

func SysMenuApi() ISysMenuApi {
	if localSysMenuApi == nil {
		panic("implement not found for interface ISysMenuApi, forgot register?")
	}
	return localSysMenuApi
}

func RegisterSysMenuApi(i ISysMenuApi) {
	localSysMenuApi = i
}

func SysMenuColumn() ISysMenuColumn {
	if localSysMenuColumn == nil {
		panic("implement not found for interface ISysMenuColumn, forgot register?")
	}
	return localSysMenuColumn
}

func RegisterSysMenuColumn(i ISysMenuColumn) {
	localSysMenuColumn = i
}

func SysOrganization() ISysOrganization {
	if localSysOrganization == nil {
		panic("implement not found for interface ISysOrganization, forgot register?")
	}
	return localSysOrganization
}

func RegisterSysOrganization(i ISysOrganization) {
	localSysOrganization = i
}

func SystemPluginsConfig() ISystemPluginsConfig {
	if localSystemPluginsConfig == nil {
		panic("implement not found for interface ISystemPluginsConfig, forgot register?")
	}
	return localSystemPluginsConfig
}

func RegisterSystemPluginsConfig(i ISystemPluginsConfig) {
	localSystemPluginsConfig = i
}

func SysPost() ISysPost {
	if localSysPost == nil {
		panic("implement not found for interface ISysPost, forgot register?")
	}
	return localSysPost
}

func RegisterSysPost(i ISysPost) {
	localSysPost = i
}

func SysRole() ISysRole {
	if localSysRole == nil {
		panic("implement not found for interface ISysRole, forgot register?")
	}
	return localSysRole
}

func RegisterSysRole(i ISysRole) {
	localSysRole = i
}

func SysUserOnline() ISysUserOnline {
	if localSysUserOnline == nil {
		panic("implement not found for interface ISysUserOnline, forgot register?")
	}
	return localSysUserOnline
}

func RegisterSysUserOnline(i ISysUserOnline) {
	localSysUserOnline = i
}

func Captcha() ICaptcha {
	if localCaptcha == nil {
		panic("implement not found for interface ICaptcha, forgot register?")
	}
	return localCaptcha
}

func RegisterCaptcha(i ICaptcha) {
	localCaptcha = i
}

func SysAuthorize() ISysAuthorize {
	if localSysAuthorize == nil {
		panic("implement not found for interface ISysAuthorize, forgot register?")
	}
	return localSysAuthorize
}

func RegisterSysAuthorize(i ISysAuthorize) {
	localSysAuthorize = i
}

func SysCertificate() ISysCertificate {
	if localSysCertificate == nil {
		panic("implement not found for interface ISysCertificate, forgot register?")
	}
	return localSysCertificate
}

func RegisterSysCertificate(i ISysCertificate) {
	localSysCertificate = i
}

func SysMenu() ISysMenu {
	if localSysMenu == nil {
		panic("implement not found for interface ISysMenu, forgot register?")
	}
	return localSysMenu
}

func RegisterSysMenu(i ISysMenu) {
	localSysMenu = i
}

func SysMenuButton() ISysMenuButton {
	if localSysMenuButton == nil {
		panic("implement not found for interface ISysMenuButton, forgot register?")
	}
	return localSysMenuButton
}

func RegisterSysMenuButton(i ISysMenuButton) {
	localSysMenuButton = i
}

func SysRoleDept() ISysRoleDept {
	if localSysRoleDept == nil {
		panic("implement not found for interface ISysRoleDept, forgot register?")
	}
	return localSysRoleDept
}

func RegisterSysRoleDept(i ISysRoleDept) {
	localSysRoleDept = i
}

func SysUser() ISysUser {
	if localSysUser == nil {
		panic("implement not found for interface ISysUser, forgot register?")
	}
	return localSysUser
}

func RegisterSysUser(i ISysUser) {
	localSysUser = i
}

func SysUserRole() ISysUserRole {
	if localSysUserRole == nil {
		panic("implement not found for interface ISysUserRole, forgot register?")
	}
	return localSysUserRole
}

func RegisterSysUserRole(i ISysUserRole) {
	localSysUserRole = i
}

func SysDept() ISysDept {
	if localSysDept == nil {
		panic("implement not found for interface ISysDept, forgot register?")
	}
	return localSysDept
}

func RegisterSysDept(i ISysDept) {
	localSysDept = i
}

func SysMessage() ISysMessage {
	if localSysMessage == nil {
		panic("implement not found for interface ISysMessage, forgot register?")
	}
	return localSysMessage
}

func RegisterSysMessage(i ISysMessage) {
	localSysMessage = i
}

func SysNotifications() ISysNotifications {
	if localSysNotifications == nil {
		panic("implement not found for interface ISysNotifications, forgot register?")
	}
	return localSysNotifications
}

func RegisterSysNotifications(i ISysNotifications) {
	localSysNotifications = i
}

func SysPlugins() ISysPlugins {
	if localSysPlugins == nil {
		panic("implement not found for interface ISysPlugins, forgot register?")
	}
	return localSysPlugins
}

func RegisterSysPlugins(i ISysPlugins) {
	localSysPlugins = i
}

func SysUserPost() ISysUserPost {
	if localSysUserPost == nil {
		panic("implement not found for interface ISysUserPost, forgot register?")
	}
	return localSysUserPost
}

func RegisterSysUserPost(i ISysUserPost) {
	localSysUserPost = i
}
