package form

import "fyne.io/fyne/v2/widget"

func GetFormText(textName *widget.Entry, textDescription *widget.Entry, text *widget.Entry) *widget.Form {
	formText := widget.NewForm(
		widget.NewFormItem("Username", textName),
		widget.NewFormItem("Description", textDescription),
		widget.NewFormItem("Data", text),
	)
	return formText
}
