package day12

type point struct {
	x     int
	y     int
	steps int
}
type dir struct {
	x int
	y int
}

var top = dir{0, 1}
var right = dir{1, 0}
var bottom = dir{0, -1}
var left = dir{-1, 0}
var dirs = []dir{top, right, left, bottom}

func validPoint(world []string, point point) bool {
	if point.x < 0 || point.x >= len(world[0]) ||
		point.y < 0 || point.y >= len(world) {

		return false
	}

	return true
}
