package envirotronics

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"github.com/sagoo-cloud/sagooiot/utility/utils"
	"regexp"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sEnvWeather struct {
}

func EnvWeather() *sEnvWeather {
	return &sEnvWeather{}
}

func init() {
	service.RegisterEnvWeather(EnvWeather())
}

// CityWeatherList 获取城市的风力及日照时长
func (a *sEnvWeather) CityWeatherList(ctx context.Context) (cityWeatherListOut []*model.CityWeatherListOut, err error) {
	info, err := service.CityData().GetAll(ctx)
	if err != nil {
		return
	}
	if info != nil {
		var cityIds []int
		for _, city := range info {
			cityIds = append(cityIds, city.Id)
		}
		if len(cityIds) > 0 {
			//绑定的数据建模
			dataInfo, _ := service.DataTemplateBusi().GetInfos(ctx, consts.Weather)
			if dataInfo != nil {
				//查询数据建模数据
				var in = new(model.TemplateDataAllInput)
				in.Id = uint64(dataInfo.DataTemplateId)
				//查询条件
				params := map[string]interface{}{
					"CREATED_AT >= ? ": gtime.Now().Format("Y-m-d 00:00:00"),
					"CREATED_AT < ? ":  gtime.Now().Format("Y-m-d H:i:s"),
				}
				in.Param = params
				out, _ := service.DataTemplate().GetAllData(ctx, in)

				for _, city := range info {
					var cityWeather = new(model.CityWeatherListOut)
					cityWeather.Id = city.Id
					cityWeather.Name = city.Name
					cityWeather.Code = city.Code
					//获取当前城市当天城市数据
					var cityOut g.List
					for _, v := range out.List {
						if strings.EqualFold(city.Code, v["adcode"].(string)) {
							cityOut = append(cityOut, v)
						}
					}
					if len(cityOut) > 0 {
						if err = gconv.Scan(cityOut[len(cityOut)-1], &cityWeather); err != nil {
							return
						}
					}
					//计算日照时长
					if cityWeather.Sunrise != "" && cityWeather.Sunset != "" {
						cityWeather.SunshineDuration = int(gtime.New(gtime.Now().Format("Y-m-d ") + cityWeather.Sunset + ":00").Sub(gtime.New(gtime.Now().Format("Y-m-d ") + cityWeather.Sunrise + ":00")).Hours())
					}

					cityWeatherListOut = append(cityWeatherListOut, cityWeather)
				}
			}
		}
	}
	return
}

// GetCityWeatherById 根据ID获取指定城市的天气
func (a *sEnvWeather) GetCityWeatherById(ctx context.Context, id int) (cityWeatherListOut *model.CityWeatherListOut, err error) {
	info, err := service.CityData().GetInfoById(ctx, id)
	if err != nil {
		return
	}
	cityWeatherListOut = new(model.CityWeatherListOut)
	if info != nil {
		//绑定的数据建模
		dataInfo, _ := service.DataTemplateBusi().GetInfo(ctx, consts.Weather)
		if dataInfo != nil {
			cityWeatherListOut.Id = info.Id
			cityWeatherListOut.Name = info.Name
			cityWeatherListOut.Code = info.Code

			var in = new(model.TemplateDataLastInput)
			in.Id = uint64(dataInfo.DataTemplateId)
			//查询条件
			params := map[string]interface{}{
				"CREATED_AT >= ? ": gtime.Now().Format("Y-m-d 00:00:00"),
				"CREATED_AT < ? ":  gtime.Now().Format("Y-m-d H:i:s"),
				"ADCODE = ? ":      info.Code,
			}
			in.Param = params
			//获取数据建模数据
			out, _ := service.DataTemplate().GetLastData(ctx, in)
			if out.Data != nil {
				if err = gconv.Scan(out.Data, &cityWeatherListOut); err != nil {
					return
				}
			}

			//计算日照时长
			if cityWeatherListOut.Sunrise != "" && cityWeatherListOut.Sunset != "" {
				cityWeatherListOut.SunshineDuration = int(gtime.New(gtime.Now().Format("Y-m-d ") + cityWeatherListOut.Sunset + ":00").Sub(gtime.New(gtime.Now().Format("Y-m-d ") + cityWeatherListOut.Sunrise + ":00")).Hours())
			}
		}
	}
	return
}

// GetCityTemperatureById 根据ID获取指定城市的温度图表
func (a *sEnvWeather) GetCityTemperatureById(ctx context.Context, id int, types int) (cityWeatherEchartOut []*model.CityWeatherEchartOut, avgCityWeatherEchartOut []*model.CityWeatherEchartOut, foreCastCityWeatherEchartOut []*model.CityWeatherEchartOut, foreCastAvgCityWeatherEchartOut []*model.CityWeatherEchartOut, err error) {
	info, err := service.CityData().GetInfoById(ctx, id)
	if err != nil {
		return
	}
	if info != nil {
		//绑定的数据建模
		dataInfo, _ := service.DataTemplateBusi().GetInfo(ctx, consts.Weather)
		if dataInfo != nil {
			//根据类型获取时间差、开始时间、结束时间
			index, begin, end := utils.GetTimeByType(types)

			//封装数据建模
			var in = new(model.TemplateDataAllInput)
			//数据建模ID
			in.Id = uint64(dataInfo.DataTemplateId)
			//查询条件
			params := map[string]interface{}{
				"CREATED_AT >= ? ": begin,
				"CREATED_AT < ? ":  end,
				"ADCODE = ? ":      info.Code,
			}
			in.Param = params

			//获取数据建模数据
			out, _ := service.DataTemplate().GetAllData(ctx, in)

			var sumValue float64
			var foreCastSumValue float64
			for i := 0; i < index; i++ {
				//实时天气
				var cityWeatherEchart = new(model.CityWeatherEchartOut)
				startTime, endTime, duration, unit := utils.GetTime(i, types, begin)
				cityWeatherEchart.Time = strconv.Itoa(duration) + unit
				var value float64
				var num float64
				for _, m := range out.List {
					if gtime.New(startTime).Before(m["created_at"].(*gtime.Time)) && m["created_at"].(*gtime.Time).Before(gtime.New(endTime)) {
						parseValue, _ := strconv.ParseFloat(m["temperature"].(string), 64)
						value += parseValue
						num++
					}
				}
				if value != 0 && num != 0 {

					cityWeatherEchart.Value = fmt.Sprintf("%.2f", value/num)
					sumValue += value / num
				} else {
					cityWeatherEchart.Value = "0.00"
				}
				cityWeatherEchartOut = append(cityWeatherEchartOut, cityWeatherEchart)

				if types == 2 || types == 3 {
					//预报天气
					var foreCastCityWeatherEchart = new(model.CityWeatherEchartOut)
					foreCastCityWeatherEchart.Time = strconv.Itoa(duration) + unit

					var foreCastValue float64
					var foreCastNum float64
					//开始时间-1天
					startTime = gtime.New(startTime).AddDate(0, 0, -1).Format("Y-m-d H:i:s")
					//结束时间-1天
					endTime = gtime.New(endTime).AddDate(0, 0, -1).Format("Y-m-d H:i:s")
					for _, m := range out.List {
						if gtime.New(startTime).Before(m["created_at"].(*gtime.Time)) && m["created_at"].(*gtime.Time).Before(gtime.New(endTime)) {
							parseValue, _ := strconv.ParseFloat(m["next_day_temp"].(string), 64)
							foreCastValue += parseValue
							foreCastNum++
						}
					}
					if foreCastValue != 0 && foreCastNum != 0 {
						foreCastCityWeatherEchart.Value = fmt.Sprintf("%.2f", foreCastValue/foreCastNum)
						foreCastSumValue += foreCastValue / foreCastNum
					} else {
						foreCastCityWeatherEchart.Value = "0.00"
					}
					foreCastCityWeatherEchartOut = append(foreCastCityWeatherEchartOut, foreCastCityWeatherEchart)
				}
			}

			//获取未来一周的时间
			if types == 2 || types == 3 {
				//获取最后一条数据
				var lastIn = new(model.TemplateDataLastInput)
				lastIn.Id = uint64(dataInfo.DataTemplateId)
				//查询条件
				lastParams := map[string]interface{}{
					"CREATED_AT >= ? ": gtime.Now().Format("Y-m-d 00:00:00"),
					"CREATED_AT < ? ":  gtime.Now().Format("Y-m-d H:i:s"),
					"ADCODE = ? ":      info.Code,
				}
				lastIn.Param = lastParams

				lastOut, _ := service.DataTemplate().GetLastData(ctx, lastIn)
				if lastOut.Data != nil {
					for i := 1; i < 7; i++ {
						//获取字段名称
						var field string
						switch i {
						case 1:
							field = "next_day_temp"
							break
						case 2:
							field = "next_three_day_temp"
							break
						case 3:
							field = "next_four_day_temp"
							break
						case 4:
							field = "next_five_day_temp"
							break
						case 5:
							field = "next_six_day_temp"
							break
						case 6:
							field = "next_seven_day_temp"
							break
						}
						parseValue, _ := strconv.ParseFloat(gconv.String(lastOut.Data[field]), 64)
						var cityWeatherEchart = new(model.CityWeatherEchartOut)
						cityWeatherEchart.Time = strconv.Itoa(gtime.New(gtime.Now()).AddDate(0, 0, i).Day()) + "日"
						cityWeatherEchart.Value = fmt.Sprintf("%.2f", parseValue)
						foreCastCityWeatherEchartOut = append(foreCastCityWeatherEchartOut, cityWeatherEchart)
						foreCastSumValue += parseValue
					}
				}
			}

			//获取实时天气平均值
			for i := 0; i < index; i++ {
				var cityWeatherEchart = new(model.CityWeatherEchartOut)
				_, _, duration, unit := utils.GetTime(i, types, begin)
				cityWeatherEchart.Time = strconv.Itoa(duration) + unit
				if sumValue != 0 {
					cityWeatherEchart.Value = fmt.Sprintf("%.2f", sumValue/float64(index))
				} else {
					cityWeatherEchart.Value = "0.00"
				}
				avgCityWeatherEchartOut = append(avgCityWeatherEchartOut, cityWeatherEchart)
				//获取预报天气值
				if types == 2 || types == 3 {
					var foreCastCityWeatherEchart = new(model.CityWeatherEchartOut)
					foreCastCityWeatherEchart.Time = strconv.Itoa(duration) + unit
					if foreCastSumValue != 0 {
						foreCastCityWeatherEchart.Value = fmt.Sprintf("%.2f", foreCastSumValue/float64(index+6))
					} else {
						foreCastCityWeatherEchart.Value = "0.00"
					}
					foreCastAvgCityWeatherEchartOut = append(foreCastAvgCityWeatherEchartOut, foreCastCityWeatherEchart)
				}
			}
			//获取未来一周的平均天气
			if types == 2 || types == 3 {
				for i := 1; i < 7; i++ {
					var cityWeatherEchart = new(model.CityWeatherEchartOut)
					cityWeatherEchart.Time = strconv.Itoa(gtime.New(gtime.Now()).AddDate(0, 0, i).Day()) + "日"
					cityWeatherEchart.Value = fmt.Sprintf("%.2f", foreCastSumValue/float64(gtime.Now().Day()+6))
					foreCastAvgCityWeatherEchartOut = append(foreCastAvgCityWeatherEchartOut, cityWeatherEchart)
				}
			}
		}
	}
	return
}

// GetCityWindpowerById 根据ID获取指定城市的风力图表
func (a *sEnvWeather) GetCityWindpowerById(ctx context.Context, id int, types int) (cityWeatherEchartOut []*model.CityWeatherEchartOut, avgCityWeatherEchartOut []*model.CityWeatherEchartOut, foreCastCityWeatherEchartOut []*model.CityWeatherEchartOut, foreCastAvgCityWeatherEchartOut []*model.CityWeatherEchartOut, err error) {
	info, err := service.CityData().GetInfoById(ctx, id)
	if err != nil {
		return
	}
	if info != nil {
		//绑定的数据建模
		dataInfo, _ := service.DataTemplateBusi().GetInfo(ctx, consts.Weather)
		if dataInfo != nil {
			//根据类型获取时间差、开始时间、结束时间
			index, begin, end := utils.GetTimeByType(types)

			//封装数据建模
			var in = new(model.TemplateDataAllInput)
			//数据建模ID
			in.Id = uint64(dataInfo.DataTemplateId)
			//查询条件
			params := map[string]interface{}{
				"CREATED_AT >= ? ": begin,
				"CREATED_AT < ? ":  end,
				"ADCODE = ? ":      info.Code,
			}
			in.Param = params

			//获取数据建模数据
			out, _ := service.DataTemplate().GetAllData(ctx, in)

			var sumValue float64
			var foreCastSumValue float64
			//获取当前数据
			for i := 0; i < index; i++ {
				var cityWeatherEchart = new(model.CityWeatherEchartOut)
				startTime, endTime, duration, unit := utils.GetTime(i, types, begin)
				cityWeatherEchart.Time = strconv.Itoa(duration) + unit
				var value float64
				var num float64

				for _, m := range out.List {
					if gtime.New(startTime).Before(m["created_at"].(*gtime.Time)) && m["created_at"].(*gtime.Time).Before(gtime.New(endTime)) {
						re := regexp.MustCompile("[0-9]+")
						if re.FindAllString(m["windpower"].(string), -1) != nil {
							parseValue, _ := strconv.ParseFloat(re.FindAllString(m["windpower"].(string), -1)[0], 64)
							value += parseValue
						} else {
							value += 0.5
						}
						num++
					}
				}
				if value != 0 && num != 0 {
					cityWeatherEchart.Value = fmt.Sprintf("%.2f", value/num)
					sumValue += value / num
				} else {
					cityWeatherEchart.Value = "0.00"
				}
				cityWeatherEchartOut = append(cityWeatherEchartOut, cityWeatherEchart)

				if types == 2 || types == 3 {
					//预报天气
					var foreCastCityWeatherEchart = new(model.CityWeatherEchartOut)
					foreCastCityWeatherEchart.Time = strconv.Itoa(duration) + unit

					var foreCastValue float64
					var foreCastNum float64
					//开始时间-1天
					startTime = gtime.New(startTime).AddDate(0, 0, -1).Format("Y-m-d H:i:s")
					//结束时间-1天
					endTime = gtime.New(endTime).AddDate(0, 0, -1).Format("Y-m-d H:i:s")
					for _, m := range out.List {
						if gtime.New(startTime).Before(m["created_at"].(*gtime.Time)) && m["created_at"].(*gtime.Time).Before(gtime.New(endTime)) {
							re := regexp.MustCompile("[0-9]+")
							if re.FindAllString(m["next_day_windpower"].(string), -1) != nil {
								parseValue, _ := strconv.ParseFloat(re.FindAllString(m["next_day_windpower"].(string), -1)[0], 64)
								foreCastValue += parseValue
							} else {
								foreCastValue += 0.5
							}
							foreCastNum++
						}
					}
					if foreCastValue != 0 && foreCastNum != 0 {
						foreCastCityWeatherEchart.Value = fmt.Sprintf("%.2f", foreCastValue/foreCastNum)
						foreCastSumValue += foreCastValue / foreCastNum
					} else {
						foreCastCityWeatherEchart.Value = "0.00"
					}
					foreCastCityWeatherEchartOut = append(foreCastCityWeatherEchartOut, foreCastCityWeatherEchart)
				}
			}

			//获取未来一周的时间
			if types == 2 || types == 3 {
				//获取最后一条数据
				var lastIn = new(model.TemplateDataLastInput)
				lastIn.Id = uint64(dataInfo.DataTemplateId)
				//查询条件
				lastParams := map[string]interface{}{
					"CREATED_AT >= ? ": gtime.Now().Format("Y-m-d 00:00:00"),
					"CREATED_AT < ? ":  gtime.Now().Format("Y-m-d H:i:s"),
					"ADCODE = ? ":      info.Code,
				}
				lastIn.Param = lastParams

				lastOut, _ := service.DataTemplate().GetLastData(ctx, lastIn)
				if lastOut.Data != nil {
					for i := 1; i < 7; i++ {
						//获取字段名称
						var field string
						switch i {
						case 1:
							field = "next_day_windpower"
							break
						case 2:
							field = "next_three_day_windpower"
							break
						case 3:
							field = "next_four_day_windpower"
							break
						case 4:
							field = "next_five_day_windpower"
							break
						case 5:
							field = "next_six_day_windpower"
							break
						case 6:
							field = "next_seven_day_windpower"
							break
						}
						var cityWeatherEchart = new(model.CityWeatherEchartOut)
						cityWeatherEchart.Time = strconv.Itoa(gtime.New(gtime.Now()).AddDate(0, 0, i).Day()) + "日"
						re := regexp.MustCompile("[0-9]+")
						var parseValue float64
						if re.FindAllString(gconv.String(lastOut.Data[field]), -1) != nil {
							parseValue, _ = strconv.ParseFloat(re.FindAllString(gconv.String(lastOut.Data[field]), -1)[0], 64)
						} else {
							parseValue += 0.5
						}
						cityWeatherEchart.Value = fmt.Sprintf("%.2f", parseValue)
						foreCastCityWeatherEchartOut = append(foreCastCityWeatherEchartOut, cityWeatherEchart)
						foreCastSumValue += parseValue
					}
				}
			}

			//获取平均值
			for i := 0; i < index; i++ {
				var cityWeatherEchart = new(model.CityWeatherEchartOut)

				_, _, duration, unit := utils.GetTime(i, types, begin)
				cityWeatherEchart.Time = strconv.Itoa(duration) + unit

				if sumValue != 0 {
					cityWeatherEchart.Value = fmt.Sprintf("%.2f", sumValue/float64(index))
				} else {
					cityWeatherEchart.Value = "0.00"
				}
				avgCityWeatherEchartOut = append(avgCityWeatherEchartOut, cityWeatherEchart)

				//获取预报天气值
				if types == 2 || types == 3 {
					var foreCastCityWeatherEchart = new(model.CityWeatherEchartOut)
					foreCastCityWeatherEchart.Time = strconv.Itoa(duration) + unit
					if sumValue != 0 {
						foreCastCityWeatherEchart.Value = fmt.Sprintf("%.2f", foreCastSumValue/float64(index+6))
					} else {
						foreCastCityWeatherEchart.Value = "0.00"
					}
					foreCastAvgCityWeatherEchartOut = append(foreCastAvgCityWeatherEchartOut, foreCastCityWeatherEchart)
				}
			}

			//获取未来一周的平均天气
			if types == 2 || types == 3 {
				for i := 1; i < 7; i++ {
					var cityWeatherEchart = new(model.CityWeatherEchartOut)
					cityWeatherEchart.Time = strconv.Itoa(gtime.New(gtime.Now()).AddDate(0, 0, i).Day()) + "日"
					cityWeatherEchart.Value = fmt.Sprintf("%.2f", foreCastSumValue/float64(gtime.Now().Day()+6))
					foreCastAvgCityWeatherEchartOut = append(foreCastAvgCityWeatherEchartOut, cityWeatherEchart)
				}
			}
		}
	}
	return
}
