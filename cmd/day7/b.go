package day7

import (
	"bufio"
	"github.com/luisya22/aoc2022/datastruct"
	"github.com/luisya22/aoc2022/fileman"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var bCmd = &cobra.Command{
	Use:   "b",
	Short: "Day 7, Problem B",
	Run: func(cmd *cobra.Command, args []string) {

		scanner, inputFile := fileman.GetFileLineBuffer("cmd/day7/input.txt")
		defer inputFile.Close()

		result := partB(scanner)

		log.Println("Result B: ", result)
	},
}

func partB(inputScanner *bufio.Scanner) int {

	minDeletableDirSize := 0
	directoryStack := datastruct.Stack[string]{}
	directorySizes := make(map[string]int)
	deviceSpace := 70_000_000
	requiredSpace := 30_000_000

	for inputScanner.Scan() {

		terminalLine := strings.Split(inputScanner.Text(), " ")
		if terminalLine[1] == "cd" {
			directoryStack.Push(terminalLine[2])
			size := calculateDirectorySize(terminalLine[2], inputScanner, &directoryStack, directorySizes)
			directorySizes[terminalLine[2]] = size
		} else {
			log.Fatal("I'm not prepared for this")
		}
	}

	neededSpace := requiredSpace - (deviceSpace - directorySizes["/"])

	for _, dirSize := range directorySizes {

		if dirSize >= neededSpace {
			if minDeletableDirSize == 0 {
				minDeletableDirSize = dirSize
			} else if dirSize < minDeletableDirSize {
				minDeletableDirSize = dirSize
			}
		}
	}

	return minDeletableDirSize
}
