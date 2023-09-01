package simulation

import (
	"log"
	"testing"
)

func TestState(t *testing.T) {
	initState := [][][]int{
		{
			{1, 0},
			{1, 1},
			{1, 2},
		},
		{
			{0, 1},
			{1, 1},
			{2, 1},
		},
	}

	expectedState := [][][]int{
		{
			{0, 1, 0},
			{0, 1, 0},
			{0, 1, 0},
		},
		{
			{0, 0, 0},
			{1, 1, 1},
			{0, 0, 0},
		},
	}

	for testId := 0; testId < len(initState); testId++ {
		canvas := newCanvas(3, 3)
		canvas.setState(initState[testId])
		resultCanvas := conwayStep(canvas)

		for row := 0; row < 3; row++ {
			for col := 0; col < 3; col++ {
				if resultCanvas[row][col] != expectedState[testId][row][col] {
					log.Fatalf(
						"Invalid state: %v, expected: %v",
						resultCanvas,
						expectedState[testId],
					)
				}
			}
		}
	}

}
