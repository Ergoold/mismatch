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

// WriteDiagnostics writes a diagnostic for each position in ps in file to writer.
func WriteDiagnostics(ps []position.Position, file *os.File, writer io.Writer) {
	showInLine := true

	for _, pos := range ps {
		pos.WriteDiagnostic(diagnostics, file, writer)

		if showInLine {
			if err := pos.ShowPositionInLine(file, writer); err != nil {
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
}
