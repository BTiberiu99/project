package main

import (
	"fmt"
	"os"
	"project/problem/puzzle"
	"project/utils"
)

func main() {
	file, err := os.Open("./test.txt")

	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	puzzle, err := puzzle.NewPuzzle(file)
	path := puzzle.DFS()

	path2 := puzzle.BFS()

	fmt.Println(len(path))
	utils.PrintPath(path2)
}
