package alg

type BFS struct {
}

func (BFS) Search(g *Graph, start int) {
	visited := make([]bool, g.len)

	queue := make([]int, 0)

	queue = append(queue, start)

	for len(queue) > 0 {

		s := queue[0]

		queue = queue[1 : len(queue)-1]

		final := g.list[s][len(g.list[s])-1]

		for i := g.list[s][0]; i != final; i++ {
			if !visited[i] {
				visited[i] = true
				queue = append(queue, i)
			}
		}
	}
}
