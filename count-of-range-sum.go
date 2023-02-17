package leetcode

import "math"

// 区间和的个数. 前缀和 + 线段树
//
// 动态开点线段树实现
//

func countRangeSum(nums []int, lower int, upper int) int {

	min := func(a ...int) int {
		res := a[0]
		for _, v := range a {
			if v < res {
				res = v
			}
		}
		return res
	}

	max := func(a ...int) int {
		res := a[0]
		for _, v := range a {
			if v > res {
				res = v
			}
		}

		return res
	}

	presum := make([]int, len(nums)+1)
	for i, v := range nums {
		presum[i+1] = presum[i] + v
	}

	lBound, rBound := math.MaxInt, math.MinInt
	for _, sum := range presum {
		lBound = min(lBound, sum, sum-lower, sum-upper)
		rBound = max(rBound, sum, sum-lower, sum-upper)
	}

	var ans int
	root := &segNode{lBound: lBound, rBound: rBound}
	for _, sum := range presum {
		left, right := sum-upper, sum-lower
		ans += root.query(left, right)
		root.insert(sum)
	}

	return ans
}

type segNode struct {
	lBound, rBound int // 左右边界
	count          int // 出现次数

	left, right *segNode
}

// insert 增加 value 出现的次数.
//
// value 对应的叶子节点可能还不存在, 需要动态增加.
func (node *segNode) insert(value int) {
	node.count++

	if node.lBound == node.rBound {
		return
	}

	mid := node.lBound + (node.rBound-node.lBound)/2
	if value <= mid {
		if node.left == nil {
			node.left = &segNode{lBound: node.lBound, rBound: mid}
		}
		node.left.insert(value)
	} else {
		if node.right == nil {
			node.right = &segNode{lBound: mid + 1, rBound: node.rBound}
		}
		node.right.insert(value)
	}
}

// query value 落在 [l,r] 区间的次数.
func (node *segNode) query(l, r int) int {
	if node == nil || l > node.rBound || r < node.lBound {
		return 0
	}

	if l <= node.lBound && node.rBound <= r {
		return node.count
	}
	return node.left.query(l, r) + node.right.query(l, r)
}
