package leetcode

/*
翻转一棵二叉树。

示例：

输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/dfs/e8gbt3/

*/
func invertTree(root *TreeNode) *TreeNode {
	var invert func(*TreeNode)
	invert = func(node *TreeNode) {
		if node == nil {
			return
		}
		node.Left, node.Right = node.Right, node.Left
		invert(node.Left)
		invert(node.Right)
	}
	invert(root)
	return root
}
