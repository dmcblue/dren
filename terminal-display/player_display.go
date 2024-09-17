package terminaldisplay

import (
	"dmcblue/dren/models"
	"fmt"

	"github.com/fatih/color"
)

func DrawPlayer(canvas Canvas, player models.Player) {
	fmt.Println(player.Turns + 1)
	WriteInHex(canvas, player.Position[0], player.Position[1], fmt.Sprintf("%d", player.Turns))
}

func DrawPlayer2(canvas Canvas, player models.Player) {
	c := color.New(color.FgCyan).SprintFunc()
	WriteInHex2(canvas, player.Position[0], player.Position[1], c("P"))
}