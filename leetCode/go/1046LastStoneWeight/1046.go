package laststoneweight

import (
	"container/heap"
	"sort"
)

func lastStoneWeight_sort(stones []int) int {
	for len(stones) > 1 {
		// increasing order
		sort.Ints(stones)
		// pop last two stone
		a := stones[len(stones)-1]
		b := stones[len(stones)-2]
		stones = stones[:len(stones)-2]
		diff := a - b
		if diff != 0 {
			stones = append(stones, diff)
		}
	}
	if len(stones) == 1 {
		return stones[0]
	}
	return 0
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
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
func lastStoneWeight(stones []int) int {
	h := &IntHeap{}
	heap.Init(h)
	for _, stone := range stones {
		heap.Push(h, stone)
	}
	for h.Len() > 1 {
		a := heap.Pop(h).(int)
		b := heap.Pop(h).(int)
		diff := a - b
		if diff != 0 {
			heap.Push(h, diff)
		}
	}
	if h.Len() == 0 {
		return 0
	}
	return heap.Pop(h).(int)
}
