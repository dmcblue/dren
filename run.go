package main

import (
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
}