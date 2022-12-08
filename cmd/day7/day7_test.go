package day7

import (
	"github.com/luisya22/aoc2022/fileman"
	"testing"
)

func TestDay7ChallengeResultA(t *testing.T) {
	testsMap := []struct {
		name     string
		filepath string
		wants    int
	}{
		{
			name:     "test1",
			filepath: "inputTest.txt",
			wants:    95437,
		},
		{
			name:     "test2",
			filepath: "inputTest2.txt",
			wants:    25,
		},
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

func TestDay7ChallengeResultB(t *testing.T) {
	testsMap := []struct {
		name     string
		filepath string
		wants    int
	}{
		{
			name:     "test1",
			filepath: "inputTest.txt",
			wants:    24933642,
		},
		//{
		//	name:     "test2",
		//	filepath: "inputTest2.txt",
		//	wants:    25,
		//},
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
