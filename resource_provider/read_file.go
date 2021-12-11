package resource_provider

import (
	"bufio"
	"fmt"
	"os"
)

func ReadAllLines(filename string) ([]string, error) {
	var lines []string
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return lines, nil
}