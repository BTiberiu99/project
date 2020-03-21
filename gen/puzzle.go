package gen

import (
	"math/rand"
	"project/problem/puzzle"
	"time"
)

type Puzzle struct {
	generated []*puzzle.Puzzle
	rand      *rand.Rand
	N         int
}

func NewPuzzle(n int) *Puzzle {
	return &Puzzle{
		N: n,
	}
}

func (p *Puzzle) Generate(nr int) {
	p.generated = make([]*puzzle.Puzzle, 0)

	noLoop := 0
	for nr > 0 {

		puzzle, err := p.nextPuzzle()

		if err != nil {
			nr--
			p.generated = append(p.generated, puzzle)
		} else {
			noLoop++
		}

		if noLoop > 20000 {
			break
		}

	}
}

func (p *Puzzle) nextPuzzle() (*puzzle.Puzzle, error) {

	initial := p.config()
	final := p.config()
	count := 0
	for !initial.IsSolvable() && count < 3000 {
		initial = p.config()
		count++
	}
	return puzzle.NewPuzzle(&puzzle.ConfigPuzzle{
		Intial: initial,
		Final:  final,
	})

}

func (p *Puzzle) config() *puzzle.Config {

	m := make([][]int8, p.N)
	for i := range m {
		m[i] = make([]int8, p.N)
	}

	i := p.N*p.N - 1

	index := p.Rand().Intn(9)

	for i > 0 {
		i--
		nr := p.randNr()
		for p.existNumber(m, nr) {
			nr = p.randNr()
		}
		if index > p.N {
			index = 0
		}
		m[index/p.N][index%p.N] = nr

		index++

	}

	return puzzle.NewConfig(m)
}

func (p *Puzzle) existNumber(mat [][]int8, nr int8) bool {

	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == nr {
				return true
			}
		}
	}

	return false
}

func (p *Puzzle) randNr() int8 {
	return int8(p.Rand().Intn(8)) + 1
}

func (p *Puzzle) Rand() *rand.Rand {
	if p.rand == nil {
		p.rand = rand.New(rand.NewSource(time.Now().Unix()))
	}
	return p.rand
}

func (p *Puzzle) Statistics() {
	for i := range p.generated {

	}
}
