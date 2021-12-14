package day3

import (
	"github.com/matryer/is"
	"testing"
)

func TestSubmarine_ConvertBinaryToNumber(t *testing.T) {
	is := is.New(t)

	t.Run("Given we have the binary string 1010 When we evaluate the decimal value Then we return the value 10", func(t *testing.T) {
		binary := "1010"
		expected := 10

		actual := convertBinaryToNumber(binary)

		is.Equal(actual, expected)
	})

	t.Run("Given we have the binary string 0101 When we evaluate the decimal value Then we return the value 5", func(t *testing.T) {
		binary := "0101"
		expected := 5

		actual := convertBinaryToNumber(binary)

		is.Equal(actual, expected)
	})

	t.Run("Given we have the binary string 0000 When we evaluate the decimal value Then we return the value 0", func(t *testing.T) {
		binary := "0000"
		expected := 0

		actual := convertBinaryToNumber(binary)

		is.Equal(actual, expected)
	})

	t.Run("Given we have the binary string 1111 When we evaluate the decimal value Then we return the value 15", func(t *testing.T) {
		binary := "1111"
		expected := 15

		actual := convertBinaryToNumber(binary)

		is.Equal(actual, expected)
	})
}
