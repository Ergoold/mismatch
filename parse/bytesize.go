package parse

import (
	"strings"
)

var (
	units = map[string]int{
		// SI standard abbreviation for kilo
		"k": 1e3,
		// SI prefixes
		"K": 1e3,
		"M": 1e6,
		"G": 1e9,
		"T": 1e12,
		// IEC binary prefixes
		"Ki": 1 << 10,
		"Mi": 1 << 20,
		"Gi": 1 << 30,
		"Ti": 1 << 40,
	}
)

func unit(s string) (int, error) {
	if len(s) == 0 {
		return 1, ErrEmpty
	}

	lastIndex := len(s) - 1
	if s[lastIndex] == 'B' {
		s = s[:lastIndex]
	}

	if value, ok := units[s]; ok {
		return value, nil
	} else {
		return 0, ErrSyntax
	}
}

// ByteSize parses a string representing a number of bytes into an int.
//
// The units supported are the SI prefix abbreviations k (kilobyte) through T (terabyte)
// and the corresponding binary prefixes. Units may contain or omit the letter B,
// representing the quantity bytes.
func ByteSize(s string) (int, error) {
	const funcName = "ByteSize"

	if len(s) == 0 {
		return 0, &Error{funcName, s, ErrEmpty}
	}

	if result, err := PositiveInteger(s); err == nil {
		return result, nil
	}

	if index := strings.LastIndexFunc(s, isDigit); index != -1 {
		before, after := s[:index+1], s[index+1:]

		quantity, err := PositiveFloat(before)
		if err != nil {
			return 0, &Error{funcName, s, err.(*Error).Unwrap()}
		}

		unit, err := unit(after)
		if err != nil {
			return 0, &Error{funcName, s, ErrSyntax}
		}

		byteSize := quantity * float64(unit)

		return int(byteSize), nil
	}

	return 0, &Error{funcName, s, ErrSyntax}
}
