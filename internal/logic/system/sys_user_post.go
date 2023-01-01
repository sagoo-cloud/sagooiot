package system

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
	"github.com/sagoo-cloud/sagooiot/internal/service"
)

type sSysUserPost struct {
}

func init() {
	service.RegisterSysUserPost(sysUserPostNew())
}

func sysUserPostNew() *sSysUserPost {
	return &sSysUserPost{}
}

//GetInfoByUserId 根据用户ID获取信息
func (s *sSysUserPost) GetInfoByUserId(ctx context.Context, userId int) (data []*entity.SysUserPost, err error) {
	var userPost []*entity.SysUserPost
	err = dao.SysUserPost.Ctx(ctx).Where(dao.SysUserPost.Columns().UserId, userId).Scan(&userPost)
	if userPost != nil {
		for _, post := range userPost {
			//判断岗位是否为已启动并未删除状态
			num, _ := dao.SysPost.Ctx(ctx).Where(g.Map{
				dao.SysPost.Columns().PostId:    post.PostId,
				dao.SysPost.Columns().Status:    1,
				dao.SysPost.Columns().IsDeleted: 0,
			}).Count()

			if num > 0 {
				data = append(data, post)
			}
		}
	}
	return
}
