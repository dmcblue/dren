package main

import (
	// "dmcblue/dren/data"
	// "dmcblue/dren/models"
	terminaldisplay "dmcblue/dren/terminal-display"
	"fmt"
	// "math/rand/v2"
)

var state gameState

func run() {
	fmt.Println("Setting up game state...")
	state = newGameState()

	// fmt.Println("North", data.Barriers[rand.IntN(5)])
	// fmt.Println("East", data.Barriers[rand.IntN(5)])
	// fmt.Println("South", data.Barriers[rand.IntN(5)])
	// fmt.Println("West", data.Barriers[rand.IntN(5)])
	DrawGameMap(state)
}

func DrawGameMap(state gameState) {
	// fmt.Println(gm)
	// gm.Draw()
	// point := [2]int{2, 2}

	// char_map := terminaldisplay.Sdraw(gm)
	// terminaldisplay.WriteInHex(char_map, point[0], point[1], "1")
	// // fmt.Println(point)
	// player.Position = gm.Move(player.Position, 1)
	// terminaldisplay.WriteInHex(char_map, point[0], point[1], "2")
	// // fmt.Println(3, point)
	// player.Position = gm.Move(player.Position, 2)
	// terminaldisplay.WriteInHex(char_map, point[0], point[1], "3")
	// // fmt.Println(2, point)
	// player.Position = gm.Move(player.Position, 0)
	// terminaldisplay.WriteInHex(char_map, point[0], point[1], "4")
	// // fmt.Println(0, point)
	// terminaldisplay.DrawCharMap(char_map)
	fmt.Printf("To the north, %v\n", state.barriers[0])
	fmt.Printf("To the east, %v\n", state.barriers[1])
	fmt.Printf("To the south, %v\n", state.barriers[2])
	fmt.Printf("To the west, %v\n", state.barriers[3])
	fmt.Println("Dren:")
	terminaldisplay.Draw(state.gameMap, state.player)
}