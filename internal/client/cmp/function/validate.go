package function

import (
	"best-goph-keeper/internal/client/service/algorithm"
	"best-goph-keeper/internal/client/service/encryption"
	"best-goph-keeper/internal/client/storage/errors"
	"best-goph-keeper/internal/client/storage/layouts"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"time"
	"unicode/utf8"
)

const userNameMaxLength = 6

func ValidateLoginForm(usernameLoginEntry *widget.Entry, passwordLoginEntry *widget.Entry) (string, bool) {
	if utf8.RuneCountInString(usernameLoginEntry.Text) < userNameMaxLength {
		return errors.ErrUsernameIncorrect, false
	}
	if !encryption.VerifyPassword(passwordLoginEntry.Text) {
		return errors.ErrPasswordIncorrect, false
	}
	return "", true
}

func ValidateRegistrationForm(usernameRegistrationEntry *widget.Entry, passwordRegistrationEntry *widget.Entry,
	passwordConfirmationRegistrationEntry *widget.Entry) (string, bool) {
	if utf8.RuneCountInString(usernameRegistrationEntry.Text) < userNameMaxLength {
		return errors.ErrUsernameIncorrect, false
	}
	if !encryption.VerifyPassword(passwordRegistrationEntry.Text) {
		return errors.ErrPasswordIncorrect, false
	}
	if passwordRegistrationEntry.Text != passwordConfirmationRegistrationEntry.Text {
		return errors.ErrPasswordDifferent, false
	}
	return "", true
}

func ValidateLoginPasswordForm(loginPasswordNameEntry *widget.Entry, loginPasswordDescriptionEntry *widget.Entry,
	loginEntry *widget.Entry, passwordEntry *widget.Entry) (string, bool) {
	if loginPasswordNameEntry.Text == "" {
		return errors.ErrNameEmpty, false
	}
	if loginPasswordDescriptionEntry.Text == "" {
		return errors.ErrDescriptionEmpty, false
	}
	if loginEntry.Text == "" {
		return errors.ErrLoginEmpty, false
	}
	if passwordEntry.Text == "" {
		return errors.ErrPasswordEmpty, false
	}
	return "", true
}

func ValidateTextForm(textNameEntry *widget.Entry, textDescriptionEntry *widget.Entry, textEntry *widget.Entry) (string, bool) {
	if textNameEntry.Text == "" {
		return errors.ErrNameEmpty, false
	}
	if textDescriptionEntry.Text == "" {
		return errors.ErrDescriptionEmpty, false
	}
	if textEntry.Text == "" {
		return errors.ErrTextEmpty, false
	}
	return "", true
}

func ValidateCardForm(cardNameEntry *widget.Entry, cardDescriptionEntry *widget.Entry, paymentSystemEntry *widget.Entry,
	numberEntry *widget.Entry, holderEntry *widget.Entry, cvcEntry *widget.Entry, endDateEntry *widget.Entry) (string, bool) {
	var err error
	if cardNameEntry.Text == "" {
		return errors.ErrNameEmpty, false
	}
	if cardDescriptionEntry.Text == "" {
		return errors.ErrDescriptionEmpty, false
	}
	if paymentSystemEntry.Text == "" {
		return errors.ErrPaymentSystemEmpty, false
	}
	if numberEntry.Text == "" {
		return errors.ErrNumberEmpty, false
	}
	intNumber, err := strconv.Atoi(numberEntry.Text)
	if err != nil {
		return errors.ErrNumberIncorrect, false
	}
	if !algorithm.ValidLuhn(intNumber) {
		return errors.ErrNumberIncorrect, false
	}
	if holderEntry.Text == "" {
		return errors.ErrHolderEmpty, false
	}
	if endDateEntry.Text == "" {
		return errors.ErrEndDateEmpty, false
	} else {
		_, err = time.Parse(string(layouts.LayoutDate), endDateEntry.Text)
		if err != nil {
			return errors.ErrEndDateIncorrect, false
		}
	}
	if cvcEntry.Text == "" {
		return errors.ErrCvcEmpty, false
	} else {
		_, err = strconv.Atoi(cvcEntry.Text)
		if err != nil {
			return errors.ErrCvcIncorrect, false
		}
	}
	return "", true
}
