// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/internal/model"
)

type (
	INoticeLog interface {
		Add(ctx context.Context, in *model.NoticeLogAddInput) (err error)
		Del(ctx context.Context, ids []uint64) (err error)
		Search(ctx context.Context, in *model.NoticeLogSearchInput) (out *model.NoticeLogSearchOutput, err error)
		ClearLogByDays(ctx context.Context, days int) (err error)
	}
	INoticeTemplate interface {
		GetNoticeTemplateList(ctx context.Context, in *model.GetNoticeTemplateListInput) (total, page int, list []*model.NoticeTemplateOutput, err error)
		GetNoticeTemplateById(ctx context.Context, id string) (out *model.NoticeTemplateOutput, err error)
		GetNoticeTemplateByConfigId(ctx context.Context, configId string) (out *model.NoticeTemplateOutput, err error)
		AddNoticeTemplate(ctx context.Context, in model.NoticeTemplateAddInput) (err error)
		EditNoticeTemplate(ctx context.Context, in model.NoticeTemplateEditInput) (err error)
		SaveNoticeTemplate(ctx context.Context, in model.NoticeTemplateAddInput) (err error)
		DeleteNoticeTemplate(ctx context.Context, Ids []string) (err error)
	}
	INoticeConfig interface {
		GetNoticeConfigList(ctx context.Context, in *model.GetNoticeConfigListInput) (total, page int, list []*model.NoticeConfigOutput, err error)
		GetNoticeConfigById(ctx context.Context, id int) (out *model.NoticeConfigOutput, err error)
		AddNoticeConfig(ctx context.Context, in model.NoticeConfigAddInput) (err error)
		EditNoticeConfig(ctx context.Context, in model.NoticeConfigEditInput) (err error)
		DeleteNoticeConfig(ctx context.Context, Ids []string) (err error)
	}
	INoticeInfo interface {
		GetNoticeInfoList(ctx context.Context, in *model.GetNoticeInfoListInput) (total, page int, list []*model.NoticeInfoOutput, err error)
		GetNoticeInfoById(ctx context.Context, id int) (out *model.NoticeInfoOutput, err error)
		AddNoticeInfo(ctx context.Context, in model.NoticeInfoAddInput) (err error)
		EditNoticeInfo(ctx context.Context, in model.NoticeInfoEditInput) (err error)
		DeleteNoticeInfo(ctx context.Context, Ids []int) (err error)
	}
)

var (
	localNoticeConfig   INoticeConfig
	localNoticeInfo     INoticeInfo
	localNoticeLog      INoticeLog
	localNoticeTemplate INoticeTemplate
)

func NoticeTemplate() INoticeTemplate {
	if localNoticeTemplate == nil {
		panic("implement not found for interface INoticeTemplate, forgot register?")
	}
	return localNoticeTemplate
}

func RegisterNoticeTemplate(i INoticeTemplate) {
	localNoticeTemplate = i
}

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
