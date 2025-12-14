package day09p1

import (
	"strings"
	"testing"

	"aoc/utils"
)

var testInput = `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`

func TestSolve(t *testing.T) {
	tests := []struct {
		input  string
		answer int
	}{
		{testInput, 50},
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
