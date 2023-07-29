package main

import (
	"io"
	"mismatch/position"
	"mismatch/report"
	"os"
)

var (
	diagnostics = map[byte]string{
		'(': "unclosed",
		')': "unopened",
	}
)

func WriteDiagnostics(ps []position.Position, file *os.File, writer io.Writer) {
	showInLine := true

	for _, pos := range ps {
		if err := writeDiagnostic(pos, file, writer, showInLine); err != nil {
			report.Warning(err)

			switch pathError := err.(type) {
			case *os.PathError:
				if pathError.Op != "seek" {
					return
				}

				showInLine = false
			default:
				return
			}
		}
	}
}

func writeDiagnostic(pos position.Position, file *os.File, writer io.Writer, showInLine bool) error {
	value := pos.Value()

	diagnostic := diagnostics[value]

	return pos.WriteDiagnostic(diagnostic, file, writer, showInLine)
}
