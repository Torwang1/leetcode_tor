package leetcode

// 题目: 三数之和
//
// 限制:
// - 不重复. 所以, 排序 + 跳过重复元素;
// - b + c = a. 所以, b,c 在有序数组上, 双指针, 相向而行. 复杂度 O(n)

import "sort"

func threeSum(nums []int) [][]int {

	sort.Ints(nums)

	ans := [][]int{}

	for first := 0; first < len(nums); first++ {
		// 去除重复数据.
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		target := -1 * nums[first]

		third := len(nums) - 1

		for second := first + 1; second < len(nums); second++ {
			// 去除重复数据.
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			// 快速移动 third 指针, 直到: 1) third == second 或者 2) a + b <= c
			for third > second && nums[second]+nums[third] > target {
				third--
			}

			// 题目要求: 指针不能重合.
			if second == third {
				break
			}

			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}

	return ans
}
