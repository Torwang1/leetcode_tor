package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 题目: 相交链表;
//
// 双指针
//
// 1) 只有当 headA 和 headB 都不为空时, 才可能相交;
//
// 当 headA 和 headB 都不为空时. pA pB 分别指向 headA headB 链表头部, 依次遍历链表:
// 1) 如果 pA == pB || (pA == nil && pB == nil) 结束循环;
// 2) 如果 pA != nil, pA 指向下一个节点. pB 同理;
// 3) 如果 pA == nil, pA 指向 headB; 同理, 如果 pB == nil, pB 指向 headA;
//

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	pA, pB := headA, headB
	for {
		// 相交
		if pA == pB {
			return pA
		}

		pA, pB = pA.Next, pB.Next

		// 不相交
		if pA == nil && pB == nil {
			return nil
		}

		if pA == nil {
			pA = headB
		}

		if pB == nil {
			pB = headA
		}
	}
}
