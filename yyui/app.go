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

func MainWindow() {
	ms := ShowInput{}
	mw := myapp.NewWindow("YY批量上链接V1.0   Xwb ALL Right Reserved QQ:543361609")
	ms.MainShow(mw)
	ms.GetInput() // 初始化对话框数据
	mw.Resize(fyne.Size{Width: 800, Height: 850})
	mw.CenterOnScreen()
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

func QRCodeWindow() {
	qw := myapp.NewWindow("扫二维码")
	qs := QRCodeInput{}
	qs.QRCodeShow(qw)
	qw.Resize(fyne.Size{Width: 300, Height: 300})
	qw.CenterOnScreen()
	qw.Show()
	//myapp.Run()
	//_ = client.PopAuthCreateToken() // 拿token
	//MainWindow()                    // 关闭二维码
	//defer qw.Close()
	defer qw.Close()
	state := strconv.FormatInt(time.Now().Unix(), 10)
	times := 0
	for times < 6 {
		c, _ := client.GetCode()
		log.Println(c.State)
		c.State = state
		if c.State == state { // 拿到最新code
			_ = client.PopAuthCreateToken() // 拿token
			MainWindow()                    // 关闭二维码
			times = 6
		}
		time.Sleep(3 * time.Second)
		times += 1
		qw.SetOnClosed(func() {
			times = 6
		})
	}
}

func NewApp() fyne.App {
	myapp = app.New()
	myapp.Settings().SetTheme(&theme.MyTheme{})
	return myapp
}
