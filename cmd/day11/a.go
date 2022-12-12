package day11

import (
	"github.com/luisya22/aoc2022/datastruct"
	"github.com/luisya22/aoc2022/fileman"
	"github.com/spf13/cobra"
	"log"
)

var aCmd = &cobra.Command{
	Use:   "a",
	Short: "Day 11, Problem A",
	Run: func(cmd *cobra.Command, args []string) {

		inputString := fileman.GetFileAsString("cmd/day11/input.txt")

		result := partA(inputString)

		log.Println("Result A: ", result)
	},
}

func partA(inputString string) int {

	monkeys := parseMonkeys(inputString)
	rounds := 20

	for rnd := 0; rnd < rounds; rnd++ {
		for _, m := range monkeys {
			if len(m.items.Data) == 0 {
				continue
			}

			itemsLen := len(m.items.Data)

			for i := 0; i < itemsLen; i++ {
				item := m.items.Dequeue()
				worryLevel := m.operation(m, *item)
				throwMonkeyIndex, itemProcessed := m.testWorryLevel(worryLevel, 1, true)
				monkeys[throwMonkeyIndex].items.Queue(itemProcessed)
				m.inspected += 1
			}

		}
	}

	monkeys = sortMonkeys(monkeys)

	monkeys = datastruct.Reverse(monkeys)

	return monkeys[0].inspected * monkeys[1].inspected
}
