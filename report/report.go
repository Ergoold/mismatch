package report

import (
	"fmt"
	"os"
)

// Fatal prints err to stderr and exits with a non-zero return value.
func Fatal(errs ...error) {
	Warning(errs...)
	os.Exit(1)
}

// Warning prints err to stderr without exiting the program.
func Warning(errs ...error) {
	for _, err := range errs {
		printErr(err)
	}
	_, _ = fmt.Fprintf(os.Stderr, "\n")
}

func printErr(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "mismatch: %v\n", err)
}
