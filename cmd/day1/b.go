/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package day1

import (
	"github.com/luisya22/aoc2022/fileman"
	"log"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// bCmd represents the b command
var bCmd = &cobra.Command{
	Use:   "b",
	Short: "Day 1, Problem B",
	Run: func(cmd *cobra.Command, args []string) {
		inputStr := fileman.GetFileAsString("cmd/day1/input.txt")
		splitStr := strings.Split(inputStr, "\n\n")

		result := partB(splitStr)

		log.Println("Result 1: ", result)
	},
}

func partB(splitStr []string) int {
	// Find the top three Elf carrying the most Calories.

	// Get Max Sum (Max Calories)
	top1 := 0
	top2 := 0
	top3 := 0

	for _, s := range splitStr {
		inv := strings.Split(s, "\n")
		sum := 0
		for _, item := range inv {
			calories, err := strconv.Atoi(item)
			if err != nil {
				log.Fatal(err)
			}

			sum += calories
		}

		if sum > top1 {
			top3 = top2
			top2 = top1
			top1 = sum
		} else if sum > top2 {
			top3 = top2
			top2 = sum
		} else if sum > top3 {
			top3 = sum
		}
	}

	top3Sum := top1 + top2 + top3

	return top3Sum
}
