package parse

// PositiveInteger parses a string into a positive integer of type int.
//
// Only base 10 is supported. The entire string is parsed. If the string begins
// with a number but contains non-numeric characters, an ErrSyntax is returned.
//
// Signs and leading zeros are not supported and cause ErrSyntax to be returned.
func PositiveInteger(s string) (int, error) {
	const funcName = "PositiveInteger"

	if s == "" {
		return 0, &Error{funcName, s, ErrEmpty}
	}

	if !isPositiveDigit(rune(s[0])) {
		return 0, &Error{funcName, s, ErrSyntax}
	}

	result := digitToInteger(rune(s[0]))

	for _, c := range s[1:] {
		if !isDigit(c) {
			return 0, &Error{funcName, s, ErrSyntax}
		}

		result *= 10
		result += digitToInteger(c)
	}

	return result, nil
}

// NonnegativeInteger parses a string into a nonnegative (i.e., positive or zero) integer of type int.
//
// Only base 10 is supported. The entire string is parsed. If the string begins
// with a number but contains non-numeric characters, an ErrSyntax is returned.
//
// Signs and leading zeros are not supported and cause ErrSyntax to be returned.
func NonnegativeInteger(s string) (int, error) {
	const funcName = "NonnegativeInteger"

	if s == "0" {
		return 0, nil
	}

	if result, err := PositiveInteger(s); err == nil {
		return result, nil
	} else {
		return result, &Error{funcName, s, err.(*Error).Unwrap()}
	}
}
