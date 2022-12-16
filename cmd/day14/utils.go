package day14

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
}

const (
	AIR      int = 1
	ROCK         = 2
	SAND         = 3
	SPAWN        = 4
	LASTPATH     = 5
)

func parse(inputScanner *bufio.Scanner) map[coord]int {

	cave := getRockCoords(inputScanner)

	c := coord{500, 0}
	cave[c] = SPAWN

	return cave
}

func getRockCoords(inputScanner *bufio.Scanner) map[coord]int {
	rocks := make(map[coord]int)

	for inputScanner.Scan() {
		line := inputScanner.Text()
		line = strings.Replace(line, " ", "", -1)
		coords := strings.Split(line, "->")
		previousX := -1
		previousY := -1

		for _, c := range coords {
			coordSplit := strings.Split(c, ",")
			x, err := strconv.Atoi(string(coordSplit[0]))
			if err != nil {
				log.Fatalf("Cannot convert %v to int", string(coordSplit[0]))
			}

			y, err := strconv.Atoi(coordSplit[1])
			if err != nil {
				log.Fatalf("Cannot convert %v to int", coordSplit[1])
			}

			if previousX == -1 && previousY == -1 {
				rock := coord{x, y}
				rocks[rock] = ROCK

				previousX = x
				previousY = y

				continue
			}

			diffX := x - previousX
			diffY := y - previousY

			if diffX != 0 {
				minX := previousX
				maxX := x

				if diffX < 0 {
					minX = x
					maxX = previousX
				}

				for i := minX; i <= maxX; i++ {
					rock := coord{i, y}
					rocks[rock] = ROCK
				}

				previousX = x
			}

			if diffY != 0 {
				minY := previousY
				maxY := y

				if diffY < 0 {
					minY = x
					maxY = previousY
				}

				for i := minY; i <= maxY; i++ {
					rock := coord{x, i}
					rocks[rock] = ROCK
				}

				previousY = y
			}

		}
	}

	return rocks
}

func finalPath(cave map[coord]int, maxY int) (coord, bool) {
	sand := coord{500, 0}
	abyss := false
	for true {

		if sand.y >= maxY {
			abyss = true
			break
		}

		down := coord{sand.x, sand.y + 1}
		downLeft := coord{sand.x - 1, sand.y + 1}
		downRight := coord{sand.x + 1, sand.y + 1}

		_, exists := cave[down]
		if !exists {
			sand = down
			cave[sand] = LASTPATH
			continue
		}

		_, exists = cave[downLeft]
		if !exists {
			sand = downLeft
			cave[sand] = LASTPATH
			continue
		}

		_, exists = cave[downRight]
		if !exists {
			sand = downRight
			cave[sand] = LASTPATH
			continue
		}

		cave[sand] = LASTPATH

		break
	}

	return sand, abyss
}

func simulateNewSand(cave map[coord]int, maxY int) (coord, bool) {
	sand := coord{500, 0}
	abyss := false
	for true {

		if sand.y >= maxY {
			abyss = true
			break
		}

		down := coord{sand.x, sand.y + 1}
		downLeft := coord{sand.x - 1, sand.y + 1}
		downRight := coord{sand.x + 1, sand.y + 1}

		_, exists := cave[down]
		if !exists {
			sand = down
			continue
		}

		_, exists = cave[downLeft]
		if !exists {
			sand = downLeft
			continue
		}

		_, exists = cave[downRight]
		if !exists {
			sand = downRight
			continue
		}

		break
	}

	return sand, abyss
}

func printResult(cave map[coord]int) {

	minX := math.MaxInt
	minY := math.MaxInt
	maxX := math.MinInt
	maxY := math.MinInt

	for o, _ := range cave {
		if o.x < minX {
			minX = o.x
		}

		if o.x > maxX {
			maxX = o.x
		}

		if o.y < minY {
			minY = o.y
		}

		if o.y > maxY {
			maxY = o.y
		}
	}

	for i := minY; i <= maxY; i++ {
		for j := minX; j <= maxX; j++ {
			object, exists := cave[coord{j, i}]

			if !exists {
				fmt.Print(".")
				continue
			}

			switch object {
			case SAND:
				fmt.Print("o")
			case ROCK:
				fmt.Print("#")
			case LASTPATH:
				fmt.Print("~")
			default:
				fmt.Print(object)
			}
		}
		fmt.Print("\n")
	}
}
