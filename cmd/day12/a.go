package day12

import (
	"github.com/luisya22/aoc2022/datastruct"
	"github.com/luisya22/aoc2022/fileman"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var aCmd = &cobra.Command{
	Use:   "a",
	Short: "Day 12, Problem A",
	Run: func(cmd *cobra.Command, args []string) {

		inputString := fileman.GetFileAsString("cmd/day12/input.txt")

		result := partA(inputString)

		log.Println("Result A: ", result)
	},
}

func partA(inputString string) int {

	//Parse Maze
	maze := strings.Split(inputString, "\n")
	seen := [][]bool{}
	startPoint := point{0, 0, 0}
	endPoint := point{0, 0, 122}
	sFlag := false
	eFlag := false

	for rowIndex, row := range maze {
		if !sFlag {
			sIndex := strings.Index(row, "S")
			if sIndex != -1 {
				startPoint.y = rowIndex
				startPoint.x = sIndex
				sFlag = true
				maze[rowIndex] = strings.Replace(maze[rowIndex], "S", "a", -1)
			}
		}

		if !eFlag {
			zIndex := strings.Index(row, "E")
			if zIndex != -1 {
				endPoint.y = rowIndex
				endPoint.x = zIndex
				eFlag = true
				maze[rowIndex] = strings.Replace(maze[rowIndex], "E", "z", -1)
			}
		}

		if eFlag && sFlag {
			break
		}
	}

	for _, row := range maze {
		var rowArr []bool
		for range row {
			rowArr = append(rowArr, false)
		}
		seen = append(seen, rowArr)
	}

	distance := bfsUp(maze, startPoint, endPoint, seen)

	return distance
}

func bfsUp(world []string, startPoint point, endPoint point, seen [][]bool) int {

	queue := datastruct.Queue[point]{}

	queue.Queue(startPoint)

	for currPoint := queue.Dequeue(); currPoint != nil; {

		if currPoint.x == endPoint.x && currPoint.y == endPoint.y {
			return currPoint.steps
		}

		// Get Neighbors
		neighbors := getValidNeighborsUp(world, currPoint, seen)
		// Push Neighbors to queue
		for _, neighbor := range neighbors {
			seen[neighbor.y][neighbor.x] = true
			queue.Queue(neighbor)
		}

		currPoint = queue.Dequeue()
	}

	return -1
}

func getValidNeighborsUp(world []string, currPoint *point, seen [][]bool) []point {

	var neighbors []point
	for _, dir := range dirs {
		neighbor := point{x: currPoint.x + dir.x, y: currPoint.y + dir.y}

		if validPoint(world, neighbor) && !seen[neighbor.y][neighbor.x] && isOneHigher(world, *currPoint, neighbor) {
			neighbor.steps = currPoint.steps + 1
			neighbors = append(neighbors, neighbor)
		}
	}

	return neighbors
}

func isOneHigher(world []string, currPoint point, nextPoint point) bool {
	if world[currPoint.y][currPoint.x]+1 >= world[nextPoint.y][nextPoint.x] {
		return true
	}

	return false
}
