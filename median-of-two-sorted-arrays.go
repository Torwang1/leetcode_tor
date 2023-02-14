package leetcode

// 题目: 寻找两个正序数组的中位数
//
// 二分查找, O(logn)
//
// 中位数:
// 1) 设两个数组的长度, 分别为: m, n;
// 2) 如果 m+n 是奇数, 中位数 := (m+n)/2;
// 3) 如果 m+n 是偶数, 中位数 := (m+n)/2, (m+n)/2-1 的平均值;
//
// 问题转化为: 在两个正序数组中, 查找第k小的数字;
//
// 目标: 第k小的数字, 比较 A[k/2-1]、B[k/2-1], 其中:
// 1) 小于 A[k/2-1] 的元素个数 := k/2 - 1;
// 2) 小于 B[k/2-1] 的元素个数 := k/2 - 1;
// 3) 对于 A[k/2-1]、B[k/2-1] 中的较小者, 只会有 (k/2 -1) + (k/2 - 1) <= k-2 个元素;
//    所以, 比‘较小者’小的元素 一定不是第 k 小的数字;
//
// 推理:
// 1) 如果 A[k/2-1] < B[k/2-1], 比 A[k/2-1] 小的元素最多只有 k-2 个. 因此, A[k/2-1] 不
//    可能是 第k小的数字. 所以, A[0 .. k/2-1] 可以全部排除.  (注意: A[k/2-1] 也可被排除)
// 2) 如果 B[k/2-1] < A[k/2-1], 同理, B[0 .. k/2-1] 可以全部排除;
// 3) 如果 A[k/2-1] = B[k/2-1], 可以归入 1) 处理;
// 4) 其中: k = k - k/2;
//
// 特殊情况:
// 1) 如果 A[k/2-1] 或 A[k/2-1] 越界, 取数组的最后一个元素, 并根据实际调整 k 值;
// 2) 其中一个数组为空, 则返回另外一个数组的第k个元素;
// 3) 如果 k = 1, 返回两个数组头部的最小值;
//
// 推理示例:
// 1) k = 5, k/2 = 2
// 2) k = 3, k/2 = 1
// 3) k = 2, k/2 = 1
// 4) k = 1, <== 结束条件;
//

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)

	count := m + n

	if count%2 == 1 {
		return float64(findMedianSortedArraysKth(nums1, nums2, count/2+1))
	} else {
		return (float64(findMedianSortedArraysKth(nums1, nums2, count/2)) +
			float64(findMedianSortedArraysKth(nums1, nums2, count/2+1))) /
			2
	}
}

// 在两个正序数组中, 查找第 k 小的数字.
func findMedianSortedArraysKth(nums1 []int, nums2 []int, kth int) int {
	// 2) 其中一个数组为空, 则返回另外一个数组的第k个元素;
	if len(nums1) == 0 || len(nums2) == 0 {
		if len(nums1) != 0 {
			return nums1[kth-1]
		} else {
			return nums2[kth-1]
		}
	}

	// 3) 如果 k = 1, 返回两个数组头部的最小值;
	if kth == 1 {
		if nums1[0] < nums2[0] {
			return nums1[0]
		} else {
			return nums2[0]
		}
	}

	// 推理
	i1 := kth/2 - 1
	if i1 >= len(nums1) {
		i1 = len(nums1) - 1
	}

	i2 := kth/2 - 1
	if i2 >= len(nums2) {
		i2 = len(nums2) - 1
	}

	if nums1[i1] <= nums2[i2] {
		return findMedianSortedArraysKth(nums1[i1+1:], nums2, kth-i1-1)
	} else {
		return findMedianSortedArraysKth(nums1, nums2[i2+1:], kth-i2-1)
	}
}
