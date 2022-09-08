package yyui

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"pddApp/common"
	"strings"
)

func CheckExcelSheet(excelPath, sheetName string) bool {
	f, err := excelize.OpenFile(excelPath)
	if err != nil {
		return false
	}
	defer f.Close()
	_, err = f.GetRows(sheetName)
	if err != nil {
		return false
	}
	return true
}
func (s *ShowInput) CheckInput() {
	log.Println("走checkInput函数")
	resultConsole := s.ConsoleResult.Text + "\n"
	// 检测登录信息
	if s.ShopId.Text == "" {
		s.ConsoleResult.SetText(resultConsole + "店铺id为空: [ERROR] 请输入店铺id")
		return
	}
	if s.ShopName.Text == "" {
		s.ConsoleResult.SetText(resultConsole + "店铺名为空: [ERROR] 请输入店铺名")
		return
	}
	// 检测店铺模板
	if s.FreightTmp.Selected == "" {
		s.ConsoleResult.SetText(resultConsole + "运费模板为空: [ERROR] 请输入运费模板")
		return
	}
	log.Println("走这里")
	// 检测公用图片
	if s.PubFileDir.Text == "" {
		s.ConsoleResult.SetText(resultConsole + "公用文件目录为空: [ERROR] 请选择或输入公用文件目录" + s.PubFileDir.Text)
		return
	}
	isPathExist, err := common.IsPathExists(s.PubFileDir.Text)
	if err != nil {
		s.ConsoleResult.SetText(resultConsole + "公用文件目录出错: [ERROR]: %s" + err.Error())
		return
	}
	if !isPathExist {
		s.ConsoleResult.SetText(resultConsole + "公用文件目录出错: [ERROR] 公用文件目录不存在" + s.PubFileDir.Text)
		return
	}
	// 检测套图
	if s.PicKitDir.Text == "" {
		s.ConsoleResult.SetText(resultConsole + "套图文件目录为空: [ERROR] 请选择或输入套图文件目录" + s.PicKitDir.Text)
		return
	}
	isPathExist, err = common.IsPathExists(s.PicKitDir.Text)
	if err != nil {
		s.ConsoleResult.SetText(resultConsole + "套图文件目录出错: [ERROR]: %s" + err.Error())
		return
	}
	if !isPathExist {
		s.ConsoleResult.SetText(resultConsole + "套图文件目录出错: [ERROR] 套图文件目录不存在" + s.PubFileDir.Text)
		return
	}
	// 检测配置文件路径
	if s.UploadedImageConfig.Text == "" {
		s.ConsoleResult.SetText(resultConsole + "已上传图片文件配置为空: [ERROR] 请选择或输入已上传图片文件配置" + s.UploadedImageConfig.Text)
		return
	}
	isPathExist, err = common.IsPathExists(s.UploadedImageConfig.Text)
	if err != nil {
		s.ConsoleResult.SetText(resultConsole + "已上传图片文件配置出错: [ERROR]: %s" + err.Error())
		return
	}
	if !isPathExist {
		s.ConsoleResult.SetText(resultConsole + "已上传图片文件配置出错: [ERROR] 已上传图片文件配置不存在" + s.UploadedImageConfig.Text)
		return
	}
	// 检测商品配置表
	if s.ShopExcel.Text == "" {
		s.ConsoleResult.SetText(resultConsole + "商品配置表为空: [ERROR] 请选择或输入商品配置表" + s.ShopExcel.Text)
		return
	}
	isPathExist, err = common.IsPathExists(s.ShopExcel.Text)
	if err != nil {
		s.ConsoleResult.SetText(resultConsole + "商品配置表出错: [ERROR]: %s" + err.Error())
		return
	}
	if !isPathExist {
		s.ConsoleResult.SetText(resultConsole + "商品配置表出错: [ERROR] 商品配置表不存在" + s.ShopExcel.Text)
		return
	}
	// 检测sku配置表
	if s.SkuExcel.Text == "" {
		s.ConsoleResult.SetText(resultConsole + "sku配置表为空: [ERROR] 请选择或输入sku配置表" + s.SkuExcel.Text)
		return
	}
	isPathExist, err = common.IsPathExists(s.SkuExcel.Text)
	if err != nil {
		s.ConsoleResult.SetText(resultConsole + "sku配置表出错: [ERROR]: %s" + err.Error())
		return
	}
	if !isPathExist {
		s.ConsoleResult.SetText(resultConsole + "sku配置表出错: [ERROR] sku配置表不存在" + s.SkuExcel.Text)
		return
	}
	// 检测型号配置表
	if s.ModelExcel.Text == "" {
		s.ConsoleResult.SetText(resultConsole + "型号对照配置表为空: [ERROR] 请选择或输入型号对照配置表" + s.ModelExcel.Text)
		return
	}
	isPathExist, err = common.IsPathExists(s.ModelExcel.Text)
	if err != nil {
		s.ConsoleResult.SetText(resultConsole + "型号对照配置表出错: [ERROR]: %s" + err.Error())
		return
	}
	if !isPathExist {
		s.ConsoleResult.SetText(resultConsole + "型号对照配置表出错: [ERROR] 型号对照配置表不存在" + s.ModelExcel.Text)
		return
	}
	// 检测表单
	if s.ShopSheetName.Text == "" {
		s.ConsoleResult.SetText(resultConsole + "商品配置表表单为空: [ERROR] 请填写")
		return
	} else {
		if !CheckExcelSheet(s.ShopExcel.Text, s.ShopSheetName.Text) {
			s.ConsoleResult.SetText(resultConsole + "商品配置表表单不存在: [ERROR] 请检查商品表单")
			return
		}
	}
	if s.ModelSheetName.Text == "" {
		s.ConsoleResult.SetText(resultConsole + "型号对照表表单为空: [ERROR] 请填写")
		return
	} else {
		if !CheckExcelSheet(s.ModelExcel.Text, s.ModelSheetName.Text) {
			log.Println("型号有没有走这里")
			s.ConsoleResult.SetText(resultConsole + "型号对照表表单不存在: [ERROR] 请检查型号对照表表单")
			return
		}
	}
	if s.SkuSheetName.Text == "" {
		s.ConsoleResult.SetText(resultConsole + "sku配置表表单为空: [ERROR] 请填写")
		return
	} else {
		if !CheckExcelSheet(s.SkuExcel.Text, s.SkuSheetName.Text) {
			s.ConsoleResult.SetText(resultConsole + "sku配置表表单不存在: [ERROR] 请检查sku配置表表单")
			return
		}
	}
	if s.AttrSheetName.Text == "" {
		s.ConsoleResult.SetText(resultConsole + "属性配置表表单为空: [ERROR] 请填写")
		return
	} else {
		if !CheckExcelSheet(s.SkuExcel.Text, s.AttrSheetName.Text) {
			s.ConsoleResult.SetText(resultConsole + "属性配置表表单不存在: [ERROR] 请检查属性配置表表单")
			return
		}
	}
	s.CheckPubImagePath()
	errs := s.CheckImagePath()
	for _, err := range errs {
		s.ConsoleResult.SetText(resultConsole + err.Error())
	}
	s.ConsoleResult.SetText(resultConsole + "检测成功")
	return
}

// 检查套图
func (s *ShowInput) CheckImagePath() (errs []error) {
	modelRows := common.GetExcel(s.ShopExcel.Text, s.ShopSheetName.Text)
	var modelList []string
	for i, row := range modelRows {
		if i > 0 && len(row) != 0 {
			modelName := strings.Trim(row[1], " ")
			if modelName != "" {
				if !common.IsEleExistsSlice(modelName, modelList) {
					modelList = append(modelList, modelName)
				}
			}
		}
	}
	compMap := common.GetComp(s.ModelExcel.Text, s.ModelSheetName.Text)
	g := common.GoodsConfig{}
	g.GetConfig(s.SkuExcel.Text, s.SkuSheetName.Text)
	for _, m := range modelList {
		val, ok := compMap[m] // 从map查找图片目录是否存在
		if !ok {
			errs = append(errs, errors.New(fmt.Sprintf("型号[%s]对应的图片目录没有找到", m)))
		} else {
			imageDir := s.PicKitDir.Text + "/" + *val.PicDir
			for _, d := range g.DetailGalleryConfigList { // 检查详情图是否完全
				if !d.IsPublic { // 不是处理公用图
					imagePath := imageDir + "/" + d.FileName + ".jpg"
					b, _ := common.IsPathExists(imagePath)
					if !b {
						errs = append(errs, errors.New(fmt.Sprintf("详情图[%s]没有找到", imagePath)))
					}
				}
			}
			for _, d := range g.SkuConfigList { // 检查详情图是否完全
				if !d.IsPublic {
					imagePath := imageDir + "/" + d.FileName + ".jpg"
					b, _ := common.IsPathExists(imagePath)
					if !b {
						errs = append(errs, errors.New(fmt.Sprintf("sku图[%s]没有找到", imagePath)))
					}
				}
			}
			for _, d := range g.CarouselGalleryConfigList { // 检查详情图是否完全
				if !d.IsPublic {
					imagePath := imageDir + "/" + d.FileName + ".jpg"
					b, _ := common.IsPathExists(imagePath)
					if !b {
						errs = append(errs, errors.New(fmt.Sprintf("轮播图[%s]没有找到", imagePath)))
					}
				}
			}
		}

	}
	return
}

// 检查公用图片
func (s *ShowInput) CheckPubImagePath() {
	resultConsole := s.ConsoleResult.Text + "\n"
	pubFilePaths := []string{s.PubFileDir.Text + "/首页.png", s.PubFileDir.Text + "/尾页.jpg"}
	for _, f := range pubFilePaths {
		isPathExist, err := common.IsPathExists(f)
		if err != nil {
			s.ConsoleResult.SetText(resultConsole + "检测图片失败: [ERROR]: %s" + err.Error())
			return
		}
		if !isPathExist {
			s.ConsoleResult.SetText(resultConsole + "检测图片失败: [ERROR] 公用文件图片不存在," + f)
			return
		}
	}
	// 检测套图
}

// 检查配置
