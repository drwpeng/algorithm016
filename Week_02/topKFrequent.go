package Week_02

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	myhash := make(map[int]int)
	for _, v := range nums {
		myhash[v]++
	}

	h := intHeap{}
	heap.Init(&h)
	for key, v := range myhash {
		heap.Push(&h, [2]int{key, v})
		if h.Len() > k {
			heap.Pop(&h)
		}
	}
	res := []int{}
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(&h).([2]int)[0])
	}
	return res
}

type intHeap [][2]int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *intHeap) delete(i int) {
	heap.Remove(h, i)
}

