package errors

import (
	"fmt"
	"github.com/pkg/errors"
)

type privateError struct {
	code Code
	err  error
}

func (e privateError) Error() string {
	return fmt.Sprintf("Code: %s, Msg: %s", e.code, e.err)
}

func Errorf(c Code, format string, a ...interface{}) error {
	if c == OK {
		return nil
	}
	return privateError{
		code: c,
		err:  errors.Errorf(format, a...), // github.com/pkg/errorsでラップする
	}
}

func GetCode(err error) Code {
	if err == nil {
		return OK
	}
	var e privateError
	if errors.As(err, &e) {
		return e.code
	}
	return Unknown
}

func StackTrace(err error) string {
	var e privateError
	if errors.As(err, &e) {
		return fmt.Sprintf("%+v\n", e.err)
	}
	return ""
}
