package parse

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	// ErrEmpty is returned by all parse functions when the input string is empty.
	ErrEmpty = errors.New("empty string")

	// ErrSyntax is returned by all parse functions when the input string contains
	// characters illegal in their position in the string for the type parsed.
	ErrSyntax = errors.New("invalid syntax")
)

// Error is the error type returned by parse functions.
//
// Func is always the name of the function which returned the error.
// Num is always the input string supplied to the function.
// Err is always one of ErrEmpty or ErrSyntax.
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
