package leetcode

// 删除有序数组中的重复项
//
// 处理方式:
// 1) 下标 i 跟踪待写入的位置;
// 2) 循环中, 找到与 nums[i-1] 不同的数字写入“下标 i”;
//

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 1

	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i-1] {
			nums[i] = nums[j]
			i++
		}
	}

	return i
}
