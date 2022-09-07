package common

import (
	"github.com/xuri/excelize/v2"
	"log"
	"strings"
)

type GoodsProperties struct {
	ExcelPath              string
	ExcelSheetName         string
	Brand                  string   //品牌
	Shape                  string   //款式
	PopularElementList     []string //PopularElement
	Style                  string   // 风格
	ProtectiveCoverTexture string   //保护套质地
}

func (g *GoodsProperties) ReadConfig() (goodProperties GoodsProperties) {
	f, err := excelize.OpenFile(g.ExcelPath)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	rows, err := f.GetRows(g.ExcelSheetName)
	if err != nil {
		log.Println(err)
		return
	}
	for i, row := range rows {
		if i > 0 {
			if i == 1 {
				goodProperties.Brand = strings.Trim(row[0], " ")
				goodProperties.Shape = strings.Trim(row[1], " ")
				goodProperties.PopularElementList = append(goodProperties.PopularElementList, strings.Trim(row[2], " "))
				goodProperties.Style = strings.Trim(row[3], " ")
				goodProperties.ProtectiveCoverTexture = strings.Trim(row[4], " ")
			} else {
				goodProperties.PopularElementList = append(goodProperties.PopularElementList, strings.Trim(row[2], " "))
			}
		}
	}
	return
}
