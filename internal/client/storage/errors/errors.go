package errors

var (
	ErrLogin              = "Неудачная попытка аутентификации"
	ErrRegistration       = "Неудачная попытка регистрации"
	ErrUserExist          = "Пользователь с таким username зарегистрирован"
	ErrUsernameIncorrect  = "Длинна username должна быть не менее шести символов"
	ErrPasswordIncorrect  = "Пароль должен состоять минимум из 8 букв, 1 цифра, 1 верхний регистр, 1 специальный символ (пример: Password-123)"
	ErrPasswordDifferent  = "Пароли не совпали"
	ErrNameEmpty          = "Name не заполнен"
	ErrTextEmpty          = "Text не заполнен"
	ErrDescriptionEmpty   = "Description не заполнен"
	ErrTextExist          = "Текст с таким name уже существует"
	ErrCartExist          = "Карта с таким name уже существует"
	ErrPaymentSystemEmpty = "Payment System не заполнен"
	ErrNumberEmpty        = "Number не заполнен"
	ErrHolderEmpty        = "Holder не заполнен"
	ErrEndDateEmpty       = "End date не заполнен"
	ErrEndDataIncorrect   = "End Date некорректный (пример: 01/02/2006)"
	ErrCvcEmpty           = "CVC не заполнен"
	ErrCvcIncorrect       = "CVC некорректный (пример: 123)"
	ErrTextAdd            = "Ошибка при добавлении текста (попробуйте обновить данные)"
)
