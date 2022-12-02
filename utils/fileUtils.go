package utils

import (
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
