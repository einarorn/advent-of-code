package day3

import (
	"math"
)

func mostOrLeastCommonBit(report []string, mostCommon bool) string {
	return ""
}

func convertBinaryToNumber(binary string) int {
	l := len(binary)
	val := 0

	for i := 0; i < l; i++ {
		pos := l-i-1
		if binary[pos:pos+1] == "1" {
			val += int(math.Pow(2, float64(i)))
		}

	}

	return val
}
