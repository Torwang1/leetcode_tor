package leetcode

// 快速排序 quicksort
//
// 优化点:
// - 随机选择下标, 避免"有序数组"场景下的性能降级.
//
func sortArray(nums []int) []int {

	quickSort(nums, func(i1, i2 int) bool { return i1 < i2 })

	return nums
}

// golang slice 原地排序.
//
// less 控制排序规则. 例如: 升序、降序.
func quickSort(nums []int, less func(int, int) bool) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}

	target := nums[0]

	i, j := 0, len(nums)-1
	for i < j {
		// 寻找比 target 小的值, 放到 i 的位置;
		for i < j && !less(nums[j], target) {
			j--
		}
		nums[i] = nums[j]

		// 寻找比 target 大的值, 放到 j 的位置;
		for i < j && less(nums[i], target) {
			i++
		}
		nums[j] = nums[i]
	}

	nums[i] = target

	// 保证:
	// - nums[x] < nums[i], x < i
	// - nums[x] > nums[i], x > i
	quickSort(nums[:i], less)
	quickSort(nums[i+1:], less)
}
