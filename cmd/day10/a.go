package day10

import (
	"github.com/luisya22/aoc2022/fileman"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
)

var aCmd = &cobra.Command{
	Use:   "a",
	Short: "Day 10, Problem A",
	Run: func(cmd *cobra.Command, args []string) {

		inputString := fileman.GetFileAsString("cmd/day10/input.txt")

		result := partA(inputString)

		log.Println("Result A: ", result)
	},
}

func partA(program string) int {

	registerX := 1
	signalStrengthSum := 0
	cyclesToCalculate := []int{20, 60, 100, 140, 180, 220}
	nextCycleToCalculatePointer := 0
	cyclesStrengthMap := make(map[int]int)
	instructionSet := strings.Split(program, "\n")
	instructionCycle := 0
	instructionIndex := 0
	var actualInstruction []string
	screenWidth := 40
	screenHeigth := 6

	for cycle := 1; cycle <= screenWidth*screenHeigth; cycle++ {

		// Get new instruction if instruction Cycle remaining == 0
		if instructionCycle == 0 {
			actualInstruction = strings.Split(instructionSet[instructionIndex], " ")

			if actualInstruction[0] == "noop" {
				instructionCycle = 1
			} else if actualInstruction[0] == "addx" {
				instructionCycle = 2
			} else {
				log.Fatalf("Syntax error at line %v", instructionIndex)
			}
		}

		// Get Signal Strength
		if cycle == cyclesToCalculate[nextCycleToCalculatePointer] {
			signalStrengthSum += cycle * registerX
			cyclesStrengthMap[cycle] = cycle * registerX

			if nextCycleToCalculatePointer == len(cyclesToCalculate)-1 {
				break
			}

			nextCycleToCalculatePointer++
		}

		// If instructionCycle == 1 execute instruction
		if instructionCycle == 1 {
			if actualInstruction[0] == "addx" {
				value, err := strconv.Atoi(actualInstruction[1])
				if err != nil {
					log.Fatalf("Syntax error at line %v", instructionIndex)
				}
				registerX += value
			}

			instructionIndex += 1
		}

		instructionCycle--
	}

	log.Println(cyclesStrengthMap)

	return signalStrengthSum
}
