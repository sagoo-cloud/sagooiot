package context

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"sagooiot/internal/consts"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
)

type sContext struct {
}

func init() {
	service.RegisterContext(New())
}

func New() *sContext {
	return &sContext{}
}

// Init 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *sContext) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(consts.ContextKey, customCtx)
}

// Get 获得上下文变量，如果没有设置，那么返回nil
func (s *sContext) Get(ctx context.Context) *model.Context {
	value := ctx.Value(consts.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

// SetUser 将上下文信息设置到上下文请求中，注意是完整覆盖
func (s *sContext) SetUser(ctx context.Context, ctxUser *model.ContextUser) {
	s.Get(ctx).User = ctxUser
}

// GetLoginUser 获取当前登陆用户信息
func (s *sContext) GetLoginUser(ctx context.Context) *model.ContextUser {
	sysContext := s.Get(ctx)
	if sysContext == nil {
		return nil
	}
	return sysContext.User
}

// GetUserId 获取当前登录用户id
func (s *sContext) GetUserId(ctx context.Context) int {
	user := s.GetLoginUser(ctx)
	if user != nil {
		return user.Id
	}
	return 0
}

// GetUserDeptId 获取当前登录用户部门ID
func (s *sContext) GetUserDeptId(ctx context.Context) int {
	user := s.GetLoginUser(ctx)
	if user != nil {
		return user.DeptId
	}
	return 0
}

// GetChildrenDeptId 获取所有子部门ID
func (s *sContext) GetChildrenDeptId(ctx context.Context) []int {
	user := s.GetLoginUser(ctx)
	if user != nil {
		return user.ChildrenDeptId
	}
	return nil
}

// GetUserName 获取当前登录用户账户
func (s *sContext) GetUserName(ctx context.Context) string {
	user := s.GetLoginUser(ctx)
	if user != nil {
		return user.UserName
	}
	return ""
}

// GetRequestWay 获取当前系统请求方式
func (s *sContext) GetRequestWay(ctx context.Context) string {
	user := s.GetLoginUser(ctx)
	if user != nil {
		return user.RequestWay
	}
	return ""
}
