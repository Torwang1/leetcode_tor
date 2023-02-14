package leetcode

// 基础算法: 二分查找.

func binarySearchII(nums []int, target int) int {

	for i, j := 0, len(nums)-1; i <= j; /**/ {
		mid := (i + j) / 2

		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] < target:
			i = mid + 1
		case nums[mid] > target:
			j = mid - 1
		}
	}

	return -1
}
