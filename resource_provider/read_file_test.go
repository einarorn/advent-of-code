package resource_provider

import (
	"github.com/matryer/is"
	"testing"
)

func TestReadFile_AllLines(t *testing.T) {
	is := is.New(t)

	t.Run("Given we have a file with 3 lines When we read all lines Then we will return them as string array", func(t *testing.T) {
		expected := []string{"Line 1", "Line 2", "Line 3"}

		allLines, err := ReadAllLines("test-all-lines.txt")

		is.NoErr(err)
		is.Equal(allLines, expected)
	})
}