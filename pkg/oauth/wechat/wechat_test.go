package wechat

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"sagooiot/pkg/oauth"
)

func Test_New(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	p := provider()

	a.Equal(p.ClientID, os.Getenv("WECHAT_KEY"))
	a.Equal(p.ClientSecret, os.Getenv("WECHAT_SECRET"))
	a.Equal(p.RedirectURL, "/foo")
}

func Test_Implements_Provider(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	a.Implements((*oauth.Provider)(nil), provider())
}

func Test_BeginAuth(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	p := provider()
	session, err := p.BeginAuth("test_state")
	s := session.(*Session)
	a.NoError(err)
	a.Contains(s.AuthURL, "open.weixin.qq.com/connect/qrconnect")
}

func Test_SessionFromJSON(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	p := provider()
	session, err := p.UnmarshalSession(`{"AuthURL":"https://open.weixin.qq.com/connect/qrconnect","AccessToken":"1234567890"}`)
	a.NoError(err)

	s := session.(*Session)
	a.Equal(s.AuthURL, "https://open.weixin.qq.com/connect/qrconnect")
	a.Equal(s.AccessToken, "1234567890")
}

func provider() *Provider {
	return New(os.Getenv("WECHAT_KEY"), os.Getenv("WECHAT_SECRET"), "/foo", WECHAT_LANG_CN)
}
