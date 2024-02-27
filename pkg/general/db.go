package general

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
)

// SelectReq 查询请求参数
type SelectReq struct {
	Param          map[string]interface{} `p:"param"`          //搜索哪些字段的数据
	KeyWords       string                 `p:"keyWords"`       //关键字搜索
	Year           string                 `p:"year"`           //提取哪一年的数据
	DateRange      []string               `p:"dateRange"`      //提取指定日期范围的数据
	Accurate       string                 `p:"accurate"`       //数据精确类型（m：月，d：天，h：小时）
	AccurateRanges string                 `p:"accurateRanges"` //当accurate精确类型为d的时候，可以指定月份，当accurate为h的时候，可以指定某日的值
	PageNum        int                    `p:"pageNum"`        //当前页码
	PageSize       int                    `p:"pageSize"`       //每页数
}

// DataListByPageRes 分页查询返回结果
type DataListByPageRes struct {
	Total       int         `json:"total"`       //总数
	CurrentPage int         `json:"currentPage"` //当前页码
	Data        interface{} `json:"data"`        //数据
}

// ListByPage 列表查询
func ListByPage(ctx context.Context, dataModel *gdb.Model, req *SelectReq, vagueField []string) (res DataListByPageRes, err error) {

	fieldsStr := "*"
	groupStr := ""

	if req != nil {

		if req.Param == nil {
			req.Param = make(map[string]interface{})
		}

		if req.Accurate != "" {
			switch req.Accurate {
			case "y": //精度按月为单位筛选数据
				fieldsStr = "*,DATE_FORMAT(created_at,'%Y') time_num,created_at"
				groupStr = "DATE_FORMAT(created_at,'%Y')"
			case "m": //精度按月为单位筛选数据
				fieldsStr = "*,DATE_FORMAT(created_at,'%m') time_num,created_at"
				groupStr = "DATE_FORMAT(created_at,'%Y-%m')"
			case "d": //精度按天为单位筛选数据
				y := gtime.Now().Format("Y")
				m := gtime.Now().Format("m")
				if req.AccurateRanges == "" {
					req.Param["DATE_FORMAT(created_at,'%Y-%m')"] = y + "-" + m
				} else {
					req.Param["DATE_FORMAT(created_at,'%Y-%m')"] = y + "-" + req.AccurateRanges
				}
				fieldsStr = "*,DATE_FORMAT(created_at,'%d') time_num,created_at"
				groupStr = "DATE_FORMAT(created_at,'%Y-%m-%d')"
			case "h": //精度按小时为单位筛选数据
				y := gtime.Now().Format("Y")
				m := gtime.Now().Format("m")
				d := gtime.Now().Format("d")
				if req.AccurateRanges == "" {
					req.Param["DATE_FORMAT(created_at,'%Y-%m-%d')"] = y + "-" + m + "-" + d
				} else {
					req.Param["DATE_FORMAT(created_at,'%Y-%m')"] = y + "-" + m + "-" + req.AccurateRanges
				}
				fieldsStr = "*,DATE_FORMAT(created_at,'%H') time_num,created_at"
				groupStr = "DATE_FORMAT('%H',created_at)"
			default:

			}

		}

		//KeyWords不为空的时候，进行模糊查询处理
		if req.KeyWords != "" && vagueField != nil {
			for _, vf := range vagueField {
				dataModel = dataModel.WhereOrLike(vf, "%"+req.KeyWords+"%")

			}
		}

		if req.Year != "" {
			req.Param["DATE_FORMAT(created_at,'%Y')"] = req.Year
		} else {
			if len(req.DateRange) > 0 {
				dataModel = dataModel.WhereBetween("created_at", req.DateRange[0], req.DateRange[1])
			}
		}

		dataModel = dataModel.Where(req.Param).Fields(fieldsStr).Group(groupStr)

	} else {
		dataModel = dataModel.Fields(fieldsStr).Group(groupStr)
		req = new(SelectReq)
	}

	countModel := dataModel
	totalData, err := countModel.All()
	res.Total = totalData.Len()
	if err != nil {
		glog.Debug(ctx, err)
		err = gerror.New("获取总行数失败")
		return
	}
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res.CurrentPage = req.PageNum
	req.PageSize = g.Cfg().MustGet(ctx, "system.DefaultPageSize").Int()
	if req.PageSize == 0 {
		req.PageSize = 1000
	}

	res.Data, err = dataModel.Page(res.CurrentPage, req.PageSize).All()
	if err != nil {
		err = gerror.New("获取数据失败")
		return
	}
	return
}
