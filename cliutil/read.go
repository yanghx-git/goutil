package cliutil

import (
	"bufio"
	"os"
	"strings"

	"github.com/gookit/color"
)

// ReadInput read user input form Stdin
func ReadInput(question string) (string, error) {
	if len(question) > 0 {
		color.Print(question)
	}

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() { // reading
		return "", scanner.Err()
	}

	answer := scanner.Text()
	return strings.TrimSpace(answer), nil
}

// ReadLine read one line from user input.
// Usage:
// 	in := ReadLine("")
// 	ans, _ := ReadLine("your name?")
func ReadLine(question string) (string, error) {
	if len(question) > 0 {
		color.Print(question)
	}

	reader := bufio.NewReader(os.Stdin)
	answer, _, err := reader.ReadLine()
	return strings.TrimSpace(string(answer)), err
}

// ReadFirst read first char
func ReadFirst(question string) (string, error) {
	answer, err := ReadLine(question)
	if len(answer) == 0 {
		return "", err
	}

	return string(answer[0]), err
}
