package form

import (
	"best-goph-keeper/internal/client/service/encryption"
	"best-goph-keeper/internal/client/storage/errors"
	"fyne.io/fyne/v2/widget"
	"log"
	"strconv"
	"time"
	"unicode/utf8"
)

func ValidateLogin(usernameLoginEntry *widget.Entry, passwordLoginEntry *widget.Entry, labelAlertAuth *widget.Label) bool {
	if utf8.RuneCountInString(usernameLoginEntry.Text) < 6 {
		labelAlertAuth.SetText(errors.ErrUsernameIncorrect)
		log.Print(labelAlertAuth.Text)
		return false
	}
	if utf8.RuneCountInString(passwordLoginEntry.Text) < 6 {
		labelAlertAuth.SetText(errors.ErrPasswordIncorrect)
		log.Print(labelAlertAuth.Text)
		return false
	}
	return true
}

func ValidateRegistration(usernameRegistrationEntry *widget.Entry, passwordRegistrationEntry *widget.Entry,
	passwordConfirmationRegistrationEntry *widget.Entry, labelAlertAuth *widget.Label) bool {
	if utf8.RuneCountInString(usernameRegistrationEntry.Text) < 6 {
		labelAlertAuth.SetText(errors.ErrUsernameIncorrect)
		log.Print(labelAlertAuth.Text)
		return false
	}
	if !encryption.VerifyPassword(passwordRegistrationEntry.Text) {
		labelAlertAuth.SetText(errors.ErrPasswordIncorrect)
		log.Print(labelAlertAuth.Text)
		return false
	}
	if passwordRegistrationEntry.Text != passwordConfirmationRegistrationEntry.Text {
		labelAlertAuth.SetText(errors.ErrPasswordDifferent)
		log.Print(labelAlertAuth.Text)
		return false
	}
	return true
}

func ValidateText(exists bool, textNameEntry *widget.Entry, textEntry *widget.Entry, textDescriptionEntry *widget.Entry, labelAlertText *widget.Label) bool {
	if exists {
		labelAlertText.SetText(errors.ErrTextExist)
		log.Print(labelAlertText)
		return false
	}
	if textNameEntry.Text == "" {
		labelAlertText.SetText(errors.ErrNameEmpty)
		log.Print(labelAlertText.Text)
		return false
	}
	if textEntry.Text == "" {
		labelAlertText.SetText(errors.ErrTextEmpty)
		log.Print(labelAlertText.Text)
		return false
	}
	if textDescriptionEntry.Text == "" {
		labelAlertText.SetText(errors.ErrDescriptionEmpty)
		log.Print(labelAlertText.Text)
		return false
	}
	return true
}

func ValidateCard(exists bool, cartNameEntry *widget.Entry, paymentSystemEntry *widget.Entry, numberEntry *widget.Entry,
	holderEntry *widget.Entry, endDateEntry *widget.Entry, cvcEntry *widget.Entry, labelAlertCart *widget.Label) bool {
	var err error
	if exists {
		labelAlertCart.SetText(errors.ErrCartExist)
		log.Print(labelAlertCart)
		return false
	}
	if cartNameEntry.Text == "" {
		labelAlertCart.SetText(errors.ErrNameEmpty)
		log.Print(labelAlertCart.Text)
		return false
	}
	if paymentSystemEntry.Text == "" {
		labelAlertCart.SetText(errors.ErrPaymentSystemEmpty)
		log.Print(labelAlertCart.Text)
		return false
	}
	if numberEntry.Text == "" {
		labelAlertCart.SetText(errors.ErrNumberEmpty)
		log.Print(labelAlertCart.Text)
		return false
	}
	if holderEntry.Text == "" {
		labelAlertCart.SetText(errors.ErrHolderEmpty)
		log.Print(labelAlertCart.Text)
		return false
	}
	if endDateEntry.Text == "" {
		labelAlertCart.SetText(errors.ErrEndDateEmpty)
		log.Print(labelAlertCart.Text)
		return false
	} else {
		layout := "01/02/2006"
		_, err = time.Parse(layout, endDateEntry.Text)
		if err != nil {
			labelAlertCart.SetText(errors.ErrEndDataIncorrect)
			log.Print(labelAlertCart.Text)
			return false
		}
	}
	if cvcEntry.Text == "" {
		labelAlertCart.SetText(errors.ErrCvcEmpty)
		log.Print(labelAlertCart.Text)
		return false
	} else {
		_, err = strconv.Atoi(cvcEntry.Text)
		if err != nil {
			labelAlertCart.SetText(errors.ErrCvcIncorrect)
			log.Print(labelAlertCart.Text)
			return false
		}
	}

	return true
}
