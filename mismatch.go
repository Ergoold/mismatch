package main

import (
	"io"
	"mismatch/position"
	"os"
)

var (
	buffer = make([]byte, 1)
	depth  = 0
)

// Depth returns the difference between the number of opening and closing parentheses in file.
func Depth(file *os.File) (int, error) {
	pos := position.Make()

ReadLoop:
	for {
		switch _, err := file.Read(buffer); err {
		case nil:
			pos, depth = processByte(buffer[0], pos, depth)
		case io.EOF:
			break ReadLoop
		default:
			return depth, err
		}
	}

	return depth, nil
}

func processByte(currentByte byte, previousPos position.Position, previousDepth int) (position.Position, int) {
	pos := previousPos.NextChar(currentByte)

	switch currentByte {
	case '(':
		return pos, previousDepth + 1
	case ')':
		return pos, previousDepth - 1
	}

	return pos, previousDepth
}
