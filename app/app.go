package app

import (
	"project/problem/puzzle"
	"project/utils"
	"strings"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func Start() {

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:     1200,
		Height:    768,
		Title:     "8 Puzzle",
		JS:        js,
		CSS:       css,
		Colour:    "#131313",
		Resizable: true,
	})
	app.Bind(Stats)
	app.Bind(TakeMoves)
	app.Bind(Page)
	app.Run()

}

var (
	items    []*puzzle.Config
	lenItems int
)

func TakeMoves(alg, configs string) int {
	p, err := puzzle.NewPuzzle(&puzzle.ConfigPuzzle{
		Reader: strings.NewReader(configs),
	})

	if err != nil {
		return 0
	}
	utils.Timeit(alg, func() {
		switch alg {
		case "bfs":

			items = p.BFS()

		case "dfs":
			items = p.DFS()

		case "astar":

			items = p.AStar()

		default:
			items = make([]*puzzle.Config, 0)

		}
	})

	lenItems = len(items)
	return lenItems
}

func Stats(configs string) []puzzle.Stats {
	p, err := puzzle.NewPuzzle(&puzzle.ConfigPuzzle{
		Reader: strings.NewReader(configs),
	})

	if err != nil {
		return []puzzle.Stats{}
	}

	return p.Statistics()
}

func Page(page, perPage int) [][]interface{} {
	m := make([][]interface{}, perPage/5)

	if items != nil {
		var index int
		for i := 0; i < perPage; i++ {
			index = (page-1)*perPage + i
			if index >= lenItems {
				break
			}
			if m[i/5] == nil {
				m[i/5] = make([]interface{}, 5)
			}

			m[i/5][i%5] = map[string]interface{}{
				"item":  items[index],
				"index": index,
			}
		}

	}
	return m
}
