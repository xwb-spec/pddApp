package yyui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type QRCodeInput struct {
	Win fyne.Window
}

func (q *QRCodeInput) QRCodeShow(w fyne.Window) {
	q.Win = w
	image := canvas.NewImageFromFile("./qrcode.png")
	image.FillMode = canvas.ImageFillOriginal
	w.SetContent(image)
}
