package main

import (
	"fmt"
	"github.com/skip2/go-qrcode"
)

type GenPinDuoduoQRcode

func main() {
	authH5 := "mai.pinduoduo.com/h5-login.html"
	clientId := ""
	redirectUri := ""
	state := ""
	qrcode.WriteFile(fmt.Sprintf("https://%s?response_type=code&client_id=%s&redirect_uri=%s&state=%s&view=h5",
		authH5, clientId, redirectUri, state),
		qrcode.Medium, 256, "./blog_qrcode.png")
}
