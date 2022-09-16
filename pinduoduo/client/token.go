package client

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"log"
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

var (
	CliId        string
	CliSecret    string
	token        string
	refreshToken string
	RedirectUri  = "https://cback.whitewolvesx.com:8088/api/v1/callback/"
)

type ReturnCodeResponse struct {
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
