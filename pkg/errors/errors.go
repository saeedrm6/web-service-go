package errors

import "net/http"

type (
	kind        uint
	serverError struct {
		kind    kind
		message string
	}
)

const (
	_ kind = iota
	KindInvalid
	KindNotFound
	KindUnAuthorized
	KindUnexpected
	KindNotAllowed
)

var (
	httpErrors = map[kind]int{
		KindInvalid:      http.StatusBadRequest,
		KindNotFound:     http.StatusNotFound,
		KindUnAuthorized: http.StatusUnauthorized,
		KindUnexpected:   http.StatusInternalServerError,
		KindNotAllowed:   http.StatusMethodNotAllowed,
	}
)

func New(kind kind, msg string) error {
	return serverError{
		kind:    kind,
		message: msg,
	}
}

func (e serverError) Error() string {
	return e.message
}

func HttpError(err error) (string, int) {
	serverErr, ok := err.(serverError)
	if !ok {
		return err.Error(), http.StatusBadRequest
	}

	code, ok := httpErrors[serverErr.kind]
	if !ok {
		return serverErr.message, http.StatusBadRequest
	}

	return serverErr.message, code
}
