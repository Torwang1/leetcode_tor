package leetcode

// 合并两个升序链表

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head, tail *ListNode

	// 从原始链表中摘除节点, 清理节点 next.
	var remove func(list **ListNode) *ListNode
	remove = func(list **ListNode) *ListNode {
		node := *list
		*list = node.Next
		node.Next = nil
		return node
	}

	// 追加到结果链表尾部.
	var append func(node *ListNode)
	append = func(node *ListNode) {
		if head == nil && tail == nil {
			head = node
			tail = node
			return
		}
		tail.Next = node
		tail = node
	}

	for list1 != nil && list2 != nil {
		var node *ListNode

		if list1.Val <= list2.Val {
			node = remove(&list1)
		} else {
			node = remove(&list2)
		}

		append(node)
	}

	if list1 != nil {
		append(list1)
	}
	if list2 != nil {
		append(list2)
	}

	return head
}
