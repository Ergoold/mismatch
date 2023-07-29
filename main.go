package main

import (
	"fmt"
	"mismatch/report"
	"os"
)

func main() {
	processFile(os.Stdin)
}

func processFile(file *os.File) {
	if stack, err := Mismatches(file); err == nil {
		if _, err = fmt.Fprintf(os.Stdout, "%v\n", len(stack)); err != nil {
			panic(err)
		}
	} else {
		report.Fatal(err)
	}
}
