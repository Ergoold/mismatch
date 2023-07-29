package main

import (
	"io"
	"mismatch/position"
	"os"
)

var (
	buffer = make([]byte, BufferSize)
	stack  = make([]position.Position, InitialStackDepth)
)

// Mismatches returns all mismatched parentheses in file.
func Mismatches(file *os.File) ([]position.Position, error) {
	stack = stack[0:0]
	pos := position.Make()

ReadLoop:
	for {
		switch read, err := file.Read(buffer); err {
		case nil:
			for i := 0; i < read; i++ {
				pos, stack = processByte(buffer[i], pos, stack)
			}
		case io.EOF:
			break ReadLoop
		default:
			return stack, err
		}
	}

	return stack, nil
}

func processByte(currentByte byte, previousPos position.Position, stack []position.Position) (position.Position, []position.Position) {
	pos := previousPos.NextChar(currentByte)

	switch currentByte {
	case '\n':
		stack = pos.UpdateLine(stack)
	case '(':
		stack = append(stack, pos)
	case ')':
		if len(stack) > 0 && stack[len(stack)-1].Value() == '(' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, pos)
		}
	}

	return pos, stack
}
