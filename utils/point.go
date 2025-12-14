package utils

import (
	"math"
	"slices"
)

// 2-D Point
type Point struct {
	X int
	Y int
}

func (p Point) xIdx() int {
	return abs(p.X)
}

func (p Point) yIdx() int {
	return abs(p.Y)
}

// transform a matrix of runes read sequentially,
// so that the origin (0, 0) is the bottom left of the matrix
func RunesToChart(r *[][]rune) {
	slices.Reverse(*r)
}

// not currently thread safe, todo improve
func PrintChart(r [][]rune) {
	slices.Reverse(r)
	for _, line := range r {
		for _, char := range line {
			print(string(char))
		}
		println()
	}
	slices.Reverse(r)
}


// +Y direction
var North = Point{0, 1}

// -Y direction
var South = Point{0, -1}

// +X direction
var East = Point{1, 0}

// -X direction
var West = Point{-1, 0}

// Slice of cardinal directions
var Directions = []Point{North, South, East, West}

// Return the sum of two points
func (p Point) Add(p2 Point) Point {
	r := Point{p.X + p2.X, p.Y + p2.Y}

	return r
}

// Multiply point by a scalar
func (p Point) Scale(factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

// Turn point right 90 degrees
func (p Point) Right() Point {
	return Point{p.Y, -p.X}
}

// Turn point left 90 degrees
func (p Point) Left() Point {
	return Point{-p.Y, p.X}
}

// Return manhattan distance (abs(x)+abs(y))
func (p Point) Manhattan() int {
	return abs(p.X) + abs(p.Y)
}

// return diagonal distance bewteen two points
func Pythagoras(a Point, b Point) float64 {
	t := math.Pow(float64(b.X - a.X), 2) + math.Pow(float64(b.Y - a.Y), 2)
	return math.Sqrt(t)
}

// return of a square between two corners
func CornerArea(a Point, b Point) float64 {
	d := Pythagoras(a, b)
	return math.Pow(d, 2) / 2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
