package system

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/do"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"github.com/sagoo-cloud/sagooiot/utility/utils"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mssola/user_agent"
)

type sSysLoginLog struct {
}

func init() {
	service.RegisterSysLoginLog(sysLoginLogNew())
}

func sysLoginLogNew() *sSysLoginLog {
	return &sSysLoginLog{}
}

func (s *sSysLoginLog) Invoke(ctx context.Context, data *model.LoginLogParams) {
	pool := grpool.New(100)
	if err := pool.Add(ctx, func(ctx context.Context) {
		//写入日志数据
		s.Add(ctx, data)
	},
	); err != nil {
		g.Log().Error(ctx, err.Error())
	}

}

// Add 记录登录日志
func (s *sSysLoginLog) Add(ctx context.Context, params *model.LoginLogParams) {
	ua := user_agent.New(params.UserAgent)
	browser, _ := ua.Browser()
	loginData := &do.SysLoginLog{
		LoginName:     params.Username,
		Ipaddr:        params.Ip,
		LoginLocation: utils.GetCityByIp(params.Ip),
		Browser:       browser,
		Os:            ua.OS(),
		Status:        params.Status,
		Msg:           params.Msg,
		LoginTime:     gtime.Now(),
		Module:        params.Module,
	}
	_, err := dao.SysLoginLog.Ctx(ctx).Insert(loginData)
	if err != nil {
		g.Log().Error(ctx, err)
	}
}

// GetList 获取访问日志数据列表
func (s *sSysLoginLog) GetList(ctx context.Context, req *model.SysLoginLogInput) (total, page int, list []*model.SysLoginLogOut, err error) {
	m := dao.SysLoginLog.Ctx(ctx)
	if req.LoginName != "" {
		m = m.WhereLike(dao.SysLoginLog.Columns().LoginName, "%"+req.LoginName+"%")
	}
	if req.Ipaddr != "" {
		m = m.WhereLike(dao.SysLoginLog.Columns().Ipaddr, "%"+req.Ipaddr+"%")
	}
	if req.LoginLocation != "" {
		m = m.WhereLike(dao.SysLoginLog.Columns().LoginLocation, "%"+req.LoginLocation+"%")
	}
	if req.Browser != "" {
		m = m.WhereLike(dao.SysLoginLog.Columns().Browser, "%"+req.Browser+"%")
	}
	if req.Os != "" {
		m = m.WhereLike(dao.SysLoginLog.Columns().Os, "%"+req.Os+"%")
	}
	if req.Status != -1 {
		m = m.Where(dao.SysLoginLog.Columns().Status, req.Status)
	}
	/*where,err := GetDataWhere(ctx, service.Context().GetUserId(ctx), new(entity.SysLoginLog))
	if err != nil {
		return
	}
	if len(where) > 0 {
		m = m.Where(where)
	}*/
	//获取总数
	total, err = m.Count()
	if err != nil {
		err = gerror.New("获取访问日志列表数据失败")
		return
	}
	page = req.PageNum
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = consts.DefaultPageSize
	}
	//获取访问日志列表信息
	err = m.Page(req.PageNum, req.PageSize).OrderDesc(dao.SysLoginLog.Columns().InfoId).Scan(&list)
	if err != nil {
		err = gerror.New("获取访问日志列表失败")
		return
	}
	return
}

// Detail 访问日志详情
func (s *sSysLoginLog) Detail(ctx context.Context, infoId int) (entity *entity.SysLoginLog, err error) {
	_ = dao.SysLoginLog.Ctx(ctx).Where(g.Map{
		dao.SysLoginLog.Columns().InfoId: infoId,
	}).Scan(&entity)
	if entity == nil {
		return nil, gerror.New("日志ID错误")
	}
	return
}

// Del 根据ID删除访问日志
func (s *sSysLoginLog) Del(ctx context.Context, infoIds []int) (err error) {
	for _, infoId := range infoIds {
		var SysLoginLog *entity.BaseDbLink
		_ = dao.SysLoginLog.Ctx(ctx).Where(g.Map{
			dao.SysLoginLog.Columns().InfoId: infoId,
		}).Scan(&SysLoginLog)
		if SysLoginLog == nil {
			return gerror.New("ID错误")
		}
	}
	//删除访问日志
	_, err = dao.SysLoginLog.Ctx(ctx).WhereIn(dao.SysLoginLog.Columns().InfoId, infoIds).
		Delete()
	return
}
