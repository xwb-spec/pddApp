package yyui

import (
	"github.com/xuri/excelize/v2"
	"pddApp/common"
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
	resultConsole := s.ConsoleResult.Text + "\n"
	// 检测登录信息
	if s.ShopId == "" {
		s.ConsoleResult.SetText(resultConsole + "店铺id为空: [ERROR] 请输入店铺id")
		return
	}
	if s.ShopName == "" {
		s.ConsoleResult.SetText(resultConsole + "店铺名为空: [ERROR] 请输入店铺名")
		return
	}
	// 检测店铺模板
	if s.FreightTmp == "" {
		s.ConsoleResult.SetText(resultConsole + "运费模板为空: [ERROR] 请输入运费模板")
		return
	}
	// 检测公用图片
	if s.PubFileDir == "" {
		s.ConsoleResult.SetText(resultConsole + "公用文件目录为空: [ERROR] 请选择或输入公用文件目录" + s.PubFileDir)
		return
	}
	isPathExist, err := common.IsPathExists(s.PubFileDir)
	if err != nil {
		s.ConsoleResult.SetText(resultConsole + "公用文件目录出错: [ERROR]: %s" + err.Error())
		return
	}
	if !isPathExist {
		s.ConsoleResult.SetText(resultConsole + "公用文件目录出错: [ERROR] 公用文件目录不存在" + s.PubFileDir)
		return
	}
	// 检测套图
	if s.PicKitDir == "" {
		s.ConsoleResult.SetText(resultConsole + "套图文件目录为空: [ERROR] 请选择或输入套图文件目录" + s.PicKitDir)
		return
	}
	isPathExist, err = common.IsPathExists(s.PicKitDir)
	if err != nil {
		s.ConsoleResult.SetText(resultConsole + "套图文件目录出错: [ERROR]: %s" + err.Error())
		return
	}
	if !isPathExist {
		s.ConsoleResult.SetText(resultConsole + "套图文件目录出错: [ERROR] 套图文件目录不存在" + s.PubFileDir)
		return
	}
	// 检测配置文件路径
	if s.UploadedImageConfig == "" {
		s.ConsoleResult.SetText(resultConsole + "已上传图片文件配置为空: [ERROR] 请选择或输入已上传图片文件配置" + s.UploadedImageConfig)
		return
	}
	isPathExist, err = common.IsPathExists(s.UploadedImageConfig)
	if err != nil {
		s.ConsoleResult.SetText(resultConsole + "已上传图片文件配置出错: [ERROR]: %s" + err.Error())
		return
	}
	if !isPathExist {
		s.ConsoleResult.SetText(resultConsole + "已上传图片文件配置出错: [ERROR] 已上传图片文件配置不存在" + s.UploadedImageConfig)
		return
	}
	// 检测商品配置表
	if s.ShopExcel == "" {
		s.ConsoleResult.SetText(resultConsole + "商品配置表为空: [ERROR] 请选择或输入商品配置表" + s.ShopExcel)
		return
	}
	isPathExist, err = common.IsPathExists(s.ShopExcel)
	if err != nil {
		s.ConsoleResult.SetText(resultConsole + "商品配置表出错: [ERROR]: %s" + err.Error())
		return
	}
	if !isPathExist {
		s.ConsoleResult.SetText(resultConsole + "商品配置表出错: [ERROR] 商品配置表不存在" + s.ShopExcel)
		return
	}
	// 检测sku配置表
	if s.SkuExcel == "" {
		s.ConsoleResult.SetText(resultConsole + "sku配置表为空: [ERROR] 请选择或输入sku配置表" + s.SkuExcel)
		return
	}
	isPathExist, err = common.IsPathExists(s.SkuExcel)
	if err != nil {
		s.ConsoleResult.SetText(resultConsole + "sku配置表出错: [ERROR]: %s" + err.Error())
		return
	}
	if !isPathExist {
		s.ConsoleResult.SetText(resultConsole + "sku配置表出错: [ERROR] sku配置表不存在" + s.SkuExcel)
		return
	}
	// 检测型号配置表
	if s.ModelExcel == "" {
		s.ConsoleResult.SetText(resultConsole + "型号对照配置表为空: [ERROR] 请选择或输入型号对照配置表" + s.ModelExcel)
		return
	}
	isPathExist, err = common.IsPathExists(s.ModelExcel)
	if err != nil {
		s.ConsoleResult.SetText(resultConsole + "型号对照配置表出错: [ERROR]: %s" + err.Error())
		return
	}
	if !isPathExist {
		s.ConsoleResult.SetText(resultConsole + "型号对照配置表出错: [ERROR] 型号对照配置表不存在" + s.ModelExcel)
		return
	}
	// 检测表单
	if s.ShopSheetName == "" {
		s.ConsoleResult.SetText(resultConsole + "商品配置表表单为空: [ERROR] 请填写")
		return
	} else {
		if !CheckExcelSheet(s.ShopExcel, s.ShopSheetName) {
			s.ConsoleResult.SetText(resultConsole + "商品配置表表单不存在: [ERROR] 请检查商品表单")
			return
		}
	}
	if s.ModelSheetName == "" {
		s.ConsoleResult.SetText(resultConsole + "型号对照表表单为空: [ERROR] 请填写")
		if !CheckExcelSheet(s.ModelExcel, s.ModelSheetName) {
			s.ConsoleResult.SetText(resultConsole + "型号对照表表单不存在: [ERROR] 请检查型号对照表表单")
			return
		}
	}
	if s.SkuSheetName == "" {
		s.ConsoleResult.SetText(resultConsole + "sku配置表表单为空: [ERROR] 请填写")
		if !CheckExcelSheet(s.SkuExcel, s.SkuSheetName) {
			s.ConsoleResult.SetText(resultConsole + "sku配置表表单不存在: [ERROR] 请检查sku配置表表单")
			return
		}
	}
	if s.AttrSheetName == "" {
		s.ConsoleResult.SetText(resultConsole + "属性配置表表单为空: [ERROR] 请填写")
		if !CheckExcelSheet(s.SkuSheetName, s.AttrSheetName) {
			s.ConsoleResult.SetText(resultConsole + "属性配置表表单不存在: [ERROR] 请检查属性配置表表单")
			return
		}
	}
}
