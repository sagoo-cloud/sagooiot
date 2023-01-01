package model

type CityWeatherListOut struct {
	Id               int    `json:"id"        description:""`
	Name             string `json:"name"      description:"名字"`
	Code             string `json:"code"      description:"编码"`
	Windpower        string `json:"windpower"      description:"风力级别"`
	Sunrise          string `json:"sunrise"      description:"日出"`
	Sunset           string `json:"sunset"      description:"日落"`
	SunshineDuration int    `json:"sunshineDuration"      description:"日照时长"`
	Temperature      int    `json:"Temperature"      description:"气温"`
	Weather          string `json:"weather"      description:"天气现象"`
	Winddirection    string `json:"winddirection"      description:"风向描述"`
	Reporttime       string `json:"reporttime"      description:"发布时间"`
}

type CityWeatherListRes struct {
	Id               int    `json:"id"        description:""`
	Name             string `json:"name"      description:"名字"`
	Code             string `json:"code"      description:"编码"`
	Windpower        string `json:"windpower"      description:"风力级别"`
	Sunrise          string `json:"sunrise"      description:"日出"`
	Sunset           string `json:"sunset"      description:"日落"`
	SunshineDuration int    `json:"sunshineDuration"      description:"日照时长"`
	Temperature      int    `json:"Temperature"      description:"气温"`
	Weather          string `json:"weather"      description:"天气现象"`
	Winddirection    string `json:"winddirection"      description:"风向描述"`
	Reporttime       string `json:"reporttime"      description:"发布时间"`
}

type CityWeatherEchartRes struct {
	Value string `json:"value"      description:"值"`
	Time  string `json:"time"      description:"时间"`
}

type CityWeatherEchartOut struct {
	Value string `json:"value"      description:"值"`
	Time  string `json:"time"      description:"时间"`
}
