package function

import (
	"best-goph-keeper/internal/client/service/algorithm"
	"best-goph-keeper/internal/client/service/encryption"
	"best-goph-keeper/internal/client/storage/errors"
	"fyne.io/fyne/v2/widget"
	"log"
	"strconv"
	"time"
	"unicode/utf8"
)

const userNameMaxLength = 6
const passwordMaxLength = 6

func ValidateLogin(usernameLoginEntry *widget.Entry, passwordLoginEntry *widget.Entry, labelAlertAuth *widget.Label) bool {
	if utf8.RuneCountInString(usernameLoginEntry.Text) < userNameMaxLength {
		labelAlertAuth.SetText(errors.ErrUsernameIncorrect)
		log.Print(labelAlertAuth.Text)
		return false
	}
	if utf8.RuneCountInString(passwordLoginEntry.Text) < passwordMaxLength {
		labelAlertAuth.SetText(errors.ErrPasswordIncorrect)
		log.Print(labelAlertAuth.Text)
		return false
	}
	return true
}

func ValidateRegistration(usernameRegistrationEntry *widget.Entry, passwordRegistrationEntry *widget.Entry,
	passwordConfirmationRegistrationEntry *widget.Entry, labelAlertAuth *widget.Label) bool {
	if utf8.RuneCountInString(usernameRegistrationEntry.Text) < userNameMaxLength {
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

func ValidateLoginPassword(exists bool, loginPasswordNameEntry *widget.Entry, loginPasswordDescriptionEntry *widget.Entry,
	loginEntry *widget.Entry, passwordEntry *widget.Entry, labelAlertLoginPassword *widget.Label) bool {
	if exists {
		labelAlertLoginPassword.SetText(errors.ErrLoginPasswordExist)
		log.Print(labelAlertLoginPassword.Text)
		return false
	}
	if loginPasswordNameEntry.Text == "" {
		labelAlertLoginPassword.SetText(errors.ErrNameEmpty)
		log.Print(labelAlertLoginPassword.Text)
		return false
	}
	if loginPasswordDescriptionEntry.Text == "" {
		labelAlertLoginPassword.SetText(errors.ErrDescriptionEmpty)
		log.Print(labelAlertLoginPassword.Text)
		return false
	}
	if loginEntry.Text == "" {
		labelAlertLoginPassword.SetText(errors.ErrLoginEmpty)
		log.Print(labelAlertLoginPassword.Text)
		return false
	}
	if passwordEntry.Text == "" {
		labelAlertLoginPassword.SetText(errors.ErrPasswordEmpty)
		log.Print(labelAlertLoginPassword.Text)
		return false
	}
	return true
}

func ValidateText(exists bool, textNameEntry *widget.Entry, textEntry *widget.Entry, textDescriptionEntry *widget.Entry,
	labelAlertText *widget.Label) bool {
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
	if textDescriptionEntry.Text == "" {
		labelAlertText.SetText(errors.ErrDescriptionEmpty)
		log.Print(labelAlertText.Text)
		return false
	}
	if textEntry.Text == "" {
		labelAlertText.SetText(errors.ErrTextEmpty)
		log.Print(labelAlertText.Text)
		return false
	}
	return true
}

func ValidateCard(exists bool, cardNameEntry *widget.Entry, cardDescriptionEntry *widget.Entry, paymentSystemEntry *widget.Entry,
	numberEntry *widget.Entry, holderEntry *widget.Entry, endDateEntry *widget.Entry, cvcEntry *widget.Entry, labelAlertCard *widget.Label) bool {
	var err error
	if exists {
		labelAlertCard.SetText(errors.ErrCardExist)
		log.Print(labelAlertCard)
		return false
	}
	if cardNameEntry.Text == "" {
		labelAlertCard.SetText(errors.ErrNameEmpty)
		log.Print(labelAlertCard.Text)
		return false
	}
	if cardDescriptionEntry.Text == "" {
		labelAlertCard.SetText(errors.ErrDescriptionEmpty)
		log.Print(labelAlertCard.Text)
		return false
	}
	if paymentSystemEntry.Text == "" {
		labelAlertCard.SetText(errors.ErrPaymentSystemEmpty)
		log.Print(labelAlertCard.Text)
		return false
	}
	if numberEntry.Text == "" {
		labelAlertCard.SetText(errors.ErrNumberEmpty)
		log.Print(labelAlertCard.Text)
		return false
	}
	intNumber, err := strconv.Atoi(numberEntry.Text)
	if err != nil {
		labelAlertCard.SetText(errors.ErrNumberIncorrect)
		log.Print(labelAlertCard.Text)
		return false
	}
	if !algorithm.ValidLuhn(intNumber) {
		labelAlertCard.SetText(errors.ErrNumberIncorrect)
		log.Print(labelAlertCard.Text)
		return false
	}
	if holderEntry.Text == "" {
		labelAlertCard.SetText(errors.ErrHolderEmpty)
		log.Print(labelAlertCard.Text)
		return false
	}
	if endDateEntry.Text == "" {
		labelAlertCard.SetText(errors.ErrEndDateEmpty)
		log.Print(labelAlertCard.Text)
		return false
	} else {
		layout := "01/02/2006"
		_, err = time.Parse(layout, endDateEntry.Text)
		if err != nil {
			labelAlertCard.SetText(errors.ErrEndDataIncorrect)
			log.Print(labelAlertCard.Text)
			return false
		}
	}
	if cvcEntry.Text == "" {
		labelAlertCard.SetText(errors.ErrCvcEmpty)
		log.Print(labelAlertCard.Text)
		return false
	} else {
		_, err = strconv.Atoi(cvcEntry.Text)
		if err != nil {
			labelAlertCard.SetText(errors.ErrCvcIncorrect)
			log.Print(labelAlertCard.Text)
			return false
		}
	}
	return true
}
