package thing

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/sagoo-cloud/sagooiot/internal/model"
)

// 物联概览统计数据
type ThingOverviewReq struct {
	g.Meta `path:"/overview" method:"get" summary:"物联概览统计数据" tags:"数据概览"`
}
type ThingOverviewRes struct {
	model.ThingOverviewOutput
}
