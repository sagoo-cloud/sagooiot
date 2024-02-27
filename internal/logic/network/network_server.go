package network

import (
	"context"
	"sagooiot/internal/consts"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/do"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"
	"sagooiot/network/core/server"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type sNetworkServer struct{}

func sNetworkServerNew() *sNetworkServer {
	return &sNetworkServer{}
}
func init() {
	service.RegisterNetworkServer(sNetworkServerNew())
}

const (
	ServerStatusOnline  = "online"
	ServerStatusOffline = "offline"
)

// GetServerList 获取列表数据
func (s *sNetworkServer) GetServerList(ctx context.Context, in *model.GetNetworkServerListInput) (total int, out []*model.NetworkServerOut, err error) {
	//g.Log().Debug(ctx, in)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.NetworkServer.Ctx(ctx)

		if in.KeyWord != "" {
			m = m.WhereLike(dao.NetworkServer.Columns().Name, "%"+in.KeyWord+"%")
		}

		total, err = m.Count()
		if err != nil {
			err = gerror.New("获取总行数失败")
			return
		}
		if in.PageNum == 0 {
			in.PageNum = 1
		}
		if in.PageSize == 0 {
			in.PageSize = consts.PageSize
		}
		err = m.Page(in.PageNum, in.PageSize).Order("created_at desc").Scan(&out)
		if err != nil {
			err = gerror.New("获取数据失败")
		}
	})

	return
}

// GetServerRunList 获取可运行的服务列表数据
func (s *sNetworkServer) GetServerRunList(ctx context.Context) (list []*model.NetworkServerRes, err error) {

	m := dao.NetworkServer.Ctx(ctx)
	err = m.Where(dao.NetworkServer.Columns().Status, 1).Scan(&list)
	if err != nil {
		err = gerror.New("获取数据失败")
	}

	return
}

// GetServerById 获取指定ID数据
func (s *sNetworkServer) GetServerById(ctx context.Context, id int) (out *model.NetworkServerOut, err error) {
	err = dao.NetworkServer.Ctx(ctx).Where("id", id).Scan(&out)
	return
}

// AddServer 添加数据 todo 需要处理
func (s *sNetworkServer) AddServer(ctx context.Context, in model.NetworkServerAddInput) (err error) {
	//查询服务器名称是否存在
	num, err := dao.NetworkServer.Ctx(ctx).Where(g.Map{dao.NetworkServer.Columns().Name: in.Name}).Count()
	if err != nil {
		return
	}
	if num > 0 {
		err = gerror.New("服务器名称已存在")
		return
	}

	insertResult, insertResultErr := dao.NetworkServer.Ctx(ctx).Data(do.NetworkServer{
		DeptId:        service.Context().GetUserDeptId(ctx),
		Name:          in.Name,
		Types:         in.Types,
		Addr:          in.Addr,
		Register:      in.Register,
		Heartbeat:     in.Heartbeat,
		Protocol:      in.Protocol,
		Devices:       in.Devices,
		Status:        in.Status,
		CreatedAt:     gtime.Now(),
		CreateBy:      in.CreateBy,
		Remark:        in.Remark,
		IsTls:         in.IsTls,
		AuthType:      in.AuthType,
		AuthUser:      in.AuthUser,
		AuthPasswd:    in.AuthPasswd,
		AccessToken:   in.AccessToken,
		CertificateId: in.CertificateId,
		Stick:         in.Stick,
	}).Insert()
	if insertResultErr != nil {
		return insertResultErr
	}
	lastId, lastIdErr := insertResult.LastInsertId()
	if lastIdErr != nil {
		return lastIdErr
	}
	if err == nil && in.Status == consts.ServerStatusOnline {
		return server.LoadServer(ctx, int(lastId))
	}
	return
}

// EditServer 修改数据 todo 需要处理
func (s *sNetworkServer) EditServer(ctx context.Context, in model.NetworkServerEditInput) (err error) {
	//根据ID获取服务数据
	var netWorkServer *entity.NetworkServer
	err = dao.NetworkServer.Ctx(ctx).Where(dao.NetworkServer.Columns().Id, in.Id).Scan(&netWorkServer)
	if err != nil {
		return err
	}
	if netWorkServer == nil {
		return gerror.New("ID错误")
	}

	//查询服务器名称是否存在
	num, err := dao.NetworkServer.Ctx(ctx).Where(g.Map{dao.NetworkServer.Columns().Name: in.Name}).WhereNot(dao.NetworkServer.Columns().Id, in.Id).Count()
	if err != nil {
		return
	}
	if num > 0 {
		err = gerror.New("服务器名称已存在")
		return
	}

	_, err = dao.NetworkServer.Ctx(ctx).FieldsEx(dao.NetworkServer.Columns().Id, dao.NetworkServer.Columns().CreateBy).Where(dao.NetworkServer.Columns().Id, in.Id).Update(in)
	if err != nil {
		return err
	}
	if err = server.RemoveServer(in.Id); err != nil {
		return err
	}
	if in.Status == 1 {
		if err = server.LoadServer(ctx, in.Id); err != nil {
			return err
		}
	}
	return
}

// 删除数据
// todo 需要处理
func (s *sNetworkServer) DeleteServer(ctx context.Context, ids []int) (err error) {

	for _, id := range ids {
		var netWorkServer *entity.NetworkServer
		err = dao.NetworkServer.Ctx(ctx).Where(dao.NetworkServer.Columns().Id, id).Scan(&netWorkServer)
		if err != nil {
			return err
		}
		if netWorkServer == nil {
			return gerror.New("ID错误")
		}
	}
	_, err = dao.NetworkServer.Ctx(ctx).Delete(dao.NetworkServer.Columns().Id+" in (?)", ids)
	if err == nil {
		for _, node := range ids {
			if err = server.RemoveServer(node); err != nil {
				return err
			}
		}
	}
	return
}

// SetServerStatus 修改状态数据 todo 需要处理
func (s *sNetworkServer) SetServerStatus(ctx context.Context, id, status int) (err error) {
	//根据ID获取服务数据
	var netWorkServer *entity.NetworkServer
	err = dao.NetworkServer.Ctx(ctx).Where(dao.NetworkServer.Columns().Id, id).Scan(&netWorkServer)
	if err != nil {
		return err
	}
	if netWorkServer == nil {
		return gerror.New("ID错误")
	}

	var data = g.Map{
		dao.NetworkServer.Columns().Status: status,
	}
	_, err = dao.NetworkServer.Ctx(ctx).Where(dao.NetworkServer.Columns().Id, id).Update(data)
	if err != nil {
		return err
	}
	if status == 1 {
		if err = server.LoadServer(ctx, id); err != nil {
			return err
		}
	} else {
		if err = server.RemoveServer(id); err != nil {
			return err
		}
	}
	return
}
