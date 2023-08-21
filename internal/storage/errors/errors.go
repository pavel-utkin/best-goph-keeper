package errors

import "errors"

var (
	ErrTypeIncorrect           = errors.New("the type incorrect")
	ErrWrongUsernameOrPassword = errors.New("wrong username or password")
	ErrRecordNotFound          = errors.New("record not found")
	ErrUserAlreadyExists       = errors.New("user with this name already exists")
	ErrBadPassword             = errors.New("password rules: at least 7 letters, 1 number, 1 upper case, 1 special character")
	ErrBadText                 = errors.New("text rules: at least 7 letters")
)
