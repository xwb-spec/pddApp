package yyui

import (
	"pddApp/common"
	"pddApp/pinduoduo/client"
	"strings"
	"sync"
)

type Goods struct {
	CatId           int64    `json:"CatId"`          //叶子类目ID
	CostTemplateId  int64    `json:"CostTemplateId"` //物流运费模板ID
	CountryId       int      `json:"CountryId"`      //地区/国家ID
	GoodsName       string   `json:"GoodsName"`      // 商品标题
	GoodsDesc       string   `json:"GoodsDesc"`      //商品描述
	GoodsType       int      `json:"GoodsType"`      //1-国内普通商品
	IsLowPrice      bool     `json:"IsLowPrice"`     //
	IsOnline        bool     `json:"IsOnline"`
	DetailGallery   []string `json:"DetailGallery"`
	CarouselGallery []string `json:"CarouselGallery"`
	SkuGallery      []string `json:"SkuGallery"`
}

func (s *ShowInput) GetImage() (imageList []string, err error) {
	goodsConfig, err := common.GetGoodsConfig(s.SkuConfigExcel.Text, s.SkuConfigSheetName.Text)
	if err != nil {
		s.ConsoleResult.SetText("[ERROR]: " + err.Error())
		return
	}
	pubDir := s.PublicDir.Text
	picDir := s.ImageDir.Text
	goods, err := common.GetGoods(s.GoodsExcel.Text, s.GoodsSheetName.Text, s.ModelImageExcel.Text, s.ModelImageSheetName.Text, picDir)
	if err != nil {
		s.ConsoleResult.SetText("[ERROR]: " + err.Error())
		return
	}
	for _, v := range goods {
		for _, i := range goodsConfig.DetailGalleryConfigList {
			if i.IsPublic {
				imageList = append(imageList, pubDir+"/"+i.FileName+".jpg")
			} else {
				imageList = append(imageList, v.ImageDir+"/"+i.FileName+".jpg")
			}
		}
		for _, i := range goodsConfig.CarouselGalleryConfigList {
			imageList = append(imageList, v.ImageDir+"/"+i.FileName+".jpg")
		}
		for _, i := range goodsConfig.SkuConfigList {
			imageList = append(imageList, v.ImageDir+"/"+i.FileName+".jpg")
		}
	}
	return
}
func (s *ShowInput) GetImageMap() (imageMap map[string]string, err error) {
	if err := common.LoadJson(s.UploadedImage.Text, &imageMap); err != nil {
		return nil, err
	}
	return imageMap, err
}
func (s *ShowInput) GetGoods() (goods []Goods, err error) {
	goodsMap, err := common.GetGoodsMap(s.GoodsExcel.Text, s.GoodsSheetName.Text)
	if err != nil {
		s.ConsoleResult.SetText("[ERROR]: " + err.Error())
		return
	}
	goodsImageMap, err := common.GetGoodsComparison(s.ModelImageExcel.Text, s.ModelImageSheetName.Text)
	if err != nil {
		s.ConsoleResult.SetText("[ERROR]: " + err.Error())
		return
	}
	goodsConfig, err := common.GetGoodsConfig(s.SkuConfigExcel.Text, s.SkuConfigSheetName.Text)
	if err != nil {
		s.ConsoleResult.SetText("[ERROR]: " + err.Error())
		return
	}
	imageMap, err := s.GetImageMap()
	pubDir := s.PublicDir.Text
	picDir := s.ImageDir.Text
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
				is, _ := common.IsPathExists(s.ImageDir.Text + "/" + *val.PicDir)
				if is {
					imageDir = *val.PicDir
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
		if common.IsEleExistsSlice("低价", isLowPriceList) {
			isLowPrice = true
		}
		if common.IsEleExistsSlice("是", isOnlineList) {
			isOnline = true
		}
		goodsDetailGallery := make([]string, len(goodsConfig.DetailGalleryConfigList))
		goodsCarouselGallery := make([]string, len(goodsConfig.CarouselGalleryConfigList))
		goodsSkulGallery := make([]string, len(goodsConfig.SkuConfigList))

		for _, i := range goodsConfig.DetailGalleryConfigList {
			if i.IsPublic {
				goodsDetailGallery[i.Num-1] = imageMap[pubDir+"/"+i.FileName+".jpg"]
			} else {
				goodsDetailGallery[i.Num-1] = imageMap[picDir+"/"+imageDir+"/"+i.FileName+".jpg"]
			}
		}
		for _, i := range goodsConfig.CarouselGalleryConfigList {
			goodsCarouselGallery[i.Num-1] = imageMap[picDir+"/"+imageDir+"/"+i.FileName+".jpg"]
		}
		for _, i := range goodsConfig.SkuConfigList {
			goodsSkulGallery[i.Num-1] = imageMap[picDir+"/"+imageDir+"/"+i.FileName+".jpg"]
		}
		goods = append(goods, Goods{
			CatId:           1234,
			CostTemplateId:  1234,
			CountryId:       1234,
			GoodsName:       k,
			GoodsDesc:       k,
			GoodsType:       1,
			IsOnline:        isOnline,
			IsLowPrice:      isLowPrice,
			DetailGallery:   goodsDetailGallery,
			CarouselGallery: goodsCarouselGallery,
			SkuGallery:      goodsSkulGallery,
		})
	}
	return goods, nil
}

func (s *ShowInput) UploadImage() (err error) {
	imageList, err := s.GetImage()
	if err != nil {
		return err
	}
	var syncMap sync.Map
	var wg sync.WaitGroup
	limit := 10 // 控制并发数
	ch := make(chan struct{}, limit)
	defer close(ch)
	for _, i := range imageList {
		_, ok := syncMap.Load(i)
		if !ok {
			ch <- struct{}{}
			go client.UploadImage(i, &wg, &syncMap, ch)
		}
	}
	wg.Wait()
	imageMap := make(map[string]string)
	// Range遍历所有sync.Map中的键值对
	syncMap.Range(func(k, v interface{}) bool {
		imageMap[k.(string)] = v.(string)
		return true
	})
	// 创建文件
	if err = common.CreateJson(s.UploadedImage.Text, imageMap); err != nil {
		return err
	}
	return nil
}
