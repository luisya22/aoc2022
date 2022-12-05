package day5

import (
	"github.com/luisya22/aoc2022/fileman"
	"strings"
	"testing"
)

func TestDay5ChallengeResult(t *testing.T) {
	type testData struct {
		name     string
		filepath string
		wantsA   string
		wantsB   string
	}

	testsMap := []testData{
		{
			name:     "prodData",
			filepath: "input.txt",
			wantsA:   "RTGWZTHLD",
			wantsB:   "STHGRZZFR",
		},
		{
			name:     "testData",
			filepath: "inputTest.txt",
			wantsA:   "CMZ",
			wantsB:   "MCD",
		},
	}

	if !testing.Short() {
		testDataLarge := testData{
			name:     "largeData",
			filepath: "inputTestLarge.txt",
			wantsA:   "GATHERING",
			wantsB:   "DEVSCHUUR",
		}
		testsMap = append(testsMap, testDataLarge)
	}

	for _, tt := range testsMap {
		t.Run(tt.name, func(t *testing.T) {
			inputStr := fileman.GetFileAsString(tt.filepath)
			splitStr := strings.Split(inputStr, "\n\n")

			result := partA(splitStr)
			if result != tt.wantsA {
				t.Errorf("got: %v; wants: %v", result, tt.wantsA)
			}

			result = partB(splitStr)
			if result != tt.wantsB {
				t.Errorf("got: %v; wants: %v", result, tt.wantsB)
			}
		})
	}
}
