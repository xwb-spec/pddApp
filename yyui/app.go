package yyui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"log"
	"pddApp/yyui/theme"
)

var myapp fyne.App

func MainWindow() {
	//n.App.Settings().SetTheme(&theme.MyTheme{})
	s := ShowInput{}
	w := myapp.NewWindow("YY批量上链接V1.0   Xwb ALL Right Reserved QQ:543361609")
	s.MainShow(w)
	s.GetInput() // 初始化对话框数据
	w.Resize(fyne.Size{Width: 800, Height: 850})
	w.CenterOnScreen()
	w.ShowAndRun()
}

func LoginWindow() {
	w := myapp.NewWindow("登录")
	s := ShowInput{}
	s.LoginShow(w)
	s.GetInput() // 初始化对话框数据
	w.Resize(fyne.Size{Width: 300, Height: 150})
	w.CenterOnScreen()
	w.Show()
	myapp.Run()
}

func QRCodeWindow() {
	//n.App = app.New()
	//a.Settings().SetTheme(&theme.MyTheme{})
	qWin := myapp.NewWindow("登录")
	q := QRCodeInput{}
	q.QRCodeShow(qWin)
	log.Println("aaa")
	//s.GetInput() // 初始化对话框数据
	qWin.Resize(fyne.Size{Width: 300, Height: 300})
	qWin.CenterOnScreen()
	qWin.Show()
}

func NewApp() {
	myapp = app.New()
	myapp.Settings().SetTheme(&theme.MyTheme{})
}
