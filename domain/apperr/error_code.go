package apperr

type ErrorCode int

const (
	ErrBadReqeust ErrorCode = iota
	ErrNotFound
	ErrUnauthorized
	ErrInternalError
)

func (e ErrorCode) String() string {
	switch e {
	case ErrBadReqeust:
		return "BadRequest"
	case ErrNotFound:
		return "NotFound"
	case ErrUnauthorized:
		return "Unauthorized"
	case ErrInternalError:
		return "InternalServerError"
	default:
		return "InternalServerError"
	}
}
