package integration

import (
	"advent-of-code/2021/day3"
	"github.com/matryer/is"
	"testing"
)

func TestSubmarine_PowerConsumption(t *testing.T) {
	is := is.New(t)

	t.Run("Given we receive sample input When we calculate the power consumption Then the value returned is 198", func(t *testing.T) {
		expected := 198

		power := day3.PowerConsumption("day3-sample.txt")

		is.Equal(power, expected)
	})

	t.Run("Given we receive Part One input When we calculate the power consumption Then the value returned is 3242606", func(t *testing.T) {
		expected := 3242606

		power := day3.PowerConsumption("day3-part1.txt")

		is.Equal(power, expected)
	})
}

func TestSubmarine_LifeSupportRating(t *testing.T) {
	is := is.New(t)

	t.Run("Given we receive sample input When we calculate the support rating Then the value returned is 230", func(t *testing.T) {
		expected := 230

		lifeSupport := day3.LifeSupportRating("day3-sample.txt")

		is.Equal(lifeSupport, expected)
	})

	t.Run("Given we receive Part Two input When we calculate the power consumption Then the value returned is 4856080", func(t *testing.T) {
		expected := 4856080

		lifeSupport := day3.LifeSupportRating("day3-part1.txt")

		is.Equal(lifeSupport, expected)
	})
}
