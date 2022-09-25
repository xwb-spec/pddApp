package main

import (
	"pddApp/yyui"
)

func main() {
	//ex := common.Comparisons{ExcelPath: "/Users/machao/Desktop/批量上商品app/批量上商品app/模板-型号图片对照表.xlsx", ExcelSheetName: "直边tpu"}
	//ex.Read()
	//fmt.Println(ex.ComparisonList)
	yyui.NewApp()
	yyui.LoginWindow()
	//fmt.Println(common.GetGoodsProperties("/Users/machao/Desktop/批量上商品app/批量上商品app/模板-sku配置.xlsx", "属性"))
}
