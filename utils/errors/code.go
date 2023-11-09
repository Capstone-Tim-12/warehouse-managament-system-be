package errors

import "net/http"

func GetCodeError(err error) int {
	switch err {
	case ErrBcryptPassword:
		return http.StatusInternalServerError
	case ErrRegisterUserDatabase:
		return http.StatusInternalServerError
	case ErrUserEmailEmpty:
		return http.StatusBadRequest
	case ErrUserPasswordEmpty:
		return http.StatusBadRequest
	case ErrUserNameEmpty:
		return http.StatusBadRequest
	case ErrNotFound:
		return http.StatusNotFound
	case ErrBadRequest:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
