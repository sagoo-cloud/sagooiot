package model

// PluginInfo 插件信息结构
type PluginInfo struct {
	Types       string   `json:"types"`
	HandleType  string   `json:"handleType"`
	Name        string   `json:"name"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Version     string   `json:"version"`
	Author      string   `json:"author"`
	Icon        string   `json:"icon"`
	Link        string   `json:"link"`
	Command     string   `json:"command"`
	Args        []string `json:"args"`
	Root        bool     `json:"root"`
	Frontend    Frontend `json:"frontend"`
}

// Frontend 插件前端配置
type Frontend struct {
	Ui            bool   `json:"ui"`
	Url           string `json:"url"`
	Configuration bool   `json:"configuration"`
}
