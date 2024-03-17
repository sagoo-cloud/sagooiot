package middleware

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/internal/consts"
	"sagooiot/internal/model"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/queues"
	"sagooiot/internal/service"
	"sagooiot/pkg/response"
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
			response.Json(r, code.Code(), "", res)
		} else {
			// 什么都不做，业务API自行处理模板渲染的成功逻辑。
			response.Json(r, code.Code(), "", res)
		}
	}
}

// Ctx 自定义上下文对象
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	ctx := r.GetCtx()
	r.SetCtx(r.GetNeverDoneCtx())

	// 初始化登录用户信息
	data, err := service.SysToken().ParseToken(r)
	if err != nil {
		// 执行下一步请求逻辑
		r.Middleware.Next()
	}
	if data != nil {
		contextModel := new(model.Context)
		err = gconv.Struct(data.Data, &contextModel.User)
		//请求方式
		contextModel.User.RequestWay = consts.TokenAuth
		if err != nil {
			g.Log().Error(ctx, err)
			// 执行下一步请求逻辑
			r.Middleware.Next()
		}
		service.Context().Init(r, contextModel)
	}

	// 执行下一步请求逻辑
	r.Middleware.Next()
}

// Auth 前台系统权限控制，用户必须登录才能访问
func (s *sMiddleware) Auth(r *ghttp.Request) {
	userId := service.Context().GetUserId(r.Context())
	if userId == 0 {
		return
	}

	//判断是否启用安全控制
	var configDataByIsSecurityControlEnabled *entity.SysConfig
	configDataByIsSecurityControlEnabled, _ = service.ConfigData().GetConfigByKey(r.Context(), consts.SysIsSecurityControlEnabled)
	sysApiSwitch := 0
	if configDataByIsSecurityControlEnabled != nil && strings.EqualFold(configDataByIsSecurityControlEnabled.ConfigValue, "1") {
		//查询API开关是否打开
		sysApiSwitchConfig, _ := service.ConfigData().GetConfigByKey(r.Context(), consts.SysApiSwitch)
		if sysApiSwitchConfig != nil {
			sysApiSwitch = gconv.Int(sysApiSwitchConfig.ConfigValue)
		}
	}

	if sysApiSwitch == 0 {
		r.Middleware.Next()
		return
	}

	//判断用户是否有访问权限
	url := r.Request.URL.Path
	if strings.EqualFold(url, "/api/v1/system/user/currentUser") || strings.EqualFold(url, "/api/v1/common/dict/data/list") {
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
	r.SetCtx(r.GetNeverDoneCtx())
	//自定义跨域限制
	corsOptions := r.Response.DefaultCORSOptions()
	corsConfig := g.Cfg().MustGet(context.Background(), "server.allowedDomains").Strings()
	if corsConfig == nil || len(corsConfig) == 0 {
		//采用默认接受所有跨域
		r.Response.CORSDefault()
	} else {
		corsOptions.AllowDomain = corsConfig
		r.Response.CORS(corsOptions)
	}
	r.Middleware.Next()
}

// OperationLog 操作日志
func (s *sMiddleware) OperationLog(r *ghttp.Request) {
	data := service.SysOperLog().AnalysisLog(r.GetCtx())

	// 写入队列
	logData, _ := json.Marshal(data)
	err := queues.ScheduledSysOperLog.Push(context.Background(), consts.QueueDeviceAlarmLogTopic, logData, 10)
	if err != nil {
		g.Log().Debug(context.TODO(), err)
	}

}

func (s *sMiddleware) Tracing(r *ghttp.Request) {
	_, span := gtrace.NewSpan(r.Context(), r.Method+"_"+r.Request.RequestURI)
	defer span.End()
	r.Middleware.Next()
}

func (s *sMiddleware) I18n(r *ghttp.Request) {
	lang := r.GetQuery("lang", "zh-CN").String()
	r.SetCtx(gi18n.WithLanguage(r.Context(), lang))
	r.Middleware.Next()
}
