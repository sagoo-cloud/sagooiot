package system

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"path/filepath"
	"regexp"
	"sagooiot/internal/consts"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/do"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"
	"sagooiot/pkg/plugins"
	"sagooiot/pkg/utility/utils"
	"strings"
)

type sSysPlugins struct{}

func sysPluginsNew() *sSysPlugins {
	return &sSysPlugins{}
}
func init() {
	service.RegisterSysPlugins(sysPluginsNew())
}

// GetSysPluginsList 获取列表数据
func (s *sSysPlugins) GetSysPluginsList(ctx context.Context, in *model.GetSysPluginsListInput) (total, page int, list []*model.GetSysPluginsListOut, err error) {
	m := dao.SysPlugins.Ctx(ctx)

	if in.KeyWord != "" {
		m = m.WhereLike(dao.SysPlugins.Columns().Name, "%"+in.KeyWord+"%")
		m = m.WhereLike(dao.SysPlugins.Columns().Title, "%"+in.KeyWord+"%")
		m = m.WhereLike(dao.SysPlugins.Columns().Description, "%"+in.KeyWord+"%")
	}

	if in.Status > 0 {
		m = m.Where(dao.SysPlugins.Columns().Status, in.Status)
	}

	total, err = m.Count()
	if err != nil {
		err = gerror.New("获取总行数失败")
		return
	}
	page = in.PageNum
	if in.PageSize == 0 {
		in.PageSize = consts.PageSize
	}
	err = m.Page(page, in.PageSize).OrderDesc(dao.SysPlugins.Columns().CreatedAt).Scan(&list)
	if err != nil {
		err = gerror.New("获取数据失败")
	}
	return
}

// GetSysPluginsById 获取指定ID数据
func (s *sSysPlugins) GetSysPluginsById(ctx context.Context, id int) (out *entity.SysPlugins, err error) {
	err = dao.SysPlugins.Ctx(ctx).Where(dao.SysPlugins.Columns().Id, id).Scan(&out)
	return
}

// GetSysPluginsByName 根据名称获取插件数据
func (s *sSysPlugins) GetSysPluginsByName(ctx context.Context, name string) (out *entity.SysPlugins, err error) {
	err = dao.SysPlugins.Ctx(ctx).Where(dao.SysPlugins.Columns().Name, name).Scan(&out)
	return
}

// GetSysPluginsByTitle 根据TITLE获取插件数据
func (s *sSysPlugins) GetSysPluginsByTitle(ctx context.Context, title string) (out *entity.SysPlugins, err error) {
	err = dao.SysPlugins.Ctx(ctx).Where(dao.SysPlugins.Columns().Title, title).Scan(&out)
	return
}

// AddSysPlugins 添加数据
func (s *sSysPlugins) AddSysPlugins(ctx context.Context, file *ghttp.UploadFile) (err error) {

	srcFileName := file.Filename
	ext := filepath.Ext(srcFileName)
	if ext != ".zip" && ext != ".zipx" {
		err = gerror.New("只允许上传zip或zipx文件！")
		return
	}

	//获取文件名(不包括后缀)
	fileNameWithoutExt := srcFileName[:len(srcFileName)-len(filepath.Ext(srcFileName))]

	pattern := `^[^\p{Han}]*$`

	match, err := regexp.MatchString(pattern, fileNameWithoutExt)
	if err != nil {
		err = gerror.New("正则表达式匹配失败")
		return
	}

	if !match {
		err = gerror.New("文件名请用英文命名")
		return
	}
	//读取压缩包内指定文件内容
	result, err := utils.ReadZipFileByFileName(file, "info.json")
	if err != nil {
		return err
	}
	var plugins *entity.SysPlugins
	if err = gconv.Scan(result, &plugins); err != nil {
		return
	}
	if plugins.Types == "" {
		err = gerror.New("插件类型不能为空")
		return
	}
	if plugins.HandleType == "" {
		err = gerror.New("处理方式类型不能为空")
		return
	}
	if plugins.Name == "" {
		err = gerror.New("名称不能为空")
		return
	}
	if plugins.Title == "" {
		err = gerror.New("标题不能为空")
		return
	}
	//针对执行程序参数做特殊处理
	var args = result["args"].([]interface{})
	var argsArr []string
	for _, v := range args {
		a, _ := v.(string)

		argsArr = append(argsArr, a)
	}
	plugins.Args = strings.Join(argsArr, ",")
	//对plugins struct赋值
	frontend, _ := result["frontend"].(map[string]interface{})

	plugins.FrontendUi = 0
	if gconv.Bool(frontend["ui"]) {
		plugins.FrontendUi = 1
	}

	plugins.FrontendUrl = frontend["url"].(string)

	plugins.FrontendConfiguration = 0
	if gconv.Bool(frontend["configuration"]) {
		plugins.FrontendConfiguration = 1
	}
	//判断名称是否重复
	pluginByName, _ := s.GetSysPluginsByName(ctx, plugins.Name)
	if pluginByName != nil {
		return gerror.New("插件名称不能重复")
	}
	//判断标题是否重复
	pluginByTitle, _ := s.GetSysPluginsByTitle(ctx, plugins.Title)
	if pluginByTitle != nil {
		return gerror.New("插件标题不能重复")
	}

	err = dao.SysPlugins.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		//添加插件
		_, err = dao.SysPlugins.Ctx(ctx).Data(do.SysPlugins{
			DeptId:                service.Context().GetUserDeptId(ctx),
			Types:                 plugins.Types,
			HandleType:            plugins.HandleType,
			Name:                  plugins.Name,
			Title:                 plugins.Title,
			Description:           plugins.Description,
			Version:               plugins.Version,
			Author:                plugins.Author,
			Icon:                  plugins.Icon,
			Link:                  plugins.Link,
			Command:               plugins.Command,
			Args:                  plugins.Args,
			Status:                0,
			FrontendUi:            plugins.FrontendUi,
			FrontendUrl:           plugins.FrontendUrl,
			FrontendConfiguration: plugins.FrontendConfiguration,
			StartTime:             plugins.StartTime,
			IsDeleted:             0,
			CreatedBy:             uint(service.Context().GetUserId(ctx)),
			CreatedAt:             gtime.Now(),
		}).Insert()
		if err != nil {
			err = gerror.New("添加插件表失败!")
			return
		}
		//上传插件
		pluginsPath := g.Cfg().MustGet(context.Background(), "system.pluginsPath").String()
		excludeFiles := []string{"info.json"} // 需要排除的文件名列表
		err = utils.UploadZip(file, pluginsPath, excludeFiles)
		if err != nil {
			return
		}
		return
	})

	return
}

func shouldExclude(filename string, excludeList []string) bool {
	baseName := filepath.Base(filename)

	for _, excludedFile := range excludeList {
		if baseName == excludedFile {
			return true
		}
	}

	return false
}

// EditSysPlugins 修改数据
func (s *sSysPlugins) EditSysPlugins(ctx context.Context, input *model.SysPluginsEditInput) (err error) {
	//根据ID查询
	var sp *entity.SysPlugins
	err = dao.SysPlugins.Ctx(ctx).Where(dao.SysPlugins.Columns().Id, input.Id).Scan(&sp)
	if sp == nil {
		err = gerror.New("ID错误")
		return
	}

	//判断名称是否重复
	pluginByName, _ := s.GetSysPluginsByName(ctx, input.Name)
	if pluginByName != nil && pluginByName.Id != sp.Id {
		return gerror.New("插件名称不能重复")
	}
	//判断标题是否重复
	pluginByTitle, _ := s.GetSysPluginsByTitle(ctx, input.Title)
	if pluginByTitle != nil && pluginByTitle.Id != sp.Id {
		return gerror.New("插件标题不能重复")
	}

	sp.Types = input.Types
	sp.HandleType = input.HandleType
	sp.Name = input.Name
	sp.Title = input.Title
	sp.Description = input.Description
	sp.Version = input.Version
	sp.Author = input.Author //strings.Join(input.Author, ",")
	sp.Icon = input.Icon
	sp.Link = input.Link
	sp.Command = input.Command
	sp.Args = strings.Join(input.Args, ",")
	sp.FrontendUi = input.FrontendUi
	sp.FrontendUrl = input.FrontendUrl
	sp.FrontendConfiguration = input.FrontendConfiguration
	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)
	sp.UpdatedBy = loginUserId
	sp.UpdatedAt = gtime.Now()
	_, err = dao.SysPlugins.Ctx(ctx).Data(sp).Where(dao.SysPlugins.Columns().Id, sp.Id).Update()
	return
}

// DeleteSysPlugins 删除数据
func (s *sSysPlugins) DeleteSysPlugins(ctx context.Context, ids []int) (err error) {
	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	for _, id := range ids {
		plugin, _ := s.GetSysPluginsById(ctx, id)
		if plugin == nil {
			return gerror.New("ID错误，请重新选择!")
		}
		_, err = dao.SysPlugins.Ctx(ctx).Data(g.Map{
			dao.SysPlugins.Columns().Status:    0,
			dao.SysPlugins.Columns().IsDeleted: 1,
			dao.SysPlugins.Columns().DeletedBy: loginUserId,
			dao.SysPlugins.Columns().DeletedAt: gtime.Now(),
		}).Where(dao.SysPlugins.Columns().Id, id).Update()
	}
	return
}

// SaveSysPlugins 存入插件数据，跟据插件类型与名称，数据中只保存一份
func (s *sSysPlugins) SaveSysPlugins(ctx context.Context, in model.SysPluginsAddInput) (err error) {
	var req = g.Map{
		dao.SysPlugins.Columns().Types: in.Types,
		dao.SysPlugins.Columns().Name:  in.Name,
	}
	res, err := dao.SysPlugins.Ctx(ctx).Where(req).One()
	if res != nil {
		_, err = dao.SysPlugins.Ctx(ctx).Data(in).Where(req).Update()

	} else {
		//获取当前登录用户ID
		loginUserId := service.Context().GetUserId(ctx)

		_, err = dao.SysPlugins.Ctx(ctx).Data(do.SysPlugins{
			DeptId:                service.Context().GetUserDeptId(ctx),
			Types:                 in.Types,
			HandleType:            in.HandleType,
			Name:                  in.Name,
			Title:                 in.Title,
			Description:           in.Description,
			Version:               in.Version,
			Author:                in.Author,
			Icon:                  in.Icon,
			Link:                  in.Link,
			Command:               in.Command,
			Args:                  in.Args,
			Status:                in.Status,
			FrontendUi:            in.FrontendUi,
			FrontendUrl:           in.FrontendUrl,
			FrontendConfiguration: in.FrontendConfiguration,
			StartTime:             in.StartTime,
			IsDeleted:             0,
			CreatedBy:             uint(loginUserId),
			CreatedAt:             gtime.Now(),
		}).Insert()
	}
	return
}

func (s *sSysPlugins) EditStatus(ctx context.Context, id int, status int) (err error) {
	//获取当前登录用户ID
	var pluginObj *entity.SysPlugins
	err = dao.SysPlugins.Ctx(ctx).Where(dao.SysPlugins.Columns().Id, id).Scan(&pluginObj)
	if err != nil {
		err = gerror.New("获取插件信息失败")
		return
	}

	switch status {
	case 0:
		if pluginObj.Types == "notice" {
			//关闭通知插件
			err = plugins.GetNoticePlugin().StopPlugin(pluginObj.Name)
		} else {
			//关闭协议插件
			err = plugins.GetProtocolPlugin().StopPlugin(pluginObj.Name)

		}
	case 1:
		if pluginObj.Types == "notice" {
			//启动通知插件
			err = plugins.GetNoticePlugin().StartPlugin(pluginObj.Name)
		} else {
			//启动协议插件
			err = plugins.GetProtocolPlugin().StartPlugin(pluginObj.Name)
		}
	}
	_, err = dao.SysPlugins.Ctx(ctx).Data("status", status).Where(dao.SysPlugins.Columns().Id, id).Update()
	return
}

// GetSysPluginsTypesAll 获取所有插件的通信方式类型
func (s *sSysPlugins) GetSysPluginsTypesAll(ctx context.Context, types string) (out []*model.SysPluginsInfoOut, err error) {
	m := dao.SysPlugins.Ctx(ctx)

	m = m.Where(g.Map{
		dao.SysPlugins.Columns().Status:    1,
		dao.SysPlugins.Columns().IsDeleted: 0,
	}).WhereIn(dao.SysPlugins.Columns().Types, types)

	//获取当前登录用户ID
	var sp []*entity.SysPlugins
	err = m.Scan(&sp)
	if err != nil {
		return
	}
	for _, plugin := range sp {
		var sysPlugins = new(model.SysPluginsInfoOut)
		sysPlugins.Types = plugin.Types
		sysPlugins.HandleType = plugin.HandleType
		sysPlugins.Name = plugin.Name
		sysPlugins.Title = plugin.Title
		out = append(out, sysPlugins)
	}
	return
}
