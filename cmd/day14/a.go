package day14

import (
	"bufio"
	"github.com/luisya22/aoc2022/fileman"
	"github.com/spf13/cobra"
	"log"
)

var aCmd = &cobra.Command{
	Use:   "a",
	Short: "Day 14, Problem A",
	Run: func(cmd *cobra.Command, args []string) {

		scanner, inputFile := fileman.GetFileLineBuffer("cmd/day14/input.txt")
		defer inputFile.Close()

		result := partA(scanner)

		log.Println("Result A: ", result)
	},
}

func partA(inputScanner *bufio.Scanner) int {
	cave := parse(inputScanner)

	maxY := 0
	for rock, _ := range cave {
		if rock.y > maxY {
			maxY = rock.y
		}
	}

	sandCount := 0
	lastCoord := coord{-1, -1}
	for lastCoord.y < maxY {
		lastCoord, fellAbyss := simulateNewSand(cave, maxY)

		if fellAbyss {
			break
		}

		cave[lastCoord] = SAND
		sandCount++
	}

	finalPath(cave, maxY)

	printResult(cave)

	return sandCount
}
