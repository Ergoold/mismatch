package parse

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	ErrEmpty  = errors.New("empty string")
	ErrSyntax = errors.New("invalid syntax")
)

type Error struct {
	Func string
	Num  string
	Err  error
}

func (e *Error) Error() string {
	return fmt.Sprintf("parse.%s: parsing %s: %s", e.Func, strconv.Quote(e.Num), e.Err)
}

func (e *Error) Unwrap() error {
	return e.Err
}
