package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"pddApp/yyui"
	"pddApp/yyui/theme"
)

func MainWindow() {
	a := app.New()
	a.Settings().SetTheme(&theme.MyTheme{})
	s := yyui.ShowInput{}
	w := a.NewWindow("YY批量上链接V1.0   Xwb ALL Right Reserved QQ:543361609")
	s.MainShow(w)
	s.GetInput() // 初始化对话框数据
	w.Resize(fyne.Size{Width: 800, Height: 850})
	w.CenterOnScreen()
	w.ShowAndRun()
}

func LoginWindow() {
	a := app.New()
	a.Settings().SetTheme(&theme.MyTheme{})
	w := a.NewWindow("登录")
	s := yyui.LoginInput{}
	s.MainShow(w)
	s.GetInput() // 初始化对话框数据
	w.Resize(fyne.Size{Width: 800, Height: 850})
	w.CenterOnScreen()
	w.ShowAndRun()
}
func main() {
	//ex := common.Comparisons{ExcelPath: "/Users/machao/Desktop/批量上商品app/批量上商品app/模板-型号图片对照表.xlsx", ExcelSheetName: "直边tpu"}
	//ex.Read()
	//fmt.Println(ex.ComparisonList)
	MainWindow()
	//fmt.Println(common.GetGoodsProperties("/Users/machao/Desktop/批量上商品app/批量上商品app/模板-sku配置.xlsx", "属性"))
}
