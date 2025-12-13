package day07p1

import (
	"io"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	manifold, err := utils.ReadChars(r)
	if err != nil {
		panic(err)
	}

	// convert chart origin
	utils.RunesToChart(&manifold)
	maxY := len(manifold) - 1
	maxX := len(manifold[0]) - 1
	_ = maxX

	split := 0
	for i:=maxY; i > 0; i-- {
		line := manifold[i]
		for j, char := range line {
			if char != 'S' && char != '|' {
				continue
			}

			current := utils.Point{X: j, Y: i}
			next := current.Add(utils.South)
			nextRune := manifold[next.Y][next.X]
			if nextRune == '^' {
				println("split")
				split++
				left := next.Add(utils.West)
				right := next.Add(utils.East)
				// TODO bounds check
				manifold[left.Y][left.X] = '|'
				manifold[right.Y][right.X] = '|'
			} else {
				println("continue")
				manifold[next.Y][next.X] = '|'
			}
		}
		
	}
	

	utils.PrintChart(manifold)

	return split
}
