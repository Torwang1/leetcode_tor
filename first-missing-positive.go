package leetcode

import "math"

// 题目: 缺失一个正整数
//
// 通俗解法:
// 1) nums 构造 hash 查找结构;
// 2) 从 1 开始查找未出现的元素;
// 最多 n 次循环. 分析:
// 1) nums 最多容纳 n 个元素;
// 2) 如果 nums 容纳 [1..N], 则第一个缺失的数字是 N;
// 3) 如果 nums 容纳了 N+x , 则 [1..N] 中必然有一个数字缺失;
//
// 题目约束:
// 1) 只能使用“常数级别”的额外空间. 所以, 利用 nums 构造 hash 查找结构.
// 构造方法:
// 1) 遍历数组: 题目中仅关注 [1..N], 所有“0”,“负数”和“>N的数”都设置为 math.MaxInt; <== 所有元素都是 > 0;
// 2) 遍历数组: 所有出现的, 在[1..N]中的数字, 将对应下标的元素 * -1; <== |x| 可以还原数值; <== 注意下标需要左移动 1 位;
// 3) 遍历元素: 第一个正数下标, 就是题解, i+1;
//
// 注意:
// 1) 题目关心的是 [1..N], 不包含 0.

func firstMissingPositive(nums []int) int {
	const sentry = math.MaxInt

	abs := func(x int) int {
		if x < 0 {
			return -1 * x
		}
		return x
	}

	// 1) 遍历数组: 题目中仅关注 [1..N], 所有“0”、“负数”和“>N的数”都设置为 math.MaxInt; <== 所有元素都是 > 0;
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] > n || nums[i] <= 0 {
			nums[i] = sentry
		}
	}

	// 2) 遍历数组: 所有出现的, 在[1..N]中的数字, 将对应下标的元素 * -1; <== |x| 可以还原数值;
	//注意: 下标需要左移动 1 位;
	for _, v := range nums {
		if v = abs(v); v != sentry {
			index := v - 1
			nums[index] = -1 * abs(nums[index])
		}
	}

	// 3) 遍历元素: 第一个正数下标, 就是题解, i+1;
	for i, v := range nums {
		if v > 0 {
			return i + 1
		}
	}
	return n + 1
}
