package leetcode

// 题目: 重排链表
//
// 通俗解法:
// 1) 拆成 l1 l2 两个链表. 其中, l2 倒序;
//
// 注意:
// 1) len(l1) >= len(l2);

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	// 获取链表长度;
	var n int
	for p := head; p != nil; /**/ {
		n++
		p = p.Next
	}

	// 拆分成 l1 和 l2.
	l1, l2 := spliteList(head, n)

	// 合并;
	head = mergeList(l1, l2)
}

// spliteList 按题目要求拆分链表. 其中:
// 1) l1 正序;
// 2) l2 逆序;
// 3) len(l1) == len(l2) 或者 len(l1) == len(l2) + 1;
func spliteList(head *ListNode, n int) (*ListNode, *ListNode) {
	len1 := (n + 1) / 2

	l1 := head
	for i := 0; i < len1-1; i++ { // 少循环 1 次;
		head = head.Next
	}
	node := head.Next // 断开 l1 和 head;
	head.Next = nil
	head = node

	var l2 *ListNode
	for head != nil {
		node := head // 摘除节点;
		head = head.Next
		node.Next = nil

		if l2 == nil { // 插入节点
			l2 = node
		} else {
			node.Next = l2
			l2 = node
		}
	}

	return l1, l2
}

// merge 合并两个链表. 调用方保证: len(l1) == len(l2) 或者 len(l1) == len(l2) + 1;
func mergeList(l1, l2 *ListNode) *ListNode {

	var head, tail *ListNode

	appendList := func(node *ListNode) {
		if head == nil && tail == nil {
			head = node
			tail = node
			return
		}

		tail.Next = node
		tail = node
	}

	// 上下文约束: l2 短;
	for l2 != nil {
		var node *ListNode

		node = l1 // 摘除节点
		l1 = l1.Next
		node.Next = nil
		appendList(node)

		node = l2 // 摘除节点
		l2 = l2.Next
		node.Next = nil
		appendList(node)
	}

	if l1 != nil {
		appendList(l1)
	}

	return head
}
