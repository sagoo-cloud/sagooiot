package model

type Base struct {
	Touser                 string `json:"touser"`
	Toparty                string `json:"toparty"`
	Totag                  string `json:"totag"`
	Msgtype                string `json:"msgtype"`
	Agentid                string `json:"agentid"`
	Safe                   int    `json:"safe"`
	EnableIDTrans          int    `json:"enable_id_trans"`
	EnableDuplicateCheck   int    `json:"enable_duplicate_check"`
	DuplicateCheckInterval int    `json:"duplicate_check_interval"`
}

//文本消息结构体
type Text struct {
	Base
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

//文本卡片消息结构
type Textcard struct {
	Base
	Textcard struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		URL         string `json:"url"`
		Btntxt      string `json:"btntxt"`
	} `json:"textcard"`
}
