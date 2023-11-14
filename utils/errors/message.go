package errors

import (
	"errors"
)

var (
	ErrUserEmailEmpty       = errors.New("user email is empty")
	ErrUserNameEmpty        = errors.New("user name is empty")
	ErrUserPasswordEmpty    = errors.New("user password is empty")
	ErrBcryptPassword       = errors.New("failed bcrypt password")
	ErrRegisterUserDatabase = errors.New("error register user database")
	ErrNotFound             = errors.New("data dot found")
	ErrBadRequest           = errors.New("bad request")
	ErrOTPWrong             = errors.New("otp is wrong")
	ErrVerifyIdIsInvalid    = errors.New("verify id is invalid")
	ErrUserHasVerfication   = errors.New("user has verify identity")
)
