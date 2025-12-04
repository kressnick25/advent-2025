package day03p1

import (
	"fmt"
	"io"
	"slices"
	"strings"
	"strconv"

	"aoc/utils"
)

func digitsToDoubleNum(first int, second int) int {
	numStr := strconv.Itoa(first) + strconv.Itoa(second)
	println("joltage is" + numStr)
	num, err := strconv.Atoi(numStr)
	if err != nil {
		panic(err)
	}
	return num
}

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	// bank
	// sort into array of arrays,
	// key is number, value is indexs of that number
	// loop through new structure
	// for all of 9, is there another 9, with an index after current
	// no? repeat

	joltage := 0
	for _, line := range lines {
		println("line is " + line)
		var bank []int
		for _, rune := range line {
			bank = append(bank, int(rune - '0'))
		}

		// matrix where k=the int value
		// value=a slice of all indexes where that k appears
		matrix := make([][]int, 9)
		for i, val := range bank {
			matrix[val] = append(matrix[val], i)
		}

		for i := 0; i < 9; i++ {
			slices.Sort(matrix[i]) // ascending
		}

		var top int
		var nextTop int

		for i := 8; i >= 0; i-- {
			// actual num we are checking is i+1
			if len(matrix[i]) > 0 {
				top = val
				for _, val := range matrix[i] {
				}
			}
		}

		joltage += digitsToDoubleNum(top, nextTop)
	}


	return joltage
}
