package etc

import (
	"strconv"
	"strings"
)

type Error struct {
	Err     error
	Code    int
	Message string
	Src     string
}

func (e Error) Error() string {
	format := strings.Builder{}
	format.WriteString("error: ")
	format.WriteString(e.Err.Error())

	if e.Code != 0 {
		format.WriteString(" | code: ")
		format.WriteString(strconv.Itoa(e.Code))
	}
	if e.Message != "" {
		format.WriteString(" | mes: ")
		format.WriteString(e.Message)
	}

	if e.Src != "" {
		format.WriteString(" | src: ")
		format.WriteString(e.Src)
	}
	return format.String()
}

// NewErr Returns new error. First optional arg is custom message second is src
func NewErr(code int, err error, mesSrc ...string) Error {
	newErr := Error{
		Code: code,
		Err:  err,
	}

	if len(mesSrc) >= 1 {
		newErr.Message = mesSrc[0]
	}

	if len(mesSrc) >= 2 {
		newErr.Src = mesSrc[1]
	}

	return newErr
}
