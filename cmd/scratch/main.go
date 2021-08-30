package main

import (
	"container/heap"
	"fmt"
)

func main() {
	nums := []int{3, 2, 20, 5, 3, 1, 2, 5, 6, 9, 10, 4}
	h := &IntHeap{}
	for _, n := range nums {
		heap.Push(h, n)
		
	
	}

	for i := 0; i < len(nums); i++ {
		fmt.Printf("%d,", heap.Pop(h).(int))
	}
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
