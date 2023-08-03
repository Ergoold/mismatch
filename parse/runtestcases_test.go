package parse

import (
	"fmt"
	"testing"
)

type testCase[T any] struct {
	in  string
	out T
	err error
}

func runTestCases[T comparable](t *testing.T, fnName string, fn func(string) (T, error), tests []testCase[T]) {
	for _, test := range tests {
		out, err := fn(test.in)
		if out == test.out && err == nil && test.err == nil {
			continue
		}

		switch e := err.(type) {
		case *Error:
			if out != test.out || e.Func != fnName || e.Num != test.in || e.Err != test.err {
				t.Errorf("parse.%v(%q) = %v, &Error{%q, %q, %v}, want %v, &Error{%q, %q, %v}",
					fnName, test.in, out, e.Func, e.Num, errorName(e.Err), test.out, fnName, test.in, errorName(test.err))
			}
		default:
			t.Errorf("parse.%v(%q) = %v, nil, want %v, &Error{%q, %q, %v}",
				fnName, test.in, out, test.out, fnName, test.in, errorName(test.err))
		}
	}
}

func errorName(err error) string {
	switch err {
	case ErrEmpty:
		return "ErrEmpty"
	case ErrSyntax:
		return "ErrSyntax"
	default:
		return fmt.Sprintf("%#v", err)
	}
}
