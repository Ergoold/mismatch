package main

import (
	"errors"
	"fmt"
	"mismatch/parse"
	"mismatch/report"
	"os"
)

var (
	// BufferSize is the size of the internal read buffer.
	BufferSize = env("BUFFER_SIZE", parse.ByteSize, 64)

	// InitialStackDepth is the initial stack depth for keeping track of mismatched parentheses.
	InitialStackDepth = env("INITIAL_STACK_DEPTH", parse.PositiveInteger, 4)
)

type parser[T any] func(string) (T, error)

func env[T any](name string, parser parser[T], defaultValue T) T {
	value := os.Getenv(name)

	if parsed, err := parser(value); err == nil {
		return parsed
	} else if !errors.Is(err, parse.ErrEmpty) {
		report.Warning(&envError{name, err})
	}

	return defaultValue
}

type envError struct {
	name string
	err  error
}

func (e *envError) Error() string {
	return fmt.Sprintf("parsing %s: %v", e.name, e.err)
}
