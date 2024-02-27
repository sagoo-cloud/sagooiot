package system

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/os/gtime"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/do"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"
	"sort"
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
					var isExist = false
					for _, postOut := range parentNodeOut {
						if postOut.PostId == parentNode.PostId {
							isExist = true
							break
						}
					}
					if !isExist {
						parentNodeOut = append(parentNodeOut, parentNode)
					}
				} else {
					//查找根节点
					parentPost := FindPostParentByChildrenId(ctx, int(v.ParentId))
					if err = gconv.Scan(parentPost, &parentNode); err != nil {
						return
					}
					var isExist = false
					for _, postOut := range parentNodeOut {
						if postOut.PostId == int64(parentPost.PostId) {
							isExist = true
							break
						}
					}
					if !isExist {
						parentNodeOut = append(parentNodeOut, parentNode)
					}
				}
			}
		}
		//对父节点进行排序
		sort.SliceStable(parentNodeOut, func(i, j int) bool {
			return parentNodeOut[i].PostSort < parentNodeOut[j].PostSort
		})
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
		//对子节点进行排序
		sort.SliceStable(v.Children, func(i, j int) bool {
			return v.Children[i].PostSort < v.Children[j].PostSort
		})
		postTree(v.Children, data)
	}
	return parentNodeOut
}

// FindPostParentByChildrenId 根据子节点获取岗位根节点
func FindPostParentByChildrenId(ctx context.Context, parentId int) *entity.SysPost {
	var post *entity.SysPost

	_ = dao.SysPost.Ctx(ctx).Where(g.Map{
		dao.SysPost.Columns().PostId: parentId,
	}).Scan(&post)

	if post.ParentId != -1 {
		return FindPostParentByChildrenId(ctx, post.ParentId)
	}
	return post
}

// Add 添加岗位
func (s *sSysPost) Add(ctx context.Context, input *model.AddPostInput) (err error) {
	var post *entity.SysPost
	//根据名称查看角色是否存在
	post = checkPostName(ctx, input.PostName, post, 0)
	if post != nil {
		return gerror.New("岗位已存在,无法添加")
	}
	//获取上级岗位信息
	if input.ParentId != -1 {
		var parentPost *entity.SysPost
		parentPost, err = s.Detail(ctx, input.ParentId)
		if err != nil {
			return
		}
		if parentPost == nil {
			err = gerror.Newf("无权限选择当前岗位")
			return
		}
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)
	post = new(entity.SysPost)
	if err = gconv.Scan(input, &post); err != nil {
		return err
	}
	post.PostCode = "G" + time.Now().Format("20060102150405")
	post.IsDeleted = 0
	post.CreatedBy = uint(loginUserId)
	_, err = dao.SysPost.Ctx(ctx).Data(do.SysPost{
		DeptId:    service.Context().GetUserDeptId(ctx),
		ParentId:  post.ParentId,
		PostCode:  post.PostCode,
		PostName:  post.PostName,
		PostSort:  post.PostSort,
		Status:    post.Status,
		Remark:    post.Remark,
		IsDeleted: 0,
		CreatedBy: post.CreatedBy,
		CreatedAt: gtime.Now(),
	}).Insert()
	if err != nil {
		return err
	}
	return
}

// Edit 修改岗位
func (s *sSysPost) Edit(ctx context.Context, input *model.EditPostInput) (err error) {
	if input.PostId == input.ParentId {
		return gerror.New("父级不能为自己")
	}
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

	//判断上级岗位是否可以选择
	if input.ParentId != -1 {
		var parentPost *entity.SysPost
		parentPost, err = s.Detail(ctx, input.ParentId)
		if err != nil {
			return
		}
		if parentPost == nil {
			err = gerror.Newf("无权限选择岗位")
			return
		}
	}

	if err = gconv.Scan(input, &post); err != nil {
		return err
	}
	post.UpdatedBy = uint(service.Context().GetUserId(ctx))
	post.UpdatedAt = gtime.Now()
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
	m := dao.SysPost.Ctx(ctx)

	_ = m.Where(g.Map{
		dao.SysPost.Columns().PostId: postId,
	}).Scan(&entity)
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
		OrderAsc(dao.SysPost.Columns().PostSort).
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
			dao.SysPost.Columns().DeletedAt: gtime.Now(),
			dao.SysPost.Columns().IsDeleted: 1,
		}).Where(dao.SysPost.Columns().PostId, postId).
		Update()
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
func (s *sSysPost) GetUsedPost(ctx context.Context) (list []*model.DetailPostOut, err error) {
	err = dao.SysPost.Ctx(ctx).Where(dao.SysPost.Columns().Status, 1).
		Order(dao.SysPost.Columns().PostSort + " ASC, " + dao.SysPost.Columns().PostId + " ASC ").Scan(&list)
	if err != nil {
		return nil, errors.New("获取岗位失败")
	}

	return
}
