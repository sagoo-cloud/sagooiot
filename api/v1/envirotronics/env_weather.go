package envirotronics

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/sagoo-cloud/sagooiot/internal/model"
)

type CityWeatherListReq struct {
	g.Meta `path:"/weather/cityWeatherList" method:"get" summary:"获取城市的风力及日照时长" tags:"天气监测"`
	Name   string `json:"name" description:"名字"`
}
type CityWeatherListRes struct {
	Info []*model.CityWeatherListRes
}

type GetCityWeatherByIdReq struct {
	g.Meta `path:"/weather/getInfoById" method:"get" summary:"根据ID获取指定城市的天气" tags:"天气监测"`
	Id     int `json:"id" description:"主键ID" v:"required#ID不能为空"`
}
type GetCityWeatherByIdRes struct {
	Info *model.CityWeatherListRes
}

type GetCityTemperatureByIdReq struct {
	g.Meta `path:"/weather/getTemperatureEchartById" method:"get" summary:"根据ID获取指定城市的温度图表" tags:"天气监测"`
	Id     int `json:"id" description:"主键ID" v:"required#ID不能为空"`
	Types  int `json:"types" description:"类型 1 日 2周 3月 4年" v:"required#类型不能为空"`
}
type GetCityTemperatureByIdRes struct {
	Info            []*model.CityWeatherEchartRes
	AvgInfo         []*model.CityWeatherEchartRes
	ForeCastInfo    []*model.CityWeatherEchartRes
	ForeCastAvgInfo []*model.CityWeatherEchartRes
}

type GetCityWindpowerByIdReq struct {
	g.Meta `path:"/weather/getWindpowerEchartById" method:"get" summary:"根据ID获取指定城市的风力图表" tags:"天气监测"`
	Id     int `json:"id" description:"主键ID" v:"required#ID不能为空"`
	Types  int `json:"types" description:"类型 1 日 2周 3月 4年" v:"required#类型不能为空"`
}
type GetCityWindpowerByIdRes struct {
	Info            []*model.CityWeatherEchartRes
	AvgInfo         []*model.CityWeatherEchartRes
	ForeCastInfo    []*model.CityWeatherEchartRes
	ForeCastAvgInfo []*model.CityWeatherEchartRes
}
