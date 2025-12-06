package day03p1

import (
	"io"
	"aoc/utils"
)

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

		topTen := 0
		max := 0
		for _, val := range bank {
			newMax := topTen*10 + val
			if newMax > max {
				max = newMax
			}
			if val > topTen {
				topTen = val
			}
		}

		joltage += max
	}


	return joltage
}
