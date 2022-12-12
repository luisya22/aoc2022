package day11

import (
	"github.com/luisya22/aoc2022/datastruct"
	"log"
	"strconv"
	"strings"
)

type monkeyOperation func(mk *monkey, old int) int

type monkey struct {
	items            datastruct.Queue[int]
	operationParts   []string
	operation        monkeyOperation
	divisibleBy      int
	monkeyTrueIndex  int
	monkeyFalseIndex int
	inspected        int
}

func (mk *monkey) testWorryLevel(itemWorry int, commonMultiple int, worryFlag bool) (int, int) {
	var worryLevel int

	if worryFlag {
		worryLevel = itemWorry / 3
	} else {
		worryLevel = itemWorry % commonMultiple
	}

	result := worryLevel % mk.divisibleBy

	if result == 0 {
		return mk.monkeyTrueIndex, worryLevel
	}

	return mk.monkeyFalseIndex, worryLevel
}

func parseMonkeys(inputString string) []*monkey {
	monkeysString := strings.Split(inputString, "\n\n")

	var monkeys []*monkey
	for _, monkeyString := range monkeysString {
		monkeyParts := strings.Split(monkeyString, "\n")
		m := monkey{}

		// Get Items
		itemStr := strings.Replace(monkeyParts[1], "Starting items: ", "", -1)
		itemStr = strings.Replace(itemStr, " ", "", -1)
		items := strings.Split(itemStr, ",")

		for i := 0; i < len(items); i++ {
			itemValue, err := strconv.Atoi(items[i])
			if err != nil {
				log.Fatal(err)
			}

			m.items.Queue(itemValue)
		}

		// Get Operation
		operationStr := strings.Replace(monkeyParts[2], "Operation: new = ", "", -1)
		operationStr = strings.TrimSpace(operationStr)
		m.operationParts = strings.Split(operationStr, " ")
		//operationPart :=

		m.operation = func(mk *monkey, old int) int {
			num := 0

			if m.operationParts[2] == "old" {
				num = old
			} else {
				num = getLastAsNumber(m.operationParts[2])
			}

			symbol := m.operationParts[1]

			switch symbol {
			case "+":
				return sum(old, num)
			case "-":
				return substraction(old, num)
			case "*":
				return multiplication(old, num)
			case "/":
				return division(old, num)
			default:
				log.Fatalf("Operation symbol not recognized: %v -> %v", symbol, m.operationParts)
			}

			return 0
		}
		// Test
		m.divisibleBy = getLastAsNumber(monkeyParts[3])

		// monkeyTrueIndex
		m.monkeyTrueIndex = getLastAsNumber(monkeyParts[4])

		// monkeyFalseIndex
		m.monkeyFalseIndex = getLastAsNumber(monkeyParts[5])

		// Inspected
		m.inspected = 0

		monkeys = append(monkeys, &m)
	}

	return monkeys
}

func sum(item1, item2 int) int {
	return item1 + item2
}

func substraction(item1, item2 int) int {
	return item1 - item2
}

func multiplication(item1, item2 int) int {
	return item1 * item2
}

func division(item1, item2 int) int {
	return item1 / item2
}

func getLastAsNumber(str string) int {
	strArr := strings.Split(str, " ")

	num, err := strconv.Atoi(strArr[len(strArr)-1])
	if err != nil {
		log.Fatal(err)
	}

	return num
}

func sortMonkeys(monkeys []*monkey) []*monkey {
	return quickSort(monkeys, 0, len(monkeys)-1)
}

func quickSort(monkeys []*monkey, low, high int) []*monkey {
	if low < high {
		var p int
		p = partition(monkeys, low, high)
		monkeys = quickSort(monkeys, low, p-1)
		monkeys = quickSort(monkeys, p+1, high)
	}

	return monkeys
}

func partition(monkeys []*monkey, low, high int) int {
	pivot := monkeys[high].inspected
	i := low
	for j := low; j < high; j++ {
		if monkeys[j].inspected < pivot {
			monkeys[i], monkeys[j] = monkeys[j], monkeys[i]
			i++
		}
	}
	monkeys[i], monkeys[high] = monkeys[high], monkeys[i]

	return i
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}
