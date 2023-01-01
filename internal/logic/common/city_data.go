package common

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
	"github.com/sagoo-cloud/sagooiot/internal/service"
)

type sCityData struct {
}

func CityData() *sCityData {
	return &sCityData{}
}

func init() {
	service.RegisterCityData(CityData())
}

// GetList 获取城市列表
func (s *sCityData) GetList(ctx context.Context, status int, name string, code string) (data []*entity.CityData, err error) {
	m := dao.CityData.Ctx(ctx)
	if status != -1 {
		m = m.Where(dao.CityData.Columns().Status, status)
	}
	if name != "" {
		m = m.Where(dao.CityData.Columns().Name, name)
	}
	if code != "" {
		m = m.Where(dao.CityData.Columns().Code, code)
	}
	m = m.Where(dao.CityData.Columns().IsDeleted, 0).OrderAsc(dao.CityData.Columns().Sort)

	//获取城市列表信息
	err = m.Scan(&data)
	if err != nil {
		err = gerror.New("获取城市列表失败")
		return
	}
	return
}

// Add 添加城市
func (s *sCityData) Add(ctx context.Context, city *entity.CityData) (err error) {
	err = dao.CityData.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		//根据名字查询城市是否存在
		num, _ := dao.CityData.Ctx(ctx).Where(g.Map{
			dao.CityData.Columns().Name:      city.Name,
			dao.CityData.Columns().IsDeleted: 0,
		}).Count()
		if num > 0 {
			return gerror.New("城市已存在")
		}
		//根据Code查询城市是否存在
		codeNum, _ := dao.CityData.Ctx(ctx).Where(g.Map{
			dao.CityData.Columns().Code:      city.Code,
			dao.CityData.Columns().IsDeleted: 0,
		}).Count()
		if codeNum > 0 {
			return gerror.New("编码已存在")
		}
		//获取当前登录用户ID
		loginUserId := service.Context().GetUserId(ctx)
		city.CreatedBy = uint(loginUserId)
		city.IsDeleted = 0
		_, addErr := dao.CityData.Ctx(ctx).Data(city).Insert()
		if addErr != nil {
			err = gerror.New("添加失败")
			return err
		}
		return err
	})

	return
}

// Edit 编辑城市
func (s *sCityData) Edit(ctx context.Context, city *entity.CityData) (err error) {
	//根据ID查询城市是否存在
	var cityInfo *entity.CityData
	err = dao.CityData.Ctx(ctx).Where(g.Map{
		dao.CityData.Columns().Id:        city.Id,
		dao.CityData.Columns().IsDeleted: 0,
	}).Scan(&cityInfo)
	if cityInfo == nil {
		return gerror.New("ID错误")
	}
	var cityInfoByName *entity.CityData
	err = dao.CityData.Ctx(ctx).Where(g.Map{
		dao.CityData.Columns().Name:      city.Name,
		dao.CityData.Columns().IsDeleted: 0,
	}).Scan(&cityInfoByName)
	if cityInfoByName != nil && cityInfoByName.Id != city.Id {
		return gerror.New("编码已存在")
	}
	var cityInfoByCode *entity.CityData
	err = dao.CityData.Ctx(ctx).Where(g.Map{
		dao.CityData.Columns().Code:      city.Code,
		dao.CityData.Columns().IsDeleted: 0,
	}).Scan(&cityInfoByCode)
	if cityInfoByCode != nil && cityInfoByCode.Id != city.Id {
		return gerror.New("编码已存在")
	}

	err = dao.CityData.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		//获取当前登录用户ID
		loginUserId := service.Context().GetUserId(ctx)
		cityInfo.UpdatedBy = uint(loginUserId)
		cityInfo.Status = city.Status
		cityInfo.Sort = city.Sort
		cityInfo.ParentId = city.ParentId
		cityInfo.Name = city.Name
		cityInfo.Code = city.Code
		_, err = dao.CityData.Ctx(ctx).Data(cityInfo).Where(dao.CityData.Columns().Id, city.Id).Update()
		if err != nil {
			return gerror.New("修改失败")
		}
		return err
	})

	return
}

// GetInfoById 根据ID获取城市
func (s *sCityData) GetInfoById(ctx context.Context, id int) (cityInfo *entity.CityData, err error) {
	err = dao.CityData.Ctx(ctx).Where(g.Map{
		dao.CityData.Columns().Id: id,
	}).Scan(&cityInfo)
	return
}

// DelById 删除城市
func (s *sCityData) DelById(ctx context.Context, id int) (err error) {
	//根据ID查询城市是否存在
	var cityInfo *entity.CityData
	err = dao.CityData.Ctx(ctx).Where(g.Map{
		dao.CityData.Columns().Id:        id,
		dao.CityData.Columns().IsDeleted: 0,
	}).Scan(&cityInfo)
	if cityInfo == nil {
		return gerror.New("ID错误")
	}
	//判断是否存在子节点
	childrenNum, err := dao.CityData.Ctx(ctx).Where(g.Map{
		dao.CityData.Columns().ParentId:  id,
		dao.CityData.Columns().IsDeleted: 0,
	}).Count()
	if childrenNum > 0 {
		return gerror.New("请先删除子节点")
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)
	cityInfo.DeletedBy = loginUserId
	cityInfo.IsDeleted = 1
	//获取当前时间
	t, err := gtime.StrToTimeFormat(gtime.Datetime(), "2006-01-02 15:04:05")
	cityInfo.DeletedAt = t
	_, err = dao.CityData.Ctx(ctx).Data(cityInfo).Where(dao.CityData.Columns().Id, id).Update()
	if err != nil {
		return gerror.New("删除失败")
	}
	return
}

// GetAll 获取所有城市
func (s *sCityData) GetAll(ctx context.Context) (data []*entity.CityData, err error) {
	err = dao.CityData.Ctx(ctx).Where(g.Map{
		dao.CityData.Columns().Status:    1,
		dao.CityData.Columns().IsDeleted: 0,
	}).OrderAsc(dao.CityData.Columns().Sort).Scan(&data)
	return
}
