package apperror

var (
	ErrBadRequest     = NewError("bad request")
	ErrNotFound       = NewError("not found")
	ErrInternalServer = NewError("internal server")
)

type Error struct {
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func NewError(message string) *Error {
	return &Error{Message: message}
}
