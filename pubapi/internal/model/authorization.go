package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/spf13/viper"
)

type AuthorizationCode struct {
	url       string
	Appid     string `json:"appid"`      // 小程序 appId
	Secret    string `json:"secret"`     // 小程序 appSecret
	JsCode    string `json:"js_code"`    // 用户登录时出传过来的 code
	GrantType string `json:"grant_type"` // 授权类型，此处只需填写 authorization_code
}

func NewAuthorizationCode(code string) AuthorizationCode {
	return AuthorizationCode{
		Appid:     viper.GetString("wxapp.appid"),
		Secret:    viper.GetString("wxapp.secret"),
		JsCode:    code,
		url:       "https://api.weixin.qq.com/sns/jscode2session",
		GrantType: "authorization_code",
	}
}

func (ac *AuthorizationCode) MakeSession() (*AuthorizationSession, error) {
	var values url.Values
	values.Set("appid", ac.Appid)
	values.Set("secret", ac.Secret)
	values.Set("js_code", ac.JsCode)
	values.Set("grant_type", ac.GrantType)
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s?%s", ac.url, values.Encode()), nil)
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var session AuthorizationSession
	if err = json.Unmarshal(body, &session); err != nil {
		return nil, err
	}
	return &session, nil
}

type AuthorizationSession struct {
	Openid     string `json:"openid"`      // 用户唯一标识
	Unionid    string `json:"unionid"`     // 用户在开放平台的唯一标识符
	SessionKey string `json:"session_key"` // 会话密钥
	Errcode    int    `json:"errcode"`     // 错误码
	Errmsg     string `json:"errmsg"`      // 错误信息
}
