package day5

import (
	"github.com/luisya22/aoc2022/fileman"
	"testing"
)

func TestDay5ChallengeResultA(t *testing.T) {
	type testData struct {
		name     string
		filepath string
		wants    string
	}

	testsMap := []testData{
		{
			name:     "prodData",
			filepath: "input.txt",
			wants:    "RTGWZTHLD",
		},
		{
			name:     "testData",
			filepath: "inputTest.txt",
			wants:    "CMZ",
		},
	}

	if !testing.Short() {
		testDataLarge := testData{
			name:     "largeData",
			filepath: "inputTestLarge.txt",
			wants:    "GATHERING",
		}
		testsMap = append(testsMap, testDataLarge)
	}

	for _, tt := range testsMap {
		t.Run(tt.name, func(t *testing.T) {
			scanner, inputFile := fileman.GetFileLineBuffer(tt.filepath)
			defer inputFile.Close()

			result := partA(scanner)
			if result != tt.wants {
				t.Errorf("got: %v; wants: %v", result, tt.wants)
			}

		})
	}
}

func TestDay5ChallengeResultB(t *testing.T) {
	type testData struct {
		name     string
		filepath string
		wants    string
	}

	testsMap := []testData{
		{
			name:     "prodData",
			filepath: "input.txt",
			wants:    "STHGRZZFR",
		},
		{
			name:     "testData",
			filepath: "inputTest.txt",
			wants:    "MCD",
		},
	}

	if !testing.Short() {
		testDataLarge := testData{
			name:     "largeData",
			filepath: "inputTestLarge.txt",
			wants:    "DEVSCHUUR",
		}
		testsMap = append(testsMap, testDataLarge)
	}

	for _, tt := range testsMap {
		t.Run(tt.name, func(t *testing.T) {
			scanner, inputFile := fileman.GetFileLineBuffer(tt.filepath)
			defer inputFile.Close()

			result := partB(scanner)
			if result != tt.wants {
				t.Errorf("got: %v; wants: %v", result, tt.wants)
			}
		})
	}
}
