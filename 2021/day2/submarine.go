package day2

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

func receiveCommands(commands []string) []Commands {
	return nil
}
