package system

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mssola/useragent"
	"sagooiot/internal/consts"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/do"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"
	"sagooiot/pkg/response"
	"sagooiot/pkg/utility"
	"sagooiot/pkg/utility/utils"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
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
	ua := useragent.New(params.UserAgent)
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

// GetList 获取登录日志数据列表
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
	if req.DateRange != nil && len(req.DateRange) > 0 {
		m = m.WhereGTE(dao.SysLoginLog.Columns().LoginTime, req.DateRange[0]+" 00:00:00")
		m = m.WhereLTE(dao.SysLoginLog.Columns().LoginTime, req.DateRange[1]+" 23:59:59")
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
		req.PageSize = consts.PageSize
	}
	//获取访问日志列表信息
	err = m.Page(req.PageNum, req.PageSize).OrderDesc(dao.SysLoginLog.Columns().InfoId).Scan(&list)
	if err != nil {
		err = gerror.New("获取访问日志列表失败")
		return
	}
	return
}

// Detail 登录日志详情
func (s *sSysLoginLog) Detail(ctx context.Context, infoId int) (entity *entity.SysLoginLog, err error) {
	_ = dao.SysLoginLog.Ctx(ctx).Where(g.Map{
		dao.SysLoginLog.Columns().InfoId: infoId,
	}).Scan(&entity)
	if entity == nil {
		return nil, gerror.New("日志ID错误")
	}
	return
}

// Del 根据ID删除登录日志
func (s *sSysLoginLog) Del(ctx context.Context, infoIds []int) (err error) {
	for _, infoId := range infoIds {
		var SysLoginLog *entity.SysLoginLog
		_ = dao.SysLoginLog.Ctx(ctx).Where(g.Map{
			dao.SysLoginLog.Columns().InfoId: infoId,
		}).Scan(&SysLoginLog)
		if SysLoginLog == nil {
			return gerror.New("ID错误")
		}
	}
	//删除登录日志
	_, err = dao.SysLoginLog.Ctx(ctx).WhereIn(dao.SysLoginLog.Columns().InfoId, infoIds).
		Delete()
	return
}

// Export 导出登录日志列表
func (s *sSysLoginLog) Export(ctx context.Context, req *model.SysLoginLogInput) (err error) {
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
	if req.DateRange != nil && len(req.DateRange) > 0 {
		m = m.WhereGTE(dao.SysLoginLog.Columns().LoginTime, req.DateRange[0]+" 00:00:00")
		m = m.WhereLTE(dao.SysLoginLog.Columns().LoginTime, req.DateRange[1]+" 23:59:59")
	}
	//获取访问日志列表信息
	var outList []*model.SysLoginLogOut
	err = m.OrderDesc(dao.SysLoginLog.Columns().InfoId).Scan(&outList)
	if err != nil {
		err = gerror.New("获取访问日志列表失败")
		return
	}

	//处理数据并导出
	var resData []interface{}
	for _, out := range outList {
		var exportOut = new(model.SysLoginLogExportOut)
		if err = gconv.Scan(out, exportOut); err != nil {
			return
		}
		if out.Status == 1 {
			exportOut.Status = "成功"
		} else if out.Status == 0 {
			exportOut.Status = "失败"
		}
		resData = append(resData, exportOut)
	}
	data := utility.ToExcel(resData)
	var request = g.RequestFromCtx(ctx)
	response.ToXls(request, data, "SysLoginLog")

	return
}
