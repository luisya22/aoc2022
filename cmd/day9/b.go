package day9

import (
	"bufio"
	"github.com/luisya22/aoc2022/fileman"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
)

var bCmd = &cobra.Command{
	Use:   "b",
	Short: "Day 9, Problem B",
	Run: func(cmd *cobra.Command, args []string) {

		scanner, inputFile := fileman.GetFileLineBuffer("cmd/day9/input.txt")
		defer inputFile.Close()

		result := partB(scanner)

		log.Println("Result B: ", result)
	},
}

func partB(inputScanner *bufio.Scanner) int {
	var rope []position
	for i := 0; i < 10; i++ {
		rope = append(rope, position{0, 0})
	}

	positionRecord := make(map[position]struct{})

	for inputScanner.Scan() {
		movement := strings.Split(inputScanner.Text(), " ")

		steps, err := strconv.Atoi(movement[1])
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < steps; i++ {
			// Move Head
			if movement[0] == "U" {
				rope[0].y += 1
			} else if movement[0] == "R" {
				rope[0].x += 1
			} else if movement[0] == "D" {
				rope[0].y -= 1
			} else if movement[0] == "L" {
				rope[0].x -= 1
			} else {
				log.Fatal("Wrong direction")
			}

			// Move Rope
			for knot := 1; knot < len(rope); knot++ {
				rope[knot] = moveRope(rope[knot-1], rope[knot])
			}

			tail := rope[len(rope)-1]
			pos := position{tail.x, tail.y}
			positionRecord[pos] = struct{}{}
		}

	}

	return len(positionRecord)
}
