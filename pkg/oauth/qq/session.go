package qq

import (
	"encoding/json"
	"errors"
	"sagooiot/pkg/oauth"
	"strings"
	"time"
)

// Session 在 qq 认证过程中存储数据
type Session struct {
	AuthURL      string
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
	Openid       string
	Unionid      string
}

var _ oauth.Session = &Session{}

// GetAuthURL 将通过调用Wepay提供程序上的' BeginAuth '函数返回URL设置
func (s Session) GetAuthURL() (string, error) {
	if s.AuthURL == "" {
		return "", errors.New(oauth.NoAuthUrlErrorMessage)
	}
	return s.AuthURL, nil
}

// Authorize 使用 Wepay 进行会话，并返回要存储的访问令牌以供将来使用。
func (s *Session) Authorize(provider oauth.Provider, params oauth.Params) (string, error) {
	p := provider.(*Provider)
	token, openid, err := p.fetchToken(params.Get("code"))

	if err != nil {
		return "", err
	}

	if !token.Valid() {
		return "", errors.New("invalid token received from provider")
	}

	s.AccessToken = token.AccessToken
	s.RefreshToken = token.RefreshToken
	s.ExpiresAt = token.Expiry
	s.Openid = openid
	return token.AccessToken, err
}

// Marshal the session into a string
func (s Session) Marshal() string {
	b, _ := json.Marshal(s)
	return string(b)
}

func (s Session) String() string {
	return s.Marshal()
}

// UnmarshalSession wil unmarshal a JSON string into a session.
func (p *Provider) UnmarshalSession(data string) (oauth.Session, error) {
	s := &Session{AuthURL: p.AuthURL}
	var err error
	if len(data) > 0 {
		err = json.NewDecoder(strings.NewReader(data)).Decode(s)
	}
	return s, err
}
