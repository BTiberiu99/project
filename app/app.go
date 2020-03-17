package app

import (
	"fmt"
	"project/problem/puzzle"
	"strings"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func Start() {

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:     1024,
		Height:    768,
		Title:     "Puzzle 8",
		JS:        js,
		CSS:       css,
		Colour:    "#131313",
		Resizable: true,
	})
	app.Bind(TakeMoves)
	app.Run()

}

func TakeMoves(alg, configs string) []*puzzle.Config {
	p, err := puzzle.NewPuzzle(&puzzle.ConfigPuzzle{
		Reader: strings.NewReader(configs),
	})
	fmt.Println(alg, "\n", configs)
	if err != nil {
		return make([]*puzzle.Config, 0)
	}

	switch alg {
	case "bfs":

		return p.BFS()
	case "dfs":
		return p.DFS()
	default:
		return make([]*puzzle.Config, 0)
	}
}
