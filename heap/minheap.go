package minheap

import (
	"fmt"
)

type MinHeap struct {
	heap []int
	size int
}

func NewMinHeap(c int) *MinHeap {
	return &MinHeap{
		heap: make([]int, c),
		size: 0,
	}
}

func (h *MinHeap) Insert(k int) int {
	if h.size >= len(h.heap) {
		return -1
	}
	h.heap[h.size] = k
	h.swim(h.size)
	h.size++
	// fmt.Println(h.heap, h.size)
	return 0
}

func (h *MinHeap) Delete() {
	if h.size == 0 {
		return
	}
	h.heap[0] = h.heap[h.size-1]
	h.size--
	h.sink(0)
}

func (h *MinHeap) Print() {
	fmt.Println(h.heap, h.size)
}

func (h *MinHeap) swim(k int) {
	for k > 0 && h.heap[(k-1)/2] > h.heap[k] {
		h.heap[(k-1)/2], h.heap[k] = h.heap[k], h.heap[(k-1)/2]
		k = (k - 1) / 2
	}
}

func (h *MinHeap) sink(k int) {
	for 2*k+1 < h.size {
		i := 2*k + 1
		if i+1 < h.size && h.heap[i+1] < h.heap[i] {
			i++
		}
		if h.heap[k] <= h.heap[i] {
			break
		}
		h.heap[k], h.heap[i] = h.heap[i], h.heap[k]
		k = i
	}
}
