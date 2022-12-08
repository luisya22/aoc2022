package day7

import (
	"bufio"
	"fmt"
	"github.com/luisya22/aoc2022/datastruct"
	"strconv"
	"strings"
)

func calculateDirectorySize(currentDirectory string, inputScanner *bufio.Scanner, directoryStack *datastruct.Stack[string], directorySizes map[string]int) int {
	directorySize := 0
	for inputScanner.Scan() {
		terminalLine := inputScanner.Text()

		if terminalLine == "$ cd .." {
			stackDir := *directoryStack.Peek()
			currDirSplit := strings.Split(currentDirectory, "/")

			if stackDir == currDirSplit[len(currDirSplit)-1] {
				directoryStack.Pop()
				break
			}

			directoryStack.Pop()
		} else if strings.HasPrefix(terminalLine, "$ cd") {
			splitCommand := strings.Split(inputScanner.Text(), " ")
			directoryStack.Push(splitCommand[2])
		} else if strings.HasPrefix(terminalLine, "$ ls") {
			dir := ""

			for _, aDir := range directoryStack.Data {
				if dir == "/" || aDir == "/" {
					dir += aDir
				} else {
					dir = fmt.Sprintf("%v/%v", dir, aDir)
				}
			}

			size := calculateDirectorySize(dir, inputScanner, directoryStack, directorySizes)
			directorySize += size

			continue
		} else if strings.HasPrefix(terminalLine, "dir") {
			continue
		} else {
			splitLine := strings.Split(inputScanner.Text(), " ")
			size, err := strconv.Atoi(splitLine[0])
			if err != nil {
				continue
			}

			directorySize += size
		}
	}

	directorySizes[currentDirectory] = directorySize

	return directorySize
}
