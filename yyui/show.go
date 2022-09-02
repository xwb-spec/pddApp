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
	ShopName            string // 店铺名
	ShopId              string // 店铺id
	FreightTmp          string // 运费模板
	PicKitDir           string // 套图文件目录
	PubFileDir          string // 公共文件目录
	UploadedImageConfig string // 已上传图片文件配置
	ShopExcel           string //商品配置表
	SkuExcel            string // sku配置表
	ModelExcel          string //型号对照表
	ShopSheetName       string // 商品表单名
	SkuSheetName        string // sku表单名
	ModelSheetName      string //型号对照表单名
	AttrSheetName       string // 属性表单名
	ModelExcelCombType  string // 型号对照表组合类型
	ModelSortType       string //型号排序类型
	SkuAutoCode         bool   // 自动生成sku编码
	IsSubmit            bool   // 是否提交
	IsOnline            bool   // 是否上架
}

func (s *ShowInput) LoginContainer() *fyne.Container {
	// 登录框
	loginShopNameLabel := widget.NewLabel("店铺名")
	loginShopNameEntry := widget.NewEntry()
	loginShopIdLabel := widget.NewLabel("店铺id")
	loginShopIdEntry := widget.NewEntry()
	loginButton := widget.NewButton("登录", func() { // //回调函数
		makeQRCode := client.QRCodeArgs{
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
	return container.New(layout.NewGridLayout(5), loginShopNameLabel, loginShopNameEntry, loginShopIdLabel, loginShopIdEntry, loginButton)
}

func (s *ShowInput) FreightTmpContainer() *fyne.Container {
	// 运费模板
	freightTmpNameLabel := widget.NewLabel("运费模板")
	options := []string{"1", "3"}
	freightTmpSelect := widget.NewSelect(options, func(s string) {
		fmt.Println("选择")
	})
	freightRefreshButton := widget.NewButton("刷新", func() {
		fmt.Println("开始刷新")
	})
	return container.New(layout.NewGridLayout(3), freightTmpNameLabel, freightTmpSelect, freightRefreshButton)
}

func (s *ShowInput) PicKitDirContainer() *fyne.Container {
	// 套图文件目录
	picKitDirLabel := widget.NewLabel("套图文件目录")
	picKitDirEntry := widget.NewEntry()
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
			picKitDirEntry.SetText(list.Path())
			s.PicKitDir = picKitDirEntry.Text
		}, s.Win)
	})
	return container.New(layout.NewGridLayout(3), picKitDirLabel, picKitDirEntry, picKitDirButton)
}

func (s *ShowInput) PublicDirContainer() *fyne.Container {
	// 公共文件目录
	publicDirLabel := widget.NewLabel("公共文件目录")
	publicDirEntry := widget.NewEntry()
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
			publicDirEntry.SetText(list.Path())
		}, s.Win)
	})
	return container.New(layout.NewGridLayout(3), publicDirLabel, publicDirEntry, publicDirButton)
}

func (s *ShowInput) UploadedPicFileContainer() *fyne.Container {
	// 已上传文件配置
	uploadedPicFileLabel := widget.NewLabel("已上传图片文件配置")
	uploadedPicFileEntry := widget.NewEntry()
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
			uploadedPicFileEntry.SetText(reader.URI().Path()) //把读取到的路径显示到输入框中
			s.UploadedImageConfig = uploadedPicFileEntry.Text
		}, s.Win)
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".txt"})) //打开的文件格式类型
		fd.Show()
	})
	return container.New(layout.NewGridLayout(3), uploadedPicFileLabel, uploadedPicFileEntry, uploadedPicFileButton)
}

func (s *ShowInput) ShopExcelContainer() *fyne.Container {
	// 商品配置表
	shopExcelLabel := widget.NewLabel("商品配置表")
	shopExcelEntry := widget.NewEntry()
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
			shopExcelEntry.SetText(reader.URI().Path()) //把读取到的路径显示到输入框中
			s.ShopExcel = shopExcelEntry.Text
		}, s.Win)
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".txt"})) //打开的文件格式类型
		fd.Show()
	})
	return container.New(layout.NewGridLayout(3), shopExcelLabel, shopExcelEntry, shopExcelButton)
}
func (s *ShowInput) SkuExcelContainer() *fyne.Container {
	// sku配置表
	skuExcelLabel := widget.NewLabel("sku配置表")
	skuExcelEntry := widget.NewEntry()
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

			skuExcelEntry.SetText(reader.URI().Path()) //把读取到的路径显示到输入框中
		}, s.Win)
		fd.SetFilter(storage.NewExtensionFileFilter([]string{"."})) //打开的文件格式类型
		fd.Show()
	})
	return container.New(layout.NewGridLayout(3), skuExcelLabel, skuExcelEntry, skuExcelButton)
}
func (s *ShowInput) ModelExcelContainer() *fyne.Container {
	// 型号对照表
	modelExcelLabel := widget.NewLabel("型号对照表")
	modelExcelEntry := widget.NewEntry()
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
			modelExcelEntry.SetText(reader.URI().Path()) //把读取到的路径显示到输入框中
		}, s.Win)
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".txt"})) //打开的文件格式类型
		fd.Show()
	})
	return container.New(layout.NewGridLayout(3), modelExcelLabel, modelExcelEntry, modelExcelButton)
}
func (s *ShowInput) Sheet1Container() *fyne.Container {
	// 商品表表单
	shopSheetLabel := widget.NewLabel("商品表表单名")
	shopSheetEntry := widget.NewEntry()
	// sku配置表表单
	skuSheetLabel := widget.NewLabel("sku配置表表单名")
	skuSheetEntry := widget.NewEntry()
	return container.New(layout.NewGridLayout(4), shopSheetLabel, shopSheetEntry, skuSheetLabel, skuSheetEntry)
}
func (s *ShowInput) Sheet2Container() *fyne.Container {
	// 型号对照表表单
	modelSheetLabel := widget.NewLabel("型号对照表表单名")
	modelSheetEntry := widget.NewEntry()
	// 属性表表单
	attrSheetLabel := widget.NewLabel("属性表表单名")
	attrSheetEntry := widget.NewEntry()
	return container.New(layout.NewGridLayout(4), modelSheetLabel, modelSheetEntry, attrSheetLabel, attrSheetEntry)
}
func (s *ShowInput) SelectContainer() *fyne.Container {
	skuCombTypeLabel := widget.NewLabel("型号对照表组合类型")
	skuCombTypeOption := []string{"款式+型号"}
	skuCombTypeSelect := widget.NewSelect(skuCombTypeOption, func(s string) {
	})
	skuSortTypeLabel := widget.NewLabel("型号排序类型")
	skuSortTypeOption := []string{"型号在前"}
	skuSortTypeSelect := widget.NewSelect(skuSortTypeOption, func(s string) {
	})
	return container.New(layout.NewGridLayout(4), skuCombTypeLabel, skuCombTypeSelect, skuSortTypeLabel, skuSortTypeSelect)
}
func (s *ShowInput) CheckContainer() *fyne.Container {
	skuAutoCode := widget.NewCheck("自动生成sku编码", func(b bool) {

	})
	isSubmit := widget.NewCheck("是否提交", func(b bool) {

	})
	isOnline := widget.NewCheck("是否上架", func(b bool) {

	})
	return container.New(layout.NewGridLayout(3), skuAutoCode, isSubmit, isOnline)
}
func (s *ShowInput) ButtonContainer() *fyne.Container {
	checkPic := widget.NewButton("检测图片", func() {

	})
	checkConfig := widget.NewButton("检测配置", func() {

	})
	startUpload := widget.NewButton("开始上传", func() {
		fmt.Println(s.ShopExcel)
	})
	return container.New(layout.NewGridLayout(3), checkPic, checkConfig, startUpload)
}
func (s *ShowInput) ResultContainer() *fyne.Container {
	consoleResult := widget.NewMultiLineEntry()
	return container.New(layout.NewGridLayout(1), consoleResult)
}
func (s *ShowInput) MainShow(w fyne.Window) {
	s.Win = w
	authInfo := widget.NewLabel("Xwb    ALL Right Reserved")
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
		s.ButtonContainer(), s.ResultContainer(), authInfo) //控制显示位置顺序
	w.SetContent(box)
}
