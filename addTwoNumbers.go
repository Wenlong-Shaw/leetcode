package leetcode

/*
* 2. 两数相加
* 给你两个 非空 的链表，表示两个非负的整数。
* 它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
* 请你将两个数相加，并以相同形式返回一个表示和的链表。

* 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
* 示例 1：
* 输入：l1 = [2,4,3], l2 = [5,6,4]
* 输出：[7,0,8]
* 解释：342 + 465 = 807.

* 示例 2：
* 输入：l1 = [0], l2 = [0]
* 输出：[0]

* 示例 3：
* 输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
* 输出：[8,9,9,9,0,0,0,1]

! 提示：
* 每个链表中的节点数在范围 [1, 100] 内
* 0 <= Node.val <= 9
* 题目数据保证列表表示的数字不含前导零

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//直接相加，但是会溢出。因为节点数在范围 [1, 100] 内，所以下面方法不行。
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var v1, v2 int
	l := []int{}
	cur := &ListNode{}
	i := 0
	for l1 != nil {
		v1 += l1.Val * pow10(i)
		l1 = l1.Next
		i++
	}
	i = 0
	for l2 != nil {
		v2 += l2.Val * pow10(i)
		l2 = l2.Next
		i++
	}
	res := v1 + v2
	if res < 10 {
		return &ListNode{res, nil}
	}
	result := cur
	for res != 0 {
		n := res % 10
		next := &ListNode{}
		l = append(l, n)
		cur.Val = n
		cur.Next = next
		if res/10 == 0 {
			break
		}
		next, cur = cur, next
		res /= 10
	}
	cur.Next = nil
	return result
}

func pow10(y int) int {
	x := 1
	if y == 0 {
		return 1
	}
	for i := 0; i < y; i++ {
		x *= 10
	}
	return x
}
