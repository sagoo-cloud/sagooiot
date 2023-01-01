package model

type PaginationInput struct {
	KeyWord   string   `json:"keyWord" dc:"搜索关键字"` //搜索关键字
	DateRange []string `p:"dateRange"`             //日期范围
	OrderBy   string   //排序方式
	PageNum   int      `json:"pageNum" in:"query" d:"1"  v:"min:0#分页号码错误"     dc:"分页号码，默认1"`
	PageSize  int      `json:"PageSize" in:"query" d:"10" v:"max:50#分页数量最大50条" dc:"分页数量，最大50"`
}

type PaginationOutput struct {
	CurrentPage int `json:"currentPage" dc:"当前页"`
	Total       int `dc:"总数"`
}
