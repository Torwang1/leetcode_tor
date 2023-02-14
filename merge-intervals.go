package leetcode

import "sort"

// 题目: 合并区间
//
// 注意: 状态放在 ans[last] 中维护. 避免单独维护 left/right, 导致状态不一致, 进而引入很多复杂性.

func mergeii(intervals [][]int) [][]int {
	const left, right = 0, 1

	// 排序, 保证左边界有序;
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][left] < intervals[j][left]
	})

	ans := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := ans[len(ans)-1]
		span := intervals[i]

		// 合并 span
		if span[left] >= last[left] && span[left] <= last[right] {
			if span[right] > last[right] {
				last[right] = span[right]
			}
			continue
		}

		// 新 span
		ans = append(ans, span)
	}

	return ans
}
