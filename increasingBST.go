package leetcode

/*
* 897. 递增顺序搜索树
* 给你一棵二叉搜索树，请你 按中序遍历 将其重新排列为一棵递增顺序搜索树，
* 使树中最左边的节点成为树的根节点，并且每个节点没有左子节点，只有一个右子节点。

* 提示：
* 树中节点数的取值范围是 [1, 100]
* 0 <= Node.val <= 1000

 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
	nodeList, curNode := []*TreeNode{}, &TreeNode{}
	var traversed func(*TreeNode)
	traversed = func(root *TreeNode) {
		if root == nil {
			return
		}
		traversed(root.Left)
		nodeList = append(nodeList, root)
		traversed(root.Right)
		return
	}
	traversed(root)
	for i := range nodeList {
		if i >= len(nodeList)-1 {
			nodeList[i].Right = nil
			nodeList[i].Left = nil
			return nodeList[0]
		}
		curNode = nodeList[i]
		curNode.Left, curNode.Right = nil, nodeList[i+1]
	}
	return nodeList[0]
}
