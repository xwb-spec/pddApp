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

type Goods struct {
	CatId          int64  `json:"CatId"`          //叶子类目ID
	CostTemplateId int64  `json:"CostTemplateId"` //物流运费模板ID
	CountryId      int    `json:"CountryId"`      //地区/国家ID
	GoodsName      string `json:"GoodsName"`      // 商品标题
	GoodsDesc      string `json:"GoodsDesc"`      //商品描述
	GoodsType      int    `json:"GoodsType"`      //1-国内普通商品
	IsLowPrice     bool   `json:"IsLowPrice"`     //
	IsOnline       bool   `json:"IsOnline"`
	ImageDir       string `json:"ImageDir"`
}
type GoodsProperties struct {
	Brand string
}

func GetExcelRows(excelPath, excelSheet string) (rows *excelize.Rows, err error) {
	f, err := excelize.OpenFile(excelPath)
	if err != nil {
		return
	}
	defer f.Close()
	rows, err = f.Rows(excelSheet)
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
	rows.Next() // 跳过第一行
	for rows.Next() {
		row, _ := rows.Columns()
		if len(row) == 6 { // 跳过空行
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
	// 关闭迭代
	if err = rows.Close(); err != nil {
		return goodsMap, err
	}
	return goodsMap, nil
}

func GetGoods(goodsExcel, goodsSheetName, modelImageExcel, modelImageSheetName, parentImageDir string) (goods []Goods, err error) {
	goodsMap, err := GetGoodsMap(goodsExcel, goodsSheetName)
	if err != nil {
		return nil, err
	}
	goodsImageMap, err := GetGoodsComparison(modelImageExcel, modelImageSheetName)
	if err != nil {
		return nil, err
	}
	for k, v := range goodsMap {
		var modelList []string
		var skuList []string
		var isLowPriceList []string
		var isOnlineList []string
		var isLowPrice bool
		var isOnline bool
		var imageDir string
		for _, l := range v {
			key := strings.ToLower(l.Model)
			val, ok := goodsImageMap[key] // 从map查找图片目录是否存在
			if ok {
				b, _ := IsPathExists(parentImageDir + "/" + *val.PicDir)
				if b {
					imageDir = parentImageDir + "/" + *val.PicDir
					break
				}
			}
		}
		for _, l := range v {
			modelList = append(modelList, l.Model)
			skuList = append(skuList, l.SkuDisplay)
			isLowPriceList = append(isLowPriceList, l.IsLowPrice)
			isOnlineList = append(isOnlineList, l.IsOnline)

		}
		if IsEleExistsSlice("低价", isLowPriceList) {
			isLowPrice = true
		}
		if IsEleExistsSlice("是", isOnlineList) {
			isOnline = true
		}
		goods = append(goods, Goods{
			CatId:          1234,
			CostTemplateId: 1234,
			CountryId:      1234,
			GoodsName:      k,
			GoodsDesc:      k,
			GoodsType:      1,
			IsOnline:       isOnline,
			IsLowPrice:     isLowPrice,
			ImageDir:       imageDir,
		})
	}
	return goods, nil
}
