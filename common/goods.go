package common

import (
	"github.com/xuri/excelize/v2"
	"strings"
)

type GoodsInfo struct {
	Model       string // 型号
	BrandPrefix string //品牌前缀
	IsLowPrice  string // 是否低价
	IsOnline    string //是否已上架
	SkuDisplay  string //sku显示
}

func GetExcelRows(excelPath, excelSheet string) (rows [][]string, err error) {
	f, err := excelize.OpenFile(excelPath)
	if err != nil {
		return
	}
	defer f.Close()
	rows, err = f.GetRows(excelSheet)
	if err != nil {
		return
	}
	return rows, nil
}

func GetExcelCols(excelPath, excelSheet string) (cols [][]string, err error) {
	f, err := excelize.OpenFile(excelPath)
	if err != nil {
		return
	}
	defer f.Close()
	cols, err = f.GetCols(excelSheet)
	if err != nil {
		return
	}
	return cols, nil
}

func GetGoodsProperties(excelPath, excelSheet string) (propertiesMap map[string][]string, err error) {
	cols, err := GetExcelCols(excelPath, excelSheet)
	if err != nil {
		return
	}
	propertiesMap = make(map[string][]string)
	for _, col := range cols {
		key := strings.Trim(col[0], " ")
		if key != "" {
			for _, c := range col[1:] {
				if strings.Trim(c, " ") != "" {
					propertiesMap[key] = append(propertiesMap[key], c)
				}
			}
		}
	}
	return propertiesMap, nil
}

func GetGoodsMap(excelPath, excelSheet string) (goodsMap map[string][]*GoodsInfo, err error) {
	rows, err := GetExcelRows(excelPath, excelSheet)
	if err != nil {
		return goodsMap, err
	}
	goodsMap = make(map[string][]*GoodsInfo)
	var key string
	for i, row := range rows {
		if i > 0 && len(row) == 5 { // 跳过空行
			title := strings.Trim(row[0], " ")
			if title != "" {
				key = title
				goodsMap[key] = append(goodsMap[key], &GoodsInfo{
					Model:       strings.Trim(row[1], " "),
					BrandPrefix: strings.Trim(row[2], " "),
					IsLowPrice:  strings.Trim(row[3], " "),
					IsOnline:    strings.Trim(row[4], " "),
					SkuDisplay:  strings.Trim(row[5], " "),
				})
			} else {
				goodsMap[key] = append(goodsMap[key], &GoodsInfo{
					Model:       strings.Trim(row[1], " "),
					BrandPrefix: strings.Trim(row[2], " "),
					IsLowPrice:  strings.Trim(row[3], " "),
					IsOnline:    strings.Trim(row[4], " "),
					SkuDisplay:  strings.Trim(row[5], " "),
				})
			}
		}
	}
	return goodsMap, nil
}
