package common

type PaginationReq struct {
	Param          map[string]interface{} `json:"param" dc:"搜索字段参数，写法param[字段名称]"`
	KeyWord        string                 `json:"keyWord" dc:"搜索关键字"`
	Year           string                 `json:"year" dc:"年份，如：2024"`
	DateRange      []string               `json:"dateRange" dc:"日期范围，数组"`
	Accurate       string                 `json:"accurate" dc:"数据精确类型（m：月，d：天，h：小时）"`
	AccurateRanges string                 `json:"accurateRanges" dc:"精确值，当accurate精确类型为d的时候，可以指定月份，当accurate为h的时候，可以指定某日的值"`
	OrderBy        string                 `json:"orderBy" dc:"排序方式"` //排序方式
	PageNum        int                    `json:"pageNum" in:"query" d:"1"  v:"min:0#分页号码错误"     dc:"分页号码，默认1"`
	PageSize       int                    `json:"PageSize" in:"query" d:"10" v:"max:500#分页数量最大500条" dc:"分页数量，最大500"`
}

type PaginationRes struct {
	CurrentPage int `json:"currentPage" dc:"当前页"`
	Total       int `dc:"总数"`
}
