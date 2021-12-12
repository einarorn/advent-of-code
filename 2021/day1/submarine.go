package day1

import (
	"advent-of-code/resource_provider"
	"fmt"
	"strconv"
)

func FindSantaSleighKeysDepth(filename string, slidingWindow int) int {
	input, err := resource_provider.ReadAllLines(filename)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	measurements := sonarSweep(input)
	increases := measurementIncreases(measurements)

	return increases
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

func measurementSlidingWindow(measurements []int, windowLen int) []int {
	var sliding []int
	arrayLen := len(measurements)

	if arrayLen < windowLen {
		return []int{}
	}

	for i := 0; i < arrayLen; i++ {
		if i+windowLen <= arrayLen {
			sum := 0
			for j := i; j < i+windowLen; j++ {
				sum += measurements[j]
			}
			sliding = append(sliding, sum)
		}
	}

	return sliding
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