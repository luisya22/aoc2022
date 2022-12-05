package day4

import (
	"github.com/luisya22/aoc2022/utils"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
)

var bCmd = &cobra.Command{
	Use:   "b",
	Short: "Day 4, Problem B",
	Run: func(cmd *cobra.Command, args []string) {

		inputStr := utils.GetFileAsString("cmd/day4/input.txt")
		splitStr := strings.Split(inputStr, "\n")

		result := partB(splitStr)

		log.Println("Result B: ", result)
	},
}

func partB(inputStr []string) int {

	total := 0
	for _, str := range inputStr {
		splitStr := strings.Split(str, ",")

		elf1 := strings.Split(strings.TrimSpace(splitStr[0]), "-")
		elf2 := strings.Split(strings.TrimSpace(splitStr[1]), "-")

		elf1Min, err := strconv.Atoi(elf1[0])
		if err != nil {
			log.Fatal(err)
		}

		elf1Max, err := strconv.Atoi(elf1[1])
		if err != nil {
			log.Fatal(err)
		}

		elf2Min, err := strconv.Atoi(elf2[0])
		if err != nil {
			log.Fatal(err)
		}

		elf2Max, err := strconv.Atoi(elf2[1])
		if err != nil {
			log.Fatal(err)
		}

		if contained(elf1Min, elf2Min, elf2Max) || contained(elf1Max, elf2Min, elf2Max) ||
			contained(elf2Min, elf1Min, elf1Max) || contained(elf2Max, elf1Min, elf1Max) {

			log.Println(elf1, elf2)
			total += 1
		}
	}

	return total
}

func contained(target, min, max int) bool {
	if target >= min && target <= max {
		return true
	}

	return false
}
