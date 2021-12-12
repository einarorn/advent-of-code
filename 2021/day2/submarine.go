package day2

import (
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
	return 0
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
