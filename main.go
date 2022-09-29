package main

import (
	"pddApp/yyui"
)

func main() {
	myapp := yyui.NewApp()
	yyui.LoginWindow()
	myapp.Run()
	defer myapp.Quit()
}
