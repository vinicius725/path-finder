package pathfinder

import (
	"math"
	"slices"

	"github.com/vinicius725/path-finder/pkg/maps"
)

type PathFinder struct {
	src       maps.Position
	dest      maps.Position
	avaliable map[maps.Position]Visit
	stepedOn  map[maps.Position]any
	terrain   maps.Map
}
type Visit struct {
	src   *Visit
	pos   maps.Position
	price int
}

func New(t maps.Map, src, dest maps.Position) *PathFinder {
	return &PathFinder{
		src:       src,
		dest:      dest,
		terrain:   t,
		stepedOn:  make(map[maps.Position]any),
		avaliable: make(map[maps.Position]Visit),
	}
}

func (v Visit) Track() (track []maps.Position) {

	var prev Visit
	for prev = v; prev.src != nil; prev = *prev.src {
		track = append(track, prev.pos)
	}
	track = append(track, prev.pos)
	slices.Reverse(track)
	return
}

func (f *PathFinder) Find() ([]maps.Position, error) {
	current := Visit{pos: f.src, price: 0}
	for {

		neigh := f.Neightbours(current)
		f.Store(neigh)
		if len(f.avaliable) == 0 {
			return nil, &ErrNoPathFound{Curr: current}
		}
		current = f.Choose()
		if current.pos == f.dest {
			return current.Track(), nil
		}
	}
}

func (f *PathFinder) Choose() (v Visit) {
	for _, a := range f.avaliable {
		v = a
		break
	}
	for _, a := range f.avaliable {
		if a.price < v.price {
			v = a
		}
		if a.price == v.price && f.Distance(a.pos, f.dest) < f.Distance(v.pos, f.dest) {
			v = a
		}
	}
	delete(f.avaliable, v.pos)
	f.stepedOn[v.pos] = nil
	return
}

func (f *PathFinder) Store(v []Visit) {
	for _, o := range v {
		if _, ok := f.stepedOn[o.pos]; !ok {
			f.avaliable[o.pos] = o
		}
	}
}

func (f *PathFinder) Neightbours(p Visit) (vs []Visit) {
	for x := p.pos.X - 1; x < p.pos.X+2; x++ {
		for y := p.pos.Y - 1; y < p.pos.Y+2; y++ {
			pos := maps.Position{X: x, Y: y}
			if p.pos != pos && f.Validate(pos) {
				vs = append(
					vs, Visit{
						src:   &p,
						pos:   pos,
						price: f.GetPrice(pos),
					},
				)
			}
		}
	}
	return
}
func (f *PathFinder) Validate(p maps.Position) (r bool) {
	if p.X >= 0 && p.X < f.terrain.Width() && p.Y >= 0 && p.Y < f.terrain.Height() && f.terrain.Tile(p) == 0 {
		r = true
	}
	return
}
func (f *PathFinder) GetPrice(p maps.Position) int {
	done := f.Distance(f.src, p)
	left := f.Distance(p, f.dest)
	return left - done
}

func (f *PathFinder) Distance(fr, t maps.Position) int {
	return Abs(fr.X-t.X) + Abs(fr.Y-t.Y)
}

func Abs(i int) int {
	return int(math.Abs(float64(i)))
}

type ErrNoPathFound struct {
	Curr Visit
}

func (e *ErrNoPathFound) Error() string {
	return "No path found"
}
