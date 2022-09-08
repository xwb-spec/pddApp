package yyui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"log"
	"pddApp/pinduoduo/client"
)

type ShowInput struct {
	Win                 fyne.Window
	ShopName            *widget.Entry  // 店铺名
	ShopId              *widget.Entry  // 店铺id
	FreightTmp          *widget.Select // 运费模板
	PicKitDir           *widget.Entry  // 套图文件目录
	PubFileDir          *widget.Entry  // 公共文件目录
	UploadedImageConfig *widget.Entry  // 已上传图片文件配置
	ShopExcel           *widget.Entry  //商品配置表
	SkuExcel            *widget.Entry  // sku配置表
	ModelExcel          *widget.Entry  //型号对照表
	ShopSheetName       *widget.Entry  // 商品表单名
	SkuSheetName        *widget.Entry  // sku表单名
	ModelSheetName      *widget.Entry  //型号对照表单名
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
	s.ShopName = widget.NewEntry()
	loginShopIdLabel := widget.NewLabel("店铺id")
	s.ShopId = widget.NewEntry()
	loginButton := widget.NewButton("登录", func() { // //回调函数
		makeQRCode := client.QRCodeRequestParam{
			Path: "/Users/machao/OptRepo/GolandProjects/awesomeProject/learning-golang/pddApp/qrcode.png",
		}
		makeQRCode.MakeQRCode()
		image := canvas.NewImageFromFile(makeQRCode.Path)
		image.FillMode = canvas.ImageFillOriginal
		win := fyne.CurrentApp().NewWindow("扫码登录")
		win.SetContent(image)
		win.Resize(fyne.NewSize(300, 300))
		win.Show()
	})
	return container.New(layout.NewGridLayout(5), loginShopNameLabel, s.ShopName, loginShopIdLabel, s.ShopId, loginButton)
}

func (s *ShowInput) FreightTmpContainer() *fyne.Container {
	// 运费模板
	freightTmpNameLabel := widget.NewLabel("运费模板")
	options := []string{"1", "3"}
	s.FreightTmp = widget.NewSelect(options, func(s string) {
		fmt.Println("选择")
	})
	freightRefreshButton := widget.NewButton("刷新", func() {
		fmt.Println("开始刷新")
	})
	return container.New(layout.NewGridLayout(3), freightTmpNameLabel, s.FreightTmp, freightRefreshButton)
}

func (s *ShowInput) PicKitDirContainer() *fyne.Container {
	// 套图文件目录
	picKitDirLabel := widget.NewLabel("套图文件目录")
	s.PicKitDir = widget.NewEntry()
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
			s.PicKitDir.SetText(list.Path())
		}, s.Win)
	})
	return container.New(layout.NewGridLayout(3), picKitDirLabel, s.PicKitDir, picKitDirButton)
}

func (s *ShowInput) PublicDirContainer() *fyne.Container {
	// 公共文件目录
	publicDirLabel := widget.NewLabel("公共文件目录")
	s.PubFileDir = widget.NewEntry()
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
			s.PubFileDir.SetText(list.Path())
		}, s.Win)
	})
	return container.New(layout.NewGridLayout(3), publicDirLabel, s.PubFileDir, publicDirButton)
}

func (s *ShowInput) UploadedPicFileContainer() *fyne.Container {
	// 已上传文件配置
	uploadedPicFileLabel := widget.NewLabel("已上传图片文件配置")
	s.UploadedImageConfig = widget.NewEntry()
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
			s.UploadedImageConfig.SetText(reader.URI().Path()) //把读取到的路径显示到输入框中
		}, s.Win)
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".json"})) //打开的文件格式类型
		fd.Show()
	})
	return container.New(layout.NewGridLayout(3), uploadedPicFileLabel, s.UploadedImageConfig, uploadedPicFileButton)
}

func (s *ShowInput) ShopExcelContainer() *fyne.Container {
	// 商品配置表
	shopExcelLabel := widget.NewLabel("商品配置表")
	s.ShopExcel = widget.NewEntry()
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
			s.ShopExcel.SetText(reader.URI().Path()) //把读取到的路径显示到输入框中
		}, s.Win)
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".xlsx"})) //打开的文件格式类型
		fd.Show()
	})
	return container.New(layout.NewGridLayout(3), shopExcelLabel, s.ShopExcel, shopExcelButton)
}
func (s *ShowInput) SkuExcelContainer() *fyne.Container {
	// sku配置表
	skuExcelLabel := widget.NewLabel("sku配置表")
	s.SkuExcel = widget.NewEntry()
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
			s.SkuExcel.SetText(reader.URI().Path()) //把读取到的路径显示到输入框中
		}, s.Win)
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".xlsx"})) //打开的文件格式类型
		fd.Show()
	})
	return container.New(layout.NewGridLayout(3), skuExcelLabel, s.SkuExcel, skuExcelButton)
}
func (s *ShowInput) ModelExcelContainer() *fyne.Container {
	// 型号对照表
	modelExcelLabel := widget.NewLabel("型号对照表")
	s.ModelExcel = widget.NewEntry()
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
			s.ModelExcel.SetText(reader.URI().Path()) //把读取到的路径显示到输入框中
		}, s.Win)
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".xlsx"})) //打开的文件格式类型
		fd.Show()
	})
	return container.New(layout.NewGridLayout(3), modelExcelLabel, s.ModelExcel, modelExcelButton)
}
func (s *ShowInput) Sheet1Container() *fyne.Container {
	// 商品表表单
	shopSheetLabel := widget.NewLabel("商品表表单名")
	s.ShopSheetName = widget.NewEntry()
	s.ShopSheetName.SetText("苹果")
	// sku配置表表单
	skuSheetLabel := widget.NewLabel("sku配置表表单名")
	s.SkuSheetName = widget.NewEntry()
	s.SkuSheetName.SetText("配置")
	return container.New(layout.NewGridLayout(4), shopSheetLabel, s.ShopSheetName, skuSheetLabel, s.SkuSheetName)
}
func (s *ShowInput) Sheet2Container() *fyne.Container {
	// 型号对照表表单
	modelSheetLabel := widget.NewLabel("型号对照表表单名")
	s.ModelSheetName = widget.NewEntry()
	s.ModelSheetName.SetText("型号")
	// 属性表表单
	attrSheetLabel := widget.NewLabel("属性表表单名")
	s.AttrSheetName = widget.NewEntry()
	s.AttrSheetName.SetText("属性")
	return container.New(layout.NewGridLayout(4), modelSheetLabel, s.ModelSheetName, attrSheetLabel, s.AttrSheetName)
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

	})
	startUpload := widget.NewButton("开始上传", func() {
		fmt.Println(s.ShopExcel)
	})
	return container.New(layout.NewGridLayout(3), checkPic, checkConfig, startUpload)
}
func (s *ShowInput) ResultContainer() *fyne.Container {
	s.ConsoleResult = widget.NewMultiLineEntry()
	s.ConsoleResult.Resize(fyne.NewSize(790, 250))
	s.ConsoleResult.SetText(s.ConsoleResult.Text)
	return container.NewWithoutLayout(s.ConsoleResult)
}
func (s *ShowInput) MainShow(w fyne.Window) {
	s.Win = w
	box := container.NewVBox(
		s.LoginContainer(),
		s.FreightTmpContainer(),
		s.PublicDirContainer(),
		s.PicKitDirContainer(),
		s.UploadedPicFileContainer(),
		s.ShopExcelContainer(),
		s.SkuExcelContainer(),
		s.ModelExcelContainer(),
		s.Sheet1Container(),
		s.Sheet2Container(),
		s.SelectContainer(),
		s.CheckContainer(),
		s.ButtonContainer(), s.ResultContainer(),
	) //控制显示位置顺序
	w.SetContent(box)
}
