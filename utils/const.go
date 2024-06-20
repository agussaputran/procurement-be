package utils

const (
	ErrDataNotFound          = "mongo: no documents in result"
	ErrIDCannotBeEmpty       = "ID cannot be empty"
	ErrUsernameAlreadyExists = "username already exists"
	ErrPhoneAlreadyExists    = "phone already exists"
	ErrEmailAlreadyExists    = "email already exists"
	ErrPasswordNotMatch      = "password not match"
	ErrInvalidPassword       = "invalid password"
	ErrTokenDoesNotExists    = "token does not exists"
	ErrHeaderDoesNotExists   = "header does not exists"
	ErrCompanyDoesNotExists  = "company does not exists"
	ErrEmailAlreadyTaken     = "email already taken"
	ErrAccountAlreadyActive  = "account already active"
	ErrCompanyAlreadyExists  = "company already exists"
)

const (
	DateLayout             = "2006-01-02"
	DatetimeLayout         = "2006-01-02 15:04:05"
	DatetimeTimezoneLayout = "2006-01-02T15:04:05"
	DatetimeRFC3339        = "2006-01-02T15:04:05Z07:00"
)
