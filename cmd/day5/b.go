package day5

import (
	"github.com/luisya22/aoc2022/fileman"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
)

var bCmd = &cobra.Command{
	Use:   "b",
	Short: "Day 5, Problem B",
	Run: func(cmd *cobra.Command, args []string) {

		inputStr := fileman.GetFileAsString("cmd/day5/input.txt")
		splitStr := strings.Split(inputStr, "\n\n")

		result := partB(splitStr)

		log.Println("Result B: ", result)
	},
}

func partB(inputStr []string) string {
	cratesData := strings.Split(inputStr[0], "\n")
	movementData := strings.Split(inputStr[1], "\n")
	var topCrate string

	// Process Crate Data
	stacks := buildStacks(cratesData)
	// Process movements
	for i, dataLine := range movementData {

		dataLine = strings.Replace(dataLine, "move ", "", -1)
		dataLine = strings.Replace(dataLine, "from ", "", -1)
		dataLine = strings.Replace(dataLine, "to ", "", -1)
		dataStr := strings.Split(dataLine, " ")

		if len(dataStr) != 3 {
			log.Fatalf("wrong movement data format on line %v: %v", i+1, dataLine)
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

		moveCratesMultiple(stacks, amount, originStack-1, targetStack-1)
	}

	for _, stack := range stacks {
		topCrate += stack.Data[len(stack.Data)-1]
	}

	return topCrate
}
