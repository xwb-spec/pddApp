package yyui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"log"
	"pddApp/pinduoduo/client"
	"pddApp/yyui/theme"
	"strconv"
	"time"
)

var myapp fyne.App

func MainWindow(w fyne.Window) {
	ms := ShowInput{}
	mw := myapp.NewWindow("YY批量上链接V1.0   Xwb ALL Right Reserved QQ:543361609")
	ms.MainShow(mw)
	ms.GetInput() // 初始化对话框数据
	mw.Resize(fyne.Size{Width: 800, Height: 850})
	mw.CenterOnScreen()
	w.Close()
	mw.Show()
	//myapp.Run()
}

func LoginWindow() {
	lw := myapp.NewWindow("登录")
	ls := ShowInput{}
	ls.LoginShow(lw)
	ls.GetAcc() // 初始化对话框数据
	lw.Resize(fyne.Size{Width: 300, Height: 150})
	lw.CenterOnScreen()
	lw.Show()
	//myapp.Run()
}

func QRCodeWindow(w fyne.Window) {
	qw := myapp.NewWindow("扫二维码")
	qs := QRCodeInput{}
	state := strconv.FormatInt(time.Now().Unix(), 10)
	client.GenerateQRCode(state)
	qw.Resize(fyne.Size{Width: 300, Height: 300})
	qw.CenterOnScreen()
	qs.QRCodeShow(qw)
	w.Close()
	qw.Show()
	qw.SetOnClosed(func() {
		qw.Close()
	})
	for {
		resp, _ := client.GetCode()
		log.Println(resp)
		if resp.State == state { // 拿到最新code
			if err := client.PopAuthCreateToken(); err != nil {
				log.Fatalln("[ERROR]: 获取token失败")
			} else {
				MainWindow(qw)
				break
			}
		}
		time.Sleep(3 * time.Second)
	}
}

func NewApp() fyne.App {
	myapp = app.New()
	myapp.Settings().SetTheme(&theme.MyTheme{})
	return myapp
}
