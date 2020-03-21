package puzzle

import (
	"container/heap"
	"sort"
)

type Helper struct {
	Key  float64
	Move int
	Root *Config
}
type Heap struct {
	sort.Interface
	configs []*Helper
}

func (h *Heap) Less(i, j int) bool {

	a, b := h.configs[i], h.configs[j]

	if a.Key < b.Key {
		return true
	}

	if a.Move < b.Move {
		return true
	}

	return a.Root.Less(b.Root)
}

// Swap swaps the elements with indexes i and j.
func (h *Heap) Swap(i, j int) {
	h.configs[i], h.configs[j] = h.configs[j], h.configs[i]
}

func (h *Heap) Push(val interface{}) {

	if config, ok := val.(*Helper); ok {

		h.configs = append(h.configs, config)
	} else {
		panic("Interface not Helper pointer")
	}

}

func (h *Heap) Pop() interface{} {

	var c *Helper

	c, h.configs = h.configs[len(h.configs)-1], h.configs[:len(h.configs)-1]

	return c
}

func (h *Heap) Index(hl *Helper) int {
	var is bool
	for i := range h.configs {
		is = h.configs[i].Key == hl.Key && h.configs[i].Move == hl.Move && h.configs[i].Root.IsSame(hl.Root)
		if is {
			return i
		}
	}

	return -1
}

func (h *Heap) Len() int {
	return len(h.configs)
}

func NewHeap() heap.Interface {

	return &Heap{
		configs: make([]*Helper, 0),
	}
}
