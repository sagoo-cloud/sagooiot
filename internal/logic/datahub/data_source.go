package datahub

import (
	"context"
	"encoding/json"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/do"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
	"github.com/sagoo-cloud/sagooiot/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type sDataSource struct{}

func init() {
	service.RegisterDataSource(dataSourceNew())
}

func dataSourceNew() *sDataSource {
	return &sDataSource{}
}

func (s *sDataSource) Add(ctx context.Context, in *model.DataSourceApiAddInput) (sourceId uint64, err error) {
	id, _ := dao.DataSource.Ctx(ctx).
		Fields(dao.DataSource.Columns().SourceId).
		Where(dao.DataSource.Columns().Key, in.Key).
		Value()
	if id.Int64() > 0 {
		err = gerror.New("数据源标识重复")
		return
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	var param *do.DataSource
	err = gconv.Scan(in, &param)
	if err != nil {
		return
	}
	param.CreateBy = uint(loginUserId)
	param.Status = 0
	param.LockKey = 0

	in.Config.Url = gstr.TrimAll(in.Config.Url)

	param.Config, err = json.Marshal(in.Config)
	if err != nil {
		err = gerror.New("数据源配置格式错误")
		return
	}

	if in.Rule != nil {
		rule, err := json.Marshal(in.Rule)
		if err != nil {
			return 0, gerror.New("规则配置格式错误")
		}
		param.Rule = rule
	}

	rs, err := dao.DataSource.Ctx(ctx).Data(param).Insert()
	if err != nil {
		return
	}

	newId, _ := rs.LastInsertId()
	sourceId = uint64(newId)

	return
}

func (s *sDataSource) Edit(ctx context.Context, in *model.DataSourceApiEditInput) (err error) {
	out, err := s.Detail(ctx, in.SourceId)
	if err != nil {
		return err
	}
	if out == nil {
		return gerror.New("数据源不存在")
	}
	if out.Status == model.DataSourceStatusOn {
		return gerror.New("数据源已发布")
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	var param *do.DataSource
	err = gconv.Scan(in, &param)
	if err != nil {
		return
	}
	param.UpdateBy = uint(loginUserId)
	param.SourceId = nil
	if out.LockKey == 1 {
		param.Key = nil
	} else {
		id, _ := dao.DataSource.Ctx(ctx).
			Fields(dao.DataSource.Columns().SourceId).
			Where(dao.DataSource.Columns().Key, in.Key).
			WhereNot(dao.DataSource.Columns().SourceId, in.SourceId).
			Value()
		if id.Int64() > 0 {
			err = gerror.New("数据源标识重复")
			return
		}
	}

	in.Config.Url = gstr.TrimAll(in.Config.Url)

	param.Config, err = json.Marshal(in.Config)
	if err != nil {
		return gerror.New("数据源配置格式错误")
	}

	if in.Rule != nil {
		rule, err := json.Marshal(in.Rule)
		if err != nil {
			return gerror.New("规则配置格式错误")
		}
		param.Rule = rule
	}

	err = dao.DataSource.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.DataSource.Ctx(ctx).
			Data(param).
			Where(dao.DataSource.Columns().SourceId, in.SourceId).
			Update()
		if err != nil {
			return err
		}

		// 同步聚合时间到定时任务管理
		job := new(model.GetJobListInput)
		job.JobGroup = "dataSourceJob"
		job.JobName = "dataSource-" + gconv.String(out.SourceId)
		job.PaginationInput = &model.PaginationInput{PageNum: 1, PageSize: 1}
		_, list, _ := service.SysJob().JobList(ctx, job)
		if len(list) > 0 {
			editJob := new(model.SysJobEditInput)
			editJob.JobName = list[0].JobName
			editJob.JobParams = list[0].JobParams
			editJob.JobGroup = list[0].JobGroup
			editJob.InvokeTarget = list[0].InvokeTarget
			editJob.CronExpression = in.Config.CronExpression
			editJob.MisfirePolicy = list[0].MisfirePolicy
			editJob.Concurrent = list[0].Concurrent
			editJob.Status = list[0].Status
			editJob.Remark = list[0].Remark

			editJob.JobId = list[0].JobId
			editJob.UpdateBy = uint64(loginUserId)
			err = service.SysJob().EditJob(ctx, editJob)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return
}

func (s *sDataSource) Del(ctx context.Context, ids []uint64) (err error) {
	var p []*entity.DataSource
	err = dao.DataSource.Ctx(ctx).WhereIn(dao.DataSource.Columns().SourceId, ids).Scan(&p)
	if len(p) == 0 {
		return gerror.New("数据源不存在")
	}
	if len(p) == 1 && p[0].Status == model.DataSourceStatusOn {
		return gerror.New("数据源已发布，请先撤回，再删除")
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	err = dao.DataNode.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 删除数据源
		var delIds []uint64
		for _, id := range ids {
			rs, err := dao.DataSource.Ctx(ctx).
				Data(do.DataSource{
					DeletedBy: uint(loginUserId),
					DeletedAt: gtime.Now(),
				}).
				Where(dao.DataSource.Columns().SourceId, id).
				Where(dao.DataSource.Columns().Status, model.DataSourceStatusOff).
				Unscoped().
				Update()
			if err != nil {
				return err
			}

			num, _ := rs.RowsAffected()
			if num > 0 {
				delIds = append(delIds, id)
			}
		}

		// 删除数据节点
		for _, id := range delIds {
			_, err = dao.DataNode.Ctx(ctx).
				Data(do.DataNode{
					DeletedBy: uint(loginUserId),
					DeletedAt: gtime.Now(),
				}).
				Where(dao.DataNode.Columns().SourceId, id).
				Unscoped().
				Update()
			if err != nil {
				return err
			}
		}

		// 删除定时任务、删除数据表
		for _, id := range delIds {
			job := new(model.GetJobListInput)
			job.JobGroup = "dataSourceJob"
			job.JobName = "dataSource-" + gconv.String(id)
			job.PaginationInput = &model.PaginationInput{PageNum: 1, PageSize: 1}
			_, list, _ := service.SysJob().JobList(ctx, job)
			if len(list) != 0 {
				err = service.SysJob().DeleteJobByIds(ctx, []int{int(list[0].JobId)})
			}

			err = dropTable(ctx, id)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return
}

func (s *sDataSource) Search(ctx context.Context, in *model.DataSourceSearchInput) (out *model.DataSourceSearchOutput, err error) {
	out = new(model.DataSourceSearchOutput)
	c := dao.DataSource.Columns()
	m := dao.DataSource.Ctx(ctx).OrderDesc(c.SourceId)

	if in.Key != "" {
		m = m.Where(c.Key, in.Key)
	}
	if in.Name != "" {
		m = m.WhereLike(c.Name, "%"+in.Name+"%")
	}
	if in.From > 0 {
		m = m.Where(c.From, in.From)
	}

	out.Total, _ = m.Count()
	out.CurrentPage = in.PageNum
	err = m.Page(in.PageNum, in.PageSize).Scan(&out.List)

	return
}

// 已发布源列表
func (s *sDataSource) List(ctx context.Context) (list []*entity.DataSource, err error) {
	err = dao.DataSource.Ctx(ctx).
		Where(dao.DataSource.Columns().Status, model.DataSourceStatusOn).
		OrderDesc(dao.DataSource.Columns().SourceId).
		Scan(&list)
	return
}

func (s *sDataSource) Detail(ctx context.Context, sourceId uint64) (out *model.DataSourceOutput, err error) {
	var p *entity.DataSource
	err = dao.DataSource.Ctx(ctx).Where(dao.DataSource.Columns().SourceId, sourceId).Scan(&p)
	if err != nil {
		return
	}
	if p == nil {
		return
	}

	// 规则配置
	var rule []*model.DataSourceRule
	if p.Rule != "" {
		j, _ := gjson.DecodeToJson(p.Rule)
		if err = j.Scan(&rule); err != nil {
			return nil, err
		}
	}

	out = new(model.DataSourceOutput)
	out.DataSource = p
	out.SourceRule = rule

	// 数据源配置
	if p.Config != "" {
		j, _ := gjson.DecodeToJson(p.Config)
		switch p.From {
		case model.DataSourceFromApi:
			// api 配置
			out.ApiConfig = &model.DataSourceConfigApi{}
			err = j.Scan(out.ApiConfig)
		case model.DataSourceFromDevice:
			// 设备配置
			out.DeviceConfig = &model.DataSourceConfigDevice{}
			err = j.Scan(out.DeviceConfig)
		case model.DataSourceFromDb:
			// 数据库配置
			out.DbConfig = &model.DataSourceConfigDb{}
			err = j.Scan(out.DbConfig)
		}
	}

	return
}

func (s *sDataSource) Deploy(ctx context.Context, sourceId uint64) (err error) {
	out, err := s.Detail(ctx, sourceId)
	if err != nil {
		return err
	}
	if out == nil {
		return gerror.New("数据源不存在")
	}
	if out.Status == model.DataSourceStatusOn {
		return gerror.New("数据源已发布")
	}

	// 获取节点
	nodeList, err := service.DataNode().List(ctx, sourceId)
	if err != nil {
		return
	}
	if len(nodeList) == 0 {
		err = gerror.New("该数据源还未创建数据节点")
		return
	}

	// 获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	err = dao.DataSource.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 创建表结构
		table := ""
		if out.From == model.DataSourceFromApi || out.From == model.DataSourceFromDb {
			if out.DataTable == "" {
				table, err = createTable(ctx, sourceId)
				if err != nil {
					return err
				}
			}
		}

		m := g.Map{
			dao.DataSource.Columns().Status:  model.DataSourceStatusOn,
			dao.DataSource.Columns().LockKey: 1,
		}
		if table != "" {
			m[dao.DataSource.Columns().DataTable] = table
		}

		_, err = dao.DataSource.Ctx(ctx).
			Data(m).
			Where(dao.DataSource.Columns().SourceId, sourceId).
			Update()

		return err
	})
	if err != nil {
		return err
	}

	if out.From != model.DataSourceFromApi && out.From != model.DataSourceFromDb {
		return
	}

	var cron string
	switch out.From {
	case model.DataSourceFromApi:
		cron = out.ApiConfig.CronExpression
	case model.DataSourceFromDb:
		cron = out.DbConfig.CronExpression
	}

	// 开启数据更新任务，任务排重
Job:
	job := new(model.GetJobListInput)
	job.JobGroup = "dataSourceJob"
	job.JobName = "dataSource-" + gconv.String(sourceId)
	job.PaginationInput = &model.PaginationInput{PageNum: 1, PageSize: 1}
	_, list, _ := service.SysJob().JobList(ctx, job)
	if len(list) == 0 {
		sysJob := new(model.SysJobAddInput)
		sysJob.JobName = "dataSource-" + gconv.String(sourceId)
		sysJob.JobParams = gconv.String(sourceId)
		sysJob.JobGroup = "dataSourceJob"
		sysJob.InvokeTarget = "dataSource"
		sysJob.CronExpression = cron
		sysJob.MisfirePolicy = 1
		sysJob.Concurrent = 0
		sysJob.Status = 0
		sysJob.Remark = out.Name
		sysJob.CreateBy = uint64(loginUserId)
		err = service.SysJob().AddJob(ctx, sysJob)
		goto Job
	}
	err = service.SysJob().JobStart(ctx, list[0])
	// 初始执行一次
	err = service.SysJob().JobRun(ctx, list[0])

	return
}

func (s *sDataSource) Undeploy(ctx context.Context, sourceId uint64) (err error) {
	out, err := s.Detail(ctx, sourceId)
	if err != nil {
		return err
	}
	if out == nil {
		return gerror.New("数据源不存在")
	}
	if out.Status == model.DataSourceStatusOff {
		return gerror.New("数据源已停用")
	}

	_, err = dao.DataSource.Ctx(ctx).
		Data(g.Map{dao.DataSource.Columns().Status: model.DataSourceStatusOff}).
		Where(dao.DataSource.Columns().SourceId, sourceId).
		Update()
	if err != nil {
		return err
	}

	if out.From != model.DataSourceFromApi && out.From != model.DataSourceFromDb {
		return
	}

	// 关闭数据更新任务
	job := new(model.GetJobListInput)
	job.JobGroup = "dataSourceJob"
	job.JobName = "dataSource-" + gconv.String(sourceId)
	job.PaginationInput = &model.PaginationInput{PageNum: 1, PageSize: 1}
	_, list, _ := service.SysJob().JobList(ctx, job)
	if len(list) == 0 {
		return
	}
	err = service.SysJob().JobStop(ctx, list[0])

	return
}

// 更新数据记录，定时任务触发
func (s *sDataSource) UpdateData(ctx context.Context, sourceId uint64) (err error) {
	out, err := s.Detail(ctx, sourceId)
	if err != nil {
		return
	}
	if out == nil {
		err = gerror.New("数据源不存在")
		return
	}

	switch out.From {
	case model.DataSourceFromApi:
		err = s.updateDataForApi(ctx, sourceId)
	case model.DataSourceFromDb:
		err = s.updateDataForDb(ctx, sourceId)
	}
	return
}

// 获取数据源的聚合数据
func (s *sDataSource) GetData(ctx context.Context, in *model.DataSourceDataInput) (out *model.DataSourceDataOutput, err error) {
	src, err := s.Detail(ctx, in.SourceId)
	if err != nil {
		return
	}
	if src == nil {
		err = gerror.New("数据源不存在")
		return
	}
	if src.Status == model.DataSourceStatusOff {
		err = gerror.New("数据源已停用")
		return
	}

	switch src.From {
	case model.DataSourceFromApi:
		out, err = s.getApiDataRecord(ctx, in, src)
	case model.DataSourceFromDevice:
		out, err = s.getDeviceDataRecord(ctx, in, src)
	case model.DataSourceFromDb:
		out, err = s.getDbDataRecord(ctx, in, src)
	}

	return
}

// 获取数据源的聚合数据，非分页
func (s *sDataSource) GetAllData(ctx context.Context, in *model.SourceDataAllInput) (out *model.SourceDataAllOutput, err error) {
	ds, err := s.Detail(ctx, in.SourceId)
	if err != nil {
		return
	}
	if ds == nil {
		err = gerror.New("数据源不存在")
		return
	}
	if ds.Status == model.DataSourceStatusOff {
		err = gerror.New("数据源已停用")
		return
	}

	switch ds.From {
	case model.DataSourceFromApi:
		out, err = s.getApiDataAllRecord(ctx, in, ds)
	case model.DataSourceFromDevice:
		out, err = s.getDeviceDataAllRecord(ctx, in, ds)
	case model.DataSourceFromDb:
		out, err = s.getDbDataAllRecord(ctx, in, ds)
	}

	return
}

// 数据源获取数据的内网方法列表，供大屏使用
func (s *sDataSource) AllSource(ctx context.Context) (out []*model.AllSourceOut, err error) {
	err = dao.DataSource.Ctx(ctx).
		Where(dao.DataSource.Columns().Status, model.DataSourceStatusOn).
		OrderDesc(dao.DataSource.Columns().SourceId).
		Scan(&out)
	if err != nil {
		return
	}

	for _, v := range out {
		v.Path = "GetAllData?pageNum=1&pageSize=10&sourceId=" + gconv.String(v.SourceId)
	}

	return
}

// 复制数据源
func (s *sDataSource) CopeSource(ctx context.Context, sourceId uint64) (err error) {
	out, err := s.Detail(ctx, sourceId)
	if err != nil {
		return
	}
	if out == nil {
		err = gerror.New("数据源不存在")
		return
	}

	switch out.From {
	case model.DataSourceFromApi:
		_, err = s.copeApiSource(ctx, out)
	case model.DataSourceFromDevice:
		_, err = s.copeDeviceSource(ctx, out)
	case model.DataSourceFromDb:
		_, err = s.copeDbSource(ctx, out)
	}

	return
}

// 复制数据源节点
func (s *sDataSource) copeNode(ctx context.Context, sourceId, newSourceId uint64) (err error) {
	nodes, err := service.DataNode().List(ctx, sourceId)
	if err != nil || len(nodes) == 0 {
		return
	}

	for _, v := range nodes {
		var in *model.DataNodeAddInput
		err = gconv.Scan(v.DataNode, &in)
		if err != nil {
			return
		}
		in.SourceId = newSourceId

		err = gconv.Scan(v.NodeRule, &in.Rule)
		if err != nil {
			return
		}

		err = service.DataNode().Add(ctx, in)
		if err != nil {
			return
		}
	}
	return
}

// 更新数据聚合时长
func (s *sDataSource) UpdateInterval(ctx context.Context, sourceId uint64, cronExpression string) (err error) {
	out, _ := s.Detail(ctx, sourceId)

	var c []byte

	switch out.From {
	case model.DataSourceFromApi:
		conf := out.ApiConfig
		conf.CronExpression = cronExpression
		c, _ = json.Marshal(conf)
	case model.DataSourceFromDb:
		conf := out.DbConfig
		conf.CronExpression = cronExpression
		c, _ = json.Marshal(conf)
	}

	_, err = dao.DataSource.Ctx(ctx).
		Data(g.Map{
			dao.DataSource.Columns().Config: c,
		}).
		Where(dao.DataSource.Columns().SourceId, sourceId).
		Update()

	return
}
