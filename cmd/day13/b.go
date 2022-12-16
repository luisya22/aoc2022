package day13

import (
	"github.com/luisya22/aoc2022/fileman"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var bCmd = &cobra.Command{
	Use:   "b",
	Short: "Day 13, Problem B",
	Run: func(cmd *cobra.Command, args []string) {

		inputString := fileman.GetFileAsString("cmd/day13/input.txt")

		result := partB(inputString)

		log.Println("Result B: ", result)
	},
}

func partB(inputString string) int {

	formattedStr := strings.Replace(inputString, "\n\n", "\n", -1)
	packets := strings.Split(formattedStr, "\n")
	packets = append(packets, "[[2]]")
	packets = append(packets, "[[6]]")
	qs(packets, 0, len(packets)-1)

	sum := 1
	for i, p := range packets {
		if p == "[[2]]" || p == "[[6]]" {
			sum += sum * i
		}
	}

	log.Println(packets)

	return sum
}

func qs(arr []string, lo, hi int) []string {
	if lo >= hi {
		return arr
	}
	arr, pivotIdx := partition(arr, lo, hi)

	arr = qs(arr, lo, pivotIdx-1)
	arr = qs(arr, pivotIdx+1, hi)

	return arr
}

func partition(arr []string, lo, hi int) ([]string, int) {
	pivot := strings.Clone(arr[hi])
	idx := lo - 1

	for i := lo; i < hi; i++ {
		copy1 := strings.Clone(arr[i])
		copy2 := strings.Clone(pivot)

		if solve(copy1, copy2) == RIGHTORDER {
			idx++
			tmp := strings.Clone(arr[i])
			arr[i] = strings.Clone(arr[idx])
			arr[idx] = strings.Clone(tmp)
		}
	}

	idx++

	arr[hi] = strings.Clone(arr[idx])
	arr[idx] = strings.Clone(pivot)

	return arr, idx
}
