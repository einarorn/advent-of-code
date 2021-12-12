package integration_test_test

import (
	"advent-of-code/2021/day1"
	"github.com/matryer/is"
	"testing"
)

func TestSubmarine_FindSantaSleighKeysDepth(t *testing.T) {
	is := is.New(t)

	t.Run("Given we have sample input When we check keys depth Then depth has increased by 7", func(t *testing.T) {
		depth := day1.FindSantaSleighKeysDepth("day1-sample.txt", 0)

		is.Equal(depth, 7)
	})

	t.Run("Given we have part 1 input When we check keys depth Then depth has increased by 1676", func(t *testing.T) {
		depth := day1.FindSantaSleighKeysDepth("day1-part1.txt", 0)

		is.Equal(depth, 1676)
	})

	t.Run("Given we have part 2 input When we check keys depth using a sliding window of 3 Then depth has increased by 1706", func(t *testing.T) {
		depth := day1.FindSantaSleighKeysDepth("day1-part1.txt", 3)

		is.Equal(depth, 1706)
	})
}
