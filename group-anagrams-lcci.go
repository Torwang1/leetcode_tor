package leetcode

import "sort"

// 异位词, 桶排序.
//
// key 策略:
// - 排序字符串;
// - 词频统计结构. 只有小写字母, 26个字段;
func groupAnagrams(strs []string) [][]string {
	bucket := map[string][]string{}

	keyFunc := func(str string) string {
		key := []byte(str)
		sort.Slice(key, func(i, j int) bool { return key[i] < key[j] })
		return string(key)
	}

	for _, str := range strs {
		key := keyFunc(str)
		bucket[key] = append(bucket[key], str)
	}

	result := make([][]string, 0, len(bucket))
	for _, v := range bucket {
		result = append(result, v)
	}

	return result
}
