package main

import (
	"advent-of-code/resource_provider"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input, _ := resource_provider.ReadAllLines("input.txt")

	partOneAnswer := 0
	partTwoAnswer := 0
	for _, line := range input {
		pair := strings.Split(line, ",")
		sectionA := strings.Split(pair[0], "-")
		sectionB := strings.Split(pair[1], "-")

		sectionAStart, _ := strconv.Atoi(sectionA[0])
		sectionAEnd, _ := strconv.Atoi(sectionA[1])
		sectionBStart, _ := strconv.Atoi(sectionB[0])
		sectionBEnd, _ := strconv.Atoi(sectionB[1])

		if DoesSectionFullyOverlap(sectionAStart, sectionAEnd, sectionBStart, sectionBEnd) {
			partOneAnswer++
		}

		if DoesSectionPartiallyOverlap(sectionAStart, sectionAEnd, sectionBStart, sectionBEnd) {
			partTwoAnswer++
		}
	}

	fmt.Printf("Part One solution: %d\n", partOneAnswer)
	fmt.Printf("Part Two solution: %d\n", partTwoAnswer)
}

func DoesSectionFullyOverlap(sectionAStart, sectionAEnd, sectionBStart, sectionBEnd int) bool {
	if sectionAStart <= sectionBStart && sectionAEnd >= sectionBEnd {
		return true
	}

	if sectionBStart <= sectionAStart && sectionBEnd >= sectionAEnd {
		return true
	}
	return false
}

func DoesSectionPartiallyOverlap(sectionAStart, sectionAEnd, sectionBStart, sectionBEnd int) bool {
	if sectionAStart <= sectionBEnd && sectionBStart <= sectionAEnd {
		return true
	}
	return false
}
