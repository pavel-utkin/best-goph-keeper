package form

import "fyne.io/fyne/v2/widget"

func GetFormRegistration(UsernameRegistration *widget.Entry, PasswordRegistration *widget.Entry, NewPasswordEntryRegistration *widget.Entry) *widget.Form {
	formRegistration := widget.NewForm(
		widget.NewFormItem("Username", UsernameRegistration),
		widget.NewFormItem("Password", PasswordRegistration),
		widget.NewFormItem("Confirm password", NewPasswordEntryRegistration),
	)
	return formRegistration
}
