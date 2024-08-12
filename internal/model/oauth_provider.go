package model

import "sagooiot/internal/model/entity"

type OauthProviderInput struct {
	entity.OauthProvider
}
type OauthProviderOutput struct {
	*entity.OauthProvider
}
type OauthProviderListOutput struct {
	List []*entity.OauthProvider `json:"list"`
}
