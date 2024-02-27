package gftoken

import "github.com/golang-jwt/jwt/v5"

const (
	//token部分
	ErrorsParseTokenFail    string = "解析token失败"
	ErrorsTokenInvalid      string = "无效的token"
	ErrorsTokenNotActiveYet string = "Token 尚未激活"
	ErrorsTokenMalFormed    string = "Token 格式不正确"

	JwtTokenOK            int = 200100  //token有效
	JwtTokenInvalid       int = -400100 //无效的token
	JwtTokenExpired       int = -400101 //过期的token
	JwtTokenFormatErrCode int = -400102 //提交的 Token 格式错误
)

type CustomClaims struct {
	Data interface{}
	jwt.RegisteredClaims
}
