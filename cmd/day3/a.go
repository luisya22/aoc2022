package day3

import (
	"github.com/luisya22/aoc2022/utils"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var aCmd = &cobra.Command{
	Use:   "a",
	Short: "Day 3, Problem A",
	Run: func(cmd *cobra.Command, args []string) {

		inputStr := utils.GetFileAsString("cmd/day3/input.txt")
		splitStr := strings.Split(inputStr, "\n")

		result := partA(splitStr)

		log.Println("Result A: ", result)
	},
}

func partA(inputStr []string) int {
	total := 0
	for _, str := range inputStr {
		compartment1 := str[0 : len(str)/2]
		compartment2 := str[len(str)/2:]

		for _, char := range compartment1 {
			if strings.Contains(compartment2, string(char)) {
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
