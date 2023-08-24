package cmp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func GetTabCards(tblCart *widget.Table, top *widget.Button, card *widget.Button) *container.TabItem {
	containerTblCart := layout.NewBorderLayout(top, card, nil, nil)
	boxCart := fyne.NewContainerWithLayout(containerTblCart, top, tblCart, card)
	return container.NewTabItem("Банковские карты", boxCart)
}
