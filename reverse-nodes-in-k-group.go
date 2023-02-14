package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 题目: K 个一组, 反转. 最后, 不足的一组保持原有顺序.
//
// 1) k 个节点一组, 拆下来;
// 2) 反转 或 原样输出;
// 3) 挂到 ans 的尾部;
//

func reverseKGroup(head *ListNode, k int) *ListNode {
	// 追加链表段
	var ans, tail *ListNode
	appendList := func(h, t *ListNode) {
		if ans == nil && tail == nil {
			ans = h
			tail = t
		} else {
			tail.Next = h
			tail = t
		}
	}

	// 反转链表
	reverse := func(head *ListNode) (h, t *ListNode) {
		node := head
		for node != nil {
			next := node.Next

			node.Next = nil

			if h == nil && t == nil {
				h = node
				t = node
			} else {
				node.Next = h
				h = node
			}

			node = next
		}
		return
	}

	for head != nil {
		h, t := head, head

		// 拆下来一组节点
		var i int
		for i = 1; i < k; i++ {
			t = t.Next
			if t == nil {
				break
			}
		}
		if t == nil {
			head = nil
		} else {
			head = t.Next
			t.Next = nil
		}

		// 反转或原样输出
		if i == k {
			h, t = reverse(h)
		}

		// 挂到 ans 输出
		appendList(h, t)
	}

	return ans
}
