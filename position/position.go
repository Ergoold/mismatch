package position

// Position is a character in a file.
type Position struct {
	line, char, pos int

	value byte

	lineStart, lineEnd int
}

// Make a new Position, initialized to the beginning of the file.
func Make() Position {
	return Position{line: 1, char: 0, pos: 0, value: 0x00, lineStart: 0, lineEnd: 0}
}

// NextChar returns a new position, that is advanced by one character.
func (p Position) NextChar(value byte) Position {
	nextChar := Position{p.line, p.char + 1, p.pos + 1, value, p.lineStart, p.lineEnd}

	if value == '\n' {
		nextChar.line++
		nextChar.char = 0

		nextChar.lineStart = nextChar.pos
	}

	return nextChar
}

// UpdateLine updates all positions in ps that don't have lineEnd to end the line at p.
//
// Assumes ps is sorted by line (earlier lines appear earlier).
func (p Position) UpdateLine(ps []Position) []Position {
	for i := len(ps) - 1; i >= 0; i-- {
		pos := ps[i]

		if pos.lineEnd != 0 {
			break
		}

		ps[i] = Position{pos.line, pos.char, pos.pos, pos.value, pos.lineStart, p.pos}
	}

	return ps
}

func (p Position) Value() byte {
	return p.value
}
