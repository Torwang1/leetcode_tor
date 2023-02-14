package leetcode

import (
	"strconv"
	"strings"
)

// 题目: 比较版本号
//
// 工程算法
// 1) sting -> []int;
// 2) 比较两个 []int;
//

func compareVersion(version1 string, version2 string) int {

	intSlice := func(str string) []int {
		var result []int

		for _, s := range strings.Split(str, ".") {
			s = strings.TrimLeft(s, "0")
			if s == "" {
				s = "0"
			}
			i, _ := strconv.Atoi(s)
			result = append(result, i)
		}

		return result
	}

	slice1 := intSlice(version1)
	slice2 := intSlice(version2)

	// 填充为相同长度;
	length := len(slice1)
	if length < len(slice2) {
		length = len(slice2)
	}
	for i, n := 0, length-len(slice1); i < n; i++ {
		slice1 = append(slice1, 0)
	}
	for i, n := 0, length-len(slice2); i < n; i++ {
		slice2 = append(slice2, 0)
	}

	// 比较版本号
	for i := 0; i < length; i++ {
		switch {
		case slice1[i] == slice2[i]:
			continue
		case slice1[i] > slice2[i]:
			return 1
		case slice2[i] > slice1[i]:
			return -1
		}
	}

	return 0
}
