package day8

import (
	"bufio"
	"github.com/luisya22/aoc2022/fileman"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

var bCmd = &cobra.Command{
	Use:   "b",
	Short: "Day 8, Problem B",
	Run: func(cmd *cobra.Command, args []string) {

		scanner, inputFile := fileman.GetFileRuneBuffer("cmd/day8/input.txt")
		defer inputFile.Close()

		result := partB(scanner)

		log.Println("Result B: ", result)
	},
}

func partB(inputScanner *bufio.Scanner) int {

	maxTreeScenicScore := 0
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

			_, topScore := scanTop(treeArray, y, x)
			_, rightScore := scanRight(treeArray, y, x)
			_, bottomScore := scanBottom(treeArray, y, x)
			_, leftScore := scanLeft(treeArray, y, x)

			scenicScore := topScore * rightScore * bottomScore * leftScore

			if scenicScore > maxTreeScenicScore {
				maxTreeScenicScore = scenicScore
			}

		}
	}

	return maxTreeScenicScore
}
