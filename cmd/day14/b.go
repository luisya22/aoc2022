package day14

import (
	"bufio"
	"github.com/luisya22/aoc2022/fileman"
	"github.com/spf13/cobra"
	"log"
)

var bCmd = &cobra.Command{
	Use:   "b",
	Short: "Day 14, Problem B",
	Run: func(cmd *cobra.Command, args []string) {

		scanner, inputFile := fileman.GetFileLineBuffer("cmd/day14/input.txt")
		defer inputFile.Close()

		result := partB(scanner)

		log.Println("Result B: ", result)
	},
}

func partB(inputScanner *bufio.Scanner) int {
	return 0
}
