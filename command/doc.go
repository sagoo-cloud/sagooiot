package command

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
	"sagooiot/internal/consts"
)

func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
	openapi.Config.CommonResponseDataField = `Data`

	// API description.
	openapi.Info.Title = `sagooAdmin Project`
	openapi.Info.Description = ``

	// Sort the tags in custom sequence.
	openapi.Tags = &goai.Tags{
		{Name: consts.OpenAPITagNameLogin},
		{Name: consts.OpenAPITagNameOrganization},
		{Name: consts.OpenAPITagNameDept},
		{Name: consts.OpenAPITagNamePost},
		{Name: consts.OpenAPITagNameRole},
		{Name: consts.OpenAPITagNameUser},
		{Name: consts.OpenAPITagNameMenu},
		{Name: consts.OpenAPITagNameApi},
		{Name: consts.OpenAPITagNameAuthorize},
	}
}
