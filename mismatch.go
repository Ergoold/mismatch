package main

import (
	"io"
	"mismatch/position"
	"os"
)

var (
	buffer = make([]byte, BufferSize)
	depth  = 0
)

// Depth returns the difference between the number of opening and closing parentheses in file.
func Depth(file *os.File) (int, error) {
	pos := position.Make()

ReadLoop:
	for {
		switch read, err := file.Read(buffer); err {
		case nil:
			for i := 0; i < read; i++ {
				pos, depth = processByte(buffer[i], pos, depth)
			}
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
