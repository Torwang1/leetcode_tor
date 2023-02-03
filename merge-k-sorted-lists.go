package leetcode

import "container/heap"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 题目: 合并K个升序链表 -> 升序链表;
//
// 小根堆的处理思路:
// - 所有链表的头部元素, 组成小根堆;
// - 每次弹出最小元素, 并将剩余元素重新入堆;
//
// 时间复杂度: 建堆 O(k), 弹出最小元素 O(n * logk)

func mergeKLists(lists []*ListNode) *ListNode {
	h := make(ListNodeHeap, 0, len(lists))

	// 初始化建堆, 跳过空链表;
	for _, list := range lists {
		if list == nil {
			continue
		}
		h = append(h, list)
	}
	heap.Init(&h)

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

	// 弹出堆顶元素, 剩余链表重新加入堆. 直到堆空为止;
	// heap.Pop() 需要保证至少有一个元素;
	for {
		if len(h) == 0 {
			break
		}

		node := heap.Pop(&h).(*ListNode)
		next := node.Next

		node.Next = nil // 从原始 List 中摘除.
		appendList(node)

		if next != nil {
			heap.Push(&h, next)
		}
	}

	return head
}

// 堆元素需要保存: 1) Val, 2) 剩余链表的指针;
// ListNode 可以直接用做堆元素;

type ListNodeHeap []*ListNode

func (h ListNodeHeap) Len() int           { return len(h) }
func (h ListNodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val } // 小根堆实现
func (h ListNodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push Pop 需要修改 slice 的长度, 而不仅仅是其中的元素. 此处使用 pointer 作为 receiver.
func (h *ListNodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}
func (h *ListNodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
