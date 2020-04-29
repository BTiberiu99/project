package puzzle

import (
	"bufio"
	"container/heap"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"project/utils"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	Empty = int8(0)
)

type Puzzle struct {
	intial         *Config
	final          *Config
	MaxDepth       int
	VisitedConfigs int
}

type Stats struct {
	Algoritm       string   `json:"algoritm"`
	FinalDepth     int      `json:"final_depth"`
	RunningTime    Duration `json:"running_time"`
	MemoryUsage    MemUsage `json:"memory_usage"`
	VisitedConfigs int      `json:"visited_configs"`
}

type Duration time.Duration

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Duration(d).String())
}

type MemUsage uint64

func (m MemUsage) MarshalJSON() ([]byte, error) {
	var um = "Kb"

	if m/1024/1024 > 0 {
		m = m / 1024 / 1024
		um = "GM"
	} else if m/1024 > 0 {
		m = m / 1024
		um = "MB"
	}
	return json.Marshal(fmt.Sprintf("%d %s", m, um))
}

type ConfigPuzzle struct {
	Reader io.Reader
	Intial *Config
	Final  *Config
}

func NewPuzzle(c *ConfigPuzzle) (*Puzzle, error) {
	puzzle := &Puzzle{}

	if c.Intial != nil && c.Final != nil {
		puzzle.intial = c.Intial
		puzzle.final = c.Final
	} else if c.Reader != nil {
		err := puzzle.FromFile(c.Reader)

		if err != nil {
			return nil, err
		}

	} else {
		return nil, errors.New("No Intial& Config matrix or Reader ")
	}

	return puzzle, nil
}

func (p *Puzzle) IsFinal(c2 *Config) bool {
	return p.final.IsSame(c2)
}

func (p *Puzzle) IsSolvable() bool {

	return p.intial.IsSolvable()
}

func (p *Puzzle) BFS() []*Config {
	p.MaxDepth = 0
	queue := NewQueue()

	queue.Append(p.intial)

	visited := map[string]bool{}
	for queue.Len() > 0 {

		node := queue.Pop()

		visited[node.Key()] = true

		if p.IsFinal(node) {
			p.VisitedConfigs = len(visited)
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
	return []*Config{}
}

func (p *Puzzle) DFS() []*Config {
	p.MaxDepth = 0
	stack := NewStack()

	p.intial.ReverseNeighbors = true

	stack.Append(p.intial)

	visited := map[string]bool{}
	for stack.Len() > 0 {

		node := stack.Pop()

		visited[node.Key()] = true

		if p.IsFinal(node) {
			p.VisitedConfigs = len(visited)
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
	return []*Config{}
}
func (p *Puzzle) H(c *Config) float64 {
	sum := float64(0)
	lenSide := p.final.Len()
	lenArr := p.final.Len() * p.final.Len()
	help := map[int][]int{}

	add := func(v, nr, index int) {
		if _, ok := help[nr]; !ok {
			help[nr] = make([]int, 2)
		}

		help[nr][index] = v
	}

	for i := 0; i < lenArr; i++ {

		//C
		add(i, int(c.mat[i/lenSide][i%lenSide]), 0)

		//Final board
		add(i, int(p.final.mat[i/lenSide][i%lenSide]), 1)
	}

	for i := range help {

		if i == 0 {
			continue
		}

		sum += math.Abs(float64(help[i][0]%lenSide-help[i][1]%lenSide)) + math.Abs(float64(help[i][0]/lenSide-help[i][1]/lenSide))
	}

	return 0
}
func (p *Puzzle) AStar() []*Config {
	p.MaxDepth = 0

	var (
		node  *Helper
		entry *Helper
		h     *Heap
	)

	hp, visited, heap_entry := NewHeap(), map[string]bool{}, map[string]*Helper{}

	p.intial.HKey = p.H(p.intial)

	entry = &Helper{
		Key:  p.intial.HKey,
		Move: p.intial.Move,
		Root: p.intial,
	}

	heap.Push(hp, entry)

	heap_entry[entry.Root.Key()] = entry

	for hp.Len() > 0 {

		node = heap.Pop(hp).(*Helper)
		delete(heap_entry, node.Root.Key())
		visited[node.Root.Key()] = true

		if p.IsFinal(node.Root) {
			p.VisitedConfigs = len(visited)
			return append(node.Root.Parents(), node.Root)
		}

		for _, neighbor := range node.Root.Neighbors() {

			neighbor.HKey = neighbor.Cost + p.H(neighbor)

			entry = &Helper{
				Key:  neighbor.HKey,
				Move: neighbor.Move,
				Root: neighbor,
			}

			if _, exist := visited[neighbor.Key()]; !exist {

				heap.Push(hp, entry)

				visited[neighbor.Key()] = true

				heap_entry[neighbor.Key()] = entry

				if neighbor.Depth > p.MaxDepth {
					p.MaxDepth = neighbor.Depth
				}

			} else if val, exist := heap_entry[neighbor.Key()]; exist && neighbor.HKey < val.Root.HKey {

				h = hp.(*Heap)

				index := h.Index(&Helper{
					Move: val.Root.Move,
					Key:  val.Root.HKey,
					Root: val.Root,
				})

				if index < 0 || index >= h.Len() {

					panic("Index not found")
				}

				h.configs[index] = entry

				heap_entry[neighbor.Key()] = entry

				heap.Init(hp)
			}
		}

	}

	return []*Config{}
}

func (p *Puzzle) Statistics() []Stats {
	return []Stats{
		p.Stat("bfs"),
		p.Stat("dfs"),
		p.Stat("astar"),
	}
}

func (p *Puzzle) Stat(algoritm string) Stats {

	var stats Stats
	switch algoritm {
	case "bfs":
		stats = Stats{
			Algoritm: "Breadth First Search",
		}
	case "dfs":
		stats = Stats{
			Algoritm: "Depth First Search",
		}
	case "astar":
		stats = Stats{
			Algoritm: "AStart",
		}
	default:
		panic("not implemented")
	}

	var memUsageKb uint64
	var finalDepth int
	t := utils.Timeit(stats.Algoritm, func() {

		var s []*Config
		f := CheckMemoryUsage()
		defer f(0)
		switch algoritm {
		case "bfs":
			s = p.BFS()

		case "dfs":
			s = p.DFS()
		case "astar":
			s = p.AStar()
		default:
			panic("not implemented")
		}
		memUsageKb = f()

		finalDepth = s[len(s)-1].Depth
	})

	stats.MemoryUsage = MemUsage(memUsageKb)
	stats.RunningTime = Duration(t)

	stats.FinalDepth = finalDepth
	stats.VisitedConfigs = p.VisitedConfigs
	return stats

}

func CheckMemoryUsage() func(...int) uint64 {
	var m runtime.MemStats

	var (
		average = uint64(0)
		run     = true
		last    = uint64(0)
		count   = uint64(0)
	)
	go func() {

		for run {
			time.Sleep(1 * time.Nanosecond)
			runtime.ReadMemStats(&m)
			if last == 0 {
				last = m.Alloc
				count = 0
			}

			dif := m.Alloc - last

			average = ((count * average) + dif) / (count + 1)

			count++
		}

	}()

	return func(cmd ...int) uint64 {
		if len(cmd) > 0 {
			switch cmd[0] {
			case 0:
				run = false
			case 1:
				average = 0
				last = 0
			}
		}

		return average / 1024
	}
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
