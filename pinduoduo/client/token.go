package client

import (
	"encoding/json"
	"fmt"
	"github.com/skip2/go-qrcode"
	"io/ioutil"
	"log"
	"net/http"
	"pddApp/pinduoduo/sdk"
)

const (
	CliId         = "11111"
	CliSecret     = "22222"
	RedirectUri   = "https://cback.whitewolvesx.com:8088/api/v1/callback/"
	OAuthEndPoint = "https://open-api.pinduoduo.com/oauth/token"
	EndPoint      = "https://gw-api.pinduoduo.com/api/router"
)

var (
	AccessToken  string
	RefreshToken string
	State        string
)

type CodeResponse struct {
	Code  string `json:"code"`  // 返回code
	State string `json:"state"` // 状态
}

// 生成授权码,商家APP扫码授权
func GenerateQRCode(state string) {
	if err := qrcode.WriteFile(fmt.Sprintf(
		"https://mai.pinduoduo.com/h5-login.html?response_type=code&client_id=%s&redirect_uri=%s&state=%s&view=h5",
		CliId, RedirectUri, state), qrcode.Medium, 256, "./qrcode.png"); err != nil {
		log.Printf("[ERROR]: 生成二维码失败, %s\n", err)
	}
}

func GetCode() (CodeResponse, error) {
	resp, err := http.Post(RedirectUri, "application/json;charset=utf-8", nil)
	if err != nil {
		log.Fatal("获取code失败")
	}
	defer resp.Body.Close()
	var code CodeResponse
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(bodyBytes, &code)
	return code, nil
}

// 获取token
func PopAuthCreateToken() (err error) {
	p := sdk.NewPdd(&sdk.Config{
		ClientId:     CliId,
		ClientSecret: CliSecret,
		EndPoint:     OAuthEndPoint,
		RetryTimes:   1, // 设置接口调用失败重试次数
	})
	code, err := GetCode()
	pdd := p.TokenAPI()
	params := sdk.NewParams()
	params.Set("code", code.Code)
	params.Set("client_secret", CliSecret)
	resp, err := pdd.PopAuthTokenCreate()
	if err != nil {
		return err
	}
	AccessToken = resp.AccessToken
	RefreshToken = resp.RefreshToken
	return nil
}

// 刷新token
func PopAuthRefreshToken() (err error) {
	p := sdk.NewPdd(&sdk.Config{
		ClientId:     CliId,
		ClientSecret: CliSecret,
		EndPoint:     OAuthEndPoint,
		RetryTimes:   3, // 设置接口调用失败重试次数
	})
	pdd := p.TokenAPI()
	params := sdk.NewParams()
	params.Set("refresh_token", RefreshToken)
	params.Set("client_secret", CliSecret)
	resp, err := pdd.PopAuthTokenRefresh()
	if err != nil {
		return
	}
	AccessToken = resp.AccessToken
	return nil
}
