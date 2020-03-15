package alg

type DFS struct {
}

func (DFS) Search(g *Graph, start int) {

	visited := make([]bool, g.len)

	var recursive func(int)
	recursive = func(v int) {
		visited[v] = true
		final := g.list[v][len(g.list[v])-1]
		for i := g.list[v][0]; i != final; i++ {
			if !visited[i] {
				recursive(i)
			}
		}
	}

	recursive(start)

}
