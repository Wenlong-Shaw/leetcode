package leetcode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var sa, sb []*ListNode
	if headA == nil || headB == nil {
		return &ListNode{Val: 0, Next: nil}
	}
	nodeA, nodeB := headA, headB
	for headA != nil || headB != nil {
		if headA != nil {
			sa = append(sa, headA)
			headA = headA.Next
		}
		if headB != nil {
			sb = append(sb, headB)
			headB = headB.Next
		}
	}
	headA, headB = nodeA, nodeB

	for {
		a := sa[len(sa)-1]
		b := sb[len(sb)-1]
		if a == b && (len(sa)-1 == 0 || len(sb)-1 == 0) {
			return a
		}
		if a != b && a.Next == b.Next {
			return a.Next
		}
		if len(sa)-1 == 0 || len(sb)-1 == 0 {
			return &ListNode{Val: 0, Next: nil}
		}
		sa = sa[:len(sa)-1]
		sb = sb[:len(sb)-1]
	}
}
