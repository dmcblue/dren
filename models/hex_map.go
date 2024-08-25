package models

import (
	// "fmt"
	// "strings"

	// "github.com/fatih/color"
)

type HexType int
const (
	Plain HexType = iota
)
type Hex struct {
	Type HexType
}
type HexMap [][]Hex


func CreateHexMap (width int, height int) HexMap {
	hexMap := make(HexMap, width)
	for x, _ := range hexMap {
		// zero is default value
		hexMap[x] = make([]Hex, height)
		for y, _ := range hexMap[x] {
			hexMap[x][y] = Hex{
				Type: Plain,
			}
		}
	}

	return hexMap
}

func (hexMap HexMap) Height() int {
	return len(hexMap[0])
}

func (hexMap HexMap) Width() int {
	return len(hexMap)
}

/*
    0
   ___
5 /   \ 1
 /     \ 
 \     /
4 \___/ 2
    3
*/
/*     ___
      /   \
  ___/ 0,0 \___
 /   \     /   \
/ 0,1 \___/ 1,1 \
\     /   \     /
 \___/ 0,2 \___/
     \     /
  0,3 \___/ 1,2
*/
func (hexMap HexMap) Move(position Point, move int) [2]int {
	nX := position[0]
	nY := position[1]
	switch move {
		case 0:
			// nX
			nY -= 2
		case 1:
			nX += 1
			nY -= 1
		case 2:
			nX += 1
			nY += 1
		case 3:
			// nX
			nY += 2
		case 4: 
			// nX
			nY += 1
		case 5:
			// nX
			nY -= 1
	}
	
	if nX < 0 || nX >= hexMap.Width() { nX = position[0] }
	if nY < 0 || nY >= hexMap.Height() { nY = position[1] }

	return [2]int{nX, nY}
}
