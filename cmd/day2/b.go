package day2

import (
	"github.com/luisya22/aoc2022/utils"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var bCmd = &cobra.Command{
	Use:   "b",
	Short: "Day 2, Problem B",
	Run: func(cmd *cobra.Command, args []string) {

		inputStr := utils.GetFileAsString("cmd/day2/input.txt")
		splitStr := strings.Split(inputStr, "\n")

		result := partB(splitStr)

		log.Println("Result B: ", result)
	},
}

func partB(inputStr []string) int {
	total := 0

	for _, str := range inputStr {
		play := strings.Split(str, " ")

		elfThrow, err := transformThrow(play[0])
		if err != nil {
			log.Println("Here 2")
			log.Fatal(err.Error())
		}

		playerThrow, err := getNeededThrow(elfThrow, play[1])
		if err != nil {
			log.Println("Here 1")
			log.Fatal(err.Error())
		}

		tempPlayPoints, err := getGamePoints(elfThrow, playerThrow)
		if err != nil {
			log.Println("Here 4")
			log.Fatal(err.Error())
		}

		tempThrowPoints, err := getThrowPoints(playerThrow)
		if err != nil {
			log.Println("Here 5")
			log.Fatal(err.Error())
		}

		total += tempPlayPoints + tempThrowPoints
	}

	return total
}

func getNeededThrow(elfThrow, neededResult string) (string, error) {

	if elfThrow == ROCK {
		if neededResult == "Y" {
			return ROCK, nil
		}

		if neededResult == "Z" {
			return PAPER, nil
		}

		if neededResult == "X" {
			return SCISSORS, nil
		}
	}

	if elfThrow == PAPER {
		if neededResult == "Y" {
			return PAPER, nil
		}

		if neededResult == "Z" {
			return SCISSORS, nil
		}

		if neededResult == "X" {
			return ROCK, nil
		}
	}

	if elfThrow == SCISSORS {
		if neededResult == "Y" {
			return SCISSORS, nil
		}

		if neededResult == "Z" {
			return ROCK, nil
		}

		if neededResult == "X" {
			return PAPER, nil
		}
	}

	return NONE, WrongInput
}
