package lesson

import (
	"container/heap"
	"fmt"
	"testing"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]

}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func TestHeap(t *testing.T) {
	t.Log("1")

	h := &IntHeap{1, 2, 3, 45}
	heap.Init(h)

	heap.Push(h, 1)
	t.Log(*h)
	fmt.Println(*h)
	//heap.Interface
}
