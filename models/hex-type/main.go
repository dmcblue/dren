package hextype

type HexType int
const (
	Plain HexType = iota
	None
	North
	East
	South
	West
)