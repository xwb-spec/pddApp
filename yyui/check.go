package yyui

import (
	"fmt"
	"pddApp/common"
)

func (s *ShowInput) CheckInput() {
	// 检测登录信息
	if s.ShopId.Text == "" {
		s.ConsoleResult.SetText("[ERROR]: 店铺id为空,请输入店铺id")
		return
	}
	if s.ShopName.Text == "" {
		s.ConsoleResult.SetText("[ERROR]: 店铺名为空,请输入店铺名")
		return
	}
	// 检测店铺模板
	if s.FreightTmp.Selected == "" {
		s.ConsoleResult.SetText("[ERROR]: 运费模板为空,请输入运费模板")
		return
	}
	// 检测公用图片
	if s.PubFileDir.Text == "" {
		s.ConsoleResult.SetText("[ERROR]: 公用文件目录为空,请选择或输入公用文件目录" + s.PubFileDir.Text)
		return
	}
	isPathExist, err := common.IsPathExists(s.PubFileDir.Text)
	if err != nil {
		s.ConsoleResult.SetText("[ERROR]: 公用文件目录出错, %s" + err.Error())
		return
	}
	if !isPathExist {
		s.ConsoleResult.SetText("[ERROR]: 公用文件目录出错,公用文件目录不存在" + s.PubFileDir.Text)
		return
	}
	// 检测套图
	if s.PicKitDir.Text == "" {
		s.ConsoleResult.SetText("[ERROR]: 套图文件目录为空,请选择或输入套图文件目录" + s.PicKitDir.Text)
		return
	}
	isPathExist, err = common.IsPathExists(s.PicKitDir.Text)
	if err != nil {
		s.ConsoleResult.SetText("[ERROR]: 套图文件目录出错, %s" + err.Error())
		return
	}
	if !isPathExist {
		s.ConsoleResult.SetText("[ERROR]: 套图文件目录出错,套图文件目录不存在" + s.PubFileDir.Text)
		return
	}
	// 检测配置文件路径
	if s.UploadedImageConfig.Text == "" {
		s.ConsoleResult.SetText("[ERROR]: 已上传图片文件配置为空,请选择或输入已上传图片文件配置" + s.UploadedImageConfig.Text)
		return
	}
	isPathExist, err = common.IsPathExists(s.UploadedImageConfig.Text)
	if err != nil {
		s.ConsoleResult.SetText("[ERROR]: 已上传图片文件配置出错, %s" + err.Error())
		return
	}
	if !isPathExist {
		s.ConsoleResult.SetText("[ERROR]: 已上传图片文件配置出错,已上传图片文件配置不存在" + s.UploadedImageConfig.Text)
		return
	}
	// 检测商品配置表
	if s.ShopExcel.Text == "" {
		s.ConsoleResult.SetText("[ERROR]: 商品配置表为空,请选择或输入商品配置表" + s.ShopExcel.Text)
		return
	}
	isPathExist, err = common.IsPathExists(s.ShopExcel.Text)
	if err != nil {
		s.ConsoleResult.SetText("[ERROR]: 商品配置表出错, %s" + err.Error())
		return
	}
	if !isPathExist {
		s.ConsoleResult.SetText("[ERROR]: 商品配置表出错,商品配置表不存在" + s.ShopExcel.Text)
		return
	}
	// 检测sku配置表
	if s.SkuExcel.Text == "" {
		s.ConsoleResult.SetText("[ERROR]: sku配置表为空,请选择或输入sku配置表" + s.SkuExcel.Text)
		return
	}
	isPathExist, err = common.IsPathExists(s.SkuExcel.Text)
	if err != nil {
		s.ConsoleResult.SetText("[ERROR]: sku配置表出错, %s" + err.Error())
		return
	}
	if !isPathExist {
		s.ConsoleResult.SetText("[ERROR]: sku配置表出错,sku配置表不存在" + s.SkuExcel.Text)
		return
	}
	// 检测型号配置表
	if s.ModelExcel.Text == "" {
		s.ConsoleResult.SetText("[ERROR]: 型号对照配置表为空,请选择或输入型号对照配置表" + s.ModelExcel.Text)
		return
	}
	isPathExist, err = common.IsPathExists(s.ModelExcel.Text)
	if err != nil {
		s.ConsoleResult.SetText("[ERROR]: 型号对照配置表出错, %s" + err.Error())
		return
	}
	if !isPathExist {
		s.ConsoleResult.SetText("[ERROR]: 型号对照配置表出错,型号对照配置表不存在" + s.ModelExcel.Text)
		return
	}
	// 检测表单
	if s.ShopSheetName.Text == "" {
		s.ConsoleResult.SetText("[ERROR]: 商品配置表表单为空,请填写")
		return
	} else {
		if !common.IsSheetExists(s.ShopExcel.Text, s.ShopSheetName.Text) {
			s.ConsoleResult.SetText("[ERROR]: 商品配置表表单不存在,请检查商品表单")
			return
		}
	}
	if s.ModelSheetName.Text == "" {
		s.ConsoleResult.SetText("[ERROR]: 型号对照表表单为空, 请填写")
		return
	} else {
		if !common.IsSheetExists(s.ModelExcel.Text, s.ModelSheetName.Text) {
			s.ConsoleResult.SetText("[ERROR]: 型号对照表表单不存在,请检查型号对照表表单")
			return
		}
	}
	if s.SkuSheetName.Text == "" {
		s.ConsoleResult.SetText("[ERROR]: sku配置表表单为空,请填写")
		return
	} else {
		if !common.IsSheetExists(s.SkuExcel.Text, s.SkuSheetName.Text) {
			s.ConsoleResult.SetText("[ERROR]: sku配置表表单不存在,请检查sku配置表表单")
			return
		}
	}
	if s.AttrSheetName.Text == "" {
		s.ConsoleResult.SetText("[ERROR]: 属性配置表表单为空,请填写")
		return
	} else {
		if !common.IsSheetExists(s.SkuExcel.Text, s.AttrSheetName.Text) {
			s.ConsoleResult.SetText("[ERROR]: 属性配置表表单不存在,请检查属性配置表表单")
			return
		}
	}
	s.CheckImagePath()
}

// 检查套图
func (s *ShowInput) CheckImagePath() {
	goodsMap, err := common.GetGoodsMap(s.ShopExcel.Text, s.ShopSheetName.Text)
	if err != nil {
		s.ConsoleResult.SetText(fmt.Sprintf("[ERROR]: 读取商品表格数据失败, %s", s.ShopExcel.Text))
		return
	}
	compMap, err := common.GetGoodsComparison(s.ModelExcel.Text, s.ModelSheetName.Text)
	if err != nil {
		s.ConsoleResult.SetText(fmt.Sprintf("[ERROR]: 读取型号图片对照表格数据失败, %s", s.ModelExcel.Text))
		return
	}
	goodsConfig, err := common.GetGoodsConfig(s.SkuExcel.Text, s.SkuSheetName.Text)
	if err != nil {
		s.ConsoleResult.SetText(fmt.Sprintf("[ERROR]: 读取sku配置表格数据失败, %s", s.ModelExcel.Text))
		return
	}
	for k, v := range goodsMap {
		isExists := false
		value := ""
		for _, d := range v {
			val, ok := compMap[d.Model] // 从map查找图片目录是否存在
			if ok {
				isExists = true
				value = *val.PicDir
				break
			}
		}
		if !isExists {
			s.ConsoleResult.SetText(fmt.Sprintf("[ERROR] :商品[%s]对应的图片目录不存在", k))
			return
		} else {
			imageDir := s.PicKitDir.Text + "/" + value
			for _, d := range goodsConfig.DetailGalleryConfigList { // 检查详情图是否完全
				if !d.IsPublic { // 不是处理公用图
					imagePath := imageDir + "/" + d.FileName + ".jpg"
					b, err := common.IsPathExists(imagePath)
					if err != nil {
						s.ConsoleResult.SetText(fmt.Sprintf("[ERROR]: 详情图[%s]", err))
						return
					}
					if !b {
						s.ConsoleResult.SetText(fmt.Sprintf("[ERROR]: 详情图[%s]不存在", imagePath))
						return
					}
				} else {
					b, err := common.IsPathExists(s.PubFileDir.Text + "/" + d.FileName + ".png")
					if err != nil {
						s.ConsoleResult.SetText(fmt.Sprintf("[ERROR]: 公用图[%s]", err))
						return
					}
					if !b {
						s.ConsoleResult.SetText(fmt.Sprintf("[ERROR]: 公用图[%s]不存在", s.PubFileDir.Text+"/"+d.FileName+".jpg"))
						return
					}
				}
			}
			for _, d := range goodsConfig.SkuConfigList { // 检查详情图是否完全
				if !d.IsPublic {
					imagePath := imageDir + "/" + d.FileName + ".jpg"
					b, err := common.IsPathExists(imagePath)
					if err != nil {
						s.ConsoleResult.SetText(fmt.Sprintf("[ERROR]: sku图[%s]", err))
						return
					}
					if !b {
						s.ConsoleResult.SetText(fmt.Sprintf("[ERROR]: sku图[%s]不存在", imagePath))
						return
					}
				}
			}
			for _, d := range goodsConfig.CarouselGalleryConfigList { // 检查详情图是否完全
				if !d.IsPublic {
					imagePath := imageDir + "/" + d.FileName + ".jpg"
					b, err := common.IsPathExists(imagePath)
					if err != nil {
						s.ConsoleResult.SetText(fmt.Sprintf("[ERROR]: 轮播图[%s]", err))
						return
					}
					if !b {
						s.ConsoleResult.SetText(fmt.Sprintf("[ERROR]: 轮播图[%s]不存在", imagePath))
						return
					}
				}
			}
		}

	}
	s.ConsoleResult.SetText("[OK]: 图片检测通过")
}

// 检查配置
func (s *ShowInput) CheckConfig() {
	goodsConfig, err := common.GetGoodsConfig(s.SkuExcel.Text, s.SkuSheetName.Text)
	if err != nil {
		s.ConsoleResult.SetText(fmt.Sprintf("[ERROR]: 读取sku配置表数据失败, %s", err.Error()))
		return
	}
	if len(goodsConfig.DetailGalleryConfigList) < 1 {
		s.ConsoleResult.SetText("[ERROR]: 详情图配置不存在")
		return
	}
	if len(goodsConfig.SkuConfigList) < 1 { // 检查详情图是否完全
		s.ConsoleResult.SetText("[ERROR]: sku图配置不存在")
		return
	}
	if len(goodsConfig.CarouselGalleryConfigList) < 0 { // 检查详情图是否完全
		s.ConsoleResult.SetText("[ERROR]: 轮播图配置不存在")
		return
	}
	s.ConsoleResult.SetText("[OK]: 配置检测通过")
}
