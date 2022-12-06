package fileman

import (
	"bufio"
	"log"
	"os"
)

func GetFileAsString(path string) string {
	inputFile, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return string(inputFile)
}

func GetFileLineBuffer(path string) (*bufio.Scanner, *os.File) {
	inputFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	return fileScanner, inputFile
}

func GetFileRuneBuffer(path string) (*bufio.Scanner, *os.File) {
	inputFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanRunes)

	return fileScanner, inputFile
}
