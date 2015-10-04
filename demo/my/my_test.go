package my

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	tests := map[int]uint64{
		1: 1,
		2: 1,
		3: 2,
		4: 3,
	}

	for idx := range tests {
		if Fib(idx) != tests[idx] {
			t.Error(fmt.Sprintf("Expected Fib(%d) == %d", idx, tests[idx]))
		}
	}
}
