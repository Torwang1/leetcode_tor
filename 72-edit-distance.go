package leetcode

import "math"

type _key struct {
	w1 string
	w2 string
}

var visited = map[_key]int{}

func minDistance(word1 string, word2 string) int {
	k := _key{word1, word2}
	if v, ok := visited[k]; ok {
		return v
	}

	if len(word1) == 0 && len(word2) == 0 {
		return 0
	}

	var I, D, R int
	I = math.MaxInt16
	if len(word2) > 0 {
		I = minDistance(word1[:], word2[1:]) + 1
	}

	D = math.MaxInt16
	if len(word1) > 0 {
		D = minDistance(word1[1:], word2[:]) + 1
	}

	R = math.MaxInt16
	if len(word2) > 0 && len(word1) > 0 {
		R = minDistance(word1[1:], word2[1:])
		if word1[0] != word2[0] {
			R = R + 1
		}
	}

	m := minInt(I, minInt(D, R))

	visited[k] = m

	return m
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
