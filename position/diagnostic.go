package position

import (
	"fmt"
	"io"
	"mismatch/report"
	"os"
)

func (p Position) WriteDiagnostic(diagnostic string, file *os.File, writer io.Writer) {
	if _, err := fmt.Fprintf(writer, "mismatch: %v:%v:%v: %v '%c'\n", file.Name(), p.line, p.char, diagnostic, p.value); err != nil {
		report.Fatal(err)
	}
}

func (p Position) ShowPositionInLine(reader io.ReadSeeker, writer io.Writer) error {
	if _, err := reader.Seek(int64(p.lineStart), io.SeekStart); err != nil {
		return err
	}

	line := make([]byte, p.lineEnd-p.lineStart)
	if _, err := reader.Read(line); err != nil {
		return err
	}

	lineNumber := fmt.Sprintf("%v", p.line)

	if _, err := fmt.Fprintf(writer, "%v | %v", lineNumber, string(line)); err != nil {
		report.Fatal(err)
	}

	for range lineNumber {
		if _, err := fmt.Fprintf(writer, "%c", ' '); err != nil {
			report.Fatal(err)
		}
	}

	if _, err := fmt.Fprintf(writer, "   "); err != nil {
		report.Fatal(err)
	}

	for i := 0; i < p.char-1; i++ {
		var char byte = ' '
		if line[i] == '\t' {
			char = '\t'
		}

		if _, err := fmt.Fprintf(writer, "%c", char); err != nil {
			report.Fatal(err)
		}
	}

	if _, err := fmt.Fprintf(writer, "^ here\n\n"); err != nil {
		report.Fatal(err)
	}

	return nil
}
