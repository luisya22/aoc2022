package day8

func scanTop(treeArray [][]int, y, x int) (bool, int) {
	scenicScore := 0

	for i := y - 1; i >= 0; i-- {
		scenicScore += 1
		if treeArray[i][x] >= treeArray[y][x] {
			return false, scenicScore
		}
	}

	return true, scenicScore
}

func scanRight(treeArray [][]int, y, x int) (bool, int) {
	scenicScore := 0
	for i := x + 1; i <= len(treeArray[y])-1; i++ {
		scenicScore += 1
		if treeArray[y][i] >= treeArray[y][x] {
			return false, scenicScore
		}
	}

	return true, scenicScore
}

func scanBottom(treeArray [][]int, y, x int) (bool, int) {

	scenicScore := 0
	for i := y + 1; i <= len(treeArray)-1; i++ {
		scenicScore += 1
		if treeArray[i][x] >= treeArray[y][x] {
			return false, scenicScore
		}
	}

	return true, scenicScore
}

func scanLeft(treeArray [][]int, y, x int) (bool, int) {
	scenicScore := 0
	for i := x - 1; i >= 0; i-- {
		scenicScore += 1
		if treeArray[y][i] >= treeArray[y][x] {
			return false, 0
		}
	}

	return true, scenicScore
}
