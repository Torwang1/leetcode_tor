package leetcode

import "container/heap"

// 滑动窗口中的最大值
// 解法一: 懒堆;
// 解法二: 单调栈;

// 懒堆实现;
func maxSlidingWindow(nums []int, k int) []int {
	h := make(slidingWindowHeap, 0, k)

	// 使用前 k 个元素建堆;
	for i := 0; i < k; i++ {
		h = append(h, &slidingWindowItem{
			Val: nums[i],
			i:   i,
		})
	}
	heap.Init(&h)

	top := 0 // 常量, 表示堆顶

	// 查询堆顶元素, 如果堆顶元素不在滑动窗口中, 则将它弹出;

	ans := []int{h[top].Val}
	for i := k; i < len(nums); i++ {
		heap.Push(&h, &slidingWindowItem{
			Val: nums[i],
			i:   i,
		})

		// 弹出不在滑动窗口中的元素
		for h[top].i <= i-k {
			heap.Pop(&h)
		}

		ans = append(ans, h[top].Val)
	}

	return ans
}

// 堆元素需要记录: 1) 原始下标, 用于判断懒堆顶是否在窗口中; 2) Val 用于构建堆;

type slidingWindowItem struct {
	Val int
	i   int
}

type slidingWindowHeap []*slidingWindowItem

func (h slidingWindowHeap) Len() int           { return len(h) }
func (h slidingWindowHeap) Less(i, j int) bool { return h[i].Val > h[j].Val } // 大根堆
func (h slidingWindowHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *slidingWindowHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*slidingWindowItem))
}
func (h *slidingWindowHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
