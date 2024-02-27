package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"sagooiot/internal/model/entity"
)

type MessageListDoInput struct {
	Types int    `json:"types"        description:"消息类型"`
	Title string `json:"title"     description:"标题"`
	*PaginationInput
}

type MessageListRes struct {
	Id          int                `json:"id"        description:""`
	UserId      int                `json:"userId"    description:"用户ID"`
	MessageId   int                `json:"messageId" description:"消息ID"`
	IsRead      int                `json:"isRead"    description:"是否已读 0 未读 1已读"`
	IsDeleted   int                `json:"isDeleted" description:"是否删除 0未删除 1已删除"`
	ReadTime    *gtime.Time        `json:"readTime"  description:"阅读时间"`
	DeletedAt   *gtime.Time        `json:"deletedAt" description:"删除时间"`
	MessageInfo *entity.SysMessage `orm:"with:id=message_id" description:"消息"`
}

type MessageListOut struct {
	Id          int                `json:"id"        description:""`
	UserId      int                `json:"userId"    description:"用户ID"`
	MessageId   int                `json:"messageId" description:"消息ID"`
	IsRead      int                `json:"isRead"    description:"是否已读 0 未读 1已读"`
	IsDeleted   int                `json:"isDeleted" description:"是否删除 0未删除 1已删除"`
	ReadTime    *gtime.Time        `json:"readTime"  description:"阅读时间"`
	DeletedAt   *gtime.Time        `json:"deletedAt" description:"删除时间"`
	MessageInfo *entity.SysMessage `orm:"with:id=message_id" description:"消息"`
}

type AddMessageInput struct {
	Title    string `json:"title"     description:"标题"`
	Types    int    `json:"types"     description:"消息类型 字典表中查询具体消息类型"`
	Scope    int    `json:"scope"     description:"消息范围 字典表中查看具体消息范围"`
	Content  string `json:"content"   description:"内容"`
	ObjectId int    `json:"objectId"   description:"推送对象ID"`
}

type SysMessagereceiveInput struct {
	UserId    int         `json:"userId"    description:"用户ID"`
	MessageId int         `json:"messageId" description:"消息ID"`
	IsRead    int         `json:"isRead"    description:"是否已读 0 未读 1已读"`
	IsPush    int         `json:"isPush"    description:"是否已经推送0 否 1是"`
	IsDeleted int         `json:"isDeleted" description:"是否删除 0未删除 1已删除"`
	ReadTime  *gtime.Time `json:"readTime"  description:"阅读时间"`
	DeletedAt *gtime.Time `json:"deletedAt" description:"删除时间"`
}
