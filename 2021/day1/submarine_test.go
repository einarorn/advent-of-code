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