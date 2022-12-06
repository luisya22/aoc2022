package day6

import (
	"bufio"
	"github.com/luisya22/aoc2022/datastruct"
	"github.com/luisya22/aoc2022/fileman"
	"github.com/spf13/cobra"
	"log"
)

var aCmd = &cobra.Command{
	Use:   "a",
	Short: "Day 6, Problem A",
	Run: func(cmd *cobra.Command, args []string) {

		scanner, inputFile := fileman.GetFileRuneBuffer("cmd/day6/input.txt")
		defer inputFile.Close()

		result := partA(scanner)

		log.Println("Result A: ", result)
	},
}

func partA(inputScanner *bufio.Scanner) int {
	markerLen := 4
	firstMarker := 0
	var dataQueue datastruct.Queue[string]

	for inputScanner.Scan() {
		if len(dataQueue.Data) < markerLen {
			dataQueue.Queue(inputScanner.Text())
			firstMarker += 1
			continue
		}

		if markerIsValid(dataQueue.Data, markerLen) {
			break
		}

		dataQueue.Dequeue()
		dataQueue.Queue(inputScanner.Text())
		firstMarker += 1
	}

	return firstMarker
}
