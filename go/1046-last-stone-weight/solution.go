/*
It is a very simple problem using a heap to implement a priority queue

Can use the 'container/heap' Go package to implement the heap
	- Create a new IntHeap from the `stones` slice
	- Using the Pop method on the heap data structure, we retrieve the item
	with highest priority (larger element) - based on the implementation of `Less()`
*/
package main

import (
	"container/heap"
)

// implement priority queue as heap
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// may need to use interface{} for LeetCode's Go version
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

// may need to use interface{} for LeetCode's Go version
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}

func lastStoneWeight(stones []int) int {
	h := IntHeap(stones)
	heap.Init(&h)
	for h.Len() > 1 {
		heap.Push(&h, heap.Pop(&h).(int)-heap.Pop(&h).(int))
	}

	return heap.Pop(&h).(int)
}
