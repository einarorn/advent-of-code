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

func calculatePositionAndDepth(commands []Commands) (int, int) {
	return -1, -1
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
