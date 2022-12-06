package day6

import (
	"bufio"
	"github.com/luisya22/aoc2022/datastruct"
	"github.com/luisya22/aoc2022/fileman"
	"github.com/spf13/cobra"
	"log"
)

var bCmd = &cobra.Command{
	Use:   "b",
	Short: "Day 6, Problem B",
	Run: func(cmd *cobra.Command, args []string) {

		scanner, inputFile := fileman.GetFileRuneBuffer("cmd/day6/input.txt")
		defer inputFile.Close()

		result := partB(scanner)

		log.Println("Result B: ", result)
	},
}

func partB(inputScanner *bufio.Scanner) int {
	markerLen := 14
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
