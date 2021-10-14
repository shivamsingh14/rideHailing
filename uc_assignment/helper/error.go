package helper

type Error struct {
	code     string
	message  string
	httpCode int
}

const (
	BAD_REQUEST = 400
)

func BadError(code, message string) *Error {
	return &Error{
		code:     code,
		message:  message,
		httpCode: 400,
	}
}
