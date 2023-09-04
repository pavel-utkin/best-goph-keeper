package form

import (
	"best-goph-keeper/internal/client/storage/labels"
	"fyne.io/fyne/v2/widget"
)

func GetFormLogin(username *widget.Entry, password *widget.Entry) *widget.Form {
	formLogin := widget.NewForm(
		widget.NewFormItem(labels.UsernameItem, username),
		widget.NewFormItem(labels.PasswordItem, password),
	)
	return formLogin
}
