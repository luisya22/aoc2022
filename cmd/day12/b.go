package day12

import (
	"github.com/luisya22/aoc2022/datastruct"
	"github.com/luisya22/aoc2022/fileman"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var bCmd = &cobra.Command{
	Use:   "b",
	Short: "Day 12, Problem B",
	Run: func(cmd *cobra.Command, args []string) {

		inputString := fileman.GetFileAsString("cmd/day12/input.txt")

		result := partB(inputString)

		log.Println("Result B: ", result)
	},
}

func partB(inputString string) int {

	//Parse Maze
	world := strings.Split(inputString, "\n")
	seen := [][]bool{}
	startPoint := point{0, 0, 0}
	sFlag := false
	eFlag := false

	for rowIndex, row := range world {

		if !sFlag {
			world[rowIndex] = strings.Replace(world[rowIndex], "S", "a", -1)
			sFlag = true
		}

		if !eFlag {
			zIndex := strings.Index(row, "E")
			if zIndex != -1 {
				startPoint.y = rowIndex
				startPoint.x = zIndex
				eFlag = true
				world[rowIndex] = strings.Replace(world[rowIndex], "E", "z", -1)
			}
		}

		if eFlag {
			break
		}
	}

	for _, row := range world {
		var rowArr []bool
		for range row {
			rowArr = append(rowArr, false)
		}
		seen = append(seen, rowArr)
	}

	distance := bfsDown(world, startPoint, seen)

	return distance
}

func bfsDown(world []string, startPoint point, seen [][]bool) int {

	queue := datastruct.Queue[point]{}

	queue.Queue(startPoint)

	for currPoint := queue.Dequeue(); currPoint != nil; {

		if world[currPoint.y][currPoint.x] == 'a' {
			return currPoint.steps
		}

		// Get Neighbors
		neighbors := getValidNeighborsDown(world, currPoint, seen)
		// Push Neighbors to queue
		for _, neighbor := range neighbors {
			seen[neighbor.y][neighbor.x] = true
			queue.Queue(neighbor)
		}

		currPoint = queue.Dequeue()
	}

	return -1
}

func getValidNeighborsDown(world []string, currPoint *point, seen [][]bool) []point {

	var neighbors []point
	for _, dir := range dirs {
		neighbor := point{x: currPoint.x + dir.x, y: currPoint.y + dir.y}

		if validPoint(world, neighbor) && !seen[neighbor.y][neighbor.x] && isOneLower(world, *currPoint, neighbor) {
			neighbor.steps = currPoint.steps + 1
			neighbors = append(neighbors, neighbor)
		}
	}

	return neighbors
}

func isOneLower(world []string, currPoint point, nextPoint point) bool {
	if world[currPoint.y][currPoint.x]-1 <= world[nextPoint.y][nextPoint.x] {
		return true
	}

	return false
}
