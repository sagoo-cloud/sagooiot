package system

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gtime"
	"sagooiot/internal/consts"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/do"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"
	"sagooiot/pkg/cache"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysApi struct {
}

func sysApiNew() *sSysApi {
	return &sSysApi{}
}

func init() {
	service.RegisterSysApi(sysApiNew())
}

// GetInfoByIds 根据接口APIID数组获取接口信息
func (s *sSysApi) GetInfoByIds(ctx context.Context, ids []int) (data []*entity.SysApi, err error) {
	//获取缓存信息
	var tmpData *gvar.Var
	tmpData, err = cache.Instance().Get(ctx, consts.CacheSysApi)

	var tmpSysApiInfo []*entity.SysApi

	var apiInfo []*entity.SysApi
	//根据菜单ID数组获取菜单列表信息
	if tmpData.Val() != nil {
		if err = json.Unmarshal([]byte(tmpData.Val().(string)), &tmpSysApiInfo); err != nil {
			return
		}

		for _, id := range ids {
			for _, menuTmp := range tmpSysApiInfo {
				if id == int(menuTmp.Id) {
					apiInfo = append(apiInfo, menuTmp)
					continue
				}
			}
		}
	}
	if apiInfo != nil && len(apiInfo) >= 0 {
		data = apiInfo
		return
	}
	err = dao.SysApi.Ctx(ctx).Where(g.Map{
		dao.SysApi.Columns().IsDeleted: 0,
		dao.SysApi.Columns().Status:    1,
	}).WhereIn(dao.SysApi.Columns().Id, ids).Scan(&data)
	return
}

// GetApiByMenuId 根据ApiID获取接口信息
func (s *sSysApi) GetApiByMenuId(ctx context.Context, apiId int) (data []*entity.SysApi, err error) {
	var apiApi []*entity.SysMenuApi
	err = dao.SysMenuApi.Ctx(ctx).Where(g.Map{
		dao.SysMenuApi.Columns().MenuId:    apiId,
		dao.SysMenuApi.Columns().IsDeleted: 0,
	}).Scan(&apiApi)
	//获取接口ID数组
	if apiApi != nil {
		var ids []int
		for _, api := range apiApi {
			ids = append(ids, api.ApiId)
		}
		err = dao.SysApi.Ctx(ctx).Where(g.Map{
			dao.SysApi.Columns().IsDeleted: 0,
			dao.SysApi.Columns().Status:    1,
		}).WhereIn(dao.SysApi.Columns().Id, ids).Scan(&data)
	}
	return
}

// GetInfoById 根据ID获取API
func (s *sSysApi) GetInfoById(ctx context.Context, id int) (entity *entity.SysApi, err error) {
	err = dao.SysApi.Ctx(ctx).Where(g.Map{
		dao.SysApi.Columns().Id: id,
	}).Scan(&entity)
	return
}

// GetApiAll 获取所有接口
func (s *sSysApi) GetApiAll(ctx context.Context, method string) (data []*entity.SysApi, err error) {
	m := dao.SysApi.Ctx(ctx)
	if method != "" {
		m = m.Where(dao.SysApi.Columns().Method, method)
	}
	err = m.Where(g.Map{
		dao.SysApi.Columns().IsDeleted: 0,
		dao.SysApi.Columns().Status:    1,
		dao.SysApi.Columns().Types:     2,
	}).Scan(&data)
	if method == "" {
		if data != nil && len(data) > 0 {
			err = cache.Instance().Set(ctx, consts.CacheSysApi, data, 0)
		} else {
			_, err = cache.Instance().Remove(ctx, consts.CacheSysApi)
		}
	}
	return
}

// GetApiTree 获取Api数结构数据
func (s *sSysApi) GetApiTree(ctx context.Context, name string, address string, status int, types int) (out []*model.SysApiTreeOut, err error) {
	var es []*model.SysApiTreeOut
	m := dao.SysApi.Ctx(ctx)
	if name != "" {
		m = m.WhereLike(dao.SysApi.Columns().Name, "%"+name+"%")
	}
	if address != "" {
		m = m.WhereLike(dao.SysApi.Columns().Address, "%"+address+"%")
	}
	if status != -1 {
		m = m.Where(dao.SysApi.Columns().Status, status)
	}
	if types != -1 {
		m = m.Where(dao.SysApi.Columns().Types, types)
	}
	m = m.Where(dao.SysApi.Columns().IsDeleted, 0)

	err = m.OrderAsc(dao.SysApi.Columns().Sort).Scan(&es)
	for _, e := range es {
		menuApiInfo, _ := service.SysMenuApi().GetInfoByApiId(ctx, int(e.Id))
		var menuIds []int
		for _, menuApi := range menuApiInfo {
			menuIds = append(menuIds, menuApi.MenuId)
		}
		e.MenuIds = append(e.MenuIds, menuIds...)
	}

	if len(es) > 0 {
		out, err = GetApiTree(es)
		if err != nil {
			return
		}
	}
	return
}

// Add 添加Api列表
func (s *sSysApi) Add(ctx context.Context, input *model.AddApiInput) (err error) {
	if input.Types == 2 {
		if input.Address == "" || len(input.MenuIds) == 0 {
			err = gerror.New("参数错误")
			return
		}
		if input.ParentId == -1 {
			err = gerror.New("接口不能为根节点")
			return
		}
	} else {
		if input.Address != "" || len(input.MenuIds) > 0 {
			err = gerror.New("参数错误")
			return
		}
	}
	if input.ParentId != -1 {
		var parentApiInfo *entity.SysApi
		err = dao.SysApi.Ctx(ctx).Where(g.Map{
			dao.SysApi.Columns().Id: input.ParentId,
		}).Scan(&parentApiInfo)
		if parentApiInfo.IsDeleted != 0 {
			return gerror.New("上级节点已删除，无法新增")
		}
		if parentApiInfo.Status != 1 {
			return gerror.New("上级节点未启用，无法新增")
		}
		if parentApiInfo.Types != 1 {
			return gerror.New("上级节点不是分类，无法新增")
		}
	}
	err = dao.SysUser.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		var apiInfo *entity.SysApi
		//根据名称查看是否存在
		apiInfo = CheckApiName(ctx, input.Name, 0, input.ParentId)
		if apiInfo != nil {
			return gerror.New("同一个分类下名称不能重复")
		}
		if input.Types == 2 {
			//根据名称查看是否存在
			apiInfo = CheckApiAddress(ctx, input.Address, 0, input.ApiTypes)
			if apiInfo != nil {
				return gerror.New("同一个服务下Api地址,无法添加")
			}
		}
		//获取当前登录用户ID
		loginUserId := service.Context().GetUserId(ctx)
		apiInfo = new(entity.SysApi)
		if apiInfoErr := gconv.Scan(input, &apiInfo); apiInfoErr != nil {
			return
		}
		apiInfo.IsDeleted = 0
		apiInfo.CreatedBy = uint(loginUserId)
		result, err := dao.SysApi.Ctx(ctx).Data(do.SysApi{
			ParentId:  apiInfo.ParentId,
			Name:      apiInfo.Name,
			Types:     apiInfo.Types,
			ApiTypes:  apiInfo.ApiTypes,
			Method:    apiInfo.Method,
			Address:   apiInfo.Address,
			Remark:    apiInfo.Remark,
			Status:    apiInfo.Status,
			Sort:      apiInfo.Sort,
			IsDeleted: apiInfo.IsDeleted,
			CreatedBy: apiInfo.CreatedBy,
			CreatedAt: gtime.Now(),
		}).Insert()
		if err != nil {
			return err
		}

		if input.Types == 2 {
			var lastInsertId int64
			//获取主键ID
			lastInsertId, err = service.Sequences().GetSequences(ctx, result, dao.SysApi.Table(), dao.SysApi.Columns().Id)
			if err != nil {
				return
			}
			//绑定菜单
			err = s.AddMenuApi(ctx, "api", []int{int(lastInsertId)}, input.MenuIds)
			if err != nil {
				return
			}
		}
		//获取所有接口并添加缓存
		_, err = s.GetApiAll(ctx, "")
		return
	})
	return
}

// Detail Api列表详情
func (s *sSysApi) Detail(ctx context.Context, id int) (out *model.SysApiOut, err error) {
	var e *entity.SysApi
	_ = dao.SysApi.Ctx(ctx).Where(g.Map{
		dao.SysApi.Columns().Id: id,
	}).Scan(&e)
	if e != nil {
		if err = gconv.Scan(e, &out); err != nil {
			return nil, err
		}
		menuApiInfo, _ := service.SysMenuApi().GetInfoByApiId(ctx, out.Id)
		var menuIds []int
		for _, menuApi := range menuApiInfo {
			menuIds = append(menuIds, menuApi.MenuId)
		}
		out.MenuIds = append(out.MenuIds, menuIds...)
	}
	return
}

func (s *sSysApi) AddMenuApi(ctx context.Context, addPageSource string, apiIds []int, menuIds []int) (err error) {
	loginUserId := service.Context().GetUserId(ctx)

	if addPageSource != "" && strings.EqualFold(addPageSource, "api") {
		//解除旧绑定关系
		_, err = dao.SysMenuApi.Ctx(ctx).Data(g.Map{
			dao.SysMenuApi.Columns().IsDeleted: 1,
			dao.SysMenuApi.Columns().DeletedBy: loginUserId,
			dao.SysMenuApi.Columns().DeletedAt: gtime.Now(),
		}).WhereIn(dao.SysMenuApi.Columns().ApiId, apiIds).Update()
	}
	if addPageSource != "" && strings.EqualFold(addPageSource, "menu") {
		//解除旧绑定关系
		_, err = dao.SysMenuApi.Ctx(ctx).Data(g.Map{
			dao.SysMenuApi.Columns().IsDeleted: 1,
			dao.SysMenuApi.Columns().DeletedBy: loginUserId,
			dao.SysMenuApi.Columns().DeletedAt: gtime.Now(),
		}).WhereIn(dao.SysMenuApi.Columns().MenuId, menuIds).Update()
	}

	if menuIds != nil && len(menuIds) > 0 && apiIds != nil && len(apiIds) > 0 {
		//添加菜单
		var sysMenuApis []*entity.SysMenuApi
		for _, menuId := range menuIds {
			var menuInfo *entity.SysMenu
			err = dao.SysMenu.Ctx(ctx).Where(dao.SysMenu.Columns().Id, menuId).Scan(&menuInfo)
			if menuInfo == nil {
				err = gerror.New("菜单ID错误")
				return
			}
			if menuInfo != nil && menuInfo.IsDeleted == 1 {
				err = gerror.New(menuInfo.Name + "已删除,无法绑定")
				return
			}
			if menuInfo != nil && menuInfo.Status == 0 {
				err = gerror.New(menuInfo.Name + "已禁用,无法绑定")
				return
			}

			for _, id := range apiIds {
				apiInfo, _ := s.Detail(ctx, id)
				if apiInfo == nil {
					return gerror.New("API接口不存在")
				}
				if apiInfo.Types != 2 {
					return gerror.New("参数错误")
				}
				var sysMenuApi = new(entity.SysMenuApi)
				sysMenuApi.MenuId = menuId
				sysMenuApi.ApiId = id
				sysMenuApi.IsDeleted = 0
				sysMenuApi.CreatedBy = uint(loginUserId)
				sysMenuApis = append(sysMenuApis, sysMenuApi)
			}
		}
		if sysMenuApis != nil {
			//添加
			var sysMenuApisInput []*model.SysMenuApiInput
			if err = gconv.Scan(sysMenuApis, &sysMenuApisInput); err != nil {
				return
			}

			_, addErr := dao.SysMenuApi.Ctx(ctx).Data(sysMenuApisInput).Insert()
			if addErr != nil {
				err = gerror.New("添加失败")
				return
			}
			//查询菜单ID绑定的所有接口ID
			var menuApiInfos []*entity.SysMenuApi
			err = dao.SysMenuApi.Ctx(ctx).Where(g.Map{
				dao.SysMenuApi.Columns().IsDeleted: 0,
			}).WhereIn(dao.SysMenuApi.Columns().MenuId, menuIds).Scan(&menuApiInfos)
			//添加缓存
			for _, menuId := range menuIds {
				var menuApi []*entity.SysMenuApi
				for _, menuApiInfo := range menuApiInfos {
					if menuId == menuApiInfo.MenuId {
						menuApi = append(menuApi, menuApiInfo)
					}
				}
				if menuApi != nil && len(menuApi) > 0 {
					err := cache.Instance().Set(ctx, consts.CacheSysMenuApi+"_"+gconv.String(menuId), menuApi, 0)
					if err != nil {
						return err
					}
				}

			}
			//获取所有信息
			var menuApiInfoAll []*entity.SysMenuApi
			err = dao.SysMenuApi.Ctx(ctx).Where(g.Map{
				dao.SysMenuApi.Columns().IsDeleted: 0,
			}).WhereIn(dao.SysMenuApi.Columns().MenuId, menuIds).Scan(&menuApiInfoAll)
			if menuApiInfoAll != nil && len(menuApiInfoAll) > 0 {
				err := cache.Instance().Set(ctx, consts.CacheSysMenuApi, menuApiInfoAll, 0)
				if err != nil {
					return err
				}
			}
		}
	}

	return
}

// Edit 修改Api列表
func (s *sSysApi) Edit(ctx context.Context, input *model.EditApiInput) (err error) {
	if input.Types == 2 {
		if input.Address == "" || len(input.MenuIds) == 0 {
			err = gerror.New("参数错误")
			return
		}
		if input.ParentId == -1 {
			err = gerror.New("接口不能为根节点")
			return
		}
	} else {
		if input.Address != "" || len(input.MenuIds) > 0 {
			err = gerror.New("参数错误")
			return
		}
	}
	if input.ParentId != -1 {
		var parentApiInfo *entity.SysApi
		err = dao.SysApi.Ctx(ctx).Where(g.Map{
			dao.SysApi.Columns().Id: input.ParentId,
		}).Scan(&parentApiInfo)
		if parentApiInfo.IsDeleted != 0 {
			return gerror.New("上级节点已删除，无法新增")
		}
		if parentApiInfo.Status != 1 {
			return gerror.New("上级节点已启用，无法新增")
		}
		if parentApiInfo.Types != 1 {
			return gerror.New("上级节点不是分类，无法新增")
		}
	}
	err = dao.SysUser.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		var apiInfo, apiInfo2 *entity.SysApi
		//根据ID查看Api列表是否存在
		apiInfo = CheckApiId(ctx, input.Id, apiInfo)
		if apiInfo == nil {
			return gerror.New("Api列表不存在")
		}
		apiInfo2 = CheckApiName(ctx, input.Name, input.Id, input.ParentId)
		if apiInfo2 != nil {
			return gerror.New("同一个分类下名称不能重复")
		}
		if input.Types == 2 {
			apiInfo2 = CheckApiAddress(ctx, input.Address, input.Id, input.ApiTypes)
			if apiInfo2 != nil {
				return gerror.New("Api地址已存在,无法修改")
			}
		}
		//获取当前登录用户ID
		loginUserId := service.Context().GetUserId(ctx)
		if apiInfoErr := gconv.Scan(input, &apiInfo); apiInfoErr != nil {
			return
		}
		apiInfo.UpdatedBy = uint(loginUserId)
		_, err = dao.SysApi.Ctx(ctx).Data(apiInfo).
			Where(dao.SysApi.Columns().Id, input.Id).Update()
		if err != nil {
			return gerror.New("修改失败")
		}
		//绑定菜单
		err = s.AddMenuApi(ctx, "api", []int{input.Id}, input.MenuIds)

		//获取所有接口并添加缓存
		_, err = s.GetApiAll(ctx, "")
		return
	})

	return
}

// Del 根据ID删除Api列表信息
func (s *sSysApi) Del(ctx context.Context, Id int) (err error) {
	var apiColumn *entity.SysApi
	_ = dao.SysApi.Ctx(ctx).Where(g.Map{
		dao.SysApi.Columns().Id: Id,
	}).Scan(&apiColumn)
	if apiColumn == nil {
		return gerror.New("ID错误")
	}
	//判断是否为分类
	if apiColumn.Types == 1 {
		//判断是否存在子节点
		num, _ := dao.SysApi.Ctx(ctx).Where(g.Map{
			dao.SysApi.Columns().ParentId:  Id,
			dao.SysApi.Columns().IsDeleted: 0,
		}).Count()
		if num > 0 {
			return gerror.New("存在子节点，无法删除")
		}
	}
	loginUserId := service.Context().GetUserId(ctx)
	//获取当前时间
	time, err := gtime.StrToTimeFormat(gtime.Datetime(), "2006-01-02 15:04:05")
	if err != nil {
		return
	}
	//开启事务管理
	err = dao.SysApi.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		//更新Api列表信息
		_, err = dao.SysApi.Ctx(ctx).
			Data(g.Map{
				dao.SysApi.Columns().DeletedBy: uint(loginUserId),
				dao.SysApi.Columns().IsDeleted: 1,
				dao.SysApi.Columns().DeletedAt: time,
			}).Where(dao.SysApi.Columns().Id, Id).
			Update()
		//删除于菜单关系绑定
		_, err = dao.SysMenuApi.Ctx(ctx).Data(g.Map{
			dao.SysMenuApi.Columns().IsDeleted: 1,
			dao.SysMenuApi.Columns().DeletedBy: loginUserId,
			dao.SysMenuApi.Columns().DeletedAt: time,
		}).Where(dao.SysMenuApi.Columns().ApiId, Id).Update()

		//获取所有接口并添加缓存
		_, err = s.GetApiAll(ctx, "")

		return
	})
	return
}

// EditStatus 修改状态
func (s *sSysApi) EditStatus(ctx context.Context, id int, status int) (err error) {
	var apiInfo *entity.SysApi
	_ = dao.SysApi.Ctx(ctx).Where(g.Map{
		dao.SysApi.Columns().Id: id,
	}).Scan(&apiInfo)
	if apiInfo == nil {
		return gerror.New("ID错误")
	}
	if apiInfo != nil && apiInfo.IsDeleted == 1 {
		return gerror.New("列表字段已删除,无法修改")
	}
	if apiInfo != nil && apiInfo.Status == status {
		return gerror.New("API已禁用或启用,无须重复修改")
	}
	loginUserId := service.Context().GetUserId(ctx)
	apiInfo.Status = status
	apiInfo.UpdatedBy = uint(loginUserId)

	_, err = dao.SysApi.Ctx(ctx).Data(apiInfo).Where(g.Map{
		dao.SysApi.Columns().Id: id,
	}).Update()

	//获取所有接口并添加缓存
	_, err = s.GetApiAll(ctx, "")

	return
}

// GetInfoByAddress 根据Address获取API
func (s *sSysApi) GetInfoByAddress(ctx context.Context, address string) (entity *entity.SysApi, err error) {
	err = dao.SysApi.Ctx(ctx).Where(g.Map{
		dao.SysApi.Columns().Address:   address,
		dao.SysApi.Columns().IsDeleted: 0,
		dao.SysApi.Columns().Status:    1,
	}).Scan(&entity)
	return
}

// CheckApiId 检查指定ID的数据是否存在
func CheckApiId(ctx context.Context, Id int, apiColumn *entity.SysApi) *entity.SysApi {
	_ = dao.SysApi.Ctx(ctx).Where(g.Map{
		dao.SysApi.Columns().Id:        Id,
		dao.SysApi.Columns().IsDeleted: 0,
	}).Scan(&apiColumn)
	return apiColumn
}

// CheckApiName 检查相同Api名称的数据是否存在
func CheckApiName(ctx context.Context, name string, tag int, parentId int) *entity.SysApi {
	var apiInfo *entity.SysApi
	m := dao.SysApi.Ctx(ctx)
	if tag > 0 {
		m = m.WhereNot(dao.SysApi.Columns().Id, tag)
	}
	m = m.Where(dao.SysApi.Columns().ParentId, parentId)
	_ = m.Where(g.Map{
		dao.SysApi.Columns().Name:      name,
		dao.SysApi.Columns().IsDeleted: 0,
	}).Scan(&apiInfo)
	return apiInfo
}

// CheckApiAddress 检查相同Api地址的数据是否存在
func CheckApiAddress(ctx context.Context, address string, tag int, apiTypes string) *entity.SysApi {
	var apiInfo *entity.SysApi
	m := dao.SysApi.Ctx(ctx)
	if tag > 0 {
		m = m.WhereNot(dao.SysApi.Columns().Id, tag)
	}
	m = m.Where(dao.SysApi.Columns().ApiTypes, apiTypes)
	_ = m.Where(g.Map{
		dao.SysApi.Columns().Address:   address,
		dao.SysApi.Columns().IsDeleted: 0,
	}).Scan(&apiInfo)
	return apiInfo
}

// GetInfoByNameAndTypes 根据名字和类型获取API
func (s *sSysApi) GetInfoByNameAndTypes(ctx context.Context, name string, types int) (entity *entity.SysApi, err error) {
	err = dao.SysApi.Ctx(ctx).Where(g.Map{
		dao.SysApi.Columns().Name:  name,
		dao.SysApi.Columns().Types: types,
	}).Scan(&entity)
	return
}

// ImportApiFile 导入API文件
func (s *sSysApi) ImportApiFile(ctx context.Context) (err error) {
	//获取服务端口号
	port := g.Cfg().MustGet(ctx, "server.address").String()
	if port == "" {
		err = gerror.New("服务地址不能为空")
		return
	}

	//获取API路径
	openApiPath := g.Cfg().MustGet(ctx, "server.openapiPath").String()
	if openApiPath == "" {
		err = gerror.New("openApi路径不能为空")
		return
	}

	url := "http://127.0.0.1" + port + openApiPath
	resp, err := g.Client().Get(ctx, url)
	if err != nil {
		return
	}
	respContent := resp.ReadAllString()
	status := resp.Status
	defer func(resp *gclient.Response) {
		if err = resp.Close(); err != nil {
			g.Log().Error(ctx, err)
		}
	}(resp)
	if !strings.Contains(status, "200") {
		err = gerror.New("请求失败,请联系管理员")
		return
	}
	var apiJsonContent map[string]interface{}
	if respContent != "" {
		err = json.Unmarshal([]byte(respContent), &apiJsonContent)
		if err != nil {
			return
		}
	}
	//封装数据、整理入库
	//获取所有的接口
	var address []string
	var apiPathsAll = make(map[string][]map[string]string)
	paths := apiJsonContent["paths"].(map[string]interface{})
	for key, value := range paths {
		//接口信息
		apiInfo := make(map[string]string)
		//获取对应的分组
		valueMap := value.(map[string]interface{})
		var valueKey string
		for valueMapKey := range valueMap {
			if !strings.EqualFold(valueMapKey, "summary") {
				valueKey = valueMapKey
				break
			}
		}
		apiInfo["address"] = key
		address = append(address, key)
		//方法
		apiInfo["method"] = valueKey
		//名字
		apiInfo["name"] = valueMap[valueKey].(map[string]interface{})["summary"].(string)

		tags := valueMap[valueKey].(map[string]interface{})["tags"].([]interface{})
		fmt.Printf("key:%s---tags:%s\n", key, tags)
		for _, tag := range tags {
			if _, ok := apiPathsAll[tag.(string)]; !ok {
				apiPathsAll[tag.(string)] = []map[string]string{apiInfo}
			} else {
				apiPathsAll[tag.(string)] = append(apiPathsAll[tag.(string)], apiInfo)
			}
		}
	}
	g.Log().Debugf(ctx, "全部API信息:%s", apiPathsAll)

	//获取数据字典类型
	dictData, err := service.DictData().GetDictDataByType(ctx, consts.ApiTypes)
	if err != nil {
		return err
	}
	apiTpyes := "IOT"
	if dictData.Data != nil && dictData.Values != nil {
		for _, value := range dictData.Values {
			if strings.EqualFold(value.DictLabel, "sagoo_iot") {
				apiTpyes = value.DictValue
			}
		}
	}
	err = dao.SysApi.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		//开始入库
		for key, value := range apiPathsAll {
			//获取分组ID
			var categoryId int
			//判断分组是否存在,不存在则创建分组
			categoryInfo, _ := s.GetInfoByNameAndTypes(ctx, key, 1)
			if categoryInfo == nil {
				//创建分组
				var result sql.Result
				result, err = dao.SysApi.Ctx(ctx).Data(do.SysApi{
					ParentId:  -1,
					Name:      key,
					Types:     1,
					Status:    1,
					IsDeleted: 0,
				}).Insert()
				if err != nil {
					return err
				}
				//获取主键ID
				var lastInsertId int64
				lastInsertId, err = service.Sequences().GetSequences(ctx, result, dao.SysApi.Table(), dao.SysApi.Columns().Id)
				if err != nil {
					return
				}
				categoryId = int(lastInsertId)
			} else {
				categoryId = int(categoryInfo.Id)
			}

			for _, apiValue := range value {
				if strings.Contains(apiValue["address"], "openapi") {
					continue
				}
				//判断接口是否存在
				var sysApi *entity.SysApi
				sysApi, err = s.GetInfoByAddress(ctx, apiValue["address"])
				if err != nil {
					return
				}
				if sysApi != nil {
					_, err = dao.SysApi.Ctx(ctx).Data(do.SysApi{
						ParentId:  categoryId,
						Name:      apiValue["name"],
						Types:     2,
						ApiTypes:  apiTpyes,
						Method:    apiValue["method"],
						Address:   apiValue["address"],
						Status:    1,
						IsDeleted: 0,
					}).Where(dao.SysApi.Columns().Address, apiValue["address"]).Update()
					if err != nil {
						return err
					}
				} else {
					_, err = dao.SysApi.Ctx(ctx).Data(do.SysApi{
						ParentId:  categoryId,
						Name:      apiValue["name"],
						Types:     2,
						ApiTypes:  apiTpyes,
						Method:    apiValue["method"],
						Address:   apiValue["address"],
						Status:    1,
						IsDeleted: 0,
					}).Insert()
					if err != nil {
						return
					}
				}
			}
		}

		//根据地址获取不存在的APIID
		var sysApiInfos []*entity.SysApi
		err = dao.SysApi.Ctx(ctx).Where(dao.SysApi.Columns().ApiTypes, apiTpyes).WhereNotIn(dao.SysApi.Columns().Address, address).Scan(&sysApiInfos)
		if err != nil {
			return
		}
		var apiIds []uint
		for _, info := range sysApiInfos {
			apiIds = append(apiIds, info.Id)
		}
		//删除API
		_, err = dao.SysApi.Ctx(ctx).Data(g.Map{
			dao.SysApi.Columns().IsDeleted: 1,
			dao.SysApi.Columns().Status:    0,
			dao.SysApi.Columns().DeletedAt: gtime.Now(),
			dao.SysApi.Columns().DeletedBy: service.Context().GetUserId(ctx),
		}).WhereIn(dao.SysApi.Columns().Id, apiIds).Update()

		if err != nil {
			return
		}
		//删除于菜单关系绑定
		_, err = dao.SysMenuApi.Ctx(ctx).Data(g.Map{
			dao.SysMenuApi.Columns().IsDeleted: 1,
			dao.SysMenuApi.Columns().DeletedAt: gtime.Now(),
			dao.SysMenuApi.Columns().DeletedBy: service.Context().GetUserId(ctx),
		}).WhereIn(dao.SysMenuApi.Columns().Id, apiIds).Update()
		if err != nil {
			return
		}
		return
	})
	return
}
