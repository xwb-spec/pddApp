package sdk

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	RedirectUri  = "https://cback.whitewolvesx.com:8088/api/v1/callback/"
	Code         = ""
	State        = ""
	CliId        = "111"
	CliSecret    string
	AccessToken  string
	RefreshToken string
)

type TokenAPI struct {
	Context *Context
}

//func NewTokenAPI(cfg *Config) *TokenAPI {
//	return &TokenAPI{Context: NewContext(cfg)}
//}
type ReturnCodeResponse struct {
	Code  string `json:"code"`  // 返回code
	State string `json:"state"` // 状态
}

func newTokenAPIWithContext(c *Context) *TokenAPI {
	return &TokenAPI{Context: c}
}

type PopAuthTokenResponse struct {
	AccessToken           string `json:"access_token"`
	ExpiresAt             string `json:"expires_at"`
	ExpiresIn             string `json:"expires_in"`
	OwnerId               string `json:"owner_id"`
	OwnerName             string `json:"owner_name"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresAt string `json:"refresh_token_expires_at"`
	RefreshTokenExpiresIn string `json:"refresh_token_expires_in"`
}

//type TokenResponse struct {
//	PopAuthTokenCreateResponse PopAuthTokenCreateResponse `json:"pop_auth_token_create_response"`
//}
// 生成授权码,商家APP扫码授权
func (g *TokenAPI) GetCode() {
	resp, err := http.Post(RedirectUri, "application/json;charset=utf-8", nil)
	if err != nil {
		log.Fatal("获取code失败")
	}
	log.Println()
	defer resp.Body.Close()
	var r ReturnCodeResponse
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(bodyBytes))
	_ = json.Unmarshal(bodyBytes, &r)
	log.Println(r.Code)
	Code = r.Code
	State = r.State
}

func (g *TokenAPI) PopAuthTokenCreate() (resp *PopAuthTokenResponse, err error) {
	g.GetCode()
	log.Println(State + "1111")
	params := NewParamsWithType("pdd.pop.auth.token.create")
	params.Set("code", Code)
	params.Set("grant_type", "authorization_code")
	params.Set("client_secret", CliSecret)
	r, err := Call(g.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r, "pop_auth_token_create_response")
	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// 刷新token
func (g *TokenAPI) PopAuthTokenRefresh(refreshToken string) (resp *PopAuthTokenResponse, err error) {
	params := NewParamsWithType("pdd.pop.auth.token.refresh")
	params.Set("refresh_token", refreshToken)
	params.Set("client_secret", CliSecret)
	r, err := Call(g.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r, "pop_auth_token_refresh_response")
	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		return nil, err
	}
	return
}

// 获取token
func PopAuthCreateToken() (string, error) {
	p := NewPdd(&Config{
		ClientId:     CliId,
		ClientSecret: CliSecret,
		EndPoint:     "https://open-api.pinduoduo.com/oauth/token",
		RetryTimes:   3, // 设置接口调用失败重试次数
	})
	pdd := p.TokenAPI()
	resp, err := pdd.PopAuthTokenCreate()
	if err != nil {
		return "", err
	}
	AccessToken = resp.AccessToken
	RefreshToken = resp.RefreshToken
	return State, nil
}

// 刷新token
func PopAuthRefreshToken() (err error) {
	p := NewPdd(&Config{
		ClientId:     CliId,
		ClientSecret: CliSecret,
		EndPoint:     "https://open-api.pinduoduo.com/oauth/token",
		RetryTimes:   3, // 设置接口调用失败重试次数
	})
	pdd := p.TokenAPI()
	resp, err := pdd.PopAuthTokenRefresh(RefreshToken)
	if err != nil {
		return
	}
	AccessToken = resp.AccessToken
	return nil
}
