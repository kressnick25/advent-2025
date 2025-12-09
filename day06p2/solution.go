package day06p2

import (
	"io"
	"strconv"
	"strings"

	"aoc/utils"
)

const add = "+"
const multiply = "*"

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	_ = lines

	var sums [][]int
	operators := strings.Fields(lines[len(lines)-1])
	for i := 0; i < len(lines[0]); i++ {
		sums = append(sums, make([]int, 0))
		for _, val := range row {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			sums[i] = append(sums[i], intVal)
		}
	}
	// utils.PrintMatrix(sums)

	total := 0
	for i := 0; i < len(sums[0]); i++ {
		subtotal := sums[0][i]
		for j := 1; j < len(sums); j++ {
			val := sums[j][i]
			switch o := operators[i]; o {
			case add:
				println("add val ", val, " to subtotal ", subtotal)
				subtotal += val
			case multiply:
				println("mult val ", val, " to subtotal ", subtotal)
				subtotal *= val
			}
		}
		println("subtotal is", subtotal)
		total += subtotal
	}

	return total
}
