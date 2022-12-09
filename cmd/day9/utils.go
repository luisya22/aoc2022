package day9

type position struct {
	x int
	y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func isNotTouching(head, tail position) bool {
	return abs(head.y-tail.y) == 2 || abs(head.x-tail.x) == 2
}

func needsTopRight(headPosition, tailPosition position) bool {
	if headPosition.y-1 > tailPosition.y && headPosition.x > tailPosition.x {
		return true
	}

	if headPosition.x-1 > tailPosition.x && headPosition.y > tailPosition.y {
		return true
	}

	return false
}

func needsBottomLeft(headPosition, tailPosition position) bool {
	if headPosition.y+1 < tailPosition.y && headPosition.x < tailPosition.x {
		return true
	}

	if headPosition.x+1 < tailPosition.x && headPosition.y < tailPosition.y {
		return true
	}

	return false
}

func needsTopLeft(headPosition, tailPosition position) bool {
	if headPosition.y-1 > tailPosition.y && headPosition.x < tailPosition.x {
		return true
	}

	if headPosition.x+1 < tailPosition.x && headPosition.y > tailPosition.y {
		return true
	}

	return false
}

func needsBottomRight(headPosition, tailPosition position) bool {
	if headPosition.y+1 < tailPosition.y && headPosition.x > tailPosition.x {
		return true
	}

	if headPosition.x-1 > tailPosition.x && headPosition.y < tailPosition.y {
		return true
	}

	return false
}

func moveRope(headPosition, tailPosition position) position {

	diffY := headPosition.y - tailPosition.y
	diffX := headPosition.x - tailPosition.x

	if isNotTouching(headPosition, tailPosition) {
		if diffY == 0 {
			if diffX > 0 {
				tailPosition.x += 1
			} else {
				tailPosition.x -= 1
			}
		} else if diffX == 0 {
			if diffY > 0 {
				tailPosition.y += 1
			} else {
				tailPosition.y -= 1
			}
		} else if diffY < 0 && diffX < 0 {
			tailPosition.x -= 1
			tailPosition.y -= 1
		} else if diffY > 0 && diffX < 0 {
			tailPosition.x -= 1
			tailPosition.y += 1
		} else if diffY > 0 && diffX > 0 {
			tailPosition.x += 1
			tailPosition.y += 1
		} else if diffY < 0 && diffX > 0 {
			tailPosition.x += 1
			tailPosition.y -= 1
		}
	}

	return tailPosition
}
