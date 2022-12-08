package day8

import (
	"bufio"
	"github.com/luisya22/aoc2022/fileman"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

var aCmd = &cobra.Command{
	Use:   "a",
	Short: "Day 8, Problem A",
	Run: func(cmd *cobra.Command, args []string) {

		inputScanner, inputFile := fileman.GetFileRuneBuffer("cmd/day8/inputTest.txt")
		defer inputFile.Close()

		result := partA(inputScanner)

		log.Println("Result A: ", result)
	},
}

func partA(inputScanner *bufio.Scanner) int {

	visibleTreeSum := 0
	treeArray := make([][]int, 0)

	tempRow := make([]int, 0)
	for inputScanner.Scan() {
		if inputScanner.Text() == "\n" {
			treeArray = append(treeArray, tempRow)
			tempRow = make([]int, 0)
			continue
		}

		treeHeight, err := strconv.Atoi(inputScanner.Text())
		if err != nil {
			log.Fatalf("%v is not a number", treeHeight)
		}

		tempRow = append(tempRow, treeHeight)
	}

	treeArray = append(treeArray, tempRow)

	for y := 1; y <= len(treeArray[0])-2; y++ {
		for x := 1; x <= len(treeArray[:][0])-2; x++ {

			top, _ := scanTop(treeArray, y, x)
			right, _ := scanRight(treeArray, y, x)
			bottom, _ := scanBottom(treeArray, y, x)
			left, _ := scanLeft(treeArray, y, x)

			if top || right || bottom || left { // Left

				visibleTreeSum += 1
				log.Printf("Yep %v,%v -> %v", y, x, treeArray[y][x])
			} else {
				log.Printf("Nop %v,%v -> %v", y, x, treeArray[y][x])
			}
		}
	}

	visibleTreeSum += (len(treeArray[0]) * 2) + (len(treeArray[:][0]) * 2) - 4 // 4 eliminates duplicates

	return visibleTreeSum
}
