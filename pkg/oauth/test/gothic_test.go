package test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"sagooiot/pkg/oauth"
	"sagooiot/pkg/oauth/wechat"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetAuthURL(t *testing.T) {
	a := assert.New(t)

	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/auth?provider=wechat", nil)
	a.NoError(err)

	u, err := oauth.GetAuthURL(res, req)
	a.NoError(err)

	// Check that we get the correct auth URL with a state parameter
	parsed, err := url.Parse(u)
	a.NoError(err)
	a.Equal("http", parsed.Scheme)
	a.Equal("example.com", parsed.Host)
	q := parsed.Query()
	a.Contains(q, "client_id")
	a.Equal("code", q.Get("response_type"))
	a.NotZero(q, "state")

	// Check that if we run GetAuthURL on another request, that request's
	// auth URL has a different state from the previous one.
	req2, err := http.NewRequest("GET", "/auth?provider=faux", nil)
	a.NoError(err)
	url2, err := oauth.GetAuthURL(httptest.NewRecorder(), req2)
	a.NoError(err)
	parsed2, err := url.Parse(url2)
	a.NoError(err)
	a.NotEqual(parsed.Query().Get("state"), parsed2.Query().Get("state"))
}

var (
	WeChatAppID     = "wx094d1e4d2dff3414"
	WeChatAppSecret = "8c6a78808d9ec58814f481b65efd9bac"
	CallbackURL     = "http://kefu.huwaishequ.com/oauth/wechat/callback"
)

func Test_CompleteUserAuth(t *testing.T) {
	a := assert.New(t)
	oauth.UseProviders(wechat.New(WeChatAppID, WeChatAppSecret, CallbackURL, wechat.WECHAT_LANG_CN))
	res := httptest.NewRecorder()
	code := "code=011sFQkl2WLRKd4jwIml2r3Pkw1sFQkx&state=mGP1eRnFsiKCp4YoDD6tatPsgmj5XWA103A31twT9JFy64XGIzoUJFrNpctqcLrwgkvZ75VFRCcut6u3ofSeog=="
	req, err := http.NewRequest("GET", "/auth/callback?provider=wechat&"+code, nil)
	a.NoError(err)

	a.NoError(err)

	user, err := oauth.CompleteUserAuth(res, req)
	a.NoError(err)

	a.Equal(user.Name, "Homer Simpson")
	a.Equal(user.Email, "homer@example.com")
}
