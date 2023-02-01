package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var ans *ListNode

	for elem := head; elem != nil; /**/ {
		next := elem.Next

		elem.Next = ans
		ans = elem

		elem = next
	}

	return ans
}
