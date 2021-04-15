package leetcode

/*
提示：

链表中节点的数目在范围 [0, 500] 内
-100 <= Node.val <= 100
0 <= k <= 2 * 10e9
*/

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	if k == 0 {
		return head
	}

	curNode := head
	nodeStack := make([]*ListNode, 0)
	nodeStack = append(nodeStack, curNode)
	l := 1
	for curNode.Next != nil {
		nodeStack = append(nodeStack, curNode.Next)
		curNode = curNode.Next
		l++
	}
	if l == 1 {
		return head
	}
	if step := k % l; step > 0 {
		//断链
		nodeStack[l-1-step].Next = nil
		//换头
		nodeStack[l-1].Next = head
		head = nodeStack[l-step]
	}

	return head
}
