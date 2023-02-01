package leetcode

// 题目: 两数之和
//
// 约束:
// - 输出下标.  即: 数组不可以被排序;
// - 只会有一个有效答案;
//
// 解法:
// - hash. 时间复杂度: O(n)
func twoSum(nums []int, target int) []int {

	// value to index list
	v2index := map[int]*[]int{}
	for i, v := range nums {
		slice := v2index[v]
		if slice == nil {
			slice = &[]int{}
			v2index[v] = slice
		}

		*slice = append(*slice, i)
	}

	// 逻辑上保证 set 内的元素不重复.
	// 返回值: 防止修改了 map 中的 slice.
	exclude := func(set *[]int, index int) []int {
		if set == nil {
			return nil
		}

		result := make([]int, 0, len(*set))
		for _, v := range *set {
			if v != index {
				result = append(result, v)
			}
		}
		return result
	}

	for i, a := range nums {
		b := target - a

		set := exclude(v2index[b], i)

		// 题目保证: 最多有一个解, 即 set 最多有一个元素.
		if len(set) == 1 {
			return []int{i, (set)[0]}
		}
	}

	return nil
}
