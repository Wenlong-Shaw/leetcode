package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//二叉搜索树的中序遍历求区间值的和。
func rangeSumBST(root *TreeNode, low int, high int) int {
	result := 0
	var traversed func(node *TreeNode)
	traversed = func(node *TreeNode) {
		if node != nil {
			traversed(node.Left)
			if node.Val >= low && node.Val <= high {
				result += node.Val
			}
			traversed(node.Right)
		}
	}
	traversed(root)
	return result
}
