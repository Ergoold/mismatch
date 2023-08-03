package parse

import (
	"testing"
)

var (
	positiveFloatTests = []testCase[float64]{
		{"", 0, ErrEmpty},
		{"0", 0, ErrSyntax},
		{"1", 1, nil},
		{"12x", 0, ErrSyntax},
		{"x0", 0, ErrSyntax},
		{"0.5", 0.5, nil},
		{"1.x", 0, ErrSyntax},
		{"x.2", 0, ErrSyntax},
		{"3.4x", 0, ErrSyntax},
		{"x5.6", 0, ErrSyntax},
		{"7x.8", 0, ErrSyntax},
		{"3.14", 3.14, nil},
		{"1.", 0, ErrSyntax},
		{".618", 0, ErrSyntax},
	}
)

func TestPositiveFloat(t *testing.T) {
	runTestCases(t, "PositiveFloat", PositiveFloat, positiveFloatTests)
}
