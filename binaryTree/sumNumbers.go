package leetcode

/*
* 给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
* 每条从根节点到叶节点的路径都代表一个数字：

* 例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
* 计算从根节点到叶节点生成的 所有数字之和 。
* 叶节点 是指没有子节点的节点。


* 示例 1：
* 输入：root = [1,2,3]
* 输出：25
* 解释：
* 从根到叶子节点路径 1->2 代表数字 12
* 从根到叶子节点路径 1->3 代表数字 13
* 因此，数字总和 = 12 + 13 = 25

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sum-root-to-leaf-numbers
*/

import "strconv"

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	var sum func(*TreeNode, string)
	sum = func(node *TreeNode, nodeStr string) {
		if node.Left == nil && node.Right == nil {
			nodeStr += strconv.Itoa(node.Val)
			num, _ := strconv.Atoi(nodeStr)
			ans += num
			return
		}
		nodeStr += strconv.Itoa(node.Val)
		if node.Left != nil {
			sum(node.Left, nodeStr)
		}
		if node.Right != nil {
			sum(node.Right, nodeStr)
		}
	}
	sum(root, "")
	return ans
}
