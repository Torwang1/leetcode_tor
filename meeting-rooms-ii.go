package leetcode

import (
	"container/heap"
	"sort"
)

// 题目: 最小会议数量;
//
// 分析: 排序 + 画时间轴;
//
// 针对**结束时间**维护小根堆, 遍历所有会议开始的时间;
//
func minMeetingRooms(intervals [][]int) int {
	begin := 0 // 开始时间(下标)
	end := 1   // 结束时间(下标)

	var ans int

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// 针对开始时间排序;
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][begin] < intervals[j][begin] })

	// 首先: 从堆中移除“已经结束”的会议;
	// 再者: 将新开始的会议加入, 计算会议室的最大数量;

	top := 0 // 堆顶元素
	h := make(IntHeap, 0)
	for _, meeting := range intervals {
		for len(h) > 0 && h[top] <= meeting[begin] {
			heap.Pop(&h)
		}

		heap.Push(&h, meeting[end])

		ans = max(ans, len(h))
	}

	return ans
}

// 题目关心“会议室数量”, 不关心具体会议. 所以, 堆内仅需要记录结束时间即可;
