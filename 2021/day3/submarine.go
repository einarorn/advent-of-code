package day3

import (
	"advent-of-code/resource_provider"
	"fmt"
	"math"
)

func PowerConsumption(filename string) int {
	input, err := resource_provider.ReadAllLines(filename)
	if err != nil {
		fmt.Println(err)
		return -1
	}

	gammaRate := mostOrLeastCommonBit(input, true)
	epsilonRate := mostOrLeastCommonBit(input, false)

	gamma := convertBinaryToNumber(gammaRate)
	epsilon := convertBinaryToNumber(epsilonRate)

	return gamma * epsilon
}

func mostOrLeastCommonBit(report []string, mostCommon bool) string {
	res := ""
	size := len(report[0])
	arr := make([]int, size)

	for _, line := range report {
		for i := 0; i < size; i++ {
			if line[i:i+1] == "1" {
				arr[i] += 1
			} else {
				arr[i] -= 1
			}
		}
	}

	bitA := "1"
	bitB := "0"

	if mostCommon == false {
		bitA = "0"
		bitB = "1"
	}

	for _, count := range arr {
		if count >= 0 {
			res += bitA
		} else {
			res += bitB
		}
	}

	return res
}

func mostOrLeastCommonBitHorizontally(report []string, mostCommon bool) string {
	return ""
}

func convertBinaryToNumber(binary string) int {
	val := 0
	l := len(binary)

	for i := 0; i < l; i++ {
		pos := l-i-1
		if binary[pos:pos+1] == "1" {
			val += int(math.Pow(2, float64(i)))
		}
	}

	return val
}
