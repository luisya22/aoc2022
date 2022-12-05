package day2

import (
	"errors"
	"github.com/luisya22/aoc2022/fileman"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var aCmd = &cobra.Command{
	Use:   "a",
	Short: "Day 2, Problem A",
	Run: func(cmd *cobra.Command, args []string) {

		inputStr := fileman.GetFileAsString("cmd/day2/input.txt")
		splitStr := strings.Split(inputStr, "\n")

		result := partA(splitStr)

		log.Println("Result A: ", result)
	},
}

var WrongInput = errors.New("wrong input")

const (
	ROCK     string = "rock"
	PAPER           = "paper"
	SCISSORS        = "scissors"
	NONE            = "none"
)

func partA(inputStr []string) int {

	total := 0

	for _, str := range inputStr {
		play := strings.Split(str, " ")

		elfThrow, err := transformThrow(play[0])
		if err != nil {
			log.Fatal(err.Error())
		}

		playerThrow, err := transformThrow(play[1])
		if err != nil {
			log.Fatal(err.Error())
		}

		tempPlayPoints, err := getGamePoints(elfThrow, playerThrow)
		if err != nil {
			log.Fatal(err.Error())
		}

		tempThrowPoints, err := getThrowPoints(playerThrow)
		if err != nil {
			log.Fatal(err.Error())
		}

		total += tempPlayPoints + tempThrowPoints
	}
	return total
}
