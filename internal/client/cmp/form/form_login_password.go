package form

import (
	"best-goph-keeper/internal/client/storage/labels"
	"fyne.io/fyne/v2/widget"
)

func GetFormLoginPassword(loginPasswordName *widget.Entry, loginPasswordDescription *widget.Entry, login *widget.Entry, password *widget.Entry) *widget.Form {
	formText := widget.NewForm(
		widget.NewFormItem(labels.NameItem, loginPasswordName),
		widget.NewFormItem(labels.DescriptionItem, loginPasswordDescription),
		widget.NewFormItem(labels.LoginItem, login),
		widget.NewFormItem(labels.PasswordItem, password),
	)
	return formText
}
