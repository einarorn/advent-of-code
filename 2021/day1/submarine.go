package day1

import (
	"fmt"
	"strconv"
)

func FindSantaSleighKeysDepth(filename string) int {
	return -1
}

func measurementIncreases(measurements []int) int {
	count := 0
	previous := measurements[0]

	for _, num := range measurements[1:] {
		if num > previous {
			count++
		}
		previous = num
	}

	return count
}

func sonarSweep(input []string) []int {
	var measurements []int

	for _, value := range input {
		number, err := strconv.Atoi(value)
		if err != nil {
			fmt.Printf("Unable to convert value to int: %v\n", err)
			return []int{}
		}
		measurements = append(measurements, number)
	}

	return measurements
}