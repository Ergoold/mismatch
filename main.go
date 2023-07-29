package main

import (
	"mismatch/report"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		processFile(os.Stdin)
	}

	for _, fileName := range os.Args[1:] {
		if file, err := os.Open(fileName); err == nil {
			processFile(file)

			if err = file.Close(); err != nil {
				report.Warning(err)
			}
		} else {
			report.Warning(err)
		}
	}
}

func processFile(file *os.File) {
	if stack, err := Mismatches(file); err == nil {
		WriteDiagnostics(stack, file, os.Stdout)
	} else {
		report.Warning(err)
	}
}
