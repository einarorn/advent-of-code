package integration_test_test

import (
	"advent-of-code/2021/day1"
	"github.com/matryer/is"
	"testing"
)

func TestSubmarine_FindSantaSleighKeysDepth(t *testing.T) {
	is := is.New(t)

	t.Run("Given we have sample input When we check keys depth Then we depth has increased by 7", func(t *testing.T) {
		depth := day1.FindSantaSleighKeysDepth("day1-sample.txt")

		is.Equal(depth, 7)
	})
}
