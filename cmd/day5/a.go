package day5

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
	Short: "Day 5, Problem A",
	Run: func(cmd *cobra.Command, args []string) {

		scanner, inputFile := fileman.GetFileLineBuffer("cmd/day5/input.txt")
		defer inputFile.Close()

		result := partA(scanner)

		log.Println("Result A: ", result)
	},
}

func partA(inputScanner *bufio.Scanner) string {
	var cratesData []string

	for inputScanner.Scan() {
		if inputScanner.Text() != "" {
			cratesData = append(cratesData, inputScanner.Text())
		} else {
			break
		}
	}

	var topCrates string

	// Process Crate Data
	stacks := buildStacks(cratesData)
	// Process movements
	for inputScanner.Scan() {

		dataLine := strings.Replace(inputScanner.Text(), "move ", "", -1)
		dataLine = strings.Replace(dataLine, "from ", "", -1)
		dataLine = strings.Replace(dataLine, "to ", "", -1)
		dataStr := strings.Split(dataLine, " ")

		if len(dataStr) != 3 {
			log.Fatalf("wrong movement data format: %v", inputScanner.Text())
		}

		amount, err := strconv.Atoi(dataStr[0])
		if err != nil {
			log.Fatal(err)
		}

		originStack, err := strconv.Atoi(dataStr[1])
		if err != nil {
			log.Fatal(err)
		}

		targetStack, err := strconv.Atoi(dataStr[2])
		if err != nil {
			log.Fatal(err)
		}

		moveCratesOne(stacks, amount, originStack-1, targetStack-1)
	}

	for _, stack := range stacks {
		topCrates += stack.Data[len(stack.Data)-1]
	}

	return topCrates
}
