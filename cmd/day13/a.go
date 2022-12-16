package day13

import (
	"github.com/luisya22/aoc2022/fileman"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var aCmd = &cobra.Command{
	Use:   "a",
	Short: "Day 13, Problem A",
	Run: func(cmd *cobra.Command, args []string) {

		inputString := fileman.GetFileAsString("cmd/day13/input.txt")

		result := partA(inputString)

		log.Println("Result A: ", result)
	},
}

const (
	RIGHTORDER int = 1
	WRONGORDER     = 2
	UNRESOLVED     = 3
)

func partA(inputString string) int {

	pairs := strings.Split(inputString, "\n\n")
	var trueIndexes []int

	for i, str := range pairs {
		pair := strings.Split(str, "\n")

		valid := solve(pair[0], pair[1])

		if valid == RIGHTORDER {
			trueIndexes = append(trueIndexes, i+1)
		}
	}

	total := 0
	for _, i := range trueIndexes {
		total += i
	}

	return total
}
