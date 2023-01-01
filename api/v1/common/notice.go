package common

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetNoticeReq struct {
	g.Meta `path:"/notice/data" tags:"公共方法" method:"get" summary:"获取实时通知信息"`
	Topic  string `p:"topic" v:"required#订阅主题不能为空"`
}

type GetNoticeRes struct {
	g.Meta `mime:"application/json"`
}

type GetNotice2Req struct {
	g.Meta `path:"/notice/data2" tags:"公共方法" method:"get" summary:"获取实时通知信息"`
	Topic  string `p:"topic" v:"required#订阅主题不能为空"`
}

type GetNotice2Res struct {
	g.Meta `mime:"application/json"`
}
