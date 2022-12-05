package day4

import (
	"github.com/luisya22/aoc2022/utils"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
)

var aCmd = &cobra.Command{
	Use:   "a",
	Short: "Day 4, Problem A",
	Run: func(cmd *cobra.Command, args []string) {

		inputStr := utils.GetFileAsString("cmd/day4/input.txt")
		splitStr := strings.Split(inputStr, "\n")

		result := partA(splitStr)

		log.Println("Result A: ", result)
	},
}

func partA(inputStr []string) int {
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

		if elf1Min >= elf2Min && elf1Max <= elf2Max {
			total += 1
		} else if elf2Min >= elf1Min && elf2Max <= elf1Max {
			total += 1
		}
	}
	return total
}
