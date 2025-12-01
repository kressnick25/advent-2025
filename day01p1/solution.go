package day01p1

import (
	"fmt"
	"io"
	"strconv"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	_ = lines

	const left = 'L'
	const right = 'R'

	dial := 50
	zeroCount := 0

	for _, line := range(lines) {
		instruction := line[0]
		fmt.Printf("line: '%s', ", line)
		
		count, err := strconv.Atoi(line[1:])
		utils.Check(err, "error parsing count")

		count = count % 100

		fmt.Printf("instruction: %c, count: %d\n", instruction, count)
		if instruction == left {
			dial -= count
			if dial < 0 {
				dial += 100
			}
		} else {
			dial += count
			dial = dial % 100
		}

		if dial == 0 {
			zeroCount ++
		}

		fmt.Printf("dial is %d\n", dial)
	}


	return zeroCount
}
