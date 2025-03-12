package main

import (
	"container/heap"
)

type MaxHeap []int

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Push(t any) {
	*h = append(*h, t.(int))
}

func (h *MaxHeap) Pop() any {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]

	return v
}

func lastStoneWeight(stones []int) int {
	mh := &MaxHeap{}
	*mh = make([]int, len(stones))
	copy(*mh, stones)
	heap.Init(mh)

	for mh.Len() > 1 {
		r := heap.Pop(mh).(int) - heap.Pop(mh).(int)

		if r != 0 {
			heap.Push(mh, r)
		}
	}

	if mh.Len() == 1 {
		return mh.Pop().(int)
	}

	return 0
}
