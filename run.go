package main

import (
	"dmcblue/dren/models"
	terminaldisplay "dmcblue/dren/terminal-display"
	"fmt"
	"math/rand/v2"
)

type Barrier string
const (
	Sea Barrier = "sea"
	Mountains Barrier = "mountains"
	WasteLands Barrier = "waste lands"
	Mists Barrier = "mists"
	Desolation Barrier = "desolation"
)
var Barriers = [5]Barrier{Sea, Mountains, WasteLands, Mists, Desolation}

func run() {
	fmt.Println("North", Barriers[rand.IntN(5)])
	fmt.Println("East", Barriers[rand.IntN(5)])
	fmt.Println("South", Barriers[rand.IntN(5)])
	fmt.Println("West", Barriers[rand.IntN(5)])
	DrawGameMap()
}

func DrawGameMap() {
	player := models.Player{
		Position: [2]int{2, 2},
		Turns: 0,
	}
	gm := models.CreateHexMap(4, 8)
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
	terminaldisplay.Draw(gm, player)
}