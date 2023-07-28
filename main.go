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
	if depth, err := Depth(file); err == nil {
		if _, err = fmt.Fprintf(os.Stdout, "%v\n", depth); err != nil {
			panic(err)
		}
	} else {
		report.Fatal(err)
	}
}
