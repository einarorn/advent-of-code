package integration

import (
	"advent-of-code/2021/day3"
	"github.com/matryer/is"
	"testing"
)

func TestSubmarine_FindSantaSleighKeysDepth(t *testing.T) {
	is := is.New(t)

	t.Run("Given we receive sample input When we calculate the power consumption Then the value returned is 198", func(t *testing.T) {
		expected := 198

		power := day3.PowerConsumption("day3-sample.txt")

		is.Equal(power, expected)
	})
}
