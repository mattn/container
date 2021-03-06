package heap

import (
	"sort"
)

// A heap for ints. The pointer to the zero value of Ints is a heap with default less
// func which compares int values on the natual order.
// Use NewInts to customize less func and initial capacity.
type Ints struct {
	// Less func given indexes of i, j in the list slice.
	less func(i, j int) bool
	list sort.IntSlice
}

// Len returns the number of elements in the current heap.
func (h *Ints) Len() int {
	return len(h.list)
}

// Push inserts an element to the heap.
func (h *Ints) Push(x int) {
	h.list = append(h.list, x)

	if h.less == nil {
		PushLastF(len(h.list), h.list.Less, h.list.Swap)
	} else {
		PushLastF(len(h.list), h.less, h.list.Swap)
	}
}

// Pop removes the top element from the heap and returns it.
func (h *Ints) Pop() int {
	if h.less == nil {
		PopToLastF(len(h.list), h.list.Less, h.list.Swap)
	} else {
		PopToLastF(len(h.list), h.less, h.list.Swap)
	}

	res := h.list[len(h.list)-1]
	h.list = h.list[:len(h.list)-1]

	return res
}

// NewInts returns a *Ints with customized less func and initial capacity.
func NewInts(less func(x, y int) bool, cap int) *Ints {
	h := &Ints{}

	if less != nil {
		h.less = func(i, j int) bool {
			return less(h.list[i], h.list[j])
		}
	}
	if cap > 0 {
		h.list = make(sort.IntSlice, 0, cap)
	}

	return h
}
