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

	quickSort(nums)

	return nums
}

// less 控制排序规则. 例如: 升序、降序.
//
// 其中:
// - rindex 随机下标. 解决: 有序数组导致快排性能下降;
// - allsame 解决: 所有元素相同导致的性能下降; rindex 无法解决这个场景.
//
func quickSort(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}

	rindex := rand.Intn(len(nums))

	nums[0], nums[rindex] = nums[rindex], nums[0]

	target := nums[0]

	allsame := true

	i, j := 0, len(nums)-1
	for i < j {
		// 寻找比 target 小的值, 放到 i 的位置;
		for i < j {
			if nums[j] > target {
				allsame = false
				j--
			} else if nums[j] == target {
				j--
			} else {
				nums[i] = nums[j]
				break
			}
		}

		// 寻找比 target 大的值, 放到 j 的位置;
		for i < j {
			if nums[i] < target {
				allsame = false
				i++
			} else if nums[j] == target {
				i++
			} else {
				nums[j] = nums[i]
				break
			}
		}
	}

	nums[i] = target

	// 保证:
	// - nums[x] < nums[i], x < i
	// - nums[x] > nums[i], x > i

	if allsame {
		return
	}

	quickSort(nums[:i])
	quickSort(nums[i+1:])
}
