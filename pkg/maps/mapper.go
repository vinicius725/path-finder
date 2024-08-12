package maps

type (
	terrain [][]uint
	Map     interface {
		Terrain() terrain
		Height() int
		Width() int
		Tile(Position) uint
		Tiles() [][]uint
	}
)

func (t terrain) Terrain() terrain {
	return t
}

func (t terrain) Tiles() [][]uint {
	return t
}

func (t terrain) validate() error {
	if len(t) == 0 {
		return &MapError{Msg: "map with height 0"}
	}
	m := len(t[0])
	for _, r := range t {
		if len(r) != m {
			return &MapError{Msg: "map is not a rectangle"}
		}
	}
	return nil
}

func NewMap(m [][]uint) (Map, error) {
	t := terrain(m)
	if err := t.validate(); err != nil {
		return nil, err
	}
	return &t, nil
}

type MapError struct {
	Msg string
}

func (e *MapError) Error() string {
	return e.Msg
}

func (e terrain) Height() int {
	return len(e)
}

func (e terrain) Width() int {
	return len(e[0])
}

func (e terrain) Tile(p Position) uint {
	return e[p.Y][p.X]
}
