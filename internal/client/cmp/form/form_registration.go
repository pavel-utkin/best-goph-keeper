package form

import (
	"best-goph-keeper/internal/client/storage/labels"
	"fyne.io/fyne/v2/widget"
)

func GetFormRegistration(UsernameRegistration *widget.Entry, PasswordRegistration *widget.Entry, NewPasswordEntryRegistration *widget.Entry) *widget.Form {
	formRegistration := widget.NewForm(
		widget.NewFormItem(labels.UsernameItem, UsernameRegistration),
		widget.NewFormItem(labels.PasswordItem, PasswordRegistration),
		widget.NewFormItem(labels.ConfirmPasswordItem, NewPasswordEntryRegistration),
	)
	return formRegistration
}
