package errorx

const (
	InternalServerError int = iota + 500000
)

const (
	InvalidRequestPayload int = iota + 400000
	DuplicateResource     int = iota
)

type ErrorCode interface {
	Code() int
}

type ErrorWithCode struct {
	code int
}

func NewErrorWithCode(code int) *ErrorWithCode {
	return &ErrorWithCode{code: code}
}

func (err *ErrorWithCode) Code() int {
	return err.code
}