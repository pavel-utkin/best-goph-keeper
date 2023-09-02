package form

import "fyne.io/fyne/v2/widget"

func GetFormLogin(username *widget.Entry, password *widget.Entry) *widget.Form {
	formLogin := widget.NewForm(
		widget.NewFormItem("Username", username),
		widget.NewFormItem("Password", password),
	)
	return formLogin
}
