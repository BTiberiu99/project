package gen

import "project/alg"

type Generator interface {
	Generate() alg.Graph
}
