package leetcode

/*
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	num1, num2 := []int{}, []int{}
	for l1 != nil || l2 != nil {
		if l1 != nil {
			num1 = append(num1, l1.Val)
			l1 = l1.Next
		}
		if l2 != nil {
			num2 = append(num2, l2.Val)
			l2 = l2.Next
		}
	}
	n, n1, n2, p := 0, 0, 0, 1
	if len(num1) > len(num2) {
		n = len(num1)
	} else {
		n = len(num2)
	}
	for i := 1; i <= n; i++ {
		if len(num1)-i >= 0 {
			n1 += num1[n-i] * p
		}
		if len(num2)-i >= 0 {
			n2 += num2[n-i] * p
		}
		p *= 10
	}
	sum := n1 + n2
	cur := &ListNode{Val: 0, Next: nil}
	for sum > 0 {
		cur.Next = &ListNode{Val: sum % 10, Next: cur.Next}
		sum /= 10
	}
	return cur.Next
}
*/

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	var s1, s2, s3 Stack
	cur, quo := 0, 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			s1.Push(l1.Val)
			l1 = l1.Next
		}
		if l2 != nil {
			s2.Push(l2.Val)
			l2 = l2.Next
		}
	}
	n := 0
	if s1.Len() > s2.Len() {
		n = s1.Len()
	} else {
		n = s2.Len()
	}
	for n > 0 {
		if s1.Len() != 0 && s2.Len() != 0 {
			cur = (s1.Top()+s2.Top())%10 + quo
			quo = (s1.Pop() + s2.Pop()) / 10
			s3.Push(cur)
		}
		if s1.Len() == 0 && s2.Len() == 0 {
			cur = quo
			if cur == 0 {
				break
			}
			s3.Push(cur)
			break
		}
		if s1.Len() == 0 {
			cur = (0+s2.Top())%10 + quo
			quo = (0 + s2.Pop()) / 10
			s3.Push(cur)
		}
		if s2.Len() == 0 {
			cur = (s1.Top()+0)%10 + quo
			quo = (s1.Pop() + 0) / 10
			s3.Push(cur)
		}
	}
	var c *ListNode
	r := c
	for s3.Len() != 0 {
		p := &ListNode{}
		p.Val = s3.Pop()
		r.Next = p
		r = p
	}
	r.Next = nil
	return c.Next
}
