package main

import (
	"bufio"
	"os"
	"strings"
)

var input string = ""

// ReadString reads user input from command line, terminated via 'enter'
func ReadString() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}
}

// PrintString prints the string entered via GetString
func PrintString() string {
	return strings.ToUpper(input)
}
