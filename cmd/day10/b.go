package day10

import (
	"fmt"
	"github.com/luisya22/aoc2022/fileman"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
)

var bCmd = &cobra.Command{
	Use:   "b",
	Short: "Day 10, Problem B",
	Run: func(cmd *cobra.Command, args []string) {

		inputString := fileman.GetFileAsString("cmd/day10/input.txt")

		result := partB(inputString)

		fmt.Println("Result B: ", result)
	},
}

func partB(program string) string {

	instructionSet := strings.Split(program, "\n")
	screenWidth := 40
	screenHeight := 6
	instructionCycle := 0
	instructionIndex := 0
	var actualInstruction []string
	registerX := 1
	frame := ""

	for cycle := 1; cycle <= screenWidth*screenHeight; cycle++ {

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

		// Print pixels
		pixel := drawPixel(cycle, registerX, screenWidth)
		frame += pixel + ""

		if cycle%screenWidth == 0 {
			frame += "\n"
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

	return frame
}

func drawPixel(cycle, registerX, screenWidth int) string {

	if (cycle-1)%screenWidth >= registerX-1 && (cycle-1)%screenWidth <= registerX+1 {
		return "#"
	}

	return "."
}
