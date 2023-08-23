package cmp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func SetDefaultColumnsWidthText(table *widget.Table) {
	colWidths := []float32{150, 500, 150, 150, 150}
	for idx, colWidth := range colWidths {
		table.SetColumnWidth(idx, colWidth)
	}
}

func GetTableText(dataTblCard [][]string) *widget.Table {
	tableDataText := widget.NewTable(
		func() (int, int) {
			return len(dataTblCard), len(dataTblCard[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(dataTblCard[i.Row][i.Col])
		})
	SetDefaultColumnsWidthText(tableDataText)
	return tableDataText
}
