package cmp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func SetDefaultColumnsWidthCard(table *widget.Table) {
	colWidths := []float32{150, 150, 150, 150, 50, 150, 150, 150, 150}
	for idx, colWidth := range colWidths {
		table.SetColumnWidth(idx, colWidth)
	}
}

func GetTableCard(dataTblCard [][]string) *widget.Table {
	tableDataCard := widget.NewTable(
		func() (int, int) {
			return len(dataTblCard), len(dataTblCard[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(dataTblCard[i.Row][i.Col])
		})
	SetDefaultColumnsWidthCard(tableDataCard)
	return tableDataCard
}
