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

//// 公共参数
//type CommonRequestParam struct {
//	Sign        string `json:"sign"` // 请求签名
//	Type        string `json:"type"` // 接口类型
//	ClientId    string `json:"client_id"`
//	TimeStamp   string `json:"timestamp"` //时间戳
//	DataType    string `json:"data_type"` // 返回数据类型JSON
//	AccessToken string `json:"access_token"`
//	Version     string `json:"version"`
//}
//type CreateTokenRequestParam struct {
//	Code string `json:"code"`
//	CommonRequestParam
//}
type QRCodeRequestParam struct {
	ClientId    string
	RedirectUri string
	State       string
	Path        string
}

//type PopAuthTokenCreateResponse struct {
//	AccessToken           string `json:"access_token"`
//	ExpiresAt             string `json:"expires_at"`
//	ExpiresIn             string `json:"expires_in"`
//	OwnerId               string `json:"owner_id"`
//	OwnerName             string `json:"owner_name"`
//	RefreshToken          string `json:"refresh_token"`
//	RefreshTokenExpiresAt string `json:"refresh_token_expires_at"`
//	RefreshTokenExpiresIn string `json:"refresh_token_expires_in"`
//}
//type TokenResponse struct {
//	PopAuthTokenCreateResponse PopAuthTokenCreateResponse `json:"pop_auth_token_create_response"`
//}

//type ErrorResponse struct {
//	ErrorMsg  string `json:"error_msg"`
//	SubMsg    string `json:"sub_msg"`
//	SubCode   int    `json:"sub_code"`
//	ErrorCode int    `json:"error_code"`
//	RequestId string `json:"request_id"`
//}
//type ErrorResultResponse struct {
//	ErrorResponse ErrorResponse `json:"error_response"`
//}

type ReturnCodeResponse struct {
	Code  string `json:"code"`  // 返回code
	State string `json:"state"` // 状态
}

// 生成授权码,商家APP扫码授权
func (q *QRCodeRequestParam) MakeQRCode() {
	qrcode.WriteFile(fmt.Sprintf(
		"https://mai.pinduoduo.com/h5-login.html?response_type=code&client_id=%s&redirect_uri=%s&state=%s&view=h5",
		q.ClientId, q.RedirectUri, q.State), qrcode.Medium, 256, q.Path)
}

// 获取返回码
func GetReturnCode() (ret ReturnCodeResponse, err error) {
	resp, err := http.Post("https://www.test.com/api/v2/callback/",
		"application/json;charset=utf-8",
		nil,
	)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	if (resp.StatusCode == 200) || (resp.StatusCode == 201) {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		_ = json.Unmarshal(bodyBytes, &ret)
		return ret, nil
	}
	return ReturnCodeResponse{}, err
}

// 获取token
func GetToken() string {
	p := sdk.NewPdd(&sdk.Config{
		ClientId:     "your client id",
		ClientSecret: "your client secret",
		RetryTimes:   3, // 设置接口调用失败重试次数
	})
	code, err := GetReturnCode()
	if err != nil {
		log.Println(err)
	}
	pdd := p.GetTokenAPI()
	resp, err := pdd.TokenGet(code.Code)
	if err != nil {
		log.Println("获取token失败")
	}
	return resp.AccessToken
}
