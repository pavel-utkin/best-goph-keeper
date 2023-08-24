package cmp

import "fyne.io/fyne/v2/widget"

func GetFormCard(name *widget.Entry, paymentSystem *widget.Entry, number *widget.Entry, holder *widget.Entry, endDate *widget.Entry, cvc *widget.Entry) *widget.Form {
	formCart := widget.NewForm(
		widget.NewFormItem("Name", name),
		widget.NewFormItem("Payment System", paymentSystem),
		widget.NewFormItem("Number", number),
		widget.NewFormItem("Holder", holder),
		widget.NewFormItem("End Date", endDate),
		widget.NewFormItem("CVC", cvc),
	)
	return formCart
}
