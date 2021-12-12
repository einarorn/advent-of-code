package day2

import (
	"advent-of-code/resource_provider"
	"fmt"
	"strconv"
	"strings"
)

type Command string

const (
	FORWARD Command = "forward"
	DOWN Command = "down"
	UP Command = "up"
)

type Commands struct {
	Action string
	Value int
}

func LocateSantaSub(filename string) int {
	input, err := resource_provider.ReadAllLines(filename)
	if err != nil {
		fmt.Println(err)
		return -1
	}

	commands := receiveCommands(input)
	position, depth := calculatePositionAndDepth(commands)

	return position*depth
}

func calculatePositionAndDepth(commands []Commands) (int, int) {
	var position, depth int

	for _, command := range commands {
		if command.Action == string(FORWARD) {
			position += command.Value
		} else {
			if command.Action == string(DOWN) {
				depth += command.Value
			} else if command.Action == string(UP) {
				depth -= command.Value
			}
		}
	}

	return position, depth
}

func calculatePositionAndDepthV2(commands []Commands) (int, int) {
	var position, depth int

	return position, depth
}

func receiveCommands(input []string) []Commands {
	var commands []Commands

	for _, line := range input {
		values := strings.Fields(line)
		action := values[0]
		value, err := strconv.Atoi(values[1])
		if err != nil {
			fmt.Println(err)
		}

		commands = append(commands, Commands{
			Action: action,
			Value:  value,
		})
	}

	return commands
}
