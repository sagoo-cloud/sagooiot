package qq

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/oauth2"
	"sagooiot/pkg/oauth"
)

const (
	AuthURL  = "https://graph.qq.com/oauth2.0/authorize"
	TokenURL = "https://graph.qq.com/oauth2.0/token"

	ScopeSnsapiLogin = "get_user_info"

	ProfileURL = "https://graph.qq.com/user/get_user_info"
)

type Provider struct {
	providerName string
	config       *oauth2.Config
	httpClient   *http.Client
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Lang         QQLangType

	AuthURL    string
	TokenURL   string
	ProfileURL string
}

type QQLangType string

const (
	QQ_LANG_CN QQLangType = "cn"
	QQ_LANG_EN QQLangType = "en"
)

// 实例化一个新的 qq 提供商，并设置重要的连接必要参数。
func New(clientID, clientSecret, redirectURL string, lang QQLangType) *Provider {
	p := &Provider{
		providerName: "qq",
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Lang:         lang,
		AuthURL:      AuthURL,
		TokenURL:     TokenURL,
		ProfileURL:   ProfileURL,
	}
	p.config = newConfig(p)
	return p
}

// 返回 授权 对象名
func (p *Provider) Name() string {
	return p.providerName
}

// 更新提供授权的名称
func (p *Provider) SetName(name string) {
	p.providerName = name
}

func (p *Provider) Client() *http.Client {
	return oauth.HTTPClientWithFallBack(p.httpClient)
}

// BeginAuth向qq请求身份验证端点
func (p *Provider) BeginAuth(state string) (oauth.Session, error) {
	params := url.Values{}
	params.Add("client_id", p.ClientID)
	params.Add("response_type", "code")
	params.Add("state", state)
	params.Add("scope", ScopeSnsapiLogin)
	params.Add("redirect_uri", p.RedirectURL)
	session := &Session{
		AuthURL: fmt.Sprintf("%s?%s", p.AuthURL, params.Encode()),
	}
	return session, nil
}

// FetchUser 将转到 Wepay并访问用户的基本信息。
func (p *Provider) FetchUser(session oauth.Session) (oauth.User, error) {
	s := session.(*Session)
	user := oauth.User{
		AccessToken:  s.AccessToken,
		Provider:     p.Name(),
		RefreshToken: s.RefreshToken,
		ExpiresAt:    s.ExpiresAt,
	}

	if user.AccessToken == "" {
		// data is not yet retrieved since accessToken is still empty
		return user, fmt.Errorf("%s cannot get user information without accessToken", p.providerName)
	}

	params := url.Values{}
	params.Add("access_token", s.AccessToken)
	params.Add("openid", s.Openid)
	params.Add("lang", string(p.Lang))

	url := fmt.Sprintf("%s?%s", p.ProfileURL, params.Encode())

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return user, err
	}
	resp, err := p.Client().Do(req)
	if err != nil {
		if resp != nil {
			resp.Body.Close()
		}
		return user, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return user, fmt.Errorf("%s responded with a %d trying to fetch user information", p.providerName, resp.StatusCode)
	}

	err = userFromReader(resp.Body, &user)
	user.UserID = s.Openid
	return user, err
}

func newConfig(provider *Provider) *oauth2.Config {
	c := &oauth2.Config{
		ClientID:     provider.ClientID,
		ClientSecret: provider.ClientSecret,
		RedirectURL:  provider.RedirectURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:  provider.AuthURL,
			TokenURL: provider.TokenURL,
		},
		Scopes: []string{},
	}

	c.Scopes = append(c.Scopes, ScopeSnsapiLogin)

	return c
}

func userFromReader(r io.Reader, user *oauth.User) error {
	u := struct {
		Openid    string `json:"openid"`
		Nickname  string `json:"nickname"`
		Sex       int    `json:"gender"`
		Province  string `json:"province"`
		City      string `json:"city"`
		Country   string `json:"year"`
		AvatarURL string `json:"figureurl_qq_1"`
		Unionid   string `json:"unionid"`
		Code      int    `json:"ret"`
		Msg       string `json:"msg"`
	}{}
	err := json.NewDecoder(r).Decode(&u)
	if err != nil {
		return err
	}

	if len(u.Msg) > 0 {
		return fmt.Errorf("CODE: %d, MSG: %s", u.Code, u.Msg)
	}

	user.Email = fmt.Sprintf("%s@wechat.com", u.Openid)
	user.Name = u.Nickname
	user.UserID = u.Openid
	user.NickName = u.Nickname
	user.Location = u.City
	user.AvatarURL = u.AvatarURL
	user.RawData = map[string]interface{}{
		"Unionid": u.Unionid,
	}
	return nil
}

// RefreshTokenAvailable 刷新令牌
func (p *Provider) RefreshTokenAvailable() bool {
	return false
}

// RefreshToken 根据刷新令牌获取新的访问令牌
func (p *Provider) RefreshToken(refreshToken string) (*oauth2.Token, error) {
	return nil, nil
}

func (p *Provider) fetchToken(code string) (*oauth2.Token, string, error) {

	params := url.Values{}
	params.Add("client_id", p.ClientID)
	params.Add("client_secret", p.ClientSecret)
	params.Add("grant_type", "authorization_code")
	params.Add("code", code)
	params.Add("need_openid", "1")
	params.Add("redirect_uri", p.RedirectURL)
	url := fmt.Sprintf("%s?%s", p.TokenURL, params.Encode())
	resp, err := p.Client().Get(url)

	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("wechat /gettoken returns code: %d", resp.StatusCode)
	}
	obj := struct {
		AccessToken  string        `json:"access_token"`
		ExpiresIn    time.Duration `json:"expires_in"`
		Openid       string        `json:"openid"`
		Code         int           `json:"errcode"`
		Msg          string        `json:"errmsg"`
		RefreshToken string        `json:"refresh_token"`
	}{}
	if err = json.NewDecoder(resp.Body).Decode(&obj); err != nil {
		return nil, "", err
	}
	if obj.Code != 0 {
		return nil, "", fmt.Errorf("CODE: %d, MSG: %s", obj.Code, obj.Msg)
	}

	token := &oauth2.Token{
		AccessToken: obj.AccessToken,
		Expiry:      time.Now().Add(obj.ExpiresIn * time.Second),
	}

	return token, obj.Openid, nil
}
