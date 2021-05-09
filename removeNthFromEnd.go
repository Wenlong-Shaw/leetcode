package leetcode

/*
删除链表的倒数第N个节点
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

进阶：你能尝试使用一趟扫描实现吗？



示例 1：


输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
示例 2：

输入：head = [1], n = 1
输出：[]
示例 3：

输入：head = [1,2], n = 1
输出：[1]


提示：

链表中结点的数目为 sz
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn2925/

*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slowstep, step := 0, 0
	pre, slow, fast := &ListNode{Val: 0, Next: nil}, &ListNode{Val: 0, Next: nil}, head
	for fast != nil {
		step++
		fast = fast.Next
		if step == n {
			slow = head
			slowstep++
		}
		if step > n {
			pre = slow
			slow = slow.Next
			slowstep++
		}
	}
	if slowstep == 1 {
		head = head.Next
	} else {
		pre.Next = slow.Next
	}
	return head
}
