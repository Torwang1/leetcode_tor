package leetcode

// 题目: 最短的无序子数组
//
// 基础算法: 单调栈
//
// 示例:
// 2, 6, 4, 8, 10, 9, 15
// 2, 4, 6, 8, 9, 10, 15
//
// 步骤记录:
// 1) left
//   a) nums[0] 压栈;
//   b) 向右, 找单调递增;
//   c) 继续向右, 弹出 cur < top 的所有元素;
// 2) right
//   a) nums[len(nums)-1] 压栈;
//   b) 向左, 找单调递减;
//   c) 继续向左, 弹出 cur > top 的所有元素;
// 3) 无序子数组的最小长度
//   a) length := len(nums) - len(left) - len(right)
//
// 特殊 case:
// 1) 有序序列: 1, 2, 3, 4, 5
//
// 两次扫描逻辑, 时间复杂度 O(n)

func findUnsortedSubarray(nums []int) int {

	left := []int{0}
	var i int
	for i = 1; i < len(nums); i++ {
		top := nums[left[len(left)-1]]
		if nums[i] >= top {
			left = append(left, i)
			continue
		}
		break
	}
	for i = i; i < len(nums); i++ {
		for len(left) > 0 && nums[i] < nums[left[len(left)-1]] {
			left = left[:len(left)-1]
		}
	}

	right := []int{len(nums) - 1}
	var j int
	for j = len(nums) - 2; j >= 0; j-- {
		top := nums[right[len(right)-1]]
		if nums[j] <= top {
			right = append(right, j)
			continue
		}
		break
	}
	for j = j; j >= 0; j-- {
		for len(right) > 0 && nums[j] > nums[right[len(right)-1]] {
			right = right[:len(right)-1]
		}
	}

	// 特殊 case: 有序序列
	if len(right) == len(nums) && len(left) == len(nums) {
		return 0
	}

	return len(nums) - len(left) - len(right)
}
