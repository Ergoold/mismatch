package report

import (
	"fmt"
	"os"
)

func Fatal(err error) {
	Warning(err)
	os.Exit(1)
}

func Warning(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "mismatch: %v\n\n", err)
}
