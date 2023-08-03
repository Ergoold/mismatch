package parse

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
