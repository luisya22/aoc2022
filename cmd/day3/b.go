package day3

import (
	"github.com/luisya22/aoc2022/fileman"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var bCmd = &cobra.Command{
	Use:   "b",
	Short: "Day 3, Problem B",
	Run: func(cmd *cobra.Command, args []string) {

		inputStr := fileman.GetFileAsString("cmd/day3/input.txt")
		splitStr := strings.Split(inputStr, "\n")

		result := partB(splitStr)

		log.Println("Result B: ", result)
	},
}

func partB(inputStr []string) int {
	total := 0

	for i := 0; i < len(inputStr); i += 3 {
		if i+1 >= len(inputStr) || i == len(inputStr) {
			break
		}
		elfSack1 := inputStr[i]
		elfSack2 := inputStr[i+1]
		elfSack3 := ""

		if i+2 < len(inputStr) {
			elfSack3 = inputStr[i+2]
		}

		for _, char := range elfSack1 {
			if strings.Contains(elfSack2, string(char)) && strings.Contains(elfSack3, string(char)) {
				priority, err := getPriority(char)
				if err != nil {
					log.Fatal(err)
				}

				total += priority
				break
			}
		}
	}

	return total
}
