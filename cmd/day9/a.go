package day9

import (
	"bufio"
	"github.com/luisya22/aoc2022/fileman"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
)

var aCmd = &cobra.Command{
	Use:   "a",
	Short: "Day 9, Problem A",
	Run: func(cmd *cobra.Command, args []string) {

		scanner, inputFile := fileman.GetFileLineBuffer("cmd/day9/input.txt")
		defer inputFile.Close()

		result := partA(scanner)

		log.Println("Result A: ", result)
	},
}

func partA(inputScanner *bufio.Scanner) int {
	headPosition := position{0, 0}
	tailPosition := position{0, 0}

	positionRecord := make(map[position]struct{})

	positionRecord[tailPosition] = struct{}{}

	for inputScanner.Scan() {
		movement := strings.Split(inputScanner.Text(), " ")

		steps, err := strconv.Atoi(movement[1])
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < steps; i++ {
			// Move Head
			if movement[0] == "U" {
				headPosition.y += 1
			} else if movement[0] == "R" {
				headPosition.x += 1
			} else if movement[0] == "D" {
				headPosition.y -= 1
			} else if movement[0] == "L" {
				headPosition.x -= 1
			} else {
				log.Fatal("Wrong direction")
			}

			tailPosition = moveRope(headPosition, tailPosition)

			// Add Tail Positions to Record
			pos := position{tailPosition.x, tailPosition.y}
			positionRecord[pos] = struct{}{}
		}

	}

	return len(positionRecord)
}
