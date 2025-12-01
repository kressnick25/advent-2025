package day01p1

import (
	"strings"
	"testing"

	"aoc/utils"
)

var testInput = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

func TestSolve(t *testing.T) {
	tests := []struct {
		input  string
		answer int
	}{
		{testInput, 6},
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
