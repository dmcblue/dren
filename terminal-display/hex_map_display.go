package terminaldisplay

import (
	"dmcblue/dren/models"
	"dmcblue/dren/utils"
	"strings"

	"github.com/fatih/color"
)

func DrawHexMap(canvas Canvas, hexMap models.HexMap) {
	for x, _ := range hexMap {
		for y, _ := range hexMap[x] {
			drawHex(canvas, hexMap[x][y], x, y)
		}
	}
}

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

func getHexColor(hex models.Hex, x int, y int) TextPrinter {
	var hexColor color.Attribute
	if x%2 == 0 && y%2 == 0 {
		hexColor = color.FgCyan
	} else if y%2 == 1 {
		hexColor = color.FgMagenta
	} else {
		hexColor = color.FgWhite
	}
	return color.New(hexColor).SprintFunc()
}

func drawHex(charMap Canvas, hex models.Hex, x int, y int) {
	c := getHexColor(hex, x, y)
	xOffset := getHexMapXOffset(x, y)
	yOffset := getHexMapYOffset(x, y)

	subMap := make([][]string, len(Hex))
	for i, hexLine := range Hex {
		chars := strings.Split(hexLine, "")
		subMap[i] = make([]string, len(chars))
		for j, ch := range chars {
			subMap[i][j] = utils.TernaryString(ch != " ", c(ch), " ")
		}
	}
	charMap.MapDraw(subMap, xOffset, yOffset)
}

func WriteInHex(charMap Canvas, x int, y int, str string) {
	var hexColor color.Attribute
	if x%2 == 0 {
		hexColor = color.FgCyan
	} else {
		hexColor = color.FgMagenta
	}
	c := color.New(hexColor).SprintFunc()
	nstr := "" + str
	if len(nstr) < 3 {
		nstr = " " + nstr
	}
	if len(nstr) < 3 {
		nstr += " "
	}
	if len(nstr) < 3 {
		nstr += " "
	}
	nstr = utils.TernaryString(len(nstr) < 3, " ", "") + nstr
	nstr = nstr + utils.TernaryString(len(nstr) < 3, " ", "")
	nstr = nstr + utils.TernaryString(len(nstr) < 3, " ", "")
	xOffset := getHexMapXOffset(x, y)
	yOffset := getHexMapYOffset(x, y)
	subMap := [][]string{{
		c(string(nstr[0])), c(string(nstr[1])), c(string(nstr[2])),
	}}

	charMap.MapDraw(subMap, xOffset+2, yOffset+2)
}

func getHexMapXOffset(x int, y int) int {
	return utils.TernaryInt(y%2 == 0, 6, 0) + x*12
}

func getHexMapYOffset(_ int, y int) int {
	return y*2 + 3 // top row
}
