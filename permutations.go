package leetcode

// 题目: 全排列
//
// 特征:
// 1) 所有可能的解;
//

func permute(nums []int) [][]int {
	ans := [][]int{}

	backtrack(
		&ans,
		nums,
		make([]bool, len(nums)),
		make([]int, 0, len(nums)),
	)

	return ans
}

func backtrack(ans *[][]int, nums []int, visited []bool, record []int) {
	nextFunc := func() []int {
		indexes := make([]int, 0)
		for i := range visited {
			if !visited[i] {
				indexes = append(indexes, i)
			}
		}
		return indexes
	}

	// 使用下标.
	next := nextFunc()

	if len(next) == 0 {
		dst := make([]int, len(record))
		copy(dst, record)
		*ans = append(*ans, dst)
	}

	for _, i := range next {
		record = append(record, nums[i])
		visited[i] = true

		backtrack(ans, nums, visited, record)

		record = record[:len(record)-1]
		visited[i] = false
	}
}
