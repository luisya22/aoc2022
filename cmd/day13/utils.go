package day13

import (
	"github.com/luisya22/aoc2022/datastruct"
	"log"
	"strconv"
	"strings"
)

func solve(packet1 string, packet2 string) int {
	if packet1[0] == '[' {
		packet1 = packet1[1 : len(packet1)-1]
	}

	if packet2[0] == '[' {
		packet2 = packet2[1 : len(packet2)-1]
	}

	queue1 := parsePacket(packet1)
	queue2 := parsePacket(packet2)

	for data1, data2 := queue1.Dequeue(), queue2.Dequeue(); ; data1, data2 = queue1.Dequeue(), queue2.Dequeue() {

		if data1 == nil && data2 != nil {
			return RIGHTORDER
		} else if data1 != nil && data2 == nil {
			return WRONGORDER
		} else if data1 == nil && data2 == nil {
			break
		}

		d1 := *data1
		d2 := *data2

		if d1[0] == '[' || d2[0] == '[' {
			valid := solve(d1, d2)
			//TODO: WORK WITH THIS

			if valid == RIGHTORDER || valid == WRONGORDER {
				return valid
			}
		} else {
			num1, err := strconv.Atoi(d1)
			if err != nil {
				log.Fatalf("Cannot convert %v to int", d1)
			}

			num2, err := strconv.Atoi(d2)
			if err != nil {
				log.Fatalf("Cannot convert %v to int", d2)
			}

			if num2 < num1 {
				return WRONGORDER
			} else if num2 > num1 {
				return RIGHTORDER
			}

		}
	}

	return UNRESOLVED
}

func parsePacket(p string) datastruct.Queue[string] {
	queue := datastruct.Queue[string]{}

	for i := 0; i < len(p); {
		if p[i] == '[' {
			index := findCloseBracket(p[i:])
			queue.Queue(p[i : i+index+1])
			i += index + 2
		} else {
			index := strings.Index(p[i:], ",")

			if index != -1 {
				queue.Queue(p[i : i+index])
				i += index + 1
			} else {
				queue.Queue(p[i:])
				i += 1
			}
		}
	}
	return queue
}

func findCloseBracket(p string) int {
	openBrackets := 0
	for i, v := range p {
		if v == '[' {
			openBrackets += 1
			continue
		}

		if v == ']' {
			if openBrackets == 1 {
				return i
			} else {
				openBrackets -= 1
			}
		}
	}

	return -1
}
