/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package day1

import (
	"github.com/luisya22/aoc2022/fileman"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
)

// aCmd represents the a command
var aCmd = &cobra.Command{
	Use:   "a",
	Short: "Day 1, Problem A",
	Run: func(cmd *cobra.Command, args []string) {

		inputStr := fileman.GetFileAsString("cmd/day1/input.txt")
		splitStr := strings.Split(inputStr, "\n\n")

		result := partA(splitStr)

		log.Println("Result 1: ", result)
	},
}

func partA(splitStr []string) int {
	// Find the Elf carrying the most Calories.

	// Get Max Sum (Max Calories)
	max := 0

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

		if sum > max {
			max = sum
		}
	}

	return max
}
