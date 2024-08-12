package maps

import (
	"math/rand"

	"github.com/vinicius725/path-finder/pkg/maps"
)

func Generate(x, y, coverage int) (maps.Map, error) {
	rows := make([][]uint, y)
	for i := range rows {
		row := make([]uint, x)
		for j := range row {
			var v uint = 0
			if i == 0 || i == y-1 || j == 0 || j == x-1 {
				row[j] = 1
				continue
			}
			if rand.Intn(100) < coverage {
				v = 1
			}
			row[j] = v
		}
		rows[i] = row
	}
	return maps.NewMap(rows)
}

func Generate2() (maps.Map, error) {
	rows := [][]uint{
		{1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 1, 0, 0, 1},
		{1, 0, 0, 1, 0, 0, 1},
		{1, 0, 0, 1, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1},
	}
	return maps.NewMap(rows)
}
