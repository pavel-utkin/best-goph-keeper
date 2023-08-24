package cmp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func GetTabTexts(tblText *widget.Table, top *widget.Button, text *widget.Button) *container.TabItem {
	containerTblText := layout.NewBorderLayout(top, text, nil, nil)
	boxText := fyne.NewContainerWithLayout(containerTblText, top, tblText, text)
	return container.NewTabItem("Текстовые данные", boxText)
}
