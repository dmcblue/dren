package main

import (
	"fmt"
	// "log"
	// "math/rand/v2"
	"os"
	// "strconv"
	// "strings"
)

const Roll = "roll"
const Run = "run"
const UI = "ui"

func help() {
	messages := [4]string{
		"dren <COMMAND>",
		"  Commands:",
		"  roll <rolls...>: '4,5 6,2' rolls a D4 five times and a D6 two times",
		"  run: Runs a game",
	}
	for _, str := range messages {
		fmt.Println(str)
	}
}

func err(message string) {
	fmt.Println(message)
	help()
	os.Exit(1)
}

func main() {
	args := os.Args[1:]

	num_args := len(args)
	if num_args > 0 {
		command := args[0]
		if command == Roll {
			make_rolls(args[1:])
		} else if command == Run {
			run()
		} else if command == UI {
			ui()
		} else {
			err(fmt.Sprintf("Unknown command '%s'", command))
		}
	} else {
		err("No command given")
	}
	
	fmt.Println("")
}
