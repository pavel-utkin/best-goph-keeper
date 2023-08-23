package service

import (
	"fyne.io/fyne/v2/widget"
	"log"
	"strconv"
	"time"
	"unicode/utf8"
)

func ValidateLogin(usernameLoginEntry *widget.Entry, passwordLoginEntry *widget.Entry, labelAlertAuth *widget.Label) bool {
	if utf8.RuneCountInString(usernameLoginEntry.Text) < 6 {
		labelAlertAuth.SetText("Длинна логина должна быть не менее шести символов")
		log.Print(labelAlertAuth.Text)
		return false
	}
	if utf8.RuneCountInString(passwordLoginEntry.Text) < 6 {
		labelAlertAuth.SetText("Длинна пароля должна быть не менее шести символов")
		log.Print(labelAlertAuth.Text)
		return false
	}
	return true
}

func ValidateRegistration(usernameRegistrationEntry *widget.Entry, passwordRegistrationEntry *widget.Entry,
	passwordConfirmationRegistrationEntry *widget.Entry, labelAlertAuth *widget.Label) bool {
	if utf8.RuneCountInString(usernameRegistrationEntry.Text) < 6 {
		labelAlertAuth.SetText("Длинна логина должна быть не менее шести символов")
		log.Print(labelAlertAuth.Text)
		return false
	}
	if utf8.RuneCountInString(passwordRegistrationEntry.Text) < 6 {
		labelAlertAuth.SetText("Длинна пароля должна быть не менее шести символов")
		log.Print(labelAlertAuth.Text)
		return false
	}
	if passwordRegistrationEntry.Text != passwordConfirmationRegistrationEntry.Text {
		labelAlertAuth.SetText("Пароли не совпали")
		log.Print(labelAlertAuth.Text)
		return false
	}
	return true
}

func ValidateText(exists bool, textNameEntry *widget.Entry, textEntry *widget.Entry, textDescriptionEntry *widget.Entry, labelAlertText *widget.Label) bool {
	if exists {
		labelAlertText.SetText("Текст с таким name уже существует")
		log.Print(labelAlertText)
		return false
	}
	if textNameEntry.Text == "" {
		labelAlertText.SetText("Username не заполнен")
		log.Print(labelAlertText.Text)
		return false
	}
	if textEntry.Text == "" {
		labelAlertText.SetText("Text не заполнен")
		log.Print(labelAlertText.Text)
		return false
	}
	return true
}

func ValidateCard(exists bool, cardNameEntry *widget.Entry, paymentSystemEntry *widget.Entry, numberEntry *widget.Entry,
	holderEntry *widget.Entry, endDateEntry *widget.Entry, cvcEntry *widget.Entry, labelAlertCard *widget.Label) bool {
	var err error

	if exists {
		labelAlertCard.SetText("Карта с таким name уже существует")
		log.Print(labelAlertCard)
		return false
	}
	if cardNameEntry.Text == "" {
		labelAlertCard.SetText("Username не заполнен")
		log.Print(labelAlertCard.Text)
		return false
	}
	if paymentSystemEntry.Text == "" {
		labelAlertCard.SetText("Payment System не заполнен")
		log.Print(labelAlertCard.Text)
		return false
	}
	if numberEntry.Text == "" {
		labelAlertCard.SetText("Number не заполнен")
		log.Print(labelAlertCard.Text)
		return false
	}
	if holderEntry.Text == "" {
		labelAlertCard.SetText("Holder не заполнен")
		log.Print(labelAlertCard.Text)
		return false
	}
	if endDateEntry.Text == "" {
		labelAlertCard.SetText("End date не заполнен")
		log.Print(labelAlertCard.Text)
		return false
	} else {
		layout := "01/02/2006"
		_, err = time.Parse(layout, endDateEntry.Text)
		if err != nil {
			labelAlertCard.SetText("End Date не корректный (пример: 01/02/2006)")
			log.Print(labelAlertCard.Text)
			return false
		}
	}
	if cvcEntry.Text == "" {
		labelAlertCard.SetText("CVC не заполнен")
		log.Print(labelAlertCard.Text)
		return false
	} else {
		_, err = strconv.Atoi(cvcEntry.Text)
		if err != nil {
			labelAlertCard.SetText("CVC не корректный (пример: 123)")
			log.Print(labelAlertCard.Text)
			return false
		}
	}

	return true
}
