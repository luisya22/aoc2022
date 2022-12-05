package day5

import (
	"github.com/luisya22/aoc2022/datastruct"
	"log"
	"strings"
)

// Helper Functions

func buildStacks(cratesData []string) []datastruct.Stack[string] {
	var stacks []datastruct.Stack[string]
	stackStringData := strings.Replace(cratesData[len(cratesData)-1], " ", "", -1)

	for _, _ = range stackStringData {
		stacks = append(stacks, datastruct.Stack[string]{})
	}
	log.Println("Stacks:", len(stacks))

	for _, crateLine := range datastruct.Reverse(cratesData[:len(cratesData)-1]) {
		for i, s := 1, 0; s < len(stacks); s++ {
			if string(crateLine[i]) != " " {
				stacks[s].Push(string(crateLine[i]))
			}

			i += 4

			if i >= len(crateLine) {
				break
			}
		}
	}

	return stacks
}

func moveCratesOne(stacks []datastruct.Stack[string], amount, originStack, targetStack int) []datastruct.Stack[string] {
	for i := 1; i <= amount; i++ {
		move(stacks, originStack, targetStack)
	}

	return stacks
}

func move(stacks []datastruct.Stack[string], originStack, targetStack int) []datastruct.Stack[string] {
	crate := stacks[originStack].Pop()
	stacks[targetStack].Push(crate)

	return stacks
}

func moveCratesMultiple(stacks []datastruct.Stack[string], amount, originStack, targetStack int) []datastruct.Stack[string] {
	var movedCrates []string

	for i := 1; i <= amount; i++ {
		crate := stacks[originStack].Pop()
		movedCrates = append(movedCrates, crate)
	}

	movedCrates = datastruct.Reverse(movedCrates)

	for _, crate := range movedCrates {
		stacks[targetStack].Push(crate)
	}

	return stacks
}
