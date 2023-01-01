package datahub

import (
	"context"
	"fmt"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/do"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"reflect"
	"strings"

	"github.com/go-gota/gota/dataframe"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sDataTemplate struct{}

func init() {
	service.RegisterDataTemplate(dataTemplateNew())
}

func dataTemplateNew() *sDataTemplate {
	return &sDataTemplate{}
}

func (s *sDataTemplate) Add(ctx context.Context, in *model.DataTemplateAddInput) (id uint64, err error) {
	tid, _ := dao.DataTemplate.Ctx(ctx).
		Fields(dao.DataTemplate.Columns().Id).
		Where(dao.DataTemplate.Columns().Key, in.Key).
		Value()
	if tid.Int64() > 0 {
		err = gerror.New("数据模型标识重复")
		return
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	var param *do.DataTemplate
	err = gconv.Scan(in, &param)
	if err != nil {
		return
	}
	param.CreateBy = uint(loginUserId)
	param.Status = 0
	param.LockKey = 0

	rs, err := dao.DataTemplate.Ctx(ctx).Data(param).Insert()
	if err != nil {
		return
	}

	newId, _ := rs.LastInsertId()
	id = uint64(newId)

	return
}

func (s *sDataTemplate) Edit(ctx context.Context, in *model.DataTemplateEditInput) (err error) {
	out, err := s.Detail(ctx, in.Id)
	if err != nil {
		return err
	}
	if out == nil {
		return gerror.New("数据模型不存在")
	}
	if out.Status == model.DataTemplateStatusOn {
		return gerror.New("数据模型已发布，请先撤回")
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	var param *do.DataTemplate
	err = gconv.Scan(in, &param)
	if err != nil {
		return
	}
	param.UpdateBy = uint(loginUserId)
	param.Id = nil
	if out.LockKey == 1 {
		param.Key = nil
	} else {
		id, _ := dao.DataTemplate.Ctx(ctx).
			Fields(dao.DataTemplate.Columns().Id).
			Where(dao.DataTemplate.Columns().Key, in.Key).
			WhereNot(dao.DataTemplate.Columns().Id, in.Id).
			Value()
		if id.Int64() > 0 {
			err = gerror.New("数据模型标识重复")
			return
		}
	}

	err = dao.DataTemplate.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.DataTemplate.Ctx(ctx).
			Data(param).
			Where(dao.DataTemplate.Columns().Id, in.Id).
			Update()
		if err != nil {
			return err
		}

		// 同步聚合时间到定时任务管理
		job := new(model.GetJobListInput)
		job.JobGroup = "dataSourceJob"
		job.JobName = "dataTemplate-" + gconv.String(out.Id)
		job.PaginationInput = &model.PaginationInput{PageNum: 1, PageSize: 1}
		_, list, _ := service.SysJob().JobList(ctx, job)
		if len(list) > 0 {
			editJob := new(model.SysJobEditInput)
			editJob.JobName = list[0].JobName
			editJob.JobParams = list[0].JobParams
			editJob.JobGroup = list[0].JobGroup
			editJob.InvokeTarget = list[0].InvokeTarget
			editJob.CronExpression = in.CronExpression
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

		// 绑定业务
		err = service.DataTemplateBusi().Add(ctx, &model.DataTemplateBusiAddInput{
			DataTemplateId: in.Id,
			BusiTypes:      in.BusiTypes,
		})

		return err
	})

	return
}

func (s *sDataTemplate) Del(ctx context.Context, ids []uint64) (err error) {
	var p []*entity.DataTemplate
	err = dao.DataTemplate.Ctx(ctx).WhereIn(dao.DataTemplate.Columns().Id, ids).Scan(&p)
	if len(p) == 0 {
		return gerror.New("数据模型不存在")
	}
	if len(p) == 1 && p[0].Status == model.DataTemplateStatusOn {
		return gerror.New("数据模型已发布，请先撤回，再删除")
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	err = dao.DataTemplate.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 删除数据模型
		var delIds []uint64
		for _, id := range ids {
			rs, err := dao.DataTemplate.Ctx(ctx).
				Data(do.DataTemplate{
					DeletedBy: uint(loginUserId),
					DeletedAt: gtime.Now(),
				}).
				Where(dao.DataTemplate.Columns().Id, id).
				Where(dao.DataTemplate.Columns().Status, model.DataTemplateStatusOff).
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

		// 删除数据模型节点
		for _, id := range delIds {
			_, err = dao.DataTemplateNode.Ctx(ctx).
				Data(do.DataTemplateNode{
					DeletedBy: uint(loginUserId),
					DeletedAt: gtime.Now(),
				}).
				Where(dao.DataTemplateNode.Columns().Tid, id).
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
			job.JobName = "dataTemplate-" + gconv.String(id)
			job.PaginationInput = &model.PaginationInput{PageNum: 1, PageSize: 1}
			_, list, _ := service.SysJob().JobList(ctx, job)
			if len(list) != 0 {
				err = service.SysJob().DeleteJobByIds(ctx, []int{int(list[0].JobId)})
			}

			if err = dropTplTable(ctx, id); err != nil {
				return err
			}
		}

		return nil
	})

	return
}

func (s *sDataTemplate) Search(ctx context.Context, in *model.DataTemplateSearchInput) (out *model.DataTemplateSearchOutput, err error) {
	out = new(model.DataTemplateSearchOutput)
	c := dao.DataTemplate.Columns()
	m := dao.DataTemplate.Ctx(ctx).WithAll().OrderDesc(c.Id)

	if in.Key != "" {
		m = m.Where(c.Key, in.Key)
	}
	if in.Name != "" {
		m = m.WhereLike(c.Name, "%"+in.Name+"%")
	}

	out.Total, _ = m.Count()
	out.CurrentPage = in.PageNum
	if err = m.Page(in.PageNum, in.PageSize).Scan(&out.List); err != nil {
		return
	}

	for k, v := range out.List {
		if len(v.DataTemplateBusi) > 0 {
			for _, bs := range v.DataTemplateBusi {
				out.List[k].BusiTypes = append(out.List[k].BusiTypes, bs.BusiTypes)
			}
		}
	}

	return
}

func (s *sDataTemplate) List(ctx context.Context) (list []*entity.DataTemplate, err error) {
	c := dao.DataTemplate.Columns()
	err = dao.DataTemplate.Ctx(ctx).Where(c.Status, model.DataTemplateStatusOn).OrderDesc(c.Id).Scan(&list)
	return
}

func (s *sDataTemplate) Detail(ctx context.Context, id uint64) (out *model.DataTemplateOutput, err error) {
	err = dao.DataTemplate.Ctx(ctx).Where(dao.DataTemplate.Columns().Id, id).Scan(&out)
	return
}

func (s *sDataTemplate) Deploy(ctx context.Context, id uint64) (err error) {
	out, err := s.Detail(ctx, id)
	if err != nil {
		return err
	}
	if out == nil {
		return gerror.New("数据模型不存在")
	}
	if out.Status == model.DataTemplateStatusOn {
		return gerror.New("数据模型已发布")
	}

	// 获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	err = dao.DataTemplate.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 创建表结构
		table := ""
		if out.DataTable == "" {
			table, err = createTplTable(ctx, id)
			if err != nil {
				return err
			}
		}

		m := g.Map{
			dao.DataTemplate.Columns().Status:  model.DataTemplateStatusOn,
			dao.DataTemplate.Columns().LockKey: 1,
		}
		if table != "" {
			m[dao.DataTemplate.Columns().DataTable] = table
		}

		_, err = dao.DataTemplate.Ctx(ctx).
			Data(m).
			Where(dao.DataTemplate.Columns().Id, id).
			Update()

		return err
	})
	if err != nil {
		return err
	}

	// 开启数据更新任务，任务排重
Job:
	job := new(model.GetJobListInput)
	job.JobGroup = "dataSourceJob"
	job.JobName = "dataTemplate-" + gconv.String(id)
	job.PaginationInput = &model.PaginationInput{PageNum: 1, PageSize: 1}
	_, list, _ := service.SysJob().JobList(ctx, job)
	if len(list) == 0 {
		sysJob := new(model.SysJobAddInput)
		sysJob.JobName = "dataTemplate-" + gconv.String(id)
		sysJob.JobParams = gconv.String(id)
		sysJob.JobGroup = "dataSourceJob"
		sysJob.InvokeTarget = "dataTemplate"
		sysJob.CronExpression = out.CronExpression
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

func (s *sDataTemplate) Undeploy(ctx context.Context, id uint64) (err error) {
	out, err := s.Detail(ctx, id)
	if err != nil {
		return err
	}
	if out == nil {
		return gerror.New("数据模型不存在")
	}
	if out.Status == model.DataTemplateStatusOff {
		return gerror.New("数据模型已停用")
	}

	_, err = dao.DataTemplate.Ctx(ctx).
		Data(g.Map{dao.DataTemplate.Columns().Status: model.DataTemplateStatusOff}).
		Where(dao.DataTemplate.Columns().Id, id).
		Update()
	if err != nil {
		return err
	}

	// 关闭数据更新任务
	job := new(model.GetJobListInput)
	job.JobGroup = "dataSourceJob"
	job.JobName = "dataTemplate-" + gconv.String(id)
	job.PaginationInput = &model.PaginationInput{PageNum: 1, PageSize: 1}
	_, list, _ := service.SysJob().JobList(ctx, job)
	if len(list) == 0 {
		return
	}
	err = service.SysJob().JobStop(ctx, list[0])

	return
}

// GetData 获取数据模型的聚合数据
func (s *sDataTemplate) GetData(ctx context.Context, in *model.DataTemplateDataInput) (out *model.DataTemplateDataOutput, err error) {
	dt, err := s.Detail(ctx, in.Id)
	if err != nil {
		return
	}
	if dt == nil {
		err = gerror.New("数据模型不存在")
		return
	}

	sortNode := "created_at desc"
	if dt.SortNodeKey != "" {
		sortDesc := "desc"
		if dt.SortDesc == 2 {
			sortDesc = "asc"
		}
		sortNode = dt.SortNodeKey + " " + sortDesc
	}

	out = new(model.DataTemplateDataOutput)

	table := getTplTableName(in.Id)

	// 搜索条件
	var exp []string
	var value []any
	if in.Param != nil {
		for k, v := range in.Param {
			exp = append(exp, k)
			if reflect.TypeOf(v).Kind() == reflect.Slice {
				s := reflect.ValueOf(v)
				for i := 0; i < s.Len(); i++ {
					ele := s.Index(i)
					value = append(value, ele.Interface())
				}
			} else {
				value = append(value, v)
			}
		}
	}
	where := ""
	if len(exp) > 0 {
		where = " where " + strings.Join(exp, " and ")
	}

	sql := "select * from " + table + where + " order by " + sortNode

	out.Total, _ = g.DB(DataCenter()).GetCount(ctx, sql, value...)
	out.CurrentPage = in.PageNum

	sql += fmt.Sprintf(" limit %d, %d", (in.PageNum-1)*in.PageSize, in.PageSize)
	rs, err := g.DB(DataCenter()).GetAll(ctx, sql, value...)
	if err != nil {
		return
	}
	out.List = rs.Json()

	return
}

// 获取数据模型的聚合数据，不分页
func (s *sDataTemplate) GetAllData(ctx context.Context, in *model.TemplateDataAllInput) (out *model.TemplateDataAllOutput, err error) {
	dt, err := s.Detail(ctx, in.Id)
	if err != nil {
		return
	}
	if dt == nil {
		err = gerror.New("数据模型不存在")
		return
	}

	sortNode := "CREATED_AT ASC"
	if dt.SortNodeKey != "" {
		sortDesc := "DESC"
		if dt.SortDesc == 2 {
			sortDesc = "ASC"
		}
		sortNode = dt.SortNodeKey + " " + sortDesc
	}

	out = new(model.TemplateDataAllOutput)

	table := getTplTableName(in.Id)

	// 搜索条件
	var exp []string
	var value []any
	if in.Param != nil {
		for k, v := range in.Param {
			exp = append(exp, k)
			if reflect.TypeOf(v).Kind() == reflect.Slice {
				s := reflect.ValueOf(v)
				for i := 0; i < s.Len(); i++ {
					ele := s.Index(i)
					value = append(value, ele.Interface())
				}
			} else {
				value = append(value, v)
			}
		}
	}
	where := ""
	if len(exp) > 0 {
		where = " WHERE " + strings.Join(exp, " AND ")
	}

	sql := "SELECT * FROM " + table + where + " ORDER BY " + sortNode
	rs, err := g.DB(DataCenter()).GetAll(ctx, sql, value...)
	if err != nil {
		return
	}
	err = gconv.Scan(rs.List(), &out.List)

	return
}

// 获取数据，返回dataframe
func (s *sDataTemplate) GetDataBySql(ctx context.Context, sql string) (df dataframe.DataFrame, err error) {
	rs, err := g.DB(DataCenter()).GetAll(ctx, sql)
	if err != nil {
		return
	}
	df = dataframe.LoadMaps(rs.List())
	return
}

// GetDataByTableName 获取数据，返回dataframe
func (s *sDataTemplate) GetDataByTableName(ctx context.Context, tableName string) (df dataframe.DataFrame, err error) {
	rs, err := g.DB(DataCenter()).Model(tableName).All()
	if err != nil {
		return
	}
	df = dataframe.LoadMaps(rs.List())
	return
}

// 获取最后一条记录
func (s *sDataTemplate) GetLastData(ctx context.Context, in *model.TemplateDataLastInput) (out *model.TemplateDataLastOutput, err error) {
	dt, err := s.Detail(ctx, in.Id)
	if err != nil {
		return
	}
	if dt == nil {
		err = gerror.New("数据模型不存在")
		return
	}

	sortNode := "CREATED_AT DESC"

	out = new(model.TemplateDataLastOutput)

	table := getTplTableName(in.Id)

	// 搜索条件
	var exp []string
	var value []any
	if in.Param != nil {
		for k, v := range in.Param {
			exp = append(exp, k)
			if reflect.TypeOf(v).Kind() == reflect.Slice {
				s := reflect.ValueOf(v)
				for i := 0; i < s.Len(); i++ {
					ele := s.Index(i)
					value = append(value, ele.Interface())
				}
			} else {
				value = append(value, v)
			}
		}
	}
	where := ""
	if len(exp) > 0 {
		where = " WHERE " + strings.Join(exp, " AND ")
	}

	sql := "SELECT * FROM " + table + where + " ORDER BY " + sortNode + " LIMIT 1"
	rs, err := g.DB(DataCenter()).GetOne(ctx, sql, value...)
	if err != nil {
		return
	}
	err = gconv.Scan(rs, &out.Data)

	return
}

// 更新数据记录，定时任务触发
func (s *sDataTemplate) UpdateData(ctx context.Context, id uint64) error {
	err := service.DataTemplateRecord().UpdateData(ctx, id)
	return err
}

func (s *sDataTemplate) GetInfoByIds(ctx context.Context, ids []uint64) (data []*entity.DataTemplate, err error) {
	var p []*entity.DataTemplate
	err = dao.DataTemplate.Ctx(ctx).WhereIn(dao.DataTemplate.Columns().Id, ids).Where(g.Map{
		dao.DataTemplate.Columns().Status: 1,
	}).Scan(&p)
	if err != nil {
		return
	}
	if p == nil {
		return
	}
	return p, err
}

// 数据模型获取数据的内网方法列表，供大屏使用
func (s *sDataTemplate) AllTemplate(ctx context.Context) (out []*model.AllTemplateOut, err error) {
	err = dao.DataTemplate.Ctx(ctx).
		Where(dao.DataTemplate.Columns().Status, model.DataTemplateStatusOn).
		OrderDesc(dao.DataTemplate.Columns().Id).
		Scan(&out)
	if err != nil {
		return
	}

	for _, v := range out {
		v.Path = "GetAllData?pageNum=1&pageSize=10&id=" + gconv.String(v.Id)
	}

	return
}

// 更新数据聚合时长
func (s *sDataTemplate) UpdateInterval(ctx context.Context, id uint64, cronExpression string) (err error) {
	_, err = dao.DataTemplate.Ctx(ctx).
		Data(g.Map{
			dao.DataTemplate.Columns().CronExpression: cronExpression,
		}).
		Where(dao.DataTemplate.Columns().Id, id).
		Update()

	return
}

// 复制数据模型
func (s *sDataTemplate) CopeTemplate(ctx context.Context, id uint64) (err error) {
	out, err := s.Detail(ctx, id)
	if err != nil {
		return err
	}
	if out == nil {
		return gerror.New("数据模型不存在")
	}

	err = dao.DataTemplate.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 复制模型
		in := new(model.DataTemplateAddInput)

		err = gconv.Scan(out, &in)
		if err != nil {
			return err
		}
		in.Key += "_" + gtime.Now().Format("YmdHis")

		newId, err := s.Add(ctx, in)
		if err != nil {
			return err
		}

		// 复制节点
		tplNodes, err := service.DataTemplateNode().List(ctx, id)
		if err != nil || len(tplNodes) == 0 {
			return err
		}

		for _, v := range tplNodes {
			var in *model.DataTemplateNodeAddInput
			err = gconv.Scan(v.DataTemplateNode, &in)
			if err != nil {
				return err
			}
			in.Tid = newId
			in.IsSorting = v.IsSorting
			in.IsDesc = v.IsDesc

			err = service.DataTemplateNode().Add(ctx, in)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return
}

// 检测数据模型是否需要设置关联
func (s *sDataTemplate) CheckRelation(ctx context.Context, id uint64) (yes bool, err error) {
	rs, err := dao.DataTemplateNode.Ctx(ctx).Fields("count(distinct source_id)").Where(dao.DataTemplateNode.Columns().Tid, id).Value()
	if err != nil || rs.Int() <= 1 {
		return
	}

	dt, err := s.Detail(ctx, id)
	if err != nil {
		return
	}
	if dt == nil {
		err = gerror.New("数据模型不存在")
		return
	}

	if dt.MainSourceId == 0 || dt.SourceNodeKey == "" {
		yes = true
	}

	return
}

// 设置主源、关联字段
func (s *sDataTemplate) SetRelation(ctx context.Context, in *model.TemplateDataRelationInput) (err error) {
	dt, err := s.Detail(ctx, in.Id)
	if err != nil {
		return
	}
	if dt == nil {
		return gerror.New("数据模型不存在")
	}
	if dt.Status == model.DataTemplateStatusOn {
		return gerror.New("数据模型已发布，请先撤回")
	}

	_, err = dao.DataTemplate.Ctx(ctx).
		Data(g.Map{
			dao.DataTemplate.Columns().MainSourceId:  in.MainSourceId,
			dao.DataTemplate.Columns().SourceNodeKey: in.SourceNodeKey,
		}).
		Where(dao.DataTemplate.Columns().Id, in.Id).
		Update()
	return
}

// 数据源列表
func (s *sDataTemplate) SourceList(ctx context.Context, id uint64) (list []*model.DataSourceOutput, err error) {
	rs, err := dao.DataTemplateNode.Ctx(ctx).Fields("source_id").
		Where(dao.DataTemplateNode.Columns().Tid, id).
		Group(dao.DataTemplateNode.Columns().SourceId).
		Array()
	if err != nil || len(rs) == 0 {
		return
	}

	for _, v := range rs {
		ds, err := service.DataSource().Detail(ctx, v.Uint64())
		if err != nil {
			return nil, err
		}
		list = append(list, ds)
	}

	return
}
