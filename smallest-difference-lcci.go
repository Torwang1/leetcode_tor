package leetcode

import (
	"math"
	"sort"
)

// 思路:
// - 排序 两个数组;
// - 双指针查找;
func smallestDifference(a []int, b []int) int {
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })

	minInt := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	absInt := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var ans int = math.MaxInt

	pa, pb := 0, 0
	for pa < len(a) && pb < len(b) {
		diff := a[pa] - b[pb]
		switch {
		case diff == 0:
			return 0
		case diff < 0:
			pa++
		case diff > 0:
			pb++
		}
		ans = minInt(ans, absInt(diff))
	}

	return ans
}
