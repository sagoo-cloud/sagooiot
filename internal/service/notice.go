// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"sagooiot/internal/model"
	"sagooiot/internal/model/entity"
)

type (
	INoticeConfig interface {
		// GetNoticeConfigList 获取列表数据
		GetNoticeConfigList(ctx context.Context, in *model.GetNoticeConfigListInput) (total, page int, list []*model.NoticeConfigOutput, err error)
		// GetNoticeConfigById 获取指定ID数据
		GetNoticeConfigById(ctx context.Context, id string) (out *model.NoticeConfigOutput, err error)
		// AddNoticeConfig 添加数据
		AddNoticeConfig(ctx context.Context, in model.NoticeConfigAddInput) (err error)
		// EditNoticeConfig 修改数据
		EditNoticeConfig(ctx context.Context, in model.NoticeConfigEditInput) (err error)
		// DeleteNoticeConfig 删除数据
		DeleteNoticeConfig(ctx context.Context, Ids []string) (err error)
	}
	INoticeInfo interface {
		// GetNoticeInfoList 获取列表数据
		GetNoticeInfoList(ctx context.Context, in *model.GetNoticeInfoListInput) (total, page int, list []*model.NoticeInfoOutput, err error)
		// GetNoticeInfoById 获取指定ID数据
		GetNoticeInfoById(ctx context.Context, id int) (out *model.NoticeInfoOutput, err error)
		// AddNoticeInfo 添加数据
		AddNoticeInfo(ctx context.Context, in model.NoticeInfoAddInput) (err error)
		// EditNoticeInfo 修改数据
		EditNoticeInfo(ctx context.Context, in model.NoticeInfoEditInput) (err error)
		// DeleteNoticeInfo 删除数据
		DeleteNoticeInfo(ctx context.Context, Ids []int) (err error)
	}
	INoticeLog interface {
		// Add 通知日志记录
		Add(ctx context.Context, in *model.NoticeLogAddInput) (err error)
		// Del 删除日志
		Del(ctx context.Context, ids []uint64) (err error)
		// GetInfoById 获取日志信息
		GetInfoById(ctx context.Context, id uint64) (out *entity.NoticeLog, err error)
		// Search 搜索
		Search(ctx context.Context, in *model.NoticeLogSearchInput) (out *model.NoticeLogSearchOutput, err error)
		// ClearLogByDays 按日期删除日志
		ClearLogByDays(ctx context.Context, days int) (err error)
	}
	INoticeTemplate interface {
		// GetNoticeTemplateList 获取列表数据
		GetNoticeTemplateList(ctx context.Context, in *model.GetNoticeTemplateListInput) (total, page int, list []*model.NoticeTemplateOutput, err error)
		// GetNoticeTemplateById 获取指定ID数据
		GetNoticeTemplateById(ctx context.Context, id string) (out *model.NoticeTemplateOutput, err error)
		// GetNoticeTemplateByConfigId 获取指定ConfigID数据
		GetNoticeTemplateByConfigId(ctx context.Context, configId string) (out *model.NoticeTemplateOutput, err error)
		// AddNoticeTemplate 添加数据
		AddNoticeTemplate(ctx context.Context, in model.NoticeTemplateAddInput) (err error)
		// EditNoticeTemplate 修改数据
		EditNoticeTemplate(ctx context.Context, in model.NoticeTemplateEditInput) (err error)
		// SaveNoticeTemplate 直接更新数据
		SaveNoticeTemplate(ctx context.Context, in model.NoticeTemplateAddInput) (err error)
		// DeleteNoticeTemplate 删除数据
		DeleteNoticeTemplate(ctx context.Context, Ids []string) (err error)
	}
)

var (
	localNoticeInfo     INoticeInfo
	localNoticeLog      INoticeLog
	localNoticeTemplate INoticeTemplate
	localNoticeConfig   INoticeConfig
)

func NoticeConfig() INoticeConfig {
	if localNoticeConfig == nil {
		panic("implement not found for interface INoticeConfig, forgot register?")
	}
	return localNoticeConfig
}

func RegisterNoticeConfig(i INoticeConfig) {
	localNoticeConfig = i
}

func NoticeInfo() INoticeInfo {
	if localNoticeInfo == nil {
		panic("implement not found for interface INoticeInfo, forgot register?")
	}
	return localNoticeInfo
}

func RegisterNoticeInfo(i INoticeInfo) {
	localNoticeInfo = i
}

func NoticeLog() INoticeLog {
	if localNoticeLog == nil {
		panic("implement not found for interface INoticeLog, forgot register?")
	}
	return localNoticeLog
}

func RegisterNoticeLog(i INoticeLog) {
	localNoticeLog = i
}

func NoticeTemplate() INoticeTemplate {
	if localNoticeTemplate == nil {
		panic("implement not found for interface INoticeTemplate, forgot register?")
	}
	return localNoticeTemplate
}

func RegisterNoticeTemplate(i INoticeTemplate) {
	localNoticeTemplate = i
}
