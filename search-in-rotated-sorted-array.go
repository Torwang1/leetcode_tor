package leetcode

// 搜索旋转后的排序数组(升序)
//

func searchII(nums []int, target int) int {

	i, j := 0, len(nums)-1

	for i <= j {
		mid := (i + j) / 2

		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] > target:
			if nums[i] <= target || nums[i] > nums[mid] {
				j = mid - 1
			} else {
				i = mid + 1
			}
		case nums[mid] < target:
			if nums[j] >= target || nums[j] < nums[mid] {
				i = mid + 1
			} else {
				j = mid - 1
			}
		}
	}

	return -1
}
