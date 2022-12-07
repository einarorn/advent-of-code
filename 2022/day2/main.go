package main

import (
	"advent-of-code/resource_provider"
	"fmt"
)

/*
A = X Rock 1
B = Y Paper 2
C = Z Scissors 3
*/

func main() {
	input, _ := resource_provider.ReadAllLines("input.txt")

	scoreChartPart1 := map[string]int{
		"A X": 1 + 3,
		"A Y": 2 + 6,
		"A Z": 3 + 0,
		"B X": 1 + 0,
		"B Y": 2 + 3,
		"B Z": 3 + 6,
		"C X": 1 + 6,
		"C Y": 2 + 0,
		"C Z": 3 + 3,
	}

	scoreChartPart2 := map[string]int{
		"A X": 3 + 0,
		"A Y": 1 + 3,
		"A Z": 2 + 6,
		"B X": 1 + 0,
		"B Y": 2 + 3,
		"B Z": 3 + 6,
		"C X": 2 + 0,
		"C Y": 3 + 3,
		"C Z": 1 + 6,
	}

	totalScorePart1 := 0
	for _, line := range input {
		totalScorePart1 += scoreChartPart1[line]
	}

	fmt.Printf("Total score part1: %d\n", totalScorePart1)

	totalScorePart2 := 0
	for _, line := range input {
		totalScorePart2 += scoreChartPart2[line]
	}

	fmt.Printf("Total score part1: %d\n", totalScorePart2)
}
