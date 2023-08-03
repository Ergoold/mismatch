package parse

import (
	"math"
	"strings"
)

func PositiveFloat(s string) (float64, error) {
	const funcName = "PositiveFloat"

	if before, after, found := strings.Cut(s, "."); found {
		result := 0.0

		if integerPart, err := NonnegativeInteger(before); err == nil {
			result = float64(integerPart)
		} else {
			return 0, &Error{funcName, s, ErrSyntax}
		}

		if len(after) == 0 {
			return 0, &Error{funcName, s, ErrSyntax}
		}

		for i, c := range after {
			if !isDigit(c) {
				return 0, &Error{funcName, s, ErrSyntax}
			}

			result += float64(digitToInteger(c)) / math.Pow10(i+1)
		}

		return result, nil
	} else {
		if result, err := PositiveInteger(s); err == nil {
			return float64(result), nil
		} else {
			return 0, &Error{funcName, s, err.(*Error).Unwrap()}
		}
	}
}
