package common

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"strings"
)

// 装载excel数据
type ExcelGoods struct {
	ExcelPath      string
	ExcelSheetName string
	Title          string   // 商品标题
	Model          []string // 商品型号
	BrandPrefix    []string // 品牌前缀
	IsLowPrice     []string // 是否低价
	IsOnline       []string // 是否已上架
	SkuDisplay     []string // sku显示
	//SkuModel []string // sku型号名 //
	//PictureDir []string // 对应的图片型号文件夹名称
	//Brand      []string // 品牌
}

type ExcelModelPictureComparison struct {
	ExcelPath      string
	ExcelSheetName string
}

func (e *ExcelModelPictureComparison) GetMe() map[string][]string {
	f, err := excelize.OpenFile(e.ExcelPath)
	if err != nil {
		log.Println(err)

	}
	defer f.Close()
	rows, err := f.GetRows(e.ExcelSheetName)
	if err != nil {
		log.Println(err)

	}
	mapData := make(map[string][]string)
	for _, row := range rows[1:] {
		if len(row) < 3 { // 跳过没有三栏的数据
			continue
		}
		if strings.Trim(row[0], " ") != "" {
			mapData[row[0]] = row[1:3]
		}
	}
	return mapData
}
func (e *ExcelGoods) GetMe() {
	f, err := excelize.OpenFile(e.ExcelPath)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	rows, err := f.GetRows(e.ExcelSheetName)
	if err != nil {
		log.Println(err)
		return
	}
	//ex2 := ModelPictureComparison{ExcelPath: "/Users/xiewenbin/Downloads/批量上商品app/模板-型号图片对照表.xlsx", ExcelSheetName: "直边tpu"}
	mapData := make(map[string][][]string)
	var key string
	for _, row := range rows[1:] {
		if len(row) != 6 { // 跳过空行
			//log.Printf("表单[%s]第%d行数据%v有问题,跳过.\n", e.ExcelSheetName, i+2, row)
			continue
		}
		if row[0] != "" {
			key = row[0]
			mapData[key] = append(mapData[key], row[1:])
		} else {
			mapData[key] = append(mapData[key], row[1:])
		}
	}
	var Goods []ExcelGoods
	for k, v := range mapData {
		var Good ExcelGoods
		Good.Title = k
		for _, r := range v {
			col1, col2, col3, col4, col5 := strings.Trim(r[0], " "), // 处理两端空格
				strings.Trim(r[1], " "), // 处理两端空格
				strings.Trim(r[2], " "), // 处理两端空格
				strings.Trim(r[3], " "), // 处理两端空格
				strings.Trim(r[4], " ") // 处理两端空格
			if col1 != "" {
				Good.Model = append(Good.Model, col1)
				// 查找图片
				//v, ok := PicData[col1]
				//if ok {
				//	Good.PictureDir = append(Good.PictureDir, v[0])
				//	Good.Brand = append(Good.Brand, v[1])
				//} else {
				//	log.Printf("型号 [%s] 没有图片信息\n", col1)
				//	Good.PictureDir = append(Good.PictureDir, "")
				//	Good.Brand = append(Good.Brand, "")
				//}
			}
			if col2 != "" {
				Good.BrandPrefix = append(Good.BrandPrefix, col2)
			}
			if col3 != "" {
				Good.IsLowPrice = append(Good.IsLowPrice, col3)
			}
			if col4 != "" {
				Good.IsOnline = append(Good.IsOnline, col4)
			}
			if col5 != "" {
				Good.SkuDisplay = append(Good.SkuDisplay, col5)
			}
		}
		Goods = append(Goods, Good)
	}
	for k, v := range Goods {
		fmt.Println(k, v)
	}
}
