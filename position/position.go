package position

// Position is a character in a file.
type Position struct {
	line, char, pos int

	value byte
}

// Make a new Position, initialized to the beginning of the file.
func Make() Position {
	return Position{line: 1, char: 0, pos: 0, value: 0x00}
}

// NextChar returns a new position, that is advanced by one character.
func (p Position) NextChar(value byte) Position {
	nextChar := Position{p.line, p.char + 1, p.pos + 1, value}

	if value == '\n' {
		nextChar.line++
		nextChar.char = 0
	}

	return nextChar
}

func (p Position) Value() byte {
	return p.value
}
