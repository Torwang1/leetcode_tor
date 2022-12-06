package leetcode

import "sort"

// “有重复字符”的排列处理思路:
// - 基础思路: 回溯; (类似: 无重复字符)
// - 保证“重复字符的出现顺序”, 避免重复结果: “qqe”: 第二个 q 出现的条件是, 第一个 q 已经出现;
func permutationII(S string) []string {

	result := []string{}

	// 有序数组: 保证相同字符连续出现.
	chars := []byte(S)
	sort.Slice(chars, func(i int, j int) bool { return chars[i] < chars[j] })

	arr := make([]byte, 0, len(S))

	// 有重复字节, 使用“数组下标”去重.
	charsVisited := make(map[int]bool)

	sequenceII(&result, &arr, chars, charsVisited)

	return result
}

// sequenceII ...
//
// arr: 状态数组, 记录递归过程选中的 字节列表;
// chars: 输入参数, 排序后的 S 字节数组;
func sequenceII(result *[]string, arr *[]byte, chars []byte, charsVisited map[int]bool) {
	candidates := candidateBytesII(chars, charsVisited)

	if len(candidates) == 0 {
		*result = append(*result, string(*arr))
		return
	}

	for _, i := range candidates {
		charsVisited[i] = true
		*arr = append(*arr, chars[i])
		sequenceII(result, arr, chars, charsVisited)
		charsVisited[i] = false
		*arr = (*arr)[:len(*arr)-1]
	}
}

// candidateBytesII 未被访问 & 保证重复字符的出现顺序.
// 返回: 数组下标集合.
func candidateBytesII(chars []byte, charsVisited map[int]bool) []int {
	result := []int{}

	// 不属于备选字符的条件:
	// - char[i] 被访问过;
	// - i>=1 && char[i] == char[i-1] && char[i-1] 未被访问过;
	for i := range chars {
		if charsVisited[i] || (i > 0 && chars[i] == chars[i-1] && charsVisited[i-1]) {
			continue
		}
		result = append(result, i)
	}

	// 属于备选字符的条件:
	// i == 0: i 未被访问过;
	// i >= 1: 如果 chars[i-1]== chars[i], chars[i-1]被访问过 && char[i]未被访问过;
	// for i := range chars {
	// 	switch i {
	// 	case 0:
	// 		if !charsVisited[i] {
	// 			result = append(result, i)
	// 		}
	// 	default:
	// 		if chars[i-1] == chars[i] {
	// 			if charsVisited[i-1] && !charsVisited[i] {
	// 				result = append(result, i)
	// 			}
	// 		} else {
	// 			if !charsVisited[i] {
	// 				result = append(result, i)
	// 			}
	// 		}
	// 	}
	// }

	return result
}
