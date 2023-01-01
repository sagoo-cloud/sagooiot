package system

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/logic/common"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
	"github.com/sagoo-cloud/sagooiot/internal/service"
)

type sSysUserOnline struct {
}

func init() {
	service.RegisterSysUserOnline(sysUserOnlineNew())
}

func sysUserOnlineNew() *sSysUserOnline {
	return &sSysUserOnline{}
}

func (s *sSysUserOnline) Invoke(ctx context.Context, data *entity.SysUserOnline) {
	pool := grpool.New(100)
	if err := pool.Add(ctx, func(ctx context.Context) {
		//写入用户在线
		s.Add(ctx, data)
	},
	); err != nil {
		g.Log().Debug(ctx, err.Error())
	}
}

// Add 记录用户在线
func (s *sSysUserOnline) Add(ctx context.Context, data *entity.SysUserOnline) {
	_, err := dao.SysUserOnline.Ctx(ctx).Data(data).Save()
	if err != nil {
		g.Log().Error(ctx, err)
	}
}

// DelByToken 根据token删除信息
func (s *sSysUserOnline) DelByToken(ctx context.Context, token string) (err error) {
	_, err = dao.SysUserOnline.Ctx(ctx).Where(dao.SysUserOnline.Columns().Token, token).Delete()
	if err != nil {
		return gerror.New("退出失败")
	}
	return
}

// GetInfoByToken 根据token获取
func (s *sSysUserOnline) GetInfoByToken(ctx context.Context, token string) (data *entity.SysUserOnline, err error) {
	err = dao.SysUserOnline.Ctx(ctx).Where(dao.SysUserOnline.Columns().Token, token).Scan(&data)
	return
}

// DelByIds 根据IDS删除信息
func (s *sSysUserOnline) DelByIds(ctx context.Context, ids []uint) (err error) {
	_, err = dao.SysUserOnline.Ctx(ctx).WhereIn(dao.SysUserOnline.Columns().Id, ids).Delete()
	if err != nil {
		return gerror.New("删除失败")
	}
	return
}

func (s *sSysUserOnline) GetAll(ctx context.Context) (data []*entity.SysUserOnline, err error) {
	err = dao.SysUserOnline.Ctx(ctx).Scan(&data)
	return
}

// UserOnlineList 在线用户列表
func (s *sSysUserOnline) UserOnlineList(ctx context.Context, input *model.UserOnlineDoListInput) (total int, out []*model.UserOnlineListOut, err error) {
	m := dao.SysUserOnline.Ctx(ctx)
	//获取总数
	total, err = m.Count()
	if err != nil {
		err = gerror.New("获取数据失败")
		return
	}
	if input.PageNum == 0 {
		input.PageNum = 1
	}
	if input.PageSize == 0 {
		input.PageSize = consts.DefaultPageSize
	}
	//获取在线用户信息
	err = m.Page(input.PageNum, input.PageSize).OrderDesc(dao.SysUser.Columns().CreatedAt).Scan(&out)
	if err != nil {
		err = gerror.New("获取在线用户列表失败")
		return
	}
	return
}

func (s *sSysUserOnline) UserOnlineStrongBack(ctx context.Context, id int) (err error) {
	var userOnline *entity.SysUserOnline
	err = dao.SysUserOnline.Ctx(ctx).Where(dao.SysUserOnline.Columns().Id, id).Scan(&userOnline)
	if userOnline == nil {
		return gerror.New("ID错误")
	}
	//删除缓存信息
	common.Cache().Remove(ctx, userOnline.Key)
	//删除在线用户
	_, err = dao.SysUserOnline.Ctx(ctx).Where(dao.SysUserOnline.Columns().Id, id).Delete()
	return
}
