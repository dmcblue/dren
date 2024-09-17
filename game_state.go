package main

import (
	// "dmcblue/dren/data"
	barrier "dmcblue/dren/data/barrier"
	"dmcblue/dren/data/pillar"
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
		edges := state.gameMap.GetEdges()
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

	pillarPoints := make([]models.Point, 0)

	current := models.Point{
		rand.IntN(14) + 3,
		rand.IntN(14) + 3,
	}
	pillarsSize := len(pillar.Pillars)
	pillarPoints = append(pillarPoints, current)

	indices := rand.Perm(pillarsSize)
	pillars := [4]pillar.Pillar{
		pillar.Pillars[indices[0]],
		pillar.Pillars[indices[1]],
		pillar.Pillars[indices[2]],
		pillar.Pillars[indices[3]],
	}
	for i := 0; i < 3; i++ {
		good := false
		for !good {
			good = true
			current = models.Point{
				rand.IntN(14) + 3,
				rand.IntN(14) + 3,
			}

			for _, pillarPoint := range pillarPoints {
				dist := models.PointDistance(current, pillarPoint)
				if dist < 5 {
					good = false
				}
			}
		}
		pillarPoints = append(pillarPoints, current)
	}

	for i, pillar := range pillars {
		point := pillarPoints[i]
		state.gameMap[point[0]][point[1]].Type = hextype.Pillar
		state.gameMap[point[0]][point[1]].Description = string(pillar)
	}

	return state
}