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
	prev := dial
	zeroCount := 0
	

	for _, line := range(lines) {
		instruction := line[0]
		fmt.Printf("line: '%s', ", line)
		
		count, err := strconv.Atoi(line[1:])
		utils.Check(err, "error parsing count")

		zeroCount += count / 100
		count = count % 100

		fmt.Printf("instruction: %c, count: %d\n", instruction, count)
		if instruction == left {
			dial -= count
			if dial < 0 {
				if prev != 0 {
					zeroCount ++
				}
				dial += 100
				println("zero")
			}
		} else {
			dial += count
			if dial > 100 {
				if prev != 0 {
					zeroCount ++ 
				}
				println("zero")
			}
			dial = dial % 100
		}

		if dial == 0 && prev != 0 {
			zeroCount ++
			println("zero")
		}

		fmt.Printf("dial is %d\n", dial)
		prev = dial
	}


	return zeroCount
}
