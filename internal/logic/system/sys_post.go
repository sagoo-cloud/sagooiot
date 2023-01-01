package system

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"github.com/sagoo-cloud/sagooiot/utility/liberr"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysPost struct {
}

func sysPostNew() *sSysPost {
	return &sSysPost{}
}

func init() {
	service.RegisterSysPost(sysPostNew())
}

// GetTree 获取全部岗位数据
func (s *sSysPost) GetTree(ctx context.Context, postName string, postCode string, status int) (data []*model.PostOut, err error) {
	postInfo, err := s.GetData(ctx, postName, postCode, status)
	if postInfo != nil {
		var parentNodeOut []*model.PostOut
		if postInfo != nil {
			//获取所有的根节点
			for _, v := range postInfo {
				var parentNode *model.PostOut
				if v.ParentId == -1 {
					if err = gconv.Scan(v, &parentNode); err != nil {
						return
					}
					parentNodeOut = append(parentNodeOut, parentNode)
				}
			}
		}
		data = postTree(parentNodeOut, postInfo)
		if len(data) == 0 {
			if err = gconv.Scan(postInfo, &data); err != nil {
				return
			}
			if err != nil {
				return
			}
		}
	}
	return
}

// Trees Tree 生成树结构
func postTree(parentNodeOut []*model.PostOut, data []*model.PostOut) (dataTree []*model.PostOut) {
	//循环所有一级菜单
	for k, v := range parentNodeOut {
		//查询所有该菜单下的所有子菜单
		for _, j := range data {
			var node *model.PostOut
			if j.ParentId == v.PostId {
				if err := gconv.Scan(j, &node); err != nil {
					return
				}
				parentNodeOut[k].Children = append(parentNodeOut[k].Children, node)
			}
		}
		postTree(v.Children, data)
	}
	return parentNodeOut
}

// Add 添加岗位
func (s *sSysPost) Add(ctx context.Context, input *model.AddPostInput) (err error) {
	var post *entity.SysPost
	//根据名称查看角色是否存在
	post = checkPostName(ctx, input.PostName, post, 0)
	if post != nil {
		return gerror.New("岗位已存在,无法添加")
	}
	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)
	post = new(entity.SysPost)
	if err := gconv.Scan(input, &post); err != nil {
		return err
	}
	post.PostCode = "G" + time.Now().Format("20060102150405")
	post.IsDeleted = 0
	post.CreatedBy = uint(loginUserId)
	_, err = dao.SysPost.Ctx(ctx).Data(post).Insert()
	if err != nil {
		return err
	}
	return
}

// Edit 修改岗位
func (s *sSysPost) Edit(ctx context.Context, input *model.EditPostInput) (err error) {
	var post, post2 *entity.SysPost
	//根据ID查看岗位是否存在
	post = checkPostId(ctx, input.PostId, post)
	if post == nil {
		return gerror.New("岗位不存在")
	}
	post2 = checkPostName(ctx, input.PostName, post2, input.PostId)
	if post2 != nil {
		return gerror.New("相同岗位已存在,无法修改")
	}
	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)
	if err := gconv.Scan(input, &post); err != nil {
		return err
	}
	post.UpdatedBy = uint(loginUserId)
	//开启事务管理
	_, err = dao.SysPost.Ctx(ctx).Data(post).
		Where(dao.SysPost.Columns().PostId, input.PostId).Update()
	if err != nil {
		return gerror.New("修改失败")
	}
	return
}

// Detail 岗位详情
func (s *sSysPost) Detail(ctx context.Context, postId int64) (entity *entity.SysPost, err error) {
	_ = dao.SysPost.Ctx(ctx).Where(g.Map{
		dao.SysPost.Columns().PostId: postId,
	}).Scan(&entity)
	if entity == nil {
		return nil, gerror.New("ID错误")
	}
	return
}

// GetData 执行获取数据操作
func (s *sSysPost) GetData(ctx context.Context, postName string, postCode string, status int) (data []*model.PostOut, err error) {
	m := dao.SysPost.Ctx(ctx)

	//模糊查询岗位名称
	if postName != "" {
		m = m.WhereLike(dao.SysPost.Columns().PostName, "%"+postName+"%")
	}
	if postCode != "" {
		m = m.WhereLike(dao.SysPost.Columns().PostCode, "%"+postCode+"%")
	}
	if status != -1 {
		m = m.Where(dao.SysPost.Columns().Status, status)
	}
	err = m.Where(dao.SysPost.Columns().IsDeleted, 0).
		OrderDesc(dao.SysPost.Columns().PostSort).
		Scan(&data)
	if err != nil {
		return
	}
	return
}

// 检查相同岗位名称的数据是否存在
func checkPostName(ctx context.Context, postName string, post *entity.SysPost, tag int64) *entity.SysPost {
	m := dao.SysPost.Ctx(ctx)
	if tag > 0 {
		m = m.WhereNot(dao.SysPost.Columns().PostId, tag)
	}
	_ = m.Where(g.Map{
		dao.SysPost.Columns().PostName:  postName,
		dao.SysPost.Columns().IsDeleted: 0,
	}).Scan(&post)
	return post
}

// Del 根据ID删除岗位信息
func (s *sSysPost) Del(ctx context.Context, postId int64) (err error) {
	var post *entity.SysPost
	_ = dao.SysPost.Ctx(ctx).Where(g.Map{
		dao.SysPost.Columns().PostId: postId,
	}).Scan(&post)
	if post == nil {
		return gerror.New("ID错误")
	}
	//查询是否有子节点
	num, err := dao.SysPost.Ctx(ctx).Where(g.Map{
		dao.SysPost.Columns().ParentId:  postId,
		dao.SysPost.Columns().IsDeleted: 0,
	}).Count()
	if err != nil {
		return err
	}
	if num > 0 {
		return gerror.New("请先删除子节点!")
	}
	loginUserId := service.Context().GetUserId(ctx)
	//更新岗位信息
	_, err = dao.SysPost.Ctx(ctx).
		Data(g.Map{
			dao.SysPost.Columns().DeletedBy: uint(loginUserId),
			dao.SysPost.Columns().IsDeleted: 1,
		}).Where(dao.SysPost.Columns().PostId, postId).
		Update()
	//删除岗位信息
	_, err = dao.SysPost.Ctx(ctx).Where(dao.SysPost.Columns().PostId, postId).
		Delete()
	return
}

// 检查指定ID的数据是否存在
func checkPostId(ctx context.Context, PostId int64, post *entity.SysPost) *entity.SysPost {
	_ = dao.SysPost.Ctx(ctx).Where(g.Map{
		dao.SysPost.Columns().PostId:    PostId,
		dao.SysPost.Columns().IsDeleted: 0,
	}).Scan(&post)
	return post
}

// GetUsedPost 获取正常状态的岗位
func (s *sSysPost) GetUsedPost(ctx context.Context) (list []*model.DetailPostRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysPost.Ctx(ctx).Where(dao.SysPost.Columns().Status, 1).
			Order(dao.SysPost.Columns().PostSort + " ASC, " + dao.SysPost.Columns().PostId + " ASC ").Scan(&list)
		liberr.ErrIsNil(ctx, err, "获取岗位数据失败")
	})
	return
}
