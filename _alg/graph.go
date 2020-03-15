package alg

type Graph struct {
	list []Edges
	len  int
}

func (g *Graph) AddEdge(from, to int) {

	if from < 0 || from >= g.len {
		panic("From dosen't match the length of the graph!")
	}

	if to < 0 || to >= g.len {
		panic("To dosen't match the length of the graph!")
	}

	if g.list[from] == nil {
		g.list[from] = Edges{}
	}

	if g.list[from].Exist(to) {
		return
	}

	g.list[from] = append(g.list[from], to)
}

func (g *Graph) Search(alg Algoritm, start int) {
	alg.Search(g, start)
}

func NewGraph(length int) *Graph {
	if length <= 0 {
		panic("Length must be greater than 0")
	}

	list := make([]Edges, length)

	return &Graph{
		list: list,
		len:  length,
	}
}
