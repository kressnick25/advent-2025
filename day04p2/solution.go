package day04p2

import (
	"fmt"
	"io"
	"aoc/utils"
)

const empty rune = '.'
const roll rune = '@'

func init2d(size int, innerSize int) *[][]byte {
	chart := make([][]byte, size)
	for i := range chart {
		chart[i] = make([]byte, innerSize)
	}
	for i := range chart {
		for j := range chart[0] {
			chart[i][j] = '.'
		}
	}
	return &chart
}

func isPaper(chart [][]rune, current []int, check []int) bool {
	// x, y
	squareToCheck := []int{current[0] + check[0], current[1] + check[1]}

	if squareToCheck[1] < 0 || 
		squareToCheck[0] < 0 ||
		squareToCheck[1] >= len(chart) || 
		squareToCheck[0] >= len(chart[0]) {
		return false
	}

	switch a := chart[squareToCheck[1]][squareToCheck[0]]; a {
	case empty:
		return false
	case roll:
		return true
	default: 
		panic("unknown square value")
	}
}

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	chart := make([][]rune, len(lines))
	for i, line := range lines {
		chart[i] = make([]rune, len(line))
		for j, c := range line {
			chart[i][j] = c
		}
	}

	adjSquares := [][]int{ // x, y
		{-1, -1}, // top left
		{0, -1}, // above
		{1, -1}, // top right
		{-1, 0}, // left
		{1, 0}, // right
		{-1, 1}, // bottom left
		{0, 1}, // bottom
		{1, 1}, // bottom right
	}
	
	rerun := 1
	total := 0
	for rerun > 0 {
		for y, line := range chart {
			for x := range line {
				paperC := 0
				current := []int {x, y}
				for _, a := range adjSquares {
					if isPaper(chart, current, a) {
						paperC++
					}
				}
				if chart[y][x] == roll {
					if paperC < 4 {
						chart[y][x] = '.'
						total++
						rerun++
					}
				}
			}
		}
		rerun--
	}

	for _, line := range chart {
		for _, val := range line {
			fmt.Printf("%c", val)
		}
		println()
	}

	return total
}
