package yyui

import (
	"encoding/json"
	"log"
	"os"
)

type CurInput struct {
	ShopName            string `json:"ShopName"`            // 店铺名
	ShopId              string `json:"ShopId"`              // 店铺id
	FreightTmp          string `json:"FreightTmp"`          // 运费模板
	PicKitDir           string `json:"PicKitDir"`           // 套图文件目录
	PubFileDir          string `json:"PubFileDir"`          // 公共文件目录
	UploadedImageConfig string `json:"UploadedImageConfig"` // 已上传图片文件配置
	ShopExcel           string `json:"ShopExcel"`           //商品配置表
	SkuExcel            string `json:"SkuExcel"`            // sku配置表
	ModelExcel          string `json:"ModelExcel"`          //型号对照表
	ShopSheetName       string `json:"ShopSheetName"`       // 商品表单名
	SkuSheetName        string `json:"SkuSheetName"`        // sku表单名
	ModelSheetName      string `json:"ModelSheetName"`      //型号对照表单名
	AttrSheetName       string `json:"AttrSheetName"`       // 属性表单名
	SkuCombType         string `json:"SkuCombType"`         // 型号对照表组合类型
	SkuSortType         string `json:"SkuSortType"`         //型号排序类型
	SkuAutoCode         bool   `json:"SkuAutoCode"`         // 自动生成sku编码
	IsSubmit            bool   `json:"IsSubmit"`            // 是否提交
	IsOnline            bool   `json:"IsOnline"`            // 是否上架
}

func (s *ShowInput) SaveInput() {
	curInput := CurInput{
		ShopName:            s.ShopName.Text,
		ShopId:              s.ShopId.Text,
		FreightTmp:          s.FreightTmp.Selected,
		PicKitDir:           s.PicKitDir.Text,
		PubFileDir:          s.PubFileDir.Text,
		UploadedImageConfig: s.UploadedImageConfig.Text,
		ShopExcel:           s.ShopExcel.Text,
		SkuExcel:            s.SkuExcel.Text,
		ModelExcel:          s.ModelExcel.Text,
		ShopSheetName:       s.SkuSheetName.Text,
		SkuSheetName:        s.SkuSheetName.Text,
		ModelSheetName:      s.ModelSheetName.Text,
		AttrSheetName:       s.AttrSheetName.Text,
		SkuCombType:         s.SkuCombType.Selected,
		SkuSortType:         s.SkuSortType.Selected,
		SkuAutoCode:         s.SkuAutoCode.Checked,
		IsSubmit:            s.IsSubmit.Checked,
		IsOnline:            s.IsOnline.Checked,
	}
	// 创建文件
	filePtr, err := os.Create("./acc.json")
	if err != nil {
		log.Println("文件创建失败 [ERROR]: " + err.Error())
		return
	}
	defer filePtr.Close()
	// 创建Json编码器
	encoder := json.NewEncoder(filePtr)
	err = encoder.Encode(curInput)
	if err != nil {
		log.Println("acc.json保存错误", err.Error())
	} else {
		log.Println("acc.json保存成功")
	}
}
func (s *ShowInput) GetInput() {
	filePtr, err := os.Open("./acc.json")
	if err != nil {
		log.Println("文件打开失败 [ERROR]: " + err.Error())
		return
	}
	defer filePtr.Close()
	var curInput CurInput
	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&curInput)
	if err != nil {
		log.Println("读取acc.json失败", err.Error())
	} else {
		log.Println("读取acc.json成功")
		s.ShopId.SetText(curInput.ShopId)
		s.ShopName.SetText(curInput.ShopName)
		s.FreightTmp.SetSelected(curInput.FreightTmp)
		s.PicKitDir.SetText(curInput.PicKitDir)
		s.PubFileDir.SetText(curInput.PubFileDir)
		s.UploadedImageConfig.SetText(curInput.UploadedImageConfig)
		s.ShopExcel.SetText(curInput.ShopExcel)
		s.SkuExcel.SetText(curInput.SkuExcel)
		s.ModelExcel.SetText(curInput.ModelExcel)
		s.ShopSheetName.SetText(curInput.ShopSheetName)
		s.SkuSheetName.SetText(curInput.SkuSheetName)
		s.ModelSheetName.SetText(curInput.ModelSheetName)
		s.AttrSheetName.SetText(curInput.AttrSheetName)
		s.SkuCombType.SetSelected(curInput.SkuCombType)
		s.SkuSortType.SetSelected(curInput.SkuSortType)
		s.SkuAutoCode.SetChecked(curInput.SkuAutoCode)
		s.IsSubmit.SetChecked(curInput.IsSubmit)
		s.IsOnline.SetChecked(curInput.IsOnline)
	}
}
