package yyui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

var (
	AccInput string
	PwdInput string
)

type LoginInput struct {
	LoginWin fyne.Window
	Acc      *widget.Entry
	Pwd      *widget.Entry
	Tips     *canvas.Text
}

func (s *ShowInput) LoginContainer() *widget.Form {
	s.Acc = widget.NewEntry()
	s.Pwd = widget.NewPasswordEntry()
	form := widget.NewForm(
		widget.NewFormItem("用户名", s.Acc),
		widget.NewFormItem("密码", s.Pwd),
	)
	s.Tips = canvas.NewText("", color.White)
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	red := color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	s.Pwd.OnSubmitted = func(str string) {
		AccInput = s.Acc.Text
		PwdInput = s.Pwd.Text
		s.SaveAcc()
		if s.Acc.Text == USER && s.Pwd.Text == PASS {
			s.Tips.Color = green
			s.Tips.Text = "登录成功"
			s.Tips.Refresh()
			QRCodeWindow()
			s.LoginWin.Close() // 登录成功关闭登录窗口
		} else {
			s.Tips.Color = red
			s.Tips.Text = "账号密码错误"
			s.Tips.Refresh()
		}
	}
	form.OnSubmit = func() {
		AccInput = s.Acc.Text
		PwdInput = s.Pwd.Text
		s.SaveAcc()
		if s.Acc.Text == USER && s.Pwd.Text == PASS {
			s.Tips.Color = green
			s.Tips.Text = "登录成功"
			s.Tips.Refresh()
			QRCodeWindow()
			s.LoginWin.Close() // 登录成功关闭登录窗口
		} else {
			s.Tips.Color = red
			s.Tips.Text = "账号密码错误"
			s.Tips.Refresh()
		}
	}
	form.SubmitText = "登录"
	return form
}
func (s *ShowInput) LoginShow(w fyne.Window) {
	s.LoginWin = w
	box := container.NewVBox(
		s.LoginContainer(),
		s.Tips,
	) //控制显示位置顺序
	w.SetContent(box)
}
