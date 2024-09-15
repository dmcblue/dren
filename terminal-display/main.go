package terminaldisplay

import (
	"dmcblue/dren/models"
	"fmt"
	"strings"
)

type Canvas [][]string

func Draw(hexMap models.HexMap, player models.Player) {
	canvas := CreateCanvas(hexMap)
	DrawHexMap2(canvas, hexMap)
	DrawPlayer2(canvas, player)
	canvas.Print()
}

func CreateCanvas(hexMap models.HexMap) Canvas {
	// w := getHexMapXOffset(hexMap.Width()+1, hexMap.Height()+1)
	// h := getHexMapYOffset(hexMap.Width()+1, hexMap.Height()+1)
	w := 40
	h := 20

	canvas := make([][]string, h)
	for i, _ := range canvas {
		canvas[i] = make([]string, w)

		for j, _ := range canvas[i] {
			canvas[i][j] = " "
		}
	}
	return canvas
}

func (canvas Canvas) Print() {
	for k, _ := range canvas {
		fmt.Println(strings.Join(canvas[k], ""))
	}
}

type TextPrinter func(...interface{}) string

func (canvas Canvas) MapDraw(subMap Canvas, xOffset int, yOffset int) {
	w := len(canvas[0])
	h := len(canvas)
	for i, _ := range subMap {
		for j, _ := range subMap[i] {
			x := xOffset+j
			y := yOffset+i
			if x < w && y < h && subMap[i][j] != " " {
				canvas[y][x] = subMap[i][j]
			}
		}
	}
}

