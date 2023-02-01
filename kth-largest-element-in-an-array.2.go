package leetcode

// 题目: 数组中第 K 个最大元素.
//
// 解题: 快速排序 的变形算法.
//
// 注意: 第 k 个, 反应在下标是 k-1
func findKthLargest2(nums []int, k int) int {

	quickSortKth(nums, 0, len(nums)-1, k-1)

	return nums[k-1]
}

// quickSortKth 非完整的快速排序, 仅保证“下标k”有最终元素.
//
// 降序实现.
func quickSortKth(nums []int, l, r int, k int) {
	// l, r 在后续递归中, 用作边界值. 所以, 不能修改.
	i, j := l, r

	// 目标: 找到“下标k”的最终元素.
	nums[i], nums[k] = nums[k], nums[i]

	target := nums[i]

	for i < j {
		// 从右向左, 找到第一个大于 target 的值, 放到 i 的位置;
		for i < j && nums[j] <= target {
			j--
		}
		nums[i] = nums[j]

		// 从左向右, 找到第一个小于 target 的值, 放在 j 的位置;
		for i < j && nums[i] >= target {
			i++
		}
		nums[j] = nums[i]
	}

	nums[i] = target

	switch {
	case i == k:
		return // 结束
	case k < i:
		quickSortKth(nums, l, i-1, k)
	case k > i:
		quickSortKth(nums, i+1, r, k)
	}
}
