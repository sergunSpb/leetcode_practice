package main

import (
	"container/heap"
	"sort"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	h *IntHeap
	k int
}

func Constructor(k int, nums []int) KthLargest {
	var h IntHeap
	if len(nums) > 0 {
		sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })

		var tn IntHeap
		if len(nums) >= k {
			tn = nums[0:k]
		} else {
			tn = nums
		}
		h = make(IntHeap, len(tn))
		copy(h, tn)
	} else {
		h = make(IntHeap, 0)
	}

	heap.Init(&h)
	return KthLargest{
		h: &h,
		k: k,
	}
}

func (this *KthLargest) Add(val int) int {
	h := *this.h
	if this.h.Len() >= this.k && val < h[0] {
		return h[0]
	}

	heap.Push(this.h, val)
	if this.h.Len() > this.k {
		heap.Pop(this.h)
	}
	h = *this.h
	return h[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
