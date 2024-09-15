package main

import (
	// "dmcblue/dren/data"
	barrier "dmcblue/dren/data/barrier"
	"dmcblue/dren/models"
	hextype "dmcblue/dren/models/hex-type"
	// "fmt"
	"math/rand/v2"
)

type gameState struct {
	player models.Player
	barriers [4]barrier.Barrier /* 0 North, 1 East, 3 South, 4 West */
	gameMap models.HexMap
}

func createGameState() gameState {
	return gameState{
		barriers: [4]barrier.Barrier{barrier.None, barrier.None, barrier.None, barrier.None},
		gameMap: models.CreateHexMap(20, 20),
		player: models.Player{
			Position: [2]int{0, 0},
			Turns: 0,
		},
	}
}

func newGameState() gameState {
	state := createGameState()
	barriersSize := len(barrier.Barriers)
	for i, _ := range state.barriers {
		state.barriers[i] = barrier.Barriers[rand.IntN(barriersSize)]
	}
	for i := 0; i < 16; i++ {
		direction := rand.IntN(4)
		edges := state.gameMap.GetEdges2()
		side := edges[direction]
		if len(side) > 0 {
			indices := rand.Perm(len(side))
			maxTimes := min(8, len(indices))
			times := rand.IntN(maxTimes)
			for j := 0; j < times; j++ {
				point := side[indices[j]]
				state.gameMap[point[0]][point[1]].Type = hextype.None
			}
		}
	}

	state.player.Position[0] = rand.IntN(10) + 5
	state.player.Position[1] = rand.IntN(10) + 5

	return state
}