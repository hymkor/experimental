package main

import (
	"os"
	"strings"

	"github.com/lxn/walk"
	W "github.com/lxn/walk/declarative"
)

func ShowSimpleOneWindow(text string) error {
	var form1 *walk.MainWindow
	var te *walk.TextEdit

	W.MainWindow{
		Title:    "goShowVer",
		AssignTo: &form1,
		MinSize:  W.Size{640, 200},
		Layout:   W.VBox{},
		Children: []W.Widget{
			W.TextEdit{
				AssignTo: &te,
				Text:     text,
				Font:     W.Font{PointSize: 12},
				VScroll: true,
				ReadOnly: false},
			W.PushButton{
				Text: "Ok",
				OnClicked: func() {
					form1.Close()
				},
			},
		},
	}.Run()
	return nil
}

func main() {
	err := ShowSimpleOneWindow(strings.Join(os.Args, " "))
	if err != nil {
		println(err.Error())
	}
}
