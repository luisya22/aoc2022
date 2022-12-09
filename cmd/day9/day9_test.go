package day9

import (
	"github.com/luisya22/aoc2022/fileman"
	"testing"
)

func TestMoveRope(t *testing.T) {
	testMaps := []struct {
		name         string
		headPosition position
		tailPosition position
		wants        position
	}{
		{
			name:         "Move Top",
			headPosition: position{0, 2},
			tailPosition: position{0, 0},
			wants:        position{0, 1},
		},
		{
			name:         "Move TopRight Far Top",
			headPosition: position{1, 2},
			tailPosition: position{0, 0},
			wants:        position{1, 1},
		},
		{
			name:         "Move TopRight Far Right",
			headPosition: position{2, 1},
			tailPosition: position{0, 0},
			wants:        position{1, 1},
		},
		{
			name:         "Move Right",
			headPosition: position{2, 0},
			tailPosition: position{0, 0},
			wants:        position{1, 0},
		},
		{
			name:         "Move BottomRight Far Right",
			headPosition: position{2, -1},
			tailPosition: position{0, 0},
			wants:        position{1, -1},
		},
		{
			name:         "Move BottomRight Far Bottom",
			headPosition: position{1, -2},
			tailPosition: position{0, 0},
			wants:        position{1, -1},
		},
		{
			name:         "Move Bottom",
			headPosition: position{0, -2},
			tailPosition: position{0, 0},
			wants:        position{0, -1},
		},
		{
			name:         "Move BottomLeft Far Bottom",
			headPosition: position{-1, -2},
			tailPosition: position{0, 0},
			wants:        position{-1, -1},
		},
		{
			name:         "Move BottomLeft Far Left",
			headPosition: position{-2, -1},
			tailPosition: position{0, 0},
			wants:        position{-1, -1},
		},
		{
			name:         "Move Left",
			headPosition: position{-2, 0},
			tailPosition: position{0, 0},
			wants:        position{-1, 0},
		},
		{
			name:         "Move TopLeft Far Left",
			headPosition: position{-2, 1},
			tailPosition: position{0, 0},
			wants:        position{-1, 1},
		},
		{
			name:         "Move TopLeft Far Top",
			headPosition: position{-1, 2},
			tailPosition: position{0, 0},
			wants:        position{-1, 1},
		},
	}

	for _, tt := range testMaps {
		t.Run(tt.name, func(t *testing.T) {

			resultPosition := moveRope(tt.headPosition, tt.tailPosition)

			if resultPosition.y != tt.wants.y || resultPosition.x != tt.wants.x {
				t.Errorf("got: %v; wants: %v", resultPosition, tt.wants)
			}
		})

	}
}

func TestDay9ChallengeResultA(t *testing.T) {
	testsMap := []struct {
		name     string
		filepath string
		wants    int
	}{
		{
			name:     "test 1",
			filepath: "inputTest.txt",
			wants:    13,
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

func TestDay9ChallengeResultB(t *testing.T) {
	testsMap := []struct {
		name     string
		filepath string
		wants    int
	}{
		{
			name:     "test 1",
			filepath: "inputTest.txt",
			wants:    1,
		},
		{
			name:     "test 2",
			filepath: "inputTest2.txt",
			wants:    36,
		},
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
