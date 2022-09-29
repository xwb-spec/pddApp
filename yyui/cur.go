package yyui

import (
	"encoding/base64"
	"log"
	"path"
	"pddApp/common"
)

type CurInput struct {
	Acc                 string `json:"acc"`
	Pwd                 string `json:"pwd"`
	MallName            string `json:"MallName"`            // 店铺名
	MallId              string `json:"MallId"`              // 店铺id
	LogisticsTemp       string `json:"LogisticsTemp"`       // 运费模板
	ImageDir            string `json:"ImageDir"`            // 套图文件目录
	PublicDir           string `json:"PublicDir"`           // 公共文件目录
	UploadedImage       string `json:"UploadedImage"`       // 已上传图片文件配置
	GoodsExcel          string `json:"GoodsExcel"`          //商品配置表
	SkuConfigExcel      string `json:"SkuConfigExcel"`      // sku配置表
	ModelImageExcel     string `json:"ModelImageExcel"`     //型号对照表
	GoodsSheetName      string `json:"GoodsSheetName"`      // 商品表单名
	SkuConfigSheetName  string `json:"SkuConfigSheetName"`  // sku表单名
	ModelImageSheetName string `json:"ModelImageSheetName"` //型号对照表单名
	AttrSheetName       string `json:"AttrSheetName"`       // 属性表单名
	SkuCombType         string `json:"SkuCombType"`         // 型号对照表组合类型
	SkuSortType         string `json:"SkuSortType"`         //型号排序类型
	SkuAutoCode         bool   `json:"SkuAutoCode"`         // 自动生成sku编码
	IsSubmit            bool   `json:"IsSubmit"`            // 是否提交
	IsOnline            bool   `json:"IsOnline"`            // 是否上架
}

func (s *ShowInput) SaveInput() {
	encodePwd := base64.StdEncoding.EncodeToString([]byte(PwdInput))
	curInput := CurInput{
		Acc:                 AccInput,
		Pwd:                 encodePwd,
		MallName:            s.MallName.Text,
		MallId:              s.MallId.Text,
		LogisticsTemp:       s.LogisticsTemp.Selected,
		ImageDir:            s.ImageDir.Text,
		PublicDir:           s.PublicDir.Text,
		UploadedImage:       s.UploadedImage.Text,
		GoodsExcel:          s.GoodsExcel.Text,
		SkuConfigExcel:      s.SkuConfigExcel.Text,
		ModelImageExcel:     s.ModelImageExcel.Text,
		GoodsSheetName:      s.GoodsSheetName.Text,
		SkuConfigSheetName:  s.SkuConfigSheetName.Text,
		ModelImageSheetName: s.ModelImageSheetName.Text,
		AttrSheetName:       s.AttrSheetName.Text,
		SkuCombType:         s.SkuCombType.Selected,
		SkuSortType:         s.SkuSortType.Selected,
		SkuAutoCode:         s.SkuAutoCode.Checked,
		IsSubmit:            s.IsSubmit.Checked,
		IsOnline:            s.IsOnline.Checked,
	}
	if err := common.CreateJson("./acc.json", curInput); err != nil {
		log.Println("acc.json保存错误", err.Error())
	}
}
func (s *ShowInput) GetInput() {
	var curInput CurInput
	if err := common.LoadJson("./acc.json", &curInput); err != nil {
		log.Println("读取acc.json失败", err.Error())
	} else {
		log.Println("读取acc.json成功")
		//s.Acc.SetText(curInput.Acc)
		//s.Pwd.SetText(curInput.Pwd)
		s.MallId.SetText(curInput.MallId)
		s.MallName.SetText(curInput.MallName)
		s.LogisticsTemp.SetSelected(curInput.LogisticsTemp)
		s.ImageDir.SetText(curInput.ImageDir)
		s.PublicDir.SetText(curInput.PublicDir)
		s.UploadedImage.SetText(curInput.UploadedImage)
		s.GoodsExcel.SetText(curInput.GoodsExcel)
		s.SkuConfigExcel.SetText(curInput.SkuConfigExcel)
		s.ModelImageExcel.SetText(curInput.ModelImageExcel)
		s.GoodsSheetName.SetText(curInput.GoodsSheetName)
		s.SkuConfigSheetName.SetText(curInput.SkuConfigSheetName)
		s.ModelImageSheetName.SetText(curInput.ModelImageSheetName)
		s.AttrSheetName.SetText(curInput.AttrSheetName)
		s.SkuCombType.SetSelected(curInput.SkuCombType)
		s.SkuSortType.SetSelected(curInput.SkuSortType)
		s.SkuAutoCode.SetChecked(curInput.SkuAutoCode)
		s.IsSubmit.SetChecked(curInput.IsSubmit)
		s.IsOnline.SetChecked(curInput.IsOnline)
	}
}

func (s *ShowInput) GetAcc() {
	var curInput CurInput
	if err := common.LoadJson(path.Join(common.GetExec(), "acc.json"), &curInput); err != nil {
		log.Println("读取acc.json失败", err.Error())
	} else {
		log.Println("读取acc.json成功")
		decodePwd, _ := base64.StdEncoding.DecodeString(curInput.Pwd)
		s.Acc.SetText(curInput.Acc)
		s.Pwd.SetText(string(decodePwd))
	}
}

func (s *ShowInput) SaveAcc() {
	var cur CurInput
	if err := common.LoadJson("./acc.json", &cur); err != nil {
		log.Println("读取acc.json失败", err.Error())
	}
	encodePwd := base64.StdEncoding.EncodeToString([]byte(PwdInput))
	curInput := CurInput{
		Acc:                 AccInput,
		Pwd:                 encodePwd,
		MallName:            cur.MallName,
		MallId:              cur.MallId,
		LogisticsTemp:       cur.LogisticsTemp,
		ImageDir:            cur.ImageDir,
		PublicDir:           cur.PublicDir,
		UploadedImage:       cur.UploadedImage,
		GoodsExcel:          cur.GoodsExcel,
		SkuConfigExcel:      cur.SkuConfigExcel,
		ModelImageExcel:     cur.ModelImageExcel,
		GoodsSheetName:      cur.GoodsSheetName,
		SkuConfigSheetName:  cur.SkuConfigSheetName,
		ModelImageSheetName: cur.ModelImageSheetName,
		AttrSheetName:       cur.AttrSheetName,
		SkuCombType:         cur.SkuCombType,
		SkuSortType:         cur.SkuSortType,
		SkuAutoCode:         cur.SkuAutoCode,
		IsSubmit:            cur.IsSubmit,
		IsOnline:            cur.IsOnline,
	}
	if err := common.CreateJson("./acc.json", curInput); err != nil {
		log.Println("acc.json保存错误", err.Error())
	}
}
