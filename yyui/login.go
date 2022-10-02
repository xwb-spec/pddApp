package yyui

import (
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/crypto/bcrypt"
	"image/color"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	AccInput string
	PwdInput string
)

type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginInput struct {
	LoginWin fyne.Window
	Acc      *widget.Entry
	Pwd      *widget.Entry
	Tips     *canvas.Text
}

func GetAccount() (Account, error) {
	resp, err := http.Post("https://cback.whitewolvesx.com:8088/api/v1/auth/", "application/json;charset=utf-8", nil)
	if err != nil {
		log.Fatal("[ERROR]: 获取用户密码失败")
	}
	defer resp.Body.Close()
	var a Account
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(bodyBytes, &a)
	return a, nil
}

// 验证密码
func PasswordVerify(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
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
	acc, _ := GetAccount()
	s.Pwd.OnSubmitted = func(str string) {
		AccInput = s.Acc.Text
		PwdInput = s.Pwd.Text
		s.SaveAcc()
		if s.Acc.Text == acc.Username && PasswordVerify(acc.Password, []byte(s.Pwd.Text)) {
			s.Tips.Color = green
			s.Tips.Text = "登录成功"
			s.Tips.Refresh()
			//QRCodeWindow()
			QRCodeWindow(s.LoginWin)
			// 登录成功关闭登录窗口
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
		if s.Acc.Text == acc.Username && PasswordVerify(acc.Password, []byte(s.Pwd.Text)) {
			s.Tips.Color = green
			s.Tips.Text = "登录成功"
			s.Tips.Refresh()
			QRCodeWindow(s.LoginWin)
			//s.LoginWin.Close() // 登录成功关闭登录窗口
		} else {
			s.Tips.Color = red
			s.Tips.Text = "账号密码错误"
			s.Tips.Refresh()
		}
	}
	form.OnCancel = func() {
		s.LoginWin.Close() // 登录成功关闭登录窗口
	}
	form.CancelText = "关闭"
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
