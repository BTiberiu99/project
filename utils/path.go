package utils

import (
	"fmt"
	"project/problem/puzzle"
)

func PrintPath(s []*puzzle.Config) {
	for i := range s {
		fmt.Println(s[i].String())
		fmt.Println("")
	}
}
