package alg

type Edges []int

func (e Edges) Exist(edge int) bool {
	for i := range e {
		if e[i] == edge {
			return true
		}
	}

	return false
}
