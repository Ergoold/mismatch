package main

import (
	"os"
	"strconv"
)

var (
	BufferSize = env("BUFFER_SIZE", strconv.Atoi, 64)
)

type parser[T any] func(string) (T, error)

func env[T any](name string, parser parser[T], defaultValue T) T {
	value := os.Getenv(name)

	if parsed, err := parser(value); err == nil {
		return parsed
	}

	return defaultValue
}
