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
	GoodsMap map[string][]*GoodsInfo
}

func GetExcel(excelPath, excelSheet string) [][]string {
	f, err := excelize.OpenFile(excelPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	rows, err := f.GetRows(excelSheet)
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

func GetGoodsProperties(excelPath, excelSheet string) {
	rows := GetExcel(excelPath, excelSheet)
	keyMap := rows[0]
	propertiesMap := make(map[string][]string)
	for idx, row := range rows {
		if idx > 0 && len(row) != 0 {
			if strings.Trim(row[0], " ") != "" {
				propertiesMap[keyMap[0]] = append(propertiesMap[keyMap[0]], row[0])
			}
			if strings.Trim(row[1], " ") != "" {
				propertiesMap[keyMap[1]] = append(propertiesMap[keyMap[1]], row[1])
			}
			if strings.Trim(row[2], " ") != "" {
				propertiesMap[keyMap[0]] = append(propertiesMap[keyMap[0]], row[0])
			}
			if strings.Trim(row[3], " ") != "" {
				propertiesMap[keyMap[0]] = append(propertiesMap[keyMap[0]], row[0])
			}
		}
	}
}

func (g *Goods) GetGoods(excelPath, excelSheet string) {
	rows := GetExcel(excelPath, excelSheet)
	g.GoodsMap = make(map[string][]*GoodsInfo)
	var key string
	for i, row := range rows {
		if i > 0 && len(row) != 0 { // 跳过空行
			title := strings.Trim(row[0], " ")
			if title != "" {
				key = title
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
