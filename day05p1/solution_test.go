package day05p1

import (
	"strings"
	"testing"

	"aoc/utils"
)

var testInput = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

func TestSolve(t *testing.T) {
	tests := []struct {
		input  string
		answer int
	}{
		{testInput, 3},
	}

	if testing.Verbose() {
		utils.Verbose = true
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)

		result := Solve(r).(int)

		if result != test.answer {
			t.Errorf("Expected %d, got %d", test.answer, result)
		}
	}
}
