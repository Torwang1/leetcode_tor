package leetcode

// 题目: 有序数组上, 求解两数之和.
//
// 条件: a + b = target
//
// 特殊:
// - 下标从 1 开始.
//
func twoSum2(numbers []int, target int) []int {

	right := len(numbers) - 1

	for left := 0; left < len(numbers); left++ {
		// 非递减数组, 快速跳过重复值.
		if left > 0 && numbers[left] == numbers[left-1] {
			continue
		}

		// 向左移动 right 指针, 直到: 1) left == right 或者 2) a + b <= target
		for right > left && numbers[right]+numbers[left] > target {
			right--
		}

		if right == left {
			break
		}

		if numbers[right]+numbers[left] == target {
			return []int{left + 1, right + 1}
		}
	}

	return nil
}
