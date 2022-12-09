package main

import (
	"advent-of-code/resource_provider"
	"fmt"
	"strings"
)

func main() {
	input, _ := resource_provider.ReadAllLines("input.txt")
	var typePriority = make(map[string]int, 52)

	asciiShift := 97
	for i := 0; i < 52; i++ {
		typePriority[string(i+asciiShift)] = i + 1
		if i == 25 {
			asciiShift = 39
		}
	}

	prioritySum := 0
	for _, line := range input {
		prioritySum += len(line) % 2
		split := len(line) / 2
		prioritySum += typePriority[commonType(line[0:split], line[split:])]
	}
	fmt.Printf("Part One: %d\n", prioritySum)

	prioritySumV2 := 0
	for i := 0; i < len(input); i += 3 {
		prioritySumV2 += typePriority[commonTypeV2(input[i], input[i+1], input[i+2])]
	}
	fmt.Printf("Part Two: %d\n", prioritySumV2)
}

func commonType(rucksackA, rucksackB string) string {
	for _, c := range rucksackA {
		if strings.Contains(rucksackB, string(c)) {
			return string(c)
		}
	}
	return ""
}

func commonTypeV2(rucksackA, rucksackB, rucksackC string) string {
	for _, c := range rucksackA {
		if strings.Contains(rucksackB, string(c)) {
			if strings.Contains(rucksackC, string(c)) {
				return string(c)
			}
		}
	}
	return ""
}
