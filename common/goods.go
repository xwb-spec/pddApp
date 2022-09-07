package common

import (
	"github.com/xuri/excelize/v2"
	"log"
	"strings"
)

type GoodsInfo struct {
	Model       string // 型号
	BrandPrefix string //品牌前缀
	IsLowPrice  string // 是否低价
	IsOnline    string //是否已上架
	SkuDisplay  string //sku显示
}

// 装载商品数据
type Goods struct {
	GoodsExcelPath      string
	GoodsExcelSheetName string
	//GoodsConfigExcelPath          string
	//GoodsConfigExcelSheetName     string
	//GoodsComparisonExcelPath      string
	//GoodsComparisonExcelSheetName string
	GoodsMap map[string][]*GoodsInfo
}

func (g *Goods) ReadGoods() {
	f, err := excelize.OpenFile(g.GoodsExcelPath)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	rows, err := f.GetRows(g.GoodsExcelSheetName)
	if err != nil {
		log.Println(err)
		return
	}
	g.GoodsMap = make(map[string][]*GoodsInfo)
	var key string
	for i, row := range rows {
		if i > 0 { // 跳过空行
			col1 := strings.Trim(row[0], " ")
			if col1 != "" {
				key = col1
				g.GoodsMap[key] = append(g.GoodsMap[key], &GoodsInfo{
					Model:       strings.Trim(row[1], " "),
					BrandPrefix: strings.Trim(row[2], " "),
					IsLowPrice:  strings.Trim(row[3], " "),
					IsOnline:    strings.Trim(row[4], " "),
					SkuDisplay:  strings.Trim(row[5], " "),
				})
			} else {
				g.GoodsMap[key] = append(g.GoodsMap[key], &GoodsInfo{
					Model:       strings.Trim(row[1], " "),
					BrandPrefix: strings.Trim(row[2], " "),
					IsLowPrice:  strings.Trim(row[3], " "),
					IsOnline:    strings.Trim(row[4], " "),
					SkuDisplay:  strings.Trim(row[5], " "),
				})
			}
		}
	}
}

func (g *Goods) GetSkuDisplayList() (skuList []string) {
	g.ReadGoods()
	for _, l := range g.GoodsMap {
		for _, v := range l {
			skuList = append(skuList, v.SkuDisplay)
		}
	}
	return
}

func (g *Goods) GetComparisonMap() map[string]*Comparison {
	goodsComparisonExcelPath := "/Users/machao/Desktop/批量上商品app/批量上商品app/模板-型号图片对照表.xlsx"
	goodsComparisonExcelSheetName := "直边tpu"
	com := Comparisons{ExcelPath: goodsComparisonExcelPath, ExcelSheetName: goodsComparisonExcelSheetName}
	com.Read()
	return com.ComparisonMap
}
