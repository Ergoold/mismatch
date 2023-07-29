package report

import (
	"fmt"
	"os"
)

// Fatal prints err to stderr and exits with a non-zero return value.
func Fatal(err error) {
	Warning(err)
	os.Exit(1)
}

// Warning prints err to stderr without exiting the program.
func Warning(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "mismatch: %v\n\n", err)
}
