package system

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"github.com/sagoo-cloud/sagooiot/utility/utils"
	"net/url"
	"strings"
)

type sSysOperLog struct {
}

func SysOperLog() *sSysOperLog {
	return &sSysOperLog{}
}

func init() {
	service.RegisterSysOperLog(SysOperLog())
}

// GetList 获取操作日志数据列表
func (s *sSysOperLog) GetList(ctx context.Context, input *model.SysOperLogDoInput) (total int, out []*model.SysOperLogOut, err error) {
	m := dao.SysOperLog.Ctx(ctx)
	if input.Title != "" {
		m = m.WhereLike(dao.SysOperLog.Columns().Title, "%"+input.Title+"%")
	}
	if input.BusinessType != "" {
		m = m.WhereLike(dao.SysOperLog.Columns().BusinessType, "%"+input.BusinessType+"%")
	}
	if input.Method != "" {
		m = m.WhereLike(dao.SysOperLog.Columns().Method, "%"+input.Method+"%")
	}
	if input.RequestMethod != "" {
		m = m.WhereLike(dao.SysOperLog.Columns().RequestMethod, "%"+input.RequestMethod+"%")
	}
	if input.OperatorType != "" {
		m = m.WhereLike(dao.SysOperLog.Columns().OperatorType, "%"+input.OperatorType+"%")
	}
	if input.OperName != "" {
		m = m.WhereLike(dao.SysOperLog.Columns().OperName, "%"+input.OperName+"%")
	}
	if input.DeptName != "" {
		m = m.WhereLike(dao.SysOperLog.Columns().DeptName, "%"+input.DeptName+"%")
	}
	if input.OperUrl != "" {
		m = m.WhereLike(dao.SysOperLog.Columns().OperUrl, "%"+input.OperUrl+"%")
	}
	if input.OperIp != "" {
		m = m.WhereLike(dao.SysOperLog.Columns().OperIp, "%"+input.OperIp+"%")
	}
	if input.OperLocation != "" {
		m = m.WhereLike(dao.SysOperLog.Columns().OperLocation, "%"+input.OperLocation+"%")
	}
	if input.Status != -1 {
		m = m.Where(dao.SysOperLog.Columns().Status, input.Status)
	}
	//获取总数
	total, err = m.Count()
	if err != nil {
		err = gerror.New("获取操作日志列表数据失败")
		return
	}
	if input.PageNum == 0 {
		input.PageNum = 1
	}
	if input.PageSize == 0 {
		input.PageSize = consts.DefaultPageSize
	}
	//获取操作日志列表信息
	err = m.Page(input.PageNum, input.PageSize).OrderDesc(dao.SysOperLog.Columns().OperId).Scan(&out)
	if err != nil {
		err = gerror.New("获取操作日志列表失败")
		return
	}
	return
}

func (s *sSysOperLog) Invoke(ctx context.Context, userId int, url *url.URL, param g.Map, method string, clientIp string, res map[string]interface{}, err error) {
	pool := grpool.New(100)
	if err := pool.Add(ctx, func(ctx context.Context) {
		//写入操作数据
		err = s.Add(ctx, userId, url, param, method, clientIp, res, err)
	},
	); err != nil {
		g.Log().Error(ctx, err.Error())
	}
}

// Add 添加操作日志
func (s *sSysOperLog) Add(ctx context.Context, userId int, url *url.URL, param g.Map, method string, clientIp string, res map[string]interface{}, erro error) (err error) {
	var operLogInfo = new(entity.SysOperLog)
	//根据用户ID查询用户信息
	var userInfo *entity.SysUser
	err = dao.SysUser.Ctx(ctx).Where(g.Map{
		dao.SysUser.Columns().Id:        userId,
		dao.SysUser.Columns().IsDeleted: 0,
		dao.SysUser.Columns().Status:    1,
	}).Scan(&userInfo)
	if userInfo != nil {
		//操作人员
		operLogInfo.OperName = userInfo.UserName
		//获取用户部门信息
		var deptInfo *entity.SysDept
		err = dao.SysDept.Ctx(ctx).Where(g.Map{
			dao.SysDept.Columns().DeptId:    userInfo.DeptId,
			dao.SysDept.Columns().IsDeleted: 0,
			dao.SysDept.Columns().Status:    1,
		}).Scan(&deptInfo)
		if deptInfo != nil {
			//部门名称
			operLogInfo.DeptName = deptInfo.DeptName
		}
	}
	//请求地址
	operLogInfo.Method = url.Path
	//根据请求地址获取请求信息
	apiInfo, _ := service.SysApi().GetInfoByAddress(ctx, url.Path)
	if apiInfo != nil {
		operLogInfo.Title = apiInfo.Name
	}
	//请求方法
	operLogInfo.RequestMethod = method
	//操作类型
	operLogInfo.OperatorType = 1
	//业务类型
	if strings.EqualFold(method, "POST") {
		operLogInfo.BusinessType = 1
	} else if strings.EqualFold(method, "PUT") {
		operLogInfo.BusinessType = 2
	} else if strings.EqualFold(method, "DELETE") {
		operLogInfo.BusinessType = 3
	} else {
		operLogInfo.BusinessType = 0
	}
	//请求地址
	rawQuery := url.RawQuery
	if rawQuery != "" {
		rawQuery = "?" + rawQuery
	}
	operLogInfo.OperUrl = url.Path + rawQuery
	//客户端IP
	operLogInfo.OperIp = clientIp
	//操作地址
	operLogInfo.OperLocation = utils.GetCityByIp(operLogInfo.OperIp)
	//获取当前时间
	time, err := gtime.StrToTimeFormat(gtime.Datetime(), "2006-01-02 15:04:05")
	if err != nil {
		return
	}
	//请求时间
	operLogInfo.OperTime = time
	//参数
	if param != nil {
		b, _ := gjson.Encode(param)
		if len(b) > 0 {
			operLogInfo.OperParam = string(b)
		}
	}
	var code gcode.Code = gcode.CodeOK
	if erro != nil {
		code = gerror.Code(erro)
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
	}
	//返回参数
	if code.Code() != 0 {
		operLogInfo.Status = 0
		errMsg := erro.Error()
		var (
			errorMsgMap = map[string]interface{}{
				"code":    code.Code(),
				"message": errMsg,
			}
		)
		errorMsg, _ := gjson.Encode(errorMsgMap)
		operLogInfo.ErrorMsg = string(errorMsg)
	} else {
		operLogInfo.Status = 1
		b, _ := gjson.Encode(res)
		if len(b) > 0 {
			operLogInfo.JsonResult = string(b)
		}
	}
	_, err = dao.SysOperLog.Ctx(ctx).Data(operLogInfo).Insert()
	return
}

// Detail 操作日志详情
func (s *sSysOperLog) Detail(ctx context.Context, operId int) (entity *entity.SysOperLog, err error) {
	_ = dao.SysOperLog.Ctx(ctx).Where(g.Map{
		dao.SysOperLog.Columns().OperId: operId,
	}).Scan(&entity)
	if entity == nil {
		return nil, gerror.New("日志ID错误")
	}
	return
}

// Del 根据ID删除操作日志
func (s *sSysOperLog) Del(ctx context.Context, operIds []int) (err error) {
	for _, operId := range operIds {
		var sysOperLog *entity.BaseDbLink
		_ = dao.SysOperLog.Ctx(ctx).Where(g.Map{
			dao.SysOperLog.Columns().OperId: operId,
		}).Scan(&sysOperLog)
		if sysOperLog == nil {
			return gerror.New("ID错误")
		}
	}
	//删除操作日志
	_, err = dao.SysOperLog.Ctx(ctx).WhereIn(dao.SysOperLog.Columns().OperId, operIds).Delete()
	return
}

func (s *sSysOperLog) ClearOperationLogByDays(ctx context.Context, days int) (err error) {
	_, err = dao.SysOperLog.Ctx(ctx).Delete("to_days(now())-to_days(`oper_time`) > ?", days+1)
	return
}
