package day09p1

import (
	"io"
	"strconv"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	_ = lines

	var corners []utils.Point
	for _, line := range lines {
		coords := strings.Split(line, ",")

		x, err := strconv.Atoi(coords[0])
		utils.Check(err, "")
		y, err := strconv.Atoi(coords[1])
		utils.Check(err, "")

		p := utils.Point{X: x, Y: y}
		corners = append(corners, p)
	}
	
	maxArea := 0
	for i, curent := range corners {
		for j, other := range corners {
			if i == j {
				continue
			}

			// TODO work out which is left and right
			area := utils.CornerArea(curent, other)
			aInt := int(area)
			if aInt > maxArea {
				maxArea = aInt
			}
		}
	}
	return maxArea
}
