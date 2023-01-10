package leetcode

import (
	"container/heap"
	"sort"
)

// 天际线, 主要包含两个部分: 1) 问题分析, 怎么找到天际线起点; 2) 算法复杂度优化, 使用堆(优先队列);
//
// 1) 怎么找到天际线起点:
// - 事实: 天际线起点, 一定在建筑物的 left 或 right;
// - 事实: 如果出现在建筑物的 right, 该建筑物的高度不计算在内;
//
// 2) 堆 优化算法复杂度:
// - 找到包含 x 的建筑物, 常规做法需要遍历, 找到所有包含 x 的建筑物列表. O(n)
// - 目标: 找到最大高度, “最大”可以用堆描述;
// - 堆元素的维护: 可以通过 x、left、right 来维护;
//
// 题解中提到的懒删除, 含义是: 堆中元素(非堆顶)可能不在 x 的投影中, 但是依然保留在堆中;
//
// 其它:
// - 题目要求, 相同高度不能连续. 直接使用 slice 最后一个元素来判断即可;

func getSkyline(buildings [][]int) [][]int {
	ans := [][]int{}

	// 建筑物边界列表, 递增排序
	boundaries := make([]int, 0, 2*len(buildings))
	for _, building := range buildings {
		boundaries = append(boundaries, building[left], building[right])
	}
	sort.Ints(boundaries)

	// 找到在 “投影 x” 中的建筑物
	i := 0
	h := make(skyLine, 0)

	for _, x := range boundaries {
		// 将投影中的建筑, 追加到堆中.
		for i < len(buildings) && buildings[i][left] <= x {
			heap.Push(&h, buildings[i])
			i++
		}

		// 保证堆顶元素在 “投影 x” 中. 注意: 把建筑物的右边界弹出.
		for len(h) > 0 && h[0][right] <= x {
			heap.Pop(&h)
		}

		maxH := 0
		if len(h) > 0 {
			maxH = h[0][height] // 投影中的最高建筑.
		}
		if len(ans) == 0 || ans[len(ans)-1][1] != maxH {
			ans = append(ans, []int{x, maxH})
		}
	}

	return ans
}

const (
	left   = 0
	right  = 1
	height = 2
)

type skyLine [][]int // []int = [left, right, height]

func (h skyLine) Len() int            { return len(h) }
func (h skyLine) Less(i, j int) bool  { return h[i][height] > h[j][height] } // 高度 的最大堆.
func (h skyLine) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *skyLine) Push(v interface{}) { *h = append(*h, v.([]int)) }
func (h *skyLine) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
