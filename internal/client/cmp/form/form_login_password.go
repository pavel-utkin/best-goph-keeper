package form

import "fyne.io/fyne/v2/widget"

func GetFormLoginPassword(loginPasswordName *widget.Entry, loginPasswordDescription *widget.Entry, login *widget.Entry, password *widget.Entry) *widget.Form {
	formText := widget.NewForm(
		widget.NewFormItem("Name", loginPasswordName),
		widget.NewFormItem("Description", loginPasswordDescription),
		widget.NewFormItem("Login", login),
		widget.NewFormItem("Password", password),
	)
	return formText
}
