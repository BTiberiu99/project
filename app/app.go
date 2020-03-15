package app

import (
	"flag"
	"fmt"
	"os"
	"project/alg"
)

func Start() {

	algFlag := *flag.String("alg", "bfs", "Accepted algoritms are bfs and dfs")

	nrGen := *flag.Uint("gen", 0, "Number of auto genereated configurations")

	fileName := *flag.String("file", "", "File that has config file")

	flag.Parse()

	required := []string{"gen", "file"}

	if nrGen == 0 && fileName == "" {
		fmt.Fprintf(os.Stderr, "missing required `-%s` or `-%s` argument/flag\n", required[0], required[1])
		os.Exit(2) // the same exit code flag.Parse uses
	}

	var algoritm alg.Algoritm

	switch algFlag {
	case "bfs":
		algoritm = alg.BFS{}
	case "dfs":
		algoritm = alg.DFS{}
	}

	if nrGen != 0 {
		RunAuto(algoritm, nrGen)
	} else {
		Run(algoritm, fileName)
	}

}

func Run(algoritm alg.Algoritm, fileName string) {

}

func RunAuto(algoritm alg.Algoritm, nr uint) {

}
