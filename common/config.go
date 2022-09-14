package common

import (
	"strconv"
	"strings"
)

type Comparison struct {
	PicDir *string
	Brand  *string
}

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
	SkuConfigList             []SkuConfig
	DetailGalleryConfigList   []DetailGalleryConfig
	CarouselGalleryConfigList []CarouselGalleryConfig
}

func GetGoodsComparison(excelPath, excelSheet string) (ComparisonMap map[string]*Comparison, err error) {
	rows, err := GetExcelRows(excelPath, excelSheet)
	if err != nil {
		return
	}
	ComparisonMap = make(map[string]*Comparison)
	rows.Next()
	for rows.Next() {
		row, _ := rows.Columns()
		if len(row) == 3 {
			key := strings.Trim(row[0], " ")
			if key != "" {
				key = strings.ToLower(key)
				picDir := strings.Trim(row[1], " ")
				brand := strings.Trim(row[2], " ")
				ComparisonMap[key] = &Comparison{
					PicDir: &picDir,
					Brand:  &brand,
				}
			}
		}
	}
	// 关闭迭代
	if err = rows.Close(); err != nil {
		return ComparisonMap, err
	}
	return ComparisonMap, nil
}

func GetGoodsConfig(excelPath, excelSheet string) (goodsConfig GoodsConfig, err error) {
	rows, err := GetExcelRows(excelPath, excelSheet)
	if err != nil {
		return
	}
	rows.Next()
	for rows.Next() {
		row, _ := rows.Columns()
		if len(row) > 0 {
			isPublicVal := strings.Trim(row[1], " ")
			var isPublic bool
			if isPublicVal == "公用" {
				isPublic = true
			} else {
				isPublic = false
			}
			num, _ := strconv.Atoi(strings.Trim(row[3], " ")) // 转换为数字类型
			configType := strings.Trim(row[0], " ")           //去除两端空格
			if configType == "sku" && len(row) == 14 {
				goodsConfig.SkuConfigList = append(goodsConfig.SkuConfigList, SkuConfig{
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
			} else if configType == "主图" && len(row) == 4 {
				goodsConfig.CarouselGalleryConfigList = append(goodsConfig.CarouselGalleryConfigList, CarouselGalleryConfig{
					IsPublic: isPublic,
					FileName: strings.Trim(row[2], " "),
					Num:      num,
				})
			} else if configType == "详情" && len(row) == 4 {
				goodsConfig.DetailGalleryConfigList = append(goodsConfig.DetailGalleryConfigList, DetailGalleryConfig{
					IsPublic: isPublic,
					FileName: strings.Trim(row[2], " "),
					Num:      num,
				})
			}
		}
	}
	// 关闭迭代
	if err = rows.Close(); err != nil {
		return goodsConfig, err
	}
	return goodsConfig, nil
}
