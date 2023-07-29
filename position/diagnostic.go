package position

import (
	"fmt"
	"io"
	"mismatch/report"
	"os"
)

func (p Position) WriteDiagnostic(diagnostic string, file *os.File, writer io.Writer) {
	if _, err := fmt.Fprintf(writer, "mismatch: %v:%v:%v: %v '%c'\n\n", file.Name(), p.line, p.char, diagnostic, p.value); err != nil {
		report.Fatal(err)
	}
}
