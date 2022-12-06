package day6

import (
	"github.com/luisya22/aoc2022/fileman"
	"testing"
)

func TestDay6ChallengeResultA(t *testing.T) {
	testsMap := []struct {
		name     string
		filepath string
		wants    int
	}{
		{
			name:     "test1",
			filepath: "inputTest1.txt",
			wants:    7,
		},
		{
			name:     "test2",
			filepath: "inputTest2.txt",
			wants:    5,
		},
		{
			name:     "test3",
			filepath: "inputTest3.txt",
			wants:    6,
		},
		{
			name:     "test4",
			filepath: "inputTest4.txt",
			wants:    10,
		},
		{
			name:     "test5",
			filepath: "inputTest5.txt",
			wants:    11,
		},
	}

	for _, tt := range testsMap {
		t.Run(tt.name, func(t *testing.T) {
			scanner, inputFile := fileman.GetFileRuneBuffer(tt.filepath)
			defer inputFile.Close()

			result := partA(scanner)
			if result != tt.wants {
				t.Errorf("got: %v; wants: %v", result, tt.wants)
			}
		})
	}
}

func TestDay6ChallengeResultB(t *testing.T) {
	testsMap := []struct {
		name     string
		filepath string
		wants    int
	}{
		{
			name:     "test1",
			filepath: "inputTest1.txt",
			wants:    19,
		},
		{
			name:     "test2",
			filepath: "inputTest2.txt",
			wants:    23,
		},
		{
			name:     "test3",
			filepath: "inputTest3.txt",
			wants:    23,
		},
		{
			name:     "test4",
			filepath: "inputTest4.txt",
			wants:    29,
		},
		{
			name:     "test5",
			filepath: "inputTest5.txt",
			wants:    26,
		},
	}

	for _, tt := range testsMap {
		t.Run(tt.name, func(t *testing.T) {
			scanner, inputFile := fileman.GetFileRuneBuffer(tt.filepath)
			defer inputFile.Close()

			result := partB(scanner)
			if result != tt.wants {
				t.Errorf("got: %v; wants: %v", result, tt.wants)
			}
		})
	}
}
