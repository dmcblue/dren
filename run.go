package main

import (
	// "dmcblue/dren/data"
	// "dmcblue/dren/models"
	hextype "dmcblue/dren/models/hex-type"
	terminaldisplay "dmcblue/dren/terminal-display"
	"fmt"
	// "math/rand/v2"
)

var state gameState

func run() {
	fmt.Println("Setting up game state...")
	state = newGameState()

	fmt.Printf("To the north, %v\n", state.barriers[0])
	fmt.Printf("To the east, %v\n", state.barriers[1])
	fmt.Printf("To the south, %v\n", state.barriers[2])
	fmt.Printf("To the west, %v\n", state.barriers[3])
	fmt.Println("Dren:")
	DrawGameMap(state)

	
	for i, row := range state.gameMap {
		for j, hex := range row {
			if hex.Type == hextype.Pillar {
				fmt.Printf("Pillar '%s' at (%d, %d)\n", hex.Description, i, j)
			}
		}
	}
}

func DrawGameMap(state gameState) {
	terminaldisplay.Draw(state.gameMap, state.player)
}