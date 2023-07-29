package main

import (
	"os"
	"strconv"
)

var (
	// BufferSize is the size of the internal read buffer.
	BufferSize = env("BUFFER_SIZE", strconv.Atoi, 64)

	// InitialStackDepth is the initial stack depth for keeping track of mismatched parentheses.
	InitialStackDepth = env("INITIAL_STACK_DEPTH", strconv.Atoi, 4)
)

type parser[T any] func(string) (T, error)

func env[T any](name string, parser parser[T], defaultValue T) T {
	value := os.Getenv(name)

	if parsed, err := parser(value); err == nil {
		return parsed
	}

	return defaultValue
}
