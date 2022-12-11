package day10

import (
	"github.com/luisya22/aoc2022/fileman"
	"testing"
)

func TestDay10ChallengeResultA(t *testing.T) {
	testsMap := []struct {
		name     string
		filepath string
		wants    int
	}{
		{
			name:     "test 1",
			filepath: "inputTest.txt",
			wants:    13140,
		},
	}

	for _, tt := range testsMap {
		t.Run(tt.name, func(t *testing.T) {
			inputString := fileman.GetFileAsString(tt.filepath)

			result := partA(inputString)

			if result != tt.wants {
				t.Errorf("got: %v; wants: %v", result, tt.wants)
			}
		})
	}
}

func TestDay10ChallengeResultB(t *testing.T) {
	testsMap := []struct {
		name     string
		filepath string
		wants    string
	}{
		//{
		//	name:     "test 1",
		//	filepath: "inputTest.txt",
		//	wants:    "##..##..##..##..##..##..##..##..##..##..\n###...###...###...###...###...###...###.\n####....####....####....####....####....\n#####.....#####.....#####.....#####.....\n######......######......######......####\n#######.......#######.......#######.....",
		//},
	}
	//TODO: Check wants string
	for _, tt := range testsMap {
		t.Run(tt.name, func(t *testing.T) {
			inputString := fileman.GetFileAsString(tt.filepath)

			result := partB(inputString)

			if result != tt.wants {
				t.Errorf("got: %v; wants: %v", result, tt.wants)
			}
		})
	}
}
