package client

import (
	"encoding/json"
	"fmt"
	"github.com/skip2/go-qrcode"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	pinAuthH5 = "mai.pinduoduo.com/h5-login.html"
	returnUrl = "http://www.test.com/api/v2/callback/"
)

// 公共参数
type CommonRequestParam struct {
	Sign        string `json:"sign"` // 请求签名
	Type        string `json:"type"` // 接口类型
	ClientId    string `json:"client_id"`
	TimeStamp   string `json:"timestamp"` //时间戳
	DataType    string `json:"data_type"` // 返回数据类型JSON
	AccessToken string `json:"access_token"`
	Version     string `json:"version"`
}
type CreateTokenRequestParam struct {
	Code string `json:"code"`
	CommonRequestParam
}
type QRCodeRequestParam struct {
	AuthH5      string
	ClientId    string
	RedirectUri string
	State       string
	Path        string
}
type PopAuthTokenCreateResponse struct {
	AccessToken           string `json:"access_token"`
	ExpiresAt             string `json:"expires_at"`
	ExpiresIn             string `json:"expires_in"`
	OwnerId               string `json:"owner_id"`
	OwnerName             string `json:"owner_name"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresAt string `json:"refresh_token_expires_at"`
	RefreshTokenExpiresIn string `json:"refresh_token_expires_in"`
}
type TokenResponse struct {
	PopAuthTokenCreateResponse PopAuthTokenCreateResponse `json:"pop_auth_token_create_response"`
}
type ErrorResponse struct {
	ErrorMsg  string `json:"error_msg"`
	SubMsg    string `json:"sub_msg"`
	SubCode   int    `json:"sub_code"`
	ErrorCode int    `json:"error_code"`
	RequestId string `json:"request_id"`
}
type ErrorResult struct {
	ErrorResponse ErrorResponse `json:"error_response"`
}

// 生成授权码,商家APP扫码授权
func (q *QRCodeRequestParam) MakeQRCode() {
	q.AuthH5 = pinAuthH5
	qrcode.WriteFile(fmt.Sprintf("https://%s?response_type=code&client_id=%s&redirect_uri=%s&state=%s&view=h5",
		q.AuthH5, q.ClientId, q.RedirectUri, q.State),
		qrcode.Medium, 256, q.Path)
}

func (q *QRCodeRequestParam) GetReturnCode() (string, error) {
	resp, err := http.Post(returnUrl, "application/json", nil) // 请求code
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		ret := make(map[string]string)
		_ = json.Unmarshal(body, ret)
		return ret["code"], nil
	}
	return "", err
}

func makeSign(args map[string]string) {

}
func getToken() (string, error) {
	c := QRCodeRequestParam{}
	code, err := c.GetReturnCode()
	if err != nil {
		log.Panic(err)
	}
	// json.Marshal
	reqParam, err := json.Marshal(&CreateTokenRequestParam{
		Code: code,
	})
	if err != nil {
		log.Fatal("Marshal RequestParam fail, err:%v", err)
		return "", err
	}
	url1 := "https://gw-api.pinduoduo.com/api/router"
	reqBody := strings.NewReader(string(reqParam))
	httpReq, err := http.NewRequest("POST", url1, reqBody)
	if err != nil {
		log.Fatalf("NewRequest fail, url: %s, reqBody: %s, err: %v", url1, reqBody, err)
		return "", err
	}
	httpReq.Header.Add("Content-Type", "application/json")
	httpRsp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		log.Fatal("do http fail, url: %s, reqBody: %s, err:%v", url1, reqBody, err)
		return "", err
	}
	defer httpRsp.Body.Close()

	// Read: HTTP结果
	rspBody, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		log.Fatalf("ReadAll failed, url: %s, reqBody: %s, err: %v", url1, reqBody, err)
		return "", err
	}
	if httpRsp.StatusCode == 200 {
		var result TokenResponse
		if err = json.Unmarshal(rspBody, &result); err != nil {
			log.Fatal("Unmarshal fail, err:%v", err)
			return result.PopAuthTokenCreateResponse.AccessToken, err
		}
	} else {
		var result ErrorResult
		if err = json.Unmarshal(rspBody, &result); err != nil {
			log.Fatal("Unmarshal fail, err:%v", err)
			return result.ErrorResponse.ErrorMsg, err
		}
	}
	return "", err
}
