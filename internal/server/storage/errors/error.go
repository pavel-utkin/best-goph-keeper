package errors

import "errors"

var (
	ErrTypeIncorrect           = errors.New("the type incorrect")
	ErrBadName                 = errors.New("name cannot be empty")
	ErrWrongUsernameOrPassword = errors.New("wrong username or password")
	ErrRecordNotFound          = errors.New("record not found")
	ErrUsernameAlreadyExists   = errors.New("username already exists")
	ErrKeyAlreadyExists        = errors.New("key already exists")
	ErrBadPassword             = errors.New("password rules: at least 8 letters, 1 number, 1 upper case, 1 special character")
	ErrBadText                 = errors.New("text rules: at least 10 letters")
	ErrNoMetadataSet           = errors.New("no metadata set")
	ErrNotValidateToken        = errors.New("not validate token")
)
