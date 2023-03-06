package leetcode

// 下一个排列
//
// nums 可以表示为: xxxa[降序部分]
//
// 处理逻辑:
// 1) 先找出[降序部分], 获得“较小数字 a”;
// 2) 在降序部分中找到第一个比 a 大的数字, 并交换;
// 3) revert 降序部分;
//

func nextPermutation(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}

	// 原地反转数组
	var revert func(nums []int)
	revert = func(nums []int) {
		for i, j := 0, len(nums)-1; i <= j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	// 1) 先找出[降序部分], 获得“较小数字 a”;
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	// 跳出循环有两种情况:
	// 1) i < 0: nums 完全降序;
	// 2) nums[i] < nums[i+1], i 是“较小数字”;

	// 2) 在降序部分中找到第一个比 a 大的数字, 并交换;
	if i >= 0 {
		j := i
		for j+1 < len(nums) && nums[j+1] > nums[i] {
			j++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	revert(nums[i+1:])
}
