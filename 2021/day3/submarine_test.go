package day3

import (
	"github.com/matryer/is"
	"testing"
)

func TestSubmarine_MostOrLeastCommonBit(t *testing.T) {
	is := is.New(t)

	t.Run("Given we receive a set of bits When we find the most common bit in each position Then we return the combined as one binary string", func(t *testing.T) {
		input := []string{"10101", "11100", "00100", "10101"}
		expected := "10101"

		actual := mostOrLeastCommonBit(input, true)

		is.Equal(actual, expected)
	})

	t.Run("Given we receive a set of bits When we find the least common bit in each position Then we return the combined as one binary string", func(t *testing.T) {
		input := []string{"10101", "11100", "00100", "10101"}
		expected := "01010"

		actual := mostOrLeastCommonBit(input, false)

		is.Equal(actual, expected)
	})
}

func TestSubmarine_MostCommonBitByValueHorizontally(t *testing.T) {
	is := is.New(t)

	t.Run("Given we receive a set of bits When we filter out horizontally by most common bit value Then we return the last remaining binary string", func(t *testing.T) {
		input := mockBitSample()
		expected := "10111"

		actual := mostOrLeastCommonBitHorizontally(input, true)

		is.Equal(actual, expected)
	})

	t.Run("Given we receive a set of bits When we filter out horizontally by least common bit value Then we return the last remaining binary string", func(t *testing.T) {
		input := mockBitSample()
		expected := "01010"

		actual := mostOrLeastCommonBitHorizontally(input, false)

		is.Equal(actual, expected)
	})
}

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

func mockBitSample() []string {
	return []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
}
