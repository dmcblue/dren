package terminaldisplay

import (
	"dmcblue/dren/models"
	"fmt"
)

func DrawPlayer(canvas Canvas, player models.Player) {
	fmt.Println(player.Turns + 1)
	WriteInHex(canvas, player.Position[0], player.Position[1], fmt.Sprintf("%d", player.Turns))
}