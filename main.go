package main

import (
	"mismatch/report"
	"os"
)

func main() {
	processFile(os.Stdin)
}

func processFile(file *os.File) {
	if stack, err := Mismatches(file); err == nil {
		WriteDiagnostics(stack, file, os.Stdout)
	} else {
		report.Fatal(err)
	}
}
