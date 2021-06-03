package leetcode

/*
* 练习：相同的树
* 给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
* 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/dfs/e8z1um/
*/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return false
}
