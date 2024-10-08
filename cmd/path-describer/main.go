package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/vinicius725/path-finder/internal/levels"
	"github.com/vinicius725/path-finder/pkg/maps"
	pathfinder "github.com/vinicius725/path-finder/pkg/path-finder"
)

func main() {
	m, err := levels.Generate(50, 50, 20)
	if err != nil {
		panic(err)
	}
	var (
		start = maps.Position{X: 1, Y: 3}
		end   = maps.Position{X: 48, Y: 48}
	)

	f := pathfinder.New(m, start, end)
	path, err := f.Find()
	if err != nil {
		tiles := make(map[maps.Position]rune)
		tiles[start] = 'O'
		tiles[end] = 'X'
		for y, row := range m.Tiles() {
			for x, tile := range row {
				if v, ok := tiles[maps.Position{X: x, Y: y}]; ok {
					color.New(color.FgRed).Print(string(v))
					continue
				}
				if tile == 0 {
					fmt.Print(" ")
				} else {
					fmt.Print("#")
				}
			}
			fmt.Print("\n")
		}
		panic(err)
	}

	tiles := make(map[maps.Position]rune)
	for _, pos := range path {
		tiles[pos] = '.'
	}
	tiles[start] = 'O'
	tiles[end] = 'X'
	for y, row := range m.Tiles() {
		for x, tile := range row {
			if v, ok := tiles[maps.Position{X: x, Y: y}]; ok {
				color.New(color.FgRed).Print(string(v))
				continue
			}
			if tile == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Print("\n")
	}

}
