package yyui

import (
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"log"
	"os"
	"pddApp/pinduoduo/client"
	"pddApp/pinduoduo/sdk"
	"strconv"
	"time"
)

type ShowInput struct {
	Win                 fyne.Window
	MallName            *widget.Entry  // 店铺名
	MallId              *widget.Entry  // 店铺id
	LogisticsTemp       *widget.Select // 运费模板
	ImageDir            *widget.Entry  // 套图文件目录
	PublicDir           *widget.Entry  // 公共文件目录
	UploadedImage       *widget.Entry  // 已上传图片文件配置
	GoodsExcel          *widget.Entry  //商品配置表
	SkuConfigExcel      *widget.Entry  // sku配置表
	ModelImageExcel     *widget.Entry  //型号对照表
	GoodsSheetName      *widget.Entry  // 商品表单名
	SkuConfigSheetName  *widget.Entry  // sku表单名
	ModelImageSheetName *widget.Entry  //型号对照表单名
	AttrSheetName       *widget.Entry  // 属性表单名
	SkuCombType         *widget.Select // sku组合类型
	SkuSortType         *widget.Select //sku排序类型
	SkuAutoCode         *widget.Check  // 自动生成sku编码
	IsSubmit            *widget.Check  // 是否提交
	IsOnline            *widget.Check  // 是否上架
	ConsoleResult       *widget.Entry
}

func (s *ShowInput) LoginContainer() *fyne.Container {
	// 登录框
	loginShopNameLabel := widget.NewLabel("店铺名")
	s.MallName = widget.NewEntry()
	loginShopIdLabel := widget.NewLabel("店铺id")
	s.MallId = widget.NewEntry()
	loginButton := widget.NewButton("二维码", func() { // //回调函数
		state := strconv.FormatInt(time.Now().Unix(), 10)
		client.GenerateQRCode(state)
		image := canvas.NewImageFromFile("./qrcode.png")
		image.FillMode = canvas.ImageFillOriginal
		win := fyne.CurrentApp().NewWindow("扫码登录")
		win.SetContent(image)
		win.Resize(fyne.NewSize(300, 300))
		win.Show()
		st, _ := sdk.PopAuthCreateToken()
		for {
			if st == state {
				win.Close()
				break
			} else {
				st, _ = sdk.PopAuthCreateToken()
				time.Sleep(10)
			}
		}
	})
	return container.New(layout.NewGridLayout(5), loginShopNameLabel, s.MallName, loginShopIdLabel, s.MallId, loginButton)
}

func (s *ShowInput) LogisticsTempContainer() *fyne.Container {
	// 运费模板
	freightTmpNameLabel := widget.NewLabel("运费模板")
	options := []string{"模板1", "模板2"}
	s.LogisticsTemp = widget.NewSelect(options, func(s string) {
		log.Println(s)
	})
	freightRefreshButton := widget.NewButton("刷新", func() {
		s.LogisticsTemp.Refresh()
	})
	return container.NewBorder(layout.NewSpacer(), layout.NewSpacer(), freightTmpNameLabel, freightRefreshButton, s.LogisticsTemp)
}

func (s *ShowInput) ImageDirContainer() *fyne.Container {
	// 套图文件目录
	picKitDirLabel := widget.NewLabel("套图文件目录")
	s.ImageDir = widget.NewEntry()
	picKitDirButton := widget.NewButton("...", func() {
		dialog.ShowFolderOpen(func(list fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, s.Win)
				return
			}
			if list == nil {
				log.Println("Cancelled")
				return
			}
			//设置输入框内容
			s.ImageDir.SetText(list.Path())
		}, s.Win)
	})
	return container.NewBorder(layout.NewSpacer(), layout.NewSpacer(), picKitDirLabel, picKitDirButton, s.ImageDir)
}

func (s *ShowInput) PublicDirContainer() *fyne.Container {
	// 公共文件目录
	publicDirLabel := widget.NewLabel("公共文件目录")
	s.PublicDir = widget.NewEntry()
	publicDirButton := widget.NewButton("...", func() {
		dialog.ShowFolderOpen(func(list fyne.ListableURI, err error) {
			if err != nil {
				dialog.ShowError(err, s.Win)
				return
			}
			if list == nil {
				log.Println("Cancelled")
				return
			}
			//设置输入框内容
			s.PublicDir.SetText(list.Path())
		}, s.Win)
	})
	return container.NewBorder(layout.NewSpacer(), layout.NewSpacer(), publicDirLabel, publicDirButton, s.PublicDir)
}

func (s *ShowInput) UploadedImageContainer() *fyne.Container {
	// 已上传文件配置
	uploadedPicFileLabel := widget.NewLabel("已上传图片文件配置")
	s.UploadedImage = widget.NewEntry()
	uploadedPicFileButton := widget.NewButton("...", func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, s.Win)
				return
			}
			if reader == nil {
				log.Println("Cancelled")
				return
			}
			s.UploadedImage.SetText(reader.URI().Path()) //把读取到的路径显示到输入框中
		}, s.Win)
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".json"})) //打开的文件格式类型
		fd.Show()
	})
	return container.NewBorder(layout.NewSpacer(), layout.NewSpacer(), uploadedPicFileLabel, uploadedPicFileButton, s.UploadedImage)
}

func (s *ShowInput) GoodsExcelContainer() *fyne.Container {
	// 商品配置表
	shopExcelLabel := widget.NewLabel("商品配置表")
	s.GoodsExcel = widget.NewEntry()
	shopExcelButton := widget.NewButton("...", func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, s.Win)
				return
			}
			if reader == nil {
				log.Println("Cancelled")
				return
			}
			s.GoodsExcel.SetText(reader.URI().Path()) //把读取到的路径显示到输入框中
		}, s.Win)
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".xlsx"})) //打开的文件格式类型
		fd.Show()
	})
	return container.NewBorder(layout.NewSpacer(), layout.NewSpacer(), shopExcelLabel, shopExcelButton, s.GoodsExcel)
}
func (s *ShowInput) SkuExcelContainer() *fyne.Container {
	// sku配置表
	skuExcelLabel := widget.NewLabel("sku配置表")
	s.SkuConfigExcel = widget.NewEntry()
	skuExcelButton := widget.NewButton("...", func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, s.Win)
				return
			}
			if reader == nil {
				log.Println("Cancelled")
				return
			}
			s.SkuConfigExcel.SetText(reader.URI().Path()) //把读取到的路径显示到输入框中
		}, s.Win)
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".xlsx"})) //打开的文件格式类型
		fd.Show()
	})
	return container.NewBorder(layout.NewSpacer(), layout.NewSpacer(), skuExcelLabel, skuExcelButton, s.SkuConfigExcel)
}
func (s *ShowInput) ModelImageExcelContainer() *fyne.Container {
	// 型号对照表
	modelExcelLabel := widget.NewLabel("型号对照表")
	s.ModelImageExcel = widget.NewEntry()
	modelExcelButton := widget.NewButton("...", func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, s.Win)
				return
			}
			if reader == nil {
				log.Println("Cancelled")
				return
			}
			s.ModelImageExcel.SetText(reader.URI().Path()) //把读取到的路径显示到输入框中
		}, s.Win)
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".xlsx"})) //打开的文件格式类型
		fd.Show()
	})
	return container.NewBorder(layout.NewSpacer(), layout.NewSpacer(), modelExcelLabel, modelExcelButton, s.ModelImageExcel)
}
func (s *ShowInput) Sheet1Container() *fyne.Container {
	// 商品表表单
	shopSheetLabel := widget.NewLabel("商品表表单名")
	s.GoodsSheetName = widget.NewEntry()
	s.GoodsSheetName.SetText("苹果")
	// sku配置表表单
	skuSheetLabel := widget.NewLabel("sku配置表表单名")
	s.SkuConfigSheetName = widget.NewEntry()
	s.SkuConfigSheetName.SetText("配置")
	return container.New(layout.NewGridLayout(4), shopSheetLabel, s.GoodsSheetName, skuSheetLabel, s.SkuConfigSheetName)
}
func (s *ShowInput) Sheet2Container() *fyne.Container {
	// 型号对照表表单
	modelSheetLabel := widget.NewLabel("型号对照表表单名")
	s.ModelImageSheetName = widget.NewEntry()
	s.ModelImageSheetName.SetText("直边tpu")
	// 属性表表单
	attrSheetLabel := widget.NewLabel("属性表表单名")
	s.AttrSheetName = widget.NewEntry()
	s.AttrSheetName.SetText("属性")
	return container.New(layout.NewGridLayout(4), modelSheetLabel, s.ModelImageSheetName, attrSheetLabel, s.AttrSheetName)
}
func (s *ShowInput) SelectContainer() *fyne.Container {
	skuCombTypeLabel := widget.NewLabel("型号对照表组合类型")
	skuCombTypeOption := []string{"款式+型号"}
	s.SkuCombType = widget.NewSelect(skuCombTypeOption, func(s string) {
	})
	skuSortTypeLabel := widget.NewLabel("型号排序类型")
	skuSortTypeOption := []string{"型号在前"}
	s.SkuSortType = widget.NewSelect(skuSortTypeOption, func(s string) {
	})
	return container.New(layout.NewGridLayout(4), skuCombTypeLabel, s.SkuCombType, skuSortTypeLabel, s.SkuSortType)
}
func (s *ShowInput) CheckContainer() *fyne.Container {
	s.SkuAutoCode = widget.NewCheck("自动生成sku编码", func(b bool) {

	})
	s.IsSubmit = widget.NewCheck("是否提交", func(b bool) {

	})
	s.IsOnline = widget.NewCheck("是否上架", func(b bool) {

	})
	return container.New(layout.NewGridLayout(3), s.SkuAutoCode, s.IsSubmit, s.IsOnline)
}
func (s *ShowInput) ButtonContainer() *fyne.Container {
	checkPic := widget.NewButton("检测图片", func() {
		s.SaveInput()
		s.CheckInput()

	})
	checkConfig := widget.NewButton("检测配置", func() {
		s.CheckConfig()
	})
	startUpload := widget.NewButton("开始上传", func() {
		a, _ := s.GetGoods()
		filePtr, err := os.Create("./test.json")
		if err != nil {
			log.Println("文件创建失败 [ERROR]: " + err.Error())
			return
		}
		defer filePtr.Close()
		// 创建Json编码器
		encoder := json.NewEncoder(filePtr)
		err = encoder.Encode(a)
		if err != nil {
			log.Println("test.json保存错误", err.Error())
		} else {
			log.Println("test.json保存成功")
		}
	})
	return container.New(layout.NewGridLayout(3), checkPic, checkConfig, startUpload)
}
func (s *ShowInput) ResultContainer() *fyne.Container {
	s.ConsoleResult = widget.NewMultiLineEntry()
	s.ConsoleResult.Resize(fyne.NewSize(790, 230))
	s.ConsoleResult.SetText(s.ConsoleResult.Text)
	return container.NewWithoutLayout(s.ConsoleResult)
}
func (s *ShowInput) MainShow(w fyne.Window) {
	s.Win = w
	box := container.NewVBox(
		s.LoginContainer(),
		s.LogisticsTempContainer(),
		s.PublicDirContainer(),
		s.ImageDirContainer(),
		s.UploadedImageContainer(),
		s.GoodsExcelContainer(),
		s.SkuExcelContainer(),
		s.ModelImageExcelContainer(),
		s.Sheet1Container(),
		s.Sheet2Container(),
		s.SelectContainer(),
		s.CheckContainer(),
		s.ButtonContainer(), s.ResultContainer(),
	) //控制显示位置顺序
	w.SetContent(box)
}
