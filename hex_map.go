package main

import (
	"fmt"
)

type HexMap [][]int

func CreateHexMap (width int, height int) HexMap {
	hex_map := make(HexMap, width)
	for x, _ := range hex_map {
		// zero is default value
		hex_map[x] = make([]int, height)
	}

	return hex_map
}

/*     ___
      /   \
  ___/ 0,0 \___
 /   \     /   \
/ 16, \___/     \
\  16 /   \     /
 \___/     \___/
     \     /
	  \___/
*/
// Only double digit sizes
// Width = w*5 + 2 
// Height = h*4 + 2 + 1 (top)
func (hex_map HexMap) Sdraw() [][]rune {
	w := hex_map.Width() * 10 + 5
	h := hex_map.Height() * 2 + 2 + 1
	char_map := make([][]rune, h)
	for i, _ := range char_map {
		char_map[i] = make([]rune, w)
		for j, _ := range char_map[i] {
			char_map[i][j] = ' '
		}
	}
	for x, _ := range hex_map {
		for y, _ := range hex_map[x] {
			drawHex(char_map, x, y)
		}
	}
	return char_map
}

func (hex_map HexMap) Draw() {
	char_map := hex_map.Sdraw()
	DrawCharMap(char_map)
}

func DrawCharMap(char_map[][]rune) {
	for k, _ := range char_map {
		fmt.Println(string(char_map[k]))
	}
}

func ternaryInt(condition bool, a int, b int) int {
	if condition { return a }
	return b
}

func ternaryString(condition bool, a string, b string) string {
	if condition { return a }
	return b
}

func drawHex(char_map [][]rune, x int, y int) {
	offset_x := offsetX(x, y)
	offset_y := offsetY(x, y)
	// top row
	char_map[offset_y - 1][offset_x + 2] = '_'
	char_map[offset_y - 1][offset_x + 3] = '_'
	char_map[offset_y - 1][offset_x + 4] = '_'
	// top sides
	char_map[offset_y    ][offset_x + 1] = '/'
	char_map[offset_y + 1][offset_x    ] = '/'
	char_map[offset_y    ][offset_x + 5] = '\\'
	char_map[offset_y + 1][offset_x + 6] = '\\'
	// bottom sides
	char_map[offset_y + 2][offset_x    ] = '\\'
	char_map[offset_y + 3][offset_x + 1] = '\\'
	char_map[offset_y + 2][offset_x + 6] = '/'
	char_map[offset_y + 3][offset_x + 5] = '/'
	// bottom
	char_map[offset_y + 3][offset_x + 2] = '_'
	char_map[offset_y + 3][offset_x + 3] = '_'
	char_map[offset_y + 3][offset_x + 4] = '_'
	// writeInHex(char_map, x, y, fmt.Sprintf("%d,%d", x, y))
}

func writeInHex(char_map [][]rune, x int, y int, str string) {
	nstr := "" + str
	if (len(nstr) < 3) { nstr = " " + nstr }
	if (len(nstr) < 3) { nstr += " " }
	if (len(nstr) < 3) { nstr += " " }
	nstr = ternaryString(len(nstr) < 3, " ", "") + nstr
	nstr = nstr + ternaryString(len(nstr) < 3, " ", "")
	nstr = nstr + ternaryString(len(nstr) < 3, " ", "")
	offset_x := offsetX(x, y)
	offset_y := offsetY(x, y)
	char_map[offset_y + 2][offset_x + 2] = rune(nstr[0])
	char_map[offset_y + 2][offset_x + 3] = rune(nstr[1])
	char_map[offset_y + 2][offset_x + 4] = rune(nstr[2])
}

func offsetX(x int, y int) int {
	return ternaryInt(y%2 == 0, 5, 0) + x*10
}

func offsetY(x int, y int) int {
	return y*2 + 1 // top row
}

func (hex_map HexMap) Height() int {
	return len(hex_map[0])
}

func (hex_map HexMap) Width() int {
	return len(hex_map)
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
func (hex_map HexMap) Move(x int, y int, move int) [2]int {
	nX := x
	nY := y
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
	
	if nX < 0 || nX >= hex_map.Width() { nX = x }
	if nY < 0 || nY >= hex_map.Height() { nY = y }

	return [2]int{nX, nY}
}
