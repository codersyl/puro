// how to write priority queue using Heap

import (
	"container/heap"
)

// 5 functions
type IntHeap []int
func (h IntHeap) Len() int           { return len(h) }

// 此为小根堆如需大根堆， h[i] > h[j]
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push 与 Pop需要改变size，所以用指针
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 可以再封一下
func (h *IntHeap) In(x int) {
	heap.Push(h, x)
}

func (h *IntHeap) Out() int {
	return heap.Pop(h).(int)
}