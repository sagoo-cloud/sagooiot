package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"github.com/sagoo-cloud/sagooiot/utility/response"
	"github.com/sagoo-cloud/sagooiot/utility/utils"
	"strings"
)

type sMiddleware struct {
	LoginUrl string // 登录路由地址
}

func init() {
	service.RegisterMiddleware(New())
}

func New() *sMiddleware {
	return &sMiddleware{
		LoginUrl: "/login",
	}
}

// ResponseHandler 返回处理中间件
func (s *sMiddleware) ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()

	// 如果已经有返回内容，那么该中间件什么也不做
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		err             = r.GetError()
		res             = r.GetHandlerResponse()
		code gcode.Code = gcode.CodeOK
	)
	if err != nil {
		code = gerror.Code(err)
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		if r.IsAjaxRequest() {
			response.JsonExit(r, code.Code(), err.Error())
		} else {
			/*service.View().Render500(r.Context(), model.View{
				Error: err.Error(),
			})*/
			response.JsonExit(r, code.Code(), err.Error())
		}
	} else {
		if r.IsAjaxRequest() {
			response.JsonExit(r, code.Code(), "", res)
		} else {
			// 什么都不做，业务API自行处理模板渲染的成功逻辑。
			response.JsonExit(r, code.Code(), "", res)
		}
	}
}

// Ctx 自定义上下文对象
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	ctx := r.GetCtx()
	// 初始化登录用户信息
	data, err := service.SysToken().ParseToken(r)
	if err != nil {
		// 执行下一步请求逻辑
		r.Middleware.Next()
	}
	if data != nil {
		context := new(model.Context)
		err = gconv.Struct(data.Data, &context.User)
		if err != nil {
			g.Log().Error(ctx, err)
			// 执行下一步请求逻辑
			r.Middleware.Next()
		}
		service.Context().Init(r, context)
	}
	// 执行下一步请求逻辑
	r.Middleware.Next()
}

// Auth 前台系统权限控制，用户必须登录才能访问
func (s *sMiddleware) Auth(r *ghttp.Request) {
	userId := service.Context().GetUserId(r.Context())
	if userId == 0 {
		response.JsonRedirectExit(r, consts.ErrorNotLogged, "未登录或会话已过期，请您登录后再继续", s.LoginUrl)
		return
	}

	//判断用户是否有访问权限
	url := r.Request.URL.Path
	if strings.EqualFold(url, "/api/v1/system/user/currentUser") {
		r.Middleware.Next()
		return
	}

	//获取用户角色信息
	userRoleInfo, err := service.SysUserRole().GetInfoByUserId(r.Context(), userId)
	if err != nil {
		response.JsonRedirectExit(r, consts.ErrorInvalidRole, "获取用户角色失败", "")
		return
	}
	if userRoleInfo == nil {
		response.JsonRedirectExit(r, consts.ErrorInvalidRole, "用户未配置角色信息,请联系管理员", "")
		return
	}

	var roleIds []int
	//判断是否为超级管理员
	var isSuperAdmin = false
	for _, userRole := range userRoleInfo {
		//获取角色ID
		if userRole.RoleId == 1 {
			isSuperAdmin = true
		}
		roleIds = append(roleIds, userRole.RoleId)
	}

	//超级管理员拥有所有访问权限
	if isSuperAdmin {
		r.Middleware.Next()
		return
	}

	//获取角色ID下所有的请求API
	authorizeInfo, authorizeErr := service.SysAuthorize().GetInfoByRoleIdsAndItemsType(r.Context(), roleIds, consts.Api)
	if authorizeErr != nil {
		response.JsonRedirectExit(r, consts.ErrorInvalidData, "获取用户权限失败", "")
		return
	}

	if authorizeInfo == nil || len(authorizeInfo) == 0 {
		response.JsonRedirectExit(r, consts.ErrorAccessDenied, "未授权接口,无访问权限!", "")
		return
	}

	//判断是否与当前访问接口一致
	var menuApiIds []int
	for _, authorize := range authorizeInfo {
		menuApiIds = append(menuApiIds, authorize.ItemsId)
	}
	//获取所有的接口API
	menuApiInfo, menuApiErr := service.SysMenuApi().GetInfoByIds(r.Context(), menuApiIds)
	if menuApiErr != nil {
		response.JsonRedirectExit(r, consts.ErrorInvalidData, "相关接口未配置", "")
		return
	}
	if menuApiInfo == nil || len(menuApiInfo) == 0 {
		response.JsonRedirectExit(r, consts.ErrorAccessDenied, "接口未绑定菜单,请联系管理员!", "")
		return
	}
	var apiIds []int
	for _, menuApi := range menuApiInfo {
		apiIds = append(apiIds, menuApi.ApiId)
	}
	//获取所有的接口
	apiInfo, apiErr := service.SysApi().GetInfoByIds(r.Context(), apiIds)
	if apiErr != nil {
		response.JsonRedirectExit(r, consts.ErrorInvalidData, "获取接口失败", "")
		return
	}
	if apiInfo == nil || len(apiInfo) == 0 {
		response.JsonRedirectExit(r, consts.ErrorInvalidData, "相关接口未配置", "")
		return
	}
	var isExist = false
	//获取请求路径
	for _, api := range apiInfo {
		if strings.EqualFold(url, api.Address) {
			isExist = true
			break
		}
	}
	if !isExist {
		response.JsonRedirectExit(r, consts.ErrorAccessDenied, "无权限访问", "")
		return
	}

	r.Middleware.Next()
}

// MiddlewareCORS 跨域处理
func (s *sMiddleware) MiddlewareCORS(r *ghttp.Request) {

	//自定义跨域限制
	//corsOptions := r.Response.DefaultCORSOptions()
	// you can set options
	//corsOptions.AllowDomain = []string{"goframe.org", "baidu.com"}
	//r.Response.CORS(corsOptions)

	//采用默认接受所有跨域
	r.Response.CORSDefault()

	r.Middleware.Next()
}

// OperationLog 操作日志
func (s *sMiddleware) OperationLog(r *ghttp.Request) {
	//获取当前登录用户信息
	loginUserId := service.Context().GetUserId(r.GetCtx())
	if loginUserId == 0 {
		return
	}
	var (
		url             = r.Request.URL //请求地址
		err             = r.GetError()
		handlerResponse = r.GetHandlerResponse()
		body            = r.GetMap()
	)

	res := gconv.Map(handlerResponse)

	service.SysOperLog().Invoke(r.GetCtx(), loginUserId, url, body, r.Method, utils.GetClientIp(r.GetCtx()), res, err)
}
