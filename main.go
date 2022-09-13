package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/flopp/go-findfont"
	"os"
	"pddApp/yyui"
	"strings"
)

//设置字体
func init() {
	fontPaths := findfont.List()
	for _, fontPath := range fontPaths {
		//fmt.Println(fontPath)
		//楷体:simkai.ttf
		//黑体:simhei.ttf
		//微软雅黑：msyh.ttc
		if strings.Contains(fontPath, "/Users/machao/OptRepo/GolandProjects/awesomeProject/learning-golang/day41/simkai.ttf") {
			err := os.Setenv("FYNE_FONT", fontPath)
			if err != nil {
				fmt.Println(err)
				return
			}
			break
		}
	}
}
func runWin() {
	err := os.Setenv("FYNE_FONT", "./simkai.ttf")
	if err != nil {
		fmt.Println(err)
		return
	}
	//新建一个app
	a := app.New()
	//设置窗口栏，任务栏图标
	//a.SetIcon(resourceIconPng)
	//新建一个窗口
	s := yyui.ShowInput{}
	w := a.NewWindow("YY批量上链接V1.0   Xwb ALL Right Reserved QQ:543361609")
	//主界面框架布局
	s.MainShow(w)
	s.GetInput() // 初始化对话框数据
	//尺寸
	w.Resize(fyne.Size{Width: 800, Height: 850})
	//w居中显示
	w.CenterOnScreen()
	//循环运行
	w.ShowAndRun()

	err = os.Unsetenv("FYNE_FONT")
	if err != nil {
		return
	}
}
func main() {
	//ex := common.Comparisons{ExcelPath: "/Users/machao/Desktop/批量上商品app/批量上商品app/模板-型号图片对照表.xlsx", ExcelSheetName: "直边tpu"}
	//ex.Read()
	//fmt.Println(ex.ComparisonList)
	runWin()
	//fmt.Println(common.GetGoodsProperties("/Users/machao/Desktop/批量上商品app/批量上商品app/模板-sku配置.xlsx", "属性"))
}
