package leetcode

/*
* 给你二叉树的根节点 root 和一个表示目标和的整数 targetSum ，
* 判断该树中是否存在 根节点到叶子节点 的路径，
* 这条路径上所有节点值相加等于目标和 targetSum 。

* 叶子节点 是指没有子节点的节点。


! 提示：
* 树中节点的数目在范围 [0, 5000] 内
* -1000 <= Node.val <= 1000
* -1000 <= targetSum <= 1000

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/dfs/e8gh3h/
*/
func hasPathSum(root *TreeNode, targetSum int) bool {
	sum := 0

	var depth func(*TreeNode, int) bool
	depth = func(node *TreeNode, sum int) bool {
		if node == nil {
			return false
		}
		sum += node.Val
		if node.Left == nil && node.Right == nil {
			return sum == targetSum
		}
		return depth(node.Left, sum) || depth(node.Right, sum)
	}

	return depth(root, sum)
}
