package day8

import (
	"github.com/luisya22/aoc2022/fileman"
	"testing"
)

func TestDay8ChallengeResultA(t *testing.T) {

	testsMap := []struct {
		name     string
		filePath string
		wants    int
	}{
		{
			name:     "test1",
			filePath: "inputTest.txt",
			wants:    21,
		},
	}

	for _, tt := range testsMap {
		t.Run(tt.name, func(t *testing.T) {
			inputScanner, inputFile := fileman.GetFileRuneBuffer(tt.filePath)
			defer inputFile.Close()

			result := partA(inputScanner)

			if result != tt.wants {
				t.Errorf("got: %v; wants: %v", result, tt.wants)
			}
		})
	}
}

func TestDay8ChallengeResultB(t *testing.T) {

	testsMap := []struct {
		name     string
		filePath string
		wants    int
	}{
		{
			name:     "test1",
			filePath: "inputTest.txt",
			wants:    8,
		},
	}

	for _, tt := range testsMap {
		t.Run(tt.name, func(t *testing.T) {
			inputScanner, inputFile := fileman.GetFileRuneBuffer(tt.filePath)
			defer inputFile.Close()

			result := partB(inputScanner)

			if result != tt.wants {
				t.Errorf("got: %v; wants: %v", result, tt.wants)
			}
		})
	}
}
