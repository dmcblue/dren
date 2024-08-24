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

func help() {
	messages := [3]string{
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
			gm := CreateHexMap(4, 8)
			// fmt.Println(gm)
			// gm.Draw()
			point := [2]int{2, 2}
			char_map := gm.Sdraw()
			writeInHex(char_map, point[0], point[1], "1")
			// fmt.Println(point)
			point = gm.Move(point[0], point[1], 1)
			writeInHex(char_map, point[0], point[1], "2")
			// fmt.Println(3, point)
			point = gm.Move(point[0], point[1], 2)
			writeInHex(char_map, point[0], point[1], "3")
			// fmt.Println(2, point)
			point = gm.Move(point[0], point[1], 0)
			writeInHex(char_map, point[0], point[1], "4")
			// fmt.Println(0, point)
			DrawCharMap(char_map)
		} else {
			err(fmt.Sprintf("Unknown command '%s'", command))
		}
	} else {
		err("No command given")
	}
	
	fmt.Println("")
}
