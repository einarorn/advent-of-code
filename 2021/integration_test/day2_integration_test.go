package integration_test

import (
	"advent-of-code/2021/day2"
	"github.com/matryer/is"
	"testing"
)

func TestSubmarine_LocateSantaSub(t *testing.T) {
	is := is.New(t)

	t.Run("Given we have sample input When we locate Santa sub Then we receive value of 150", func(t *testing.T) {
		location := day2.LocateSantaSub("day2-sample.txt", false)

		is.Equal(location, 150)
	})

	t.Run("Given we have Part One input When we locate Santa sub Then we receive value of 2039912", func(t *testing.T) {
		location := day2.LocateSantaSub("day2-part1.txt", false)

		is.Equal(location, 2039912)
	})

	t.Run("Given we have sample input When we locate Santa sub using complex calculation Then we receive value of 900", func(t *testing.T) {
		location := day2.LocateSantaSub("day2-sample.txt", true)

		is.Equal(location, 900)
	})

	t.Run("Given we have Part Two input When we locate Santa sub using complex calculation Then we receive value of 1942068080", func(t *testing.T) {
		location := day2.LocateSantaSub("day2-part1.txt", true)

		is.Equal(location, 1942068080)
	})
}
