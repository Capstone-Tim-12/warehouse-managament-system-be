package errors

type ApplicationError struct {
	ErrorCode int
	Message   string
}

func New(errorCode int, message string) error {
	return &ApplicationError{
		ErrorCode: errorCode,
		Message:   message,
	}
}

func (e *ApplicationError) Error() string {
	return e.Message
}
