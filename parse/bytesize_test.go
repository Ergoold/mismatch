package parse

import (
	"testing"
)

var (
	byteSizeTests = []testCase[int]{
		{"", 0, ErrEmpty},
		{"1", 1, nil},
		{"2.72", 0, ErrSyntax},
		{"1k", 1e3, nil},
		{"0.5K", 500, nil},
		{"1.5MB", 1_500_000, nil},
		{"0.1G", 1e8, nil},
		{"0.02TB", 2e10, nil},
		{"3KiB", (1 << 10) * 3, nil},
		{"0.1Mi", (1 << 20) / 10, nil},
		{"0.x1GiB", 0, ErrSyntax},
		{"TiB", 0, ErrSyntax},
		{"16E", 0, ErrSyntax},
	}
)

func TestByteSize(t *testing.T) {
	runTestCases(t, "ByteSize", ByteSize, byteSizeTests)
}
