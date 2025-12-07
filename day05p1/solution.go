package day05p1

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	_ = lines

	pastMiddle := false
	var ranges [][]int
	var ingredients []int
	for _, line := range lines {
		if len(line) == 0 {
			pastMiddle = true
			continue
		}
		
		if !pastMiddle {
			s := strings.Split(line, "-")
			start, _ := strconv.Atoi(s[0])
			end, _ := strconv.Atoi(s[1])
			ranges = append(ranges, []int{start, end})
		} else {
			val, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			ingredients = append(ingredients, val)
		}
	}
	fmt.Printf("%v\n", ranges)
	fmt.Printf("%v\n", ingredients)


	fresh := 0
	for _, ing := range ingredients {
		for _, r := range ranges {
			if ing >= r[0] && ing <= r[1] {
				fresh++
				fmt.Printf("%d is in range %v\n", ing, r)
				break
			}
		}
	}

	return fresh
}
