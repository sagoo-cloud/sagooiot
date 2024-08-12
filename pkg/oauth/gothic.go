package oauth

import (
	"bytes"
	"compress/gzip"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/sessions"
)

// 参数用于将数据传递给会话进行授权
type Params interface {
	Get(string) string
}

// Session需要作为 provider 的一部分来实现。
type Session interface {
	// GetAuthURL 返回提供程序的身份验证端点的URL。
	GetAuthURL() (string, error)
	// Marshal 生成会话的字符串表示形式，以便在请求之间存储。
	Marshal() string
	// Authorize 验证来自提供程序的数据并返回访问令牌
	Authorize(Provider, Params) (string, error)
}

func BeginAuthHandler(res http.ResponseWriter, req *http.Request) {
	url, err := GetAuthURL(res, req)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(res, err)
		return
	}

	http.Redirect(res, req, url, http.StatusTemporaryRedirect)
}

// SetState state不存在,则生成 state 返回
func SetState(req *http.Request) string {
	state := req.URL.Query().Get("state")
	if len(state) > 0 {
		return state
	}
	nonceBytes := make([]byte, 64)
	_, err := io.ReadFull(rand.Reader, nonceBytes)
	if err != nil {
		panic("gothic: source of randomness unavailable: " + err.Error())
	}
	return base64.URLEncoding.EncodeToString(nonceBytes)
}

// GetState 获取提供程序在回调期间返回的状态。
func GetState(req *http.Request) string {
	params := req.URL.Query()
	if params.Encode() == "" && req.Method == http.MethodPost {
		return req.FormValue("state")
	}
	return params.Get("state")
}

func GetAuthURL(res http.ResponseWriter, req *http.Request) (string, error) {
	providerName, err := getProviderName(req)
	if err != nil {
		return "", err
	}

	provider, err := GetProvider(providerName)
	if err != nil {
		return "", err
	}
	sess, err := provider.BeginAuth(SetState(req))
	if err != nil {
		return "", err
	}

	url, err := sess.GetAuthURL()
	if err != nil {
		return "", err
	}

	return url, err
}

// CompleteUserAuth做它在罐头上说的。它完成身份验证过程，并从提供程序获取有关用户的所有基本信息。
func CompleteUserAuth(res http.ResponseWriter, req *http.Request) (User, error) {

	providerName, err := getProviderName(req)
	if err != nil {
		return User{}, err
	}

	provider, err := GetProvider(providerName)
	if err != nil {
		return User{}, err
	}
	sess, err := provider.UnmarshalSession("")
	if err != nil {
		return User{}, err
	}

	err = validateState(req, sess)
	if err != nil {
		return User{}, err
	}

	params := req.URL.Query()
	if params.Encode() == "" && req.Method == "POST" {
		req.ParseForm()
		params = req.Form
	}

	// get new token and retry fetch
	_, err = sess.Authorize(provider, params)
	if err != nil {
		return User{}, err
	}

	gu, err := provider.FetchUser(sess)
	return gu, err
}

// validateState 确保状态令牌参数来自原始的
func validateState(req *http.Request, sess Session) error {
	rawAuthURL, err := sess.GetAuthURL()
	if err != nil {
		return err
	}

	authURL, err := url.Parse(rawAuthURL)
	if err != nil {
		return err
	}

	reqState := GetState(req)

	originalState := authURL.Query().Get("state")
	if originalState != "" && (originalState != reqState) {
		return errors.New("state token mismatch")
	}
	return nil
}

func getProviderName(req *http.Request) (string, error) {
	// 试着从url参数"provider"中获取它
	if p := req.URL.Query().Get("provider"); p != "" {
		return p, nil
	}

	// 试着从url参数":provider"中获取它
	if p := req.URL.Query().Get(":provider"); p != "" {
		return p, nil
	}

	return "", errors.New("you must select a provider")
}

func getSessionValue(session *sessions.Session, key string) (string, error) {
	value := session.Values[key]
	if value == nil {
		return "", fmt.Errorf("could not find a matching session for this request")
	}

	rdata := strings.NewReader(value.(string))
	r, err := gzip.NewReader(rdata)
	if err != nil {
		return "", err
	}
	s, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	return string(s), nil
}

func updateSessionValue(session *sessions.Session, key, value string) error {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte(value)); err != nil {
		return err
	}
	if err := gz.Flush(); err != nil {
		return err
	}
	if err := gz.Close(); err != nil {
		return err
	}

	session.Values[key] = b.String()
	return nil
}
