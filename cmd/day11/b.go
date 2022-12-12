package day11

import (
	"github.com/luisya22/aoc2022/datastruct"
	"github.com/luisya22/aoc2022/fileman"
	"github.com/spf13/cobra"
	"log"
)

var bCmd = &cobra.Command{
	Use:   "b",
	Short: "Day 11, Problem B",
	Run: func(cmd *cobra.Command, args []string) {

		inputString := fileman.GetFileAsString("cmd/day11/input.txt")

		result := partB(inputString)

		log.Println("Result B: ", result)
	},
}

func partB(inputString string) int {

	monkeys := parseMonkeys(inputString)
	rounds := 10_000

	var divisibles []int

	for _, monkey := range monkeys {
		divisibles = append(divisibles, monkey.divisibleBy)
	}

	commonMultiple := lcm(divisibles[0], divisibles[1], divisibles[2:]...)

	for rnd := 0; rnd < rounds; rnd++ {
		for _, m := range monkeys {
			if len(m.items.Data) == 0 {
				continue
			}

			itemsLen := len(m.items.Data)

			for i := 0; i < itemsLen; i++ {
				item := m.items.Dequeue()
				worryLevel := m.operation(m, *item)
				throwMonkeyIndex, itemProcessed := m.testWorryLevel(worryLevel, commonMultiple, false)
				monkeys[throwMonkeyIndex].items.Queue(itemProcessed)
				m.inspected += 1
			}

		}
	}

	monkeys = sortMonkeys(monkeys)

	monkeys = datastruct.Reverse(monkeys)

	for _, monkey := range monkeys {
		log.Println(monkey.inspected, monkey.items.Data)
	}

	return monkeys[0].inspected * monkeys[1].inspected
}
