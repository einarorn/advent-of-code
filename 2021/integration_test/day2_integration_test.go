package integration_test

import (
	"advent-of-code/2021/day2"
	"github.com/matryer/is"
	"testing"
)

func TestSubmarine_LocateSantaSub(t *testing.T) {
	is := is.New(t)

	t.Run("Given we have sample input When we locate Santa sub Then we receive value of 150", func(t *testing.T) {
		location := day2.LocateSantaSub("day2-sample.txt")

		is.Equal(location, 150)
	})
}
