package day1

import (
	"github.com/matryer/is"
	"testing"
)

func TestSubmarine_MeasurementIncreases(t *testing.T) {
	is := is.New(t)

	t.Run("Given we receive 10 measurements When 7 of them increase from previous value Then we return 7", func(t *testing.T) {
		measurements := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

		increase := measurementIncreases(measurements)

		is.Equal(increase, 7)
	})

	t.Run("Given we receive 10 measurements When none of them increase from previous value Then we return 0", func(t *testing.T) {
		measurements := []int{199, 190, 190, 188, 176, 170, 168, 150, 145, 132}

		increase := measurementIncreases(measurements)

		is.Equal(increase, 0)
	})

	t.Run("Given we receive 10 measurements When all of them increase from previous value Then we return 9", func(t *testing.T) {
		measurements := []int{199, 200, 208, 210, 212, 222, 240, 269, 272, 277}

		increase := measurementIncreases(measurements)

		is.Equal(increase, 9)
	})
}

func TestSubmarine_MeasurementSlidingWindow(t *testing.T) {
	is := is.New(t)

	t.Run("Given we receive a set of measurements When they have been converted into a three-measurement sliding window Then we return the corresponding sums", func(t *testing.T) {
		measurements := []int{100, 200, 300, 400, 500}
		expected := []int{600, 900, 1200}

		slidingWindow := measurementSlidingWindow(measurements, 3)

		is.Equal(slidingWindow, expected)
	})

	t.Run("Given we receive a set measurements When they are less then the sliding window Then we return an empty array", func(t *testing.T) {
		measurements := []int{10, 20}

		slidingWindow := measurementSlidingWindow(measurements, 3)

		is.Equal(slidingWindow, []int{})
	})
}

func TestSubmarine_SonarSweep(t *testing.T) {
	is := is.New(t)

	t.Run("Given we receive an input of 10 lines When we read all of them and they are all numeric Then we return corresponding array", func(t *testing.T) {
		input := []string{"199", "200", "208", "212"}
		expected := []int{199, 200, 208, 212}

		measurements := sonarSweep(input)

		is.Equal(measurements, expected)
	})
}