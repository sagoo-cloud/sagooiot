// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"database/sql"
	"sagooiot/api/v1/common"
	"sagooiot/internal/model"
	"sagooiot/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IUpload interface {
		// UploadFiles 上传多文件
		UploadFiles(ctx context.Context, files []*ghttp.UploadFile, checkFileType string, source int) (result common.UploadMultipleRes, err error)
		// UploadFile 上传单文件
		UploadFile(ctx context.Context, file *ghttp.UploadFile, checkFileType string, source int) (result common.UploadResponse, err error)
		// UploadTencent 上传至腾讯云
		UploadTencent(ctx context.Context, file *ghttp.UploadFile) (result common.UploadResponse, err error)
		// UploadLocal 上传本地
		UploadLocal(ctx context.Context, file *ghttp.UploadFile) (result common.UploadResponse, err error)
		// CheckSize 检查上传文件大小
		CheckSize(ctx context.Context, checkFileType string, file *ghttp.UploadFile) (err error)
		// CheckType 检查上传文件类型
		CheckType(ctx context.Context, checkFileType string, file *ghttp.UploadFile) (err error)
		// UploadMinIO 上传至MinIO
		UploadMinIO(ctx context.Context, file *ghttp.UploadFile) (result common.UploadResponse, err error)
	}
	ICheckAuth interface {
		// IsToken 验证TOKEN是否正确
		IsToken(ctx context.Context) (isToken bool, expiresAt int64, isAuth string, err error)
		// CheckAccessAuth 验证访问权限
		CheckAccessAuth(ctx context.Context, address string) (isAllow bool, err error)
	}
	IConfigData interface {
		// List 系统参数列表
		List(ctx context.Context, input *model.ConfigDoInput) (total int, out []*model.SysConfigOut, err error)
		Add(ctx context.Context, input *model.AddConfigInput, userId int) (err error)
		// CheckConfigKeyUnique 验证参数键名是否存在
		CheckConfigKeyUnique(ctx context.Context, configKey string, configId ...int) (err error)
		// Get 获取系统参数
		Get(ctx context.Context, id int) (out *model.SysConfigOut, err error)
		// Edit 修改系统参数
		Edit(ctx context.Context, input *model.EditConfigInput, userId int) (err error)
		// Delete 删除系统参数 //TODO 转为KEY处理
		Delete(ctx context.Context, ids []int) (err error)
		// GetConfigByKey 通过key获取参数（从缓存获取）
		GetConfigByKey(ctx context.Context, key string) (config *entity.SysConfig, err error)
		// GetConfigByKeys 通过key数组获取参数（从缓存获取）
		GetConfigByKeys(ctx context.Context, keys []string) (out []*entity.SysConfig, err error)
		// GetByKey 通过key获取参数（从数据库获取）
		GetByKey(ctx context.Context, key string) (config *entity.SysConfig, err error)
		// GetByKeys 通过keys获取参数（从数据库获取）
		GetByKeys(ctx context.Context, keys []string) (config []*entity.SysConfig, err error)
		GetSysConfigSetting(ctx context.Context, types int) (out []*entity.SysConfig, err error)
		// EditSysConfigSetting 修改系统配置设置
		EditSysConfigSetting(ctx context.Context, inputs []*model.EditConfigInput) (err error)
		// GetLoadCache 获取本地缓存配置
		GetLoadCache(ctx context.Context) (conf *model.CacheConfig, err error)
	}
	IDictData interface {
		// GetDictWithDataByType 通过字典键类型获取选项
		GetDictWithDataByType(ctx context.Context, input *model.GetDictInput) (dict *model.GetDictOut, err error)
		// List 获取字典数据
		List(ctx context.Context, input *model.SysDictSearchInput) (total int, out []*model.SysDictDataOut, err error)
		Add(ctx context.Context, input *model.AddDictDataInput, userId int) (err error)
		// Get 获取字典数据
		Get(ctx context.Context, dictCode uint) (out *model.SysDictDataOut, err error)
		// Edit 修改字典数据
		Edit(ctx context.Context, input *model.EditDictDataInput, userId int) (err error)
		// Delete 删除字典数据
		Delete(ctx context.Context, ids []int) (err error)
		// GetDictDataByType 通过字典键类型获取选项
		GetDictDataByType(ctx context.Context, dictType string) (dict *model.GetDictOut, err error)
	}
	IDictType interface {
		// List 字典类型列表
		List(ctx context.Context, input *model.DictTypeDoInput) (total int, out []*model.SysDictTypeInfoOut, err error)
		// Add 添加字典类型
		Add(ctx context.Context, input *model.AddDictTypeInput, userId int) (err error)
		// Edit 修改字典类型
		Edit(ctx context.Context, input *model.EditDictTypeInput, userId int) (err error)
		Get(ctx context.Context, req *common.DictTypeGetReq) (dictType *model.SysDictTypeOut, err error)
		// ExistsDictType 检查类型是否已经存在
		ExistsDictType(ctx context.Context, dictType string, dictId ...int) (err error)
		// Delete 删除字典类型
		Delete(ctx context.Context, dictIds []int) (err error)
	}
	IPgSequences interface {
		// GetPgSequences 获取PG指定表序列信息
		GetPgSequences(ctx context.Context, tableName string, primaryKey string) (out *model.PgSequenceOut, err error)
	}
	ISequences interface {
		// GetSequences 获取主键ID
		GetSequences(ctx context.Context, result sql.Result, tableName string, primaryKey string) (lastInsertId int64, err error)
	}
	ISysInfo interface {
		GetSysInfo(ctx context.Context) (out g.Map, err error)
		// ServerInfoEscalation 客户端服务信息上报
		ServerInfoEscalation(ctx context.Context) (err error)
	}
)

var (
	localSysInfo     ISysInfo
	localUpload      IUpload
	localCheckAuth   ICheckAuth
	localConfigData  IConfigData
	localDictData    IDictData
	localDictType    IDictType
	localPgSequences IPgSequences
	localSequences   ISequences
)

func DictData() IDictData {
	if localDictData == nil {
		panic("implement not found for interface IDictData, forgot register?")
	}
	return localDictData
}

func RegisterDictData(i IDictData) {
	localDictData = i
}

func DictType() IDictType {
	if localDictType == nil {
		panic("implement not found for interface IDictType, forgot register?")
	}
	return localDictType
}

func RegisterDictType(i IDictType) {
	localDictType = i
}

func PgSequences() IPgSequences {
	if localPgSequences == nil {
		panic("implement not found for interface IPgSequences, forgot register?")
	}
	return localPgSequences
}

func RegisterPgSequences(i IPgSequences) {
	localPgSequences = i
}

func Sequences() ISequences {
	if localSequences == nil {
		panic("implement not found for interface ISequences, forgot register?")
	}
	return localSequences
}

func RegisterSequences(i ISequences) {
	localSequences = i
}

func SysInfo() ISysInfo {
	if localSysInfo == nil {
		panic("implement not found for interface ISysInfo, forgot register?")
	}
	return localSysInfo
}

func RegisterSysInfo(i ISysInfo) {
	localSysInfo = i
}

func Upload() IUpload {
	if localUpload == nil {
		panic("implement not found for interface IUpload, forgot register?")
	}
	return localUpload
}

func RegisterUpload(i IUpload) {
	localUpload = i
}

func CheckAuth() ICheckAuth {
	if localCheckAuth == nil {
		panic("implement not found for interface ICheckAuth, forgot register?")
	}
	return localCheckAuth
}

func RegisterCheckAuth(i ICheckAuth) {
	localCheckAuth = i
}

func ConfigData() IConfigData {
	if localConfigData == nil {
		panic("implement not found for interface IConfigData, forgot register?")
	}
	return localConfigData
}

func RegisterConfigData(i IConfigData) {
	localConfigData = i
}
