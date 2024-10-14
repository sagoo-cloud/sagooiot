package oauth

import (
	"context"
	oauthV1 "sagooiot/api/v1/oauth"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"
)

type cProvider struct{}

var OProvider = cProvider{}

func (c *cProvider) ListProvider(ctx context.Context, req *oauthV1.ListProviderReq) (*oauthV1.ListProviderRes, error) {
	providers, err := service.OauthProvider().List(ctx, &entity.OauthProvider{})
	if err != nil {
		return nil, err
	}
	providersRes := []*oauthV1.Provider{}
	for _, value := range providers {
		providersRes = append(providersRes, &oauthV1.Provider{
			Name:   value.Name,
			Logo:   value.Logo,
			Status: value.Status,
		})
	}
	return &oauthV1.ListProviderRes{
		Providers: providersRes,
	}, nil
}
