package leetcode

/*ListNode
203. 移除链表元素
给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。

示例 1
输入：head = [1,2,6,3,4,5,6], val = 6
输出：[1,2,3,4,5]

示例 2：
输入：head = [], val = 1
输出：[]

示例 3：
输入：head = [7,7,7,7], val = 7
输出：[]


提示：
列表中的节点在范围 [0, 10^4] 内
1 <= Node.val <= 50
0 <= k <= 50

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-linked-list-elements
*/

func removeElements(head *ListNode, val int) *ListNode {
	if val == 0 || head == nil {
		return head
	}
	pre := &ListNode{0, head}
	tmp := &ListNode{}
	tmp1 := tmp
	i := 0
	for ; head != nil; head = head.Next {
		if i == 0 && head.Val != val {
			tmp = head
			i++
		}
		if head.Val == val {
			pre.Next = head.Next
			continue
		}
		pre = head
	}
	if tmp1 == tmp {
		return nil
	}
	return tmp
}
