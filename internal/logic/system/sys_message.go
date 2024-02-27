package system

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/internal/consts"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/do"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"
	"strings"
)

type sSysMessage struct {
}

func sysMessageNew() *sSysMessage {
	return &sSysMessage{}
}

func init() {
	service.RegisterSysMessage(sysMessageNew())
}

// GetList 获取列表数据
func (s *sSysMessage) GetList(ctx context.Context, input *model.MessageListDoInput) (total int, out []*model.MessageListOut, err error) {
	m := g.Model(dao.SysMessagereceive.Table() + " l")
	m = m.LeftJoin(dao.SysMessage.Table()+" s", "s."+dao.SysMessage.Columns().Id+"=l."+dao.SysMessagereceive.Columns().MessageId)
	if input.Title != "" {
		m = m.WhereLike("s."+dao.SysMessage.Columns().Title, "%"+input.Title+"%")
	}
	if input.Types != -1 {
		m = m.Where("s."+dao.SysMessage.Columns().Types, input.Types)
	}
	//获取当前用户信息
	loginUserId := service.Context().GetUserId(ctx)
	m = m.Where("l."+dao.SysMessagereceive.Columns().UserId, loginUserId)

	m = m.Where("l."+dao.SysMessagereceive.Columns().IsDeleted, 0)

	//获取总数
	total, err = m.Count()
	if err != nil {
		err = gerror.New("获取消息列表数据失败")
		return
	}
	if input.PageNum == 0 {
		input.PageNum = 1
	}
	if input.PageSize == 0 {
		input.PageSize = consts.PageSize
	}
	err = m.Page(input.PageNum, input.PageSize).Fields("l.*").WithAll().OrderDesc(dao.SysMessage.Columns().CreatedAt).Scan(&out)

	if err != nil {
		err = gerror.New("获取消息列表失败")
		return
	}
	return
}

// Add 新增
func (s *sSysMessage) Add(ctx context.Context, messageInfo *model.AddMessageInput) (err error) {
	if messageInfo.Title == "" {
		err = gerror.New("标题不能为空")
		return
	}
	if messageInfo.Content == "" {
		err = gerror.New("内容不能为空")
		return
	}
	if messageInfo.Scope == 0 {
		err = gerror.New("消息范围不能为空")
		return
	}
	if messageInfo.Types == 0 {
		err = gerror.New("消息类型不能为空")
		return
	}
	err = dao.SysMessage.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		//新增消息
		/*var message = new(entity.SysMessage)
		message.Types = messageInfo.Types
		message.Title = messageInfo.Title
		message.Scope = messageInfo.Scope
		message.Content = messageInfo.Content*/
		/*message.IsDeleted = 0*/
		loginUserId := service.Context().GetUserId(ctx)
		/*message.CreatedBy = uint(loginUserId)*/
		result, err := dao.SysMessage.Ctx(ctx).Data(do.SysMessage{
			Title:     messageInfo.Title,
			Types:     messageInfo.Types,
			Scope:     messageInfo.Scope,
			Content:   messageInfo.Content,
			IsDeleted: 0,
			CreatedBy: uint(loginUserId),
			CreatedAt: gtime.Now(),
		}).Insert()
		if err != nil {
			return
		}

		//获取主键ID
		lastInsertId, err := service.Sequences().GetSequences(ctx, result, dao.SysMessage.Table(), dao.SysMessage.Columns().Id)
		if err != nil {
			return
		}

		//判断消息范围
		var sysDictData *entity.SysDictData
		err = dao.SysDictData.Ctx(ctx).Where(dao.SysDictData.Columns().DictCode, messageInfo.Scope).Scan(&sysDictData)
		if sysDictData == nil {
			err = gerror.New("类型错误")
			return
		}
		var messagereceives []*entity.SysMessagereceive
		if strings.EqualFold(sysDictData.DictValue, "1") {
			//系统消息
			//查询所有用户
			var user []*entity.SysUser
			err = dao.SysUser.Ctx(ctx).Where(g.Map{
				dao.SysUser.Columns().IsDeleted: 0,
				dao.SysUser.Columns().Status:    1,
			}).Scan(&user)
			if user != nil && len(user) > 0 {
				for _, v := range user {
					var messagereceive = new(entity.SysMessagereceive)
					messagereceive.MessageId = int(lastInsertId)
					messagereceive.UserId = int(v.Id)
					messagereceive.IsRead = 0
					messagereceive.IsPush = 0
					messagereceive.IsDeleted = 0
					messagereceives = append(messagereceives, messagereceive)
				}
			}
		} else if strings.EqualFold(sysDictData.DictValue, "2") {
			if messageInfo.ObjectId == 0 {
				err = gerror.New("推送组织不能为空")
				return
			}
			//组织消息
			var dept []*entity.SysDept
			err = dao.SysDept.Ctx(ctx).Where(g.Map{
				dao.SysDept.Columns().OrganizationId: messageInfo.ObjectId,
				dao.SysDept.Columns().IsDeleted:      0,
				dao.SysDept.Columns().Status:         1,
			}).Scan(&dept)
			var deptId []int64
			if dept != nil && len(dept) > 0 {
				for _, v := range dept {
					deptId = append(deptId, v.DeptId)
				}
			}
			//根据部门ID获取用户信息
			var user []*entity.SysUser
			err = dao.SysUser.Ctx(ctx).Where(g.Map{
				dao.SysUser.Columns().IsDeleted: 0,
				dao.SysUser.Columns().Status:    1,
			}).WhereIn(dao.SysUser.Columns().DeptId, deptId).Scan(&user)

			if user != nil && len(user) > 0 {
				for _, v := range user {
					var messagereceive = new(entity.SysMessagereceive)
					messagereceive.MessageId = int(lastInsertId)
					messagereceive.UserId = int(v.Id)
					messagereceive.IsRead = 0
					messagereceive.IsPush = 0
					messagereceive.IsDeleted = 0
					messagereceives = append(messagereceives, messagereceive)
				}
			}

		} else if strings.EqualFold(sysDictData.DictValue, "3") {
			if messageInfo.ObjectId == 0 {
				err = gerror.New("推送部门不能为空")
				return
			}
			//部门消息
			//根据部门ID获取用户信息
			var user []*entity.SysUser
			err = dao.SysUser.Ctx(ctx).Where(g.Map{
				dao.SysUser.Columns().IsDeleted: 0,
				dao.SysUser.Columns().Status:    1,
				dao.SysUser.Columns().DeptId:    messageInfo.ObjectId,
			}).Scan(&user)
			if user != nil && len(user) > 0 {
				for _, v := range user {
					var messagereceive = new(entity.SysMessagereceive)
					messagereceive.MessageId = int(lastInsertId)
					messagereceive.UserId = int(v.Id)
					messagereceive.IsRead = 0
					messagereceive.IsPush = 0
					messagereceive.IsDeleted = 0
					messagereceives = append(messagereceives, messagereceive)
				}
			}
		} else if strings.EqualFold(sysDictData.DictValue, "4") {
			if messageInfo.ObjectId == 0 {
				err = gerror.New("推送用户不能为空")
				return
			}
			//用户消息
			num, _ := dao.SysUser.Ctx(ctx).Where(g.Map{
				dao.SysUser.Columns().IsDeleted: 0,
				dao.SysUser.Columns().Status:    1,
				dao.SysUser.Columns().Id:        messageInfo.ObjectId,
			}).Count()

			if num == 0 {
				err = gerror.New("用户不存在")
				return
			}
			var messagereceive = new(entity.SysMessagereceive)
			messagereceive.MessageId = int(lastInsertId)
			messagereceive.UserId = messageInfo.ObjectId
			messagereceive.IsRead = 0
			messagereceive.IsPush = 0
			messagereceive.IsDeleted = 0
			messagereceives = append(messagereceives, messagereceive)
		}
		if messagereceives != nil && len(messagereceives) > 0 {
			var messagereceiveInput []*model.SysMessagereceiveInput
			if err = gconv.Scan(messagereceives, &messagereceiveInput); err != nil {
				return
			}
			//添加推送消息
			_, err = dao.SysMessagereceive.Ctx(ctx).Data(messagereceiveInput).Insert()
			if err != nil {
				return
			}
		}
		return
	})
	return
}

// GetUnReadMessageAll 获取所有未读消息
func (s *sSysMessage) GetUnReadMessageAll(ctx context.Context, input *model.MessageListDoInput) (total int, out []*model.MessageListOut, err error) {
	m := g.Model(dao.SysMessagereceive.Table() + " l")
	m = m.LeftJoin(dao.SysMessage.Table()+" s", "s."+dao.SysMessage.Columns().Id+"=l."+dao.SysMessagereceive.Columns().MessageId)
	//获取当前用户信息
	loginUserId := service.Context().GetUserId(ctx)
	m = m.Where("l."+dao.SysMessagereceive.Columns().UserId, loginUserId)
	m = m.Where("l."+dao.SysMessagereceive.Columns().IsRead, 0)
	m = m.Where("l."+dao.SysMessagereceive.Columns().IsDeleted, 0)
	//获取总数
	total, err = m.Count()
	if err != nil {
		err = gerror.New("获取消息列表数据失败")
		return
	}
	if input.PageNum == 0 {
		input.PageNum = 1
	}
	if input.PageSize == 0 {
		input.PageSize = consts.PageSize
	}

	err = m.Page(input.PageNum, input.PageSize).Fields("l.*").WithAll().OrderDesc(dao.SysMessage.Columns().CreatedAt).Scan(&out)
	if err != nil {
		err = gerror.New("获取消息列表失败")
		return
	}
	return
}

// GetUnReadMessageCount 获取所有未读消息数量
func (s *sSysMessage) GetUnReadMessageCount(ctx context.Context) (out int, err error) {
	m := dao.SysMessagereceive.Ctx(ctx)
	m = m.LeftJoin(dao.SysMessage.Table(), "base_messagereceive.message_id = base_message.id")
	//获取当前用户信息
	loginUserId := service.Context().GetUserId(ctx)
	m = m.Where(dao.SysMessagereceive.Columns().UserId, loginUserId)
	m = m.Where(dao.SysMessagereceive.Columns().IsRead, 0)
	m = m.Where(dao.SysMessagereceive.Columns().IsDeleted, 0)
	out, err = m.WithAll().Count()
	if err != nil {
		err = gerror.New("获取消息数量失败")
		return
	}
	return
}

// DelMessage 删除消息
func (s *sSysMessage) DelMessage(ctx context.Context, ids []int) (err error) {
	var memberreceives []*entity.SysMessagereceive
	err = dao.SysMessagereceive.Ctx(ctx).Where(g.Map{
		dao.SysMessagereceive.Columns().IsDeleted: 0,
	}).WhereIn(dao.SysMessagereceive.Columns().Id, ids).Scan(&memberreceives)

	for _, memberreceive := range memberreceives {
		time, _ := gtime.StrToTimeFormat(gtime.Datetime(), "2006-01-02 15:04:05")
		memberreceive.IsDeleted = 1
		memberreceive.DeletedAt = time
		_, err = dao.SysMessagereceive.Ctx(ctx).Where(dao.SysMessagereceive.Columns().Id, memberreceive.Id).Data(memberreceive).Update()
		if err != nil {
			return
		}
	}
	return
}

// ClearMessage 一键清空消息
func (s *sSysMessage) ClearMessage(ctx context.Context) (err error) {
	time, _ := gtime.StrToTimeFormat(gtime.Datetime(), "2006-01-02 15:04:05")
	_, err = dao.SysMessagereceive.Ctx(ctx).Data(g.Map{dao.SysMessagereceive.Columns().IsDeleted: 1,
		dao.SysMessagereceive.Columns().DeletedAt: time}).Update()
	if err != nil {
		return
	}
	return
}

// ReadMessage 阅读消息
func (s *sSysMessage) ReadMessage(ctx context.Context, id int) (err error) {
	var memberreceive *entity.SysMessagereceive
	err = dao.SysMessagereceive.Ctx(ctx).Where(g.Map{
		dao.SysMessagereceive.Columns().Id:        id,
		dao.SysMessagereceive.Columns().IsDeleted: 0,
	}).Scan(&memberreceive)
	if memberreceive == nil {
		err = gerror.New("ID错误")
		return
	}
	time, err := gtime.StrToTimeFormat(gtime.Datetime(), "2006-01-02 15:04:05")
	if err != nil {
		return
	}
	memberreceive.IsRead = 1
	memberreceive.ReadTime = time
	_, err = dao.SysMessagereceive.Ctx(ctx).Where(dao.SysMessagereceive.Columns().Id, id).Data(memberreceive).Update()
	if err != nil {
		return
	}
	return
}

// ReadMessageAll 全部阅读消息
func (s *sSysMessage) ReadMessageAll(ctx context.Context) (err error) {
	loginUserId := service.Context().GetUserId(ctx)
	time, err := gtime.StrToTimeFormat(gtime.Datetime(), "2006-01-02 15:04:05")
	if err != nil {
		return
	}
	_, err = dao.SysMessagereceive.Ctx(ctx).Where(dao.SysMessagereceive.Columns().UserId, loginUserId).Data(g.Map{
		dao.SysMessagereceive.Columns().IsRead:   1,
		dao.SysMessagereceive.Columns().ReadTime: time,
	}).Update()
	if err != nil {
		return
	}
	return
}

// GetUnReadMessageLast 获取用户最后一条未读消息
func (s *sSysMessage) GetUnReadMessageLast(ctx context.Context, userId int) (out []*model.MessageListOut, err error) {
	//TODO 这个地方从缓存（Redis）中取 ===============
	m := g.Model(dao.SysMessagereceive.Table() + " l")
	m = m.LeftJoin(dao.SysMessage.Table()+" s", "s."+dao.SysMessage.Columns().Id+"=l."+dao.SysMessagereceive.Columns().MessageId)
	//获取当前用户信息
	m = m.Where("l."+dao.SysMessagereceive.Columns().UserId, userId)
	m = m.Where("l."+dao.SysMessagereceive.Columns().IsRead, 0)
	m = m.Where("l."+dao.SysMessagereceive.Columns().IsDeleted, 0)
	m = m.Where("l."+dao.SysMessagereceive.Columns().IsPush, 0)

	err = m.Fields("l.*").WithAll().Scan(&out)
	if err != nil {
		err = gerror.New("获取消息列表失败")
		return
	}
	//修改消息
	var ids []int
	if out != nil && len(out) > 0 {
		for _, v := range out {
			ids = append(ids, v.Id)
		}
	}
	if len(ids) > 0 {
		_, err = dao.SysMessagereceive.Ctx(ctx).Data(g.Map{
			dao.SysMessagereceive.Columns().IsPush: 1,
		}).WhereIn(dao.SysMessagereceive.Columns().Id, ids).Update()
	}

	return
}
