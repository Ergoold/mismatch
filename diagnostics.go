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
		writeDiagnostic(pos, file, writer)

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

func writeDiagnostic(pos position.Position, file *os.File, writer io.Writer) {
	value := pos.Value()

	diagnostic := diagnostics[value]

	pos.WriteDiagnostic(diagnostic, file, writer)
}
