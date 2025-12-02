package day02p1

import (
	"io"
	"strconv"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	_ = lines
	
	line := lines[0]
	ranges := strings.Split(line, ",")

	total := 0
	for _, r := range(ranges) {
		s := strings.Split(r, "-")

		start, err := strconv.Atoi(s[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}
		
		for i := start; i <= end; i++ {
			runes := strconv.Itoa(i)

			isEvenLen := (len(runes) % 2) == 0
			if !isEvenLen {
				continue
			}
			
			mid := len(runes) / 2

			firstHalf := string(runes[:mid])
			secondHalf := string(runes[mid:])

			if firstHalf == secondHalf {
				total += i
			}	
		}

	}

	return total
}
