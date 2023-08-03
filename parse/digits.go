package parse

func isDigit(c rune) bool {
	return '0' <= c && c <= '9'
}

func isPositiveDigit(c rune) bool {
	return '1' <= c && c <= '9'
}

func digitToInteger(c rune) int {
	return int(c - '0')
}
