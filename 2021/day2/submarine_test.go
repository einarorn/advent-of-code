package day2

import (
	"github.com/matryer/is"
	"testing"
)

func TestSubmarine_SonarSweep(t *testing.T) {
	is := is.New(t)

	t.Run("Given we receive an input of 10 lines When we read all of them and they are all numeric Then we return corresponding array", func(t *testing.T) {
		input := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
		expected := []Commands{
			{
				Action: string(FORWARD),
				Value:  5,
			},
			{
				Action: string(DOWN),
				Value:  5,
			},
			{
				Action: string(FORWARD),
				Value:  8,
			},
			{
				Action: string(UP),
				Value:  3,
			},
			{
				Action: string(DOWN),
				Value:  8,
			},
			{
				Action: string(FORWARD),
				Value:  2,
			}}

		measurements := receiveCommands(input)

		is.Equal(measurements, expected)
	})
}
