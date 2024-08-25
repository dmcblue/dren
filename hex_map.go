package main

import (
	// "fmt"
	"strings"

	"github.com/fatih/color"
)

type HexMap [][]int
var Hex = [4]string{
	` /⎺⎺⎺\ `,
	`/     \`,
	`\     /`,
	` \___/ `,
	// ` /ΞΞΞ\ `,
	// `//   \\`,
	// `\\   //`,
	// ` \ΞΞΞ/ `,
}

func CreateHexMap (width int, height int) HexMap {
	hexMap := make(HexMap, width)
	for x, _ := range hexMap {
		// zero is default value
		hexMap[x] = make([]int, height)
	}

	return hexMap
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
func (hexMap HexMap) Sdraw() [][]string {
	w := hexMap.Width() * 10 + 5 + 30
	h := hexMap.Height() * 2 + 2 + 1 + 30
	// charMap := make([][]rune, h)
	
	charMap := make([][]string, h)
	for i, _ := range charMap {
		// charMap[i] = make([]rune, w)
		charMap[i] = make([]string, w)

		for j, _ := range charMap[i] {
			charMap[i][j] = " "
		}
	}
	for x, _ := range hexMap {
		for y, _ := range hexMap[x] {
			drawHex(charMap, x, y)
		}
	}
	return charMap
}

func (hexMap HexMap) Draw() {
	charMap := hexMap.Sdraw()
	DrawCharMap(charMap)
}

func DrawCharMap(charMap[][]string) {
	c := color.New(color.FgCyan) //.Add(color.Underline)
	for k, _ := range charMap {
		c.Println(strings.Join(charMap[k], ""))
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

func drawHex(charMap [][]string, x int, y int) {
	var hexColor color.Attribute
	if x%2 == 0 {
		hexColor = color.FgCyan
	} else {
		hexColor = color.FgMagenta
	}
	c := color.New(hexColor) .SprintFunc()
	xOffset := xOffset(x, y)
	yOffset := yOffset(x, y)


	// // top row
	// charMap[yOffset - 1][xOffset + 2] = c("_")
	// // charMap[yOffset - 1][xOffset + 3] = c("_")
	// charMap[yOffset - 1][xOffset + 4] = c("_")
	// // top sides
	// charMap[yOffset    ][xOffset + 1] = c("/")
	// charMap[yOffset + 1][xOffset    ] = c("/")
	// charMap[yOffset    ][xOffset + 5] = c("\\")
	// charMap[yOffset + 1][xOffset + 6] = c("\\")
	// // bottom sides
	// charMap[yOffset + 2][xOffset    ] = c("\\")
	// charMap[yOffset + 3][xOffset + 1] = c("\\")
	// charMap[yOffset + 2][xOffset + 6] = c("/")
	// charMap[yOffset + 3][xOffset + 5] = c("/")
	// // bottom
	// charMap[yOffset + 3][xOffset + 2] = c("_")
	// charMap[yOffset + 3][xOffset + 3] = c("_")
	// charMap[yOffset + 3][xOffset + 4] = c("_")
	// writeInHex(charMap, x, y, fmt.Sprintf("%d,%d", x, y))
	
	subMap := make([][]string, len(Hex))
	for i, hexLine := range Hex {
		chars := strings.Split(hexLine, "")
		subMap[i] = make([]string, len(chars))
		for j, ch := range chars {
			subMap[i][j] = ternaryString(ch != " ", c(ch), " ")
		}
	}
	mapDraw(charMap, subMap, xOffset, yOffset)
}

type TextPrinter func(...interface{}) string
func mapDraw(charMap [][]string, subMap[][]string, xOffset int, yOffset int) {
	for i, _ := range subMap {
		for j, _ := range subMap[i] {
			if subMap[i][j] != " " {
				charMap[yOffset + i][xOffset + j] = subMap[i][j]
			}
		}
	}
}

func writeInHex(charMap [][]string, x int, y int, str string) {
	var hexColor color.Attribute
	if x%2 == 0 {
		hexColor = color.FgCyan
	} else {
		hexColor = color.FgMagenta
	}
	c := color.New(hexColor) .SprintFunc()
	nstr := "" + str
	if (len(nstr) < 3) { nstr = " " + nstr }
	if (len(nstr) < 3) { nstr += " " }
	if (len(nstr) < 3) { nstr += " " }
	nstr = ternaryString(len(nstr) < 3, " ", "") + nstr
	nstr = nstr + ternaryString(len(nstr) < 3, " ", "")
	nstr = nstr + ternaryString(len(nstr) < 3, " ", "")
	xOffset := xOffset(x, y)
	yOffset := yOffset(x, y)
	charMap[yOffset + 2][xOffset + 2] = c(string(nstr[0]))
	charMap[yOffset + 2][xOffset + 3] = c(string(nstr[1]))
	charMap[yOffset + 2][xOffset + 4] = c(string(nstr[2]))
}

func xOffset(x int, y int) int {
	return ternaryInt(y%2 == 0, 6, 0) + x*12
}

func yOffset(_ int, y int) int {
	return y*2 + 3 // top row
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
func (hexMap HexMap) Move(x int, y int, move int) [2]int {
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
	
	if nX < 0 || nX >= hexMap.Width() { nX = x }
	if nY < 0 || nY >= hexMap.Height() { nY = y }

	return [2]int{nX, nY}
}
