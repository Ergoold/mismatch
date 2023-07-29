package main

import (
	"io"
	"mismatch/position"
	"os"
)

var (
	diagnostics = map[byte]string{
		'(': "unclosed",
		')': "unopened",
	}
)

func WriteDiagnostics(ps []position.Position, file *os.File, writer io.Writer) {
	for _, pos := range ps {
		writeDiagnostic(pos, file, writer)
	}
}

func writeDiagnostic(pos position.Position, file *os.File, writer io.Writer) {
	value := pos.Value()

	diagnostic := diagnostics[value]

	pos.WriteDiagnostic(diagnostic, file, writer)
}
