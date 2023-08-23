package cmp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func GetTabCards(tblCard *widget.Table, top *widget.Button, card *widget.Button) *container.TabItem {
	containerTblCard := layout.NewBorderLayout(top, card, nil, nil)
	boxCard := fyne.NewContainerWithLayout(containerTblCard, top, tblCard, card)
	return container.NewTabItem("Банковские карты", boxCard)
}
