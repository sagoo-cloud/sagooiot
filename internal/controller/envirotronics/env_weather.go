package envirotronics

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/api/v1/envirotronics"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
)

var Weather = cWeather{}

type cWeather struct{}

// CityWeatherList 获取城市的风力及日照时长
func (a *cWeather) CityWeatherList(ctx context.Context, req *envirotronics.CityWeatherListReq) (res *envirotronics.CityWeatherListRes, err error) {
	out, err := service.EnvWeather().CityWeatherList(ctx)
	if err != nil {
		return
	}
	var cityWeatherListRes []*model.CityWeatherListRes
	if out != nil {
		if err = gconv.Scan(out, &cityWeatherListRes); err != nil {
			return
		}
	}
	res = &envirotronics.CityWeatherListRes{
		Info: cityWeatherListRes,
	}
	return
}

// GetCityWeatherById 根据ID获取指定城市的天气
func (a *cWeather) GetCityWeatherById(ctx context.Context, req *envirotronics.GetCityWeatherByIdReq) (res *envirotronics.GetCityWeatherByIdRes, err error) {
	out, err := service.EnvWeather().GetCityWeatherById(ctx, req.Id)
	if err != nil {
		return
	}
	var cityWeatherListRes *model.CityWeatherListRes
	if out != nil {
		if err = gconv.Scan(out, &cityWeatherListRes); err != nil {
			return
		}
	}

	res = &envirotronics.GetCityWeatherByIdRes{
		Info: cityWeatherListRes,
	}
	return
}

// GetCityTemperatureById 根据ID获取指定城市的温度图表
func (a *cWeather) GetCityTemperatureById(ctx context.Context, req *envirotronics.GetCityTemperatureByIdReq) (res *envirotronics.GetCityTemperatureByIdRes, err error) {
	cityWeatherEchartOut, avgCityWeatherEchartOut, foreCastCityWeatherHighEchartOut, foreCastCityWeatherLowEchartOut, err := service.EnvWeather().GetCityTemperatureById(ctx, req.Id, req.Types)
	if err != nil {
		return
	}
	var cityWeatherEchartRes []*model.CityWeatherEchartRes
	if cityWeatherEchartOut != nil {
		if err = gconv.Scan(cityWeatherEchartOut, &cityWeatherEchartRes); err != nil {
			return
		}
	}
	var avgCityWeatherEchartRes []*model.CityWeatherEchartRes
	if avgCityWeatherEchartOut != nil {
		if err = gconv.Scan(avgCityWeatherEchartOut, &avgCityWeatherEchartRes); err != nil {
			return
		}
	}
	var foreCastCityWeatherHighEchartRes []*model.CityWeatherEchartRes
	if foreCastCityWeatherHighEchartOut != nil {
		if err = gconv.Scan(foreCastCityWeatherHighEchartOut, &foreCastCityWeatherHighEchartRes); err != nil {
			return
		}
	}
	var foreCastCityWeatherLowEchartRes []*model.CityWeatherEchartRes
	if foreCastCityWeatherLowEchartOut != nil {
		if err = gconv.Scan(foreCastCityWeatherLowEchartOut, &foreCastCityWeatherLowEchartRes); err != nil {
			return
		}
	}
	res = &envirotronics.GetCityTemperatureByIdRes{
		Info:             cityWeatherEchartRes,
		AvgInfo:          avgCityWeatherEchartRes,
		ForeCastHighInfo: foreCastCityWeatherHighEchartRes,
		ForeCastLowInfo:  foreCastCityWeatherLowEchartRes,
	}
	return
}

// GetCityWindpowerById 根据ID获取指定城市的风力图表
func (a *cWeather) GetCityWindpowerById(ctx context.Context, req *envirotronics.GetCityWindpowerByIdReq) (res *envirotronics.GetCityWindpowerByIdRes, err error) {
	cityWeatherEchartOut, avgCityWeatherEchartOut, foreCastCityWeatherEchartOut, foreCastAvgCityWeatherEchartOut, err := service.EnvWeather().GetCityWindpowerById(ctx, req.Id, req.Types)
	if err != nil {
		return
	}
	var cityWeatherEchartRes []*model.CityWeatherEchartRes
	if cityWeatherEchartOut != nil {
		if err = gconv.Scan(cityWeatherEchartOut, &cityWeatherEchartRes); err != nil {
			return
		}
	}
	var avgCityWeatherEchartRes []*model.CityWeatherEchartRes
	if avgCityWeatherEchartOut != nil {
		if err = gconv.Scan(avgCityWeatherEchartOut, &avgCityWeatherEchartRes); err != nil {
			return
		}
	}
	var foreCastCityWeatherEchartRes []*model.CityWeatherEchartRes
	if foreCastCityWeatherEchartOut != nil {
		if err = gconv.Scan(foreCastCityWeatherEchartOut, &foreCastCityWeatherEchartRes); err != nil {
			return
		}
	}
	var foreCastAvgCityWeatherEchartRes []*model.CityWeatherEchartRes
	if foreCastAvgCityWeatherEchartOut != nil {
		if err = gconv.Scan(foreCastAvgCityWeatherEchartOut, &foreCastAvgCityWeatherEchartRes); err != nil {
			return
		}
	}
	res = &envirotronics.GetCityWindpowerByIdRes{
		Info:            cityWeatherEchartRes,
		AvgInfo:         avgCityWeatherEchartRes,
		ForeCastInfo:    foreCastCityWeatherEchartRes,
		ForeCastAvgInfo: foreCastAvgCityWeatherEchartRes,
	}
	return
}

// GetCityWeatherHistoryById 根据ID获取指定城市的天气历史数据
func (a *cWeather) GetCityWeatherHistoryById(ctx context.Context, req *envirotronics.GetCityWeatherHistoryByIdReq) (res *envirotronics.GetCityWeatherHistoryByIdRes, err error) {
	out, err := service.EnvWeather().GetCityWeatherHistoryById(ctx, req.Id, req.DateTime)
	if err != nil {
		return
	}
	var cityWeatherHistoryListRes []*model.CityWeatherHistoryListRes
	if out != nil {
		if err = gconv.Scan(out, &cityWeatherHistoryListRes); err != nil {
			return
		}
	}

	res = &envirotronics.GetCityWeatherHistoryByIdRes{
		Info: cityWeatherHistoryListRes,
	}
	return
}

// GetCityWeatherHistoryByIdExport 根据ID获取指定城市的天气历史数据导出
func (a *cWeather) GetCityWeatherHistoryByIdExport(ctx context.Context, req *envirotronics.GetCityWeatherHistoryByIdExportReq) (res *envirotronics.GetCityWeatherHistoryByIdExportRes, err error) {
	err = service.EnvWeather().GetCityWeatherHistoryByIdExport(ctx, req.Id, req.DateTime)
	if err != nil {
		return
	}
	return
}
