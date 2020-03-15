package puzzle

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	Empty = int8(0)
)

type Puzzle struct {
	intial   *Config
	final    *Config
	MaxDepth int
}

func NewPuzzle(file io.Reader) (*Puzzle, error) {
	puzzle := &Puzzle{}

	err := puzzle.FromFile(file)

	if err != nil {
		return nil, err
	}

	return puzzle, nil
}

func (p *Puzzle) IsFinal(c2 *Config) bool {
	return p.final.IsSame(c2)
}

func (p *Puzzle) BFS() []*Config {

	queue := NewQueue()

	queue.Append(p.intial)

	visited := map[string]bool{}
	for queue.Len() > 0 {

		node := queue.Pop()

		visited[node.Key()] = true

		if p.IsFinal(node) {

			return append(node.Parents(), node)
		}

		for _, neighbor := range node.Neighbors() {
			if _, exist := visited[neighbor.Key()]; !exist {

				visited[neighbor.Key()] = true
				queue.Append(neighbor)

				if neighbor.Depth > p.MaxDepth {
					p.MaxDepth = neighbor.Depth
				}

			}
		}

	}
	return queue.Configs()
}

func (p *Puzzle) DFS() []*Config {

	stack := NewStack()

	p.intial.ReverseNeighbors = true

	stack.Append(p.intial)

	visited := map[string]bool{}
	for stack.Len() > 0 {

		node := stack.Pop()

		visited[node.Key()] = true

		if p.IsFinal(node) {

			return append(node.Parents(), node)
		}

		for _, neighbor := range node.Neighbors() {
			if _, exist := visited[neighbor.Key()]; !exist {

				visited[neighbor.Key()] = true
				stack.Append(neighbor)

				if neighbor.Depth > p.MaxDepth {
					p.MaxDepth = neighbor.Depth
				}

			}
		}

	}
	return stack.Configs()
}

//First should be the intial matrix
//Last should be final matrix
//Between them should be an empty row
func (p *Puzzle) FromFile(file io.Reader) error {

	scanner := bufio.NewScanner(file)
	matrixs := make([]string, 2)

	//Split matrixes
	str := ""

	for scanner.Scan() {

		row := strings.TrimSpace(scanner.Text())

		if scanner.Err() != nil {
			return scanner.Err()
		}

		if row == "" {
			matrixs[0] = str
			str = ""

		} else {
			str += row + "\n"
		}

	}

	matrixs[1] = str

	//End Split

	//Make Intial Matrix
	matrix, err := p.Read(strings.NewReader(matrixs[0]))

	if err != nil {
		return err
	}
	p.intial = NewConfig(matrix)

	//Make Final Matrix
	matrix, err = p.Read(strings.NewReader(matrixs[1]))

	if err != nil {
		return err
	}
	p.final = NewConfig(matrix)

	if p.final.Len() != p.intial.Len() {
		return errors.New("The two matrix length dosen'\t match")
	}

	return nil
}

func (p *Puzzle) Read(matrixBuffer io.Reader) ([][]int8, error) {
	scanner := bufio.NewScanner(matrixBuffer)

	first := true
	lenR := 0
	nrRow := 1

	matrix := make([][]int8, 0)

	for scanner.Scan() {

		row := scanner.Text()
		split := ""
		if strings.Contains(row, " ") {
			split = " "
		}
		nrs := strings.Split(row, split)
		if first {
			first = false
			lenR = len(nrs)
		}

		if len(nrs) != lenR {
			return nil, errors.New(fmt.Sprintf("Number of numbers on row %d is not %d", nrRow, lenR))
		}

		nrsInt := make([]int8, len(nrs))

		for i := range nrs {
			nrsInt[i] = stringToInt(nrs[i])
		}

		matrix = append(matrix, nrsInt)

		nrRow++

	}

	if len(matrix) != lenR {
		return nil, errors.New("Matrix is not of type n x n")
	}

	return matrix, scanner.Err()
}

func stringToInt(str string) int8 {

	str = strings.TrimSpace(str)

	nr, err := strconv.ParseInt(str, 10, 8)

	if err != nil {
		panic(err.Error())
	}

	return int8(nr)
}
