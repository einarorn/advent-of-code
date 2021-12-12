package day2

import (
	"github.com/matryer/is"
	"testing"
)

func TestSubmarine_CalculatePositionAndDepth(t *testing.T) {
	is := is.New(t)

	t.Run("Given we receive series of commands When we have calculated them Then we will have the correct position and depth", func(t *testing.T) {
		commands := sampleCommands()

		position, depth := calculatePositionAndDepth(commands)

		is.Equal(position, 15)
		is.Equal(depth, 10)
	})
}

func TestSubmarine_CalculatePositionAndDepthV2(t *testing.T) {
	is := is.New(t)

	t.Run("Given we receive series of commands When we have calculated them Then we will have the correct position and depth", func(t *testing.T) {
		commands := sampleCommands()

		position, depth := calculatePositionAndDepthV2(commands)

		is.Equal(position, 15)
		is.Equal(depth, 60)
	})
}

func TestSubmarine_ReceiveCommands(t *testing.T) {
	is := is.New(t)

	t.Run("Given we receive command input When we have processed them Then we return the corresponding series of commands", func(t *testing.T) {
		input := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
		expected := sampleCommands()

		measurements := receiveCommands(input)

		is.Equal(measurements, expected)
	})
}

func sampleCommands() []Commands {
	commands := []Commands{
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

	return commands
}
