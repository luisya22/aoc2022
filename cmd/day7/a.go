package day7

import (
	"bufio"
	"github.com/luisya22/aoc2022/datastruct"
	"github.com/luisya22/aoc2022/fileman"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var aCmd = &cobra.Command{
	Use:   "a",
	Short: "Day 7, Problem A",
	Run: func(cmd *cobra.Command, args []string) {

		scanner, inputFile := fileman.GetFileLineBuffer("cmd/day7/input.txt")
		defer inputFile.Close()

		result := partA(scanner)

		log.Println("Result A: ", result)
	},
}

func partA(inputScanner *bufio.Scanner) int {

	fileSum := 0
	directoryStack := datastruct.Stack[string]{}
	directorySizes := make(map[string]int)

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

	for _, dirSize := range directorySizes {

		if dirSize <= 100000 {
			fileSum += dirSize
		}
	}

	return fileSum
}
