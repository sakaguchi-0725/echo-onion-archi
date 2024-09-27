package apperr

type ApplicationError struct {
	Code    ErrorCode
	Message string
	Err     error
}

func NewApplicationError(code ErrorCode, message string, err error) *ApplicationError {
	return &ApplicationError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func (e *ApplicationError) Error() string {
	return e.Err.Error()
}
