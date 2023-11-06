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
)
