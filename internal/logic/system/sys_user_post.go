package system

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/internal/dao"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"
)

type sSysUserPost struct {
}

func init() {
	service.RegisterSysUserPost(sysUserPostNew())
}

func sysUserPostNew() *sSysUserPost {
	return &sSysUserPost{}
}

// GetInfoByUserId 根据用户ID获取信息
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

// BindUserAndPost 添加用户与岗位绑定关系
func (s *sSysUserPost) BindUserAndPost(ctx context.Context, userId int, postIds []int) (err error) {
	if len(postIds) > 0 {
		//删除原有用户与岗位绑定管理
		_, err = dao.SysUserPost.Ctx(ctx).Where(dao.SysUserPost.Columns().UserId, userId).Delete()
		if err != nil {
			return gerror.New("删除用户与岗位绑定关系失败")
		}

		var sysUserPosts []*entity.SysUserPost
		//查询用户与岗位是否存在
		for _, postId := range postIds {
			/*var sysUserPost *entity.SysUserPost
			err = dao.SysUserPost.Ctx(ctx).Where(g.Map{
				dao.SysUserPost.Columns().UserId: userId,
				dao.SysUserPost.Columns().PostId: postId,
			}).Scan(&sysUserPost)
			if sysUserPost == nil {

			}*/
			//添加用户与岗位绑定管理
			var sysUserPost = new(entity.SysUserPost)
			sysUserPost.UserId = userId
			sysUserPost.PostId = postId
			sysUserPosts = append(sysUserPosts, sysUserPost)
		}
		_, err = dao.SysUserPost.Ctx(ctx).Data(sysUserPosts).Insert()
		if err != nil {
			return gerror.New("绑定岗位失败")
		}
	}
	return
}
