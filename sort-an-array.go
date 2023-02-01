package leetcode

import (
	"math/rand"
	"time"
)

// 快速排序 quicksort
//
// 优化点:
// - 有序场景: target 使用随机下标;
// - 相同元素: quickSort 记录是否有序;
//
func sortArray(nums []int) []int {
	rand.Seed(time.Now().Unix())

	if len(nums) > 1 {
		quickSort(nums, func(i1, i2 int) bool { return i1 < i2 })
	}

	return nums
}

// less 控制排序规则. 例如: 升序、降序.
//
// 其中:
// - rindex 随机下标. 解决: 有序数组导致快排性能下降;
func quickSort(nums []int, less func(int, int) bool) {
	rindex := rand.Intn(len(nums))

	nums[0], nums[rindex] = nums[rindex], nums[0]

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

	if slice := nums[:i]; len(slice) > 1 {
		quickSort(slice, less)
	}
	if slice := nums[i+1:]; len(slice) > 1 {
		quickSort(nums[i+1:], less)
	}
}
