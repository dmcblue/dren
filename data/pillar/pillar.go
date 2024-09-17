package pillar

type Pillar string
const (
	None Pillar = ""
	Marinus Pillar = "Marinus"
	Charon Pillar = "Charon"
	TheEyes Pillar = "TheEyes"
	TheBuriedKing Pillar = "TheBuriedKing"
	Keth Pillar = "keth"
	Amarth Pillar = "Amarth"
	Fylus Pillar = "Fylus"
	Horseman Pillar = "Horseman"
)
var Pillars = [8]Pillar{
	Marinus,
	Charon,
	TheEyes,
	TheBuriedKing,
	Keth,
	Amarth,
	Fylus,
	Horseman,
}