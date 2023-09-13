package form

import (
	"best-goph-keeper/internal/client/storage/labels"
	"fyne.io/fyne/v2/widget"
)

func GetFormText(textName *widget.Entry, textDescription *widget.Entry, text *widget.Entry) *widget.Form {
	formText := widget.NewForm(
		widget.NewFormItem(labels.NameItem, textName),
		widget.NewFormItem(labels.DescriptionItem, textDescription),
		widget.NewFormItem(labels.DataItem, text),
	)
	return formText
}
