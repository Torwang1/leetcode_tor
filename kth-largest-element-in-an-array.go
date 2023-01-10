package leetcode

import "container/heap"

func findKthLargest(nums []int, k int) int {
	h := IntBigHeap(nums)

	heap.Init(&h)

	// 弹出前 k-1 个元素;
	for i := 0; i < k-1; i++ {
		heap.Pop(&h)
	}

	// 弹出第 k 个元素;
	return heap.Pop(&h).(int)
}

type IntBigHeap []int

func (h IntBigHeap) Len() int           { return len(h) }
func (h IntBigHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntBigHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push 会修改 slice 的长度, 而不仅仅修改其元素. 所以, 使用 pointer 作为 receiver.
func (h *IntBigHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop 移除数组中的最后一个元素.
// 参考 heap.Pop 实现: 首先将 [0] 和 [n-1] 交换, 再 down 修正堆.  [n-1]表示最小元素.
func (h *IntBigHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
