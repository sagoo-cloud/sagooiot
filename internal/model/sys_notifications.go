package model

type GetNotificationsListInput struct {
	*PaginationInput
}

type NotificationsRes struct {
	Id int `json:"id"          description:"ID"`
}

type NotificationsOut struct {
	Id int `json:"id"          description:"ID"`
}

type NotificationsAddInput struct {
	Title     string `json:"title"          description:"标题"`
	Doc       string `json:"doc"          description:"描述"`
	Source    string `json:"source"          description:"消息来源"`
	Types     string `json:"types"          description:"类型"`
	CreatedAt string `json:"createdAt"          description:"发送时间"`
	Status    string `json:"status"          description:"0，未读，1，已读"`
}
type NotificationsEditInput struct {
	Id int `json:"id"          description:"ID"`
	NotificationsAddInput
}
