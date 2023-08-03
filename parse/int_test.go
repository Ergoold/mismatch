package parse

import (
	"testing"
)

var (
	positiveIntegerTests = []testCase[int]{
		{"", 0, ErrEmpty},
		{"1", 1, nil},
		{"12", 12, nil},
		{"01", 0, ErrSyntax},
		{"0", 0, ErrSyntax},
		{"12x", 0, ErrSyntax},
		{"0x", 0, ErrSyntax},
		{"x0", 0, ErrSyntax},
	}
	nonnegativeIntegerTests = []testCase[int]{
		{"", 0, ErrEmpty},
		{"1", 1, nil},
		{"12", 12, nil},
		{"01", 0, ErrSyntax},
		{"0", 0, nil},
		{"12x", 0, ErrSyntax},
		{"0x", 0, ErrSyntax},
		{"x0", 0, ErrSyntax},
	}
)

func TestPositiveInteger(t *testing.T) {
	runTestCases(t, "PositiveInteger", PositiveInteger, positiveIntegerTests)
}

func TestNonnegativeInteger(t *testing.T) {
	runTestCases(t, "NonnegativeInteger", NonnegativeInteger, nonnegativeIntegerTests)
}
