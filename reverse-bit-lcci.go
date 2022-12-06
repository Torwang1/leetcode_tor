package leetcode

func reverseBits(num int) int {

	// 当前位置, 连续 1 的个数;
	cur := 0

	// 当前位置为 0, 转换成 1 后, 连续 1 的个数;
	insert := 0

	// 最大连续 1 的个数;
	result := 0

	for i := 0; i < 32; i++ {
		if num&(1<<i) == 1 {
			insert = insert + 1
			cur = cur + 1
		} else {
			insert = cur + 1
			cur = 0
		}

		result = max(result, insert)

	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
