package game

type Menu struct {
	Items []MenuItem
}

type MenuItem struct {
	Label  string
	Action func()
}

const (
	StartGame = iota
	ShipSelection
)

func NewMenu() *Menu {
	return &Menu{
		Items: nil,
	}
}
