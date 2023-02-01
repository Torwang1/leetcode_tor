package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 题目: K 个一组, 反转. 最后, 不足的一组保持原有顺序.
func reverseKGroup(head *ListNode, k int) *ListNode {
	enode := &ListNode{}

	// 辅助节点.
	enode.Next = head
	head = enode

	tail := head
	next := tail.Next // 游标, 指向第一个有效元素.

	for {
		var i int
		for i = 0; i < k && next != nil; i++ {
			next = next.Next
		}
		if i < k {
			break // 最后一组不满足 K 个元素, 保持原有顺序.
		}

		var tmp *ListNode // 临时 header

		for node := tail.Next; node != next; /**/ {
			n := node.Next

			if tmp == nil {
				tmp = node
				node.Next = next
			} else {
				node.Next = tmp
				tmp = node
			}

			node = n
		}

		tail.Next = tmp

		// tail 指向 next 的前序节点.
		for i := 0; i < k; i++ {
			tail = tail.Next
		}
	}

	return head.Next
}
