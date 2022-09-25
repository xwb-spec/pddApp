package yyui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"log"
)

type QRCodeInput struct {
	Win fyne.Window
}

func (q *QRCodeInput) QRCodeShow(w fyne.Window) {
	log.Println("aa")
	q.Win = w
	image := canvas.NewImageFromFile("/Users/xiewenbin/GolandProjects/pddApp/qrcode.png")
	image.FillMode = canvas.ImageFillOriginal
	w.SetContent(image)
}
