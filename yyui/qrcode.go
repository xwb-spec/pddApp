package yyui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"path"
	"pddApp/common"
)

type QRCodeInput struct {
	Win fyne.Window
}

func (q *QRCodeInput) QRCodeShow(w fyne.Window) {
	q.Win = w
	image := canvas.NewImageFromFile(path.Join(common.GetExec(), "qrcode.png"))
	image.FillMode = canvas.ImageFillOriginal
	w.SetContent(image)
}
