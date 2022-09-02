package client

import (
	"fmt"
	"github.com/skip2/go-qrcode"
)

var (
	pinAuthH5 = "mai.pinduoduo.com/h5-login.html"
)

type QRCodeArgs struct {
	AuthH5      string
	ClientId    string
	RedirectUri string
	State       string
	Path        string
}

// 生成授权码,商家APP扫码授权
func (q *QRCodeArgs) MakeQRCode() {
	q.AuthH5 = pinAuthH5
	qrcode.WriteFile(fmt.Sprintf("https://%s?response_type=code&client_id=%s&redirect_uri=%s&state=%s&view=h5",
		q.AuthH5, q.ClientId, q.RedirectUri, q.State),
		qrcode.Medium, 256, q.Path)
}
