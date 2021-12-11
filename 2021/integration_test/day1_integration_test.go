package integration_test_test

import (
	"advent-of-code/2021/day1"
	"github.com/matryer/is"
	"testing"
)

func TestSubmarine_FindSantaSleighKeysDepth(t *testing.T) {
	is := is.New(t)

	t.Run("Given we have sample input When we check keys depth Then depth has increased by 7", func(t *testing.T) {
		depth := day1.FindSantaSleighKeysDepth("day1-sample.txt")

		is.Equal(depth, 7)
	})

	t.Run("Given we have part 1 input When we check keys depth Then depth has increased by 100", func(t *testing.T) {
		depth := day1.FindSantaSleighKeysDepth("day1-part1.txt")

		is.Equal(depth, 1676)
	})
}
