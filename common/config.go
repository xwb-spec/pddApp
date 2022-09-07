package common

import (
	"github.com/xuri/excelize/v2"
	"log"
	"strconv"
	"strings"
)

// 商品轮播图,主图
type CarouselGalleryConfig struct {
	IsPublic bool
	FileName string
	Num      int
}

// 商品详情图
type DetailGalleryConfig struct {
	IsPublic bool
	FileName string
	Num      int
}
type SkuConfig struct {
	IsPublic         bool
	FileName         string
	Num              int
	SkuName          string
	ImageCode        string
	MatCode          string
	ColorCode        string
	GiftCode         string
	ListPrice        string // 拼单价
	SinglePrice      string // 单买价
	Stock            string // 库存
	LowPriceDrainage string // 低价引流
	MarketPrice      string // 市场价
}

type GoodsConfig struct {
	ExcelPath                 string
	ExcelSheetName            string
	SkuConfigList             []SkuConfig
	DetailGalleryConfigList   []DetailGalleryConfig
	CarouselGalleryConfigList []CarouselGalleryConfig
}

func (g *GoodsConfig) ReadConfig() (goodsConfig GoodsConfig) {
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
			isPublicVal := strings.Trim(row[1], " ")
			var isPublic bool
			if isPublicVal == "公用" {
				isPublic = true
			} else {
				isPublic = false
			}
			num, _ := strconv.Atoi(strings.Trim(row[3], " ")) // 转换为数字类型
			configType := strings.Trim(row[0], " ")           //去除两端空格
			if configType == "sku" {
				g.SkuConfigList = append(g.SkuConfigList, SkuConfig{
					IsPublic:         isPublic,
					FileName:         strings.Trim(row[2], " "),
					Num:              num,
					SkuName:          strings.Trim(row[4], " "),
					ImageCode:        strings.Trim(row[5], " "),
					MatCode:          strings.Trim(row[6], " "),
					ColorCode:        strings.Trim(row[7], " "),
					GiftCode:         strings.Trim(row[8], " "),
					ListPrice:        strings.Trim(row[9], " "),
					SinglePrice:      strings.Trim(row[10], " "),
					Stock:            strings.Trim(row[11], " "),
					LowPriceDrainage: strings.Trim(row[12], " "),
					MarketPrice:      strings.Trim(row[13], " "),
				})
			} else if configType == "主图" {
				g.CarouselGalleryConfigList = append(g.CarouselGalleryConfigList, CarouselGalleryConfig{
					IsPublic: isPublic,
					FileName: strings.Trim(row[2], " "),
					Num:      num,
				})
			} else if configType == "详情" {
				g.DetailGalleryConfigList = append(g.DetailGalleryConfigList, DetailGalleryConfig{
					IsPublic: isPublic,
					FileName: strings.Trim(row[2], " "),
					Num:      num,
				})
			}
		}
	}
	return
}
