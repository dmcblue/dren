package barrier

type Barrier string
const (
	None Barrier = ""
	Sea Barrier = "sea"
	Mountains Barrier = "mountains"
	WasteLands Barrier = "waste lands"
	Mists Barrier = "mists"
	Desolation Barrier = "desolation"
)
var Barriers = [5]Barrier{Sea, Mountains, WasteLands, Mists, Desolation}