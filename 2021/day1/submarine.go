package day1

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

func sonarSweep(filename string) []int {
	return []int{}
}