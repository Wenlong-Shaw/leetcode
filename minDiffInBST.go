package leetcode

import (
	"math"
)

/*
* 783. 二叉搜索树节点最小距离
* 给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。

! 注意：本题与 530：https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/ 相同

! key point: 二叉搜索树有个性质为：二叉搜索树中序遍历得到的值序列是递增有序的.
*/

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

func MinDiffInBST(root *TreeNode) int {
	nodeVal := []int{}
	ans := math.MaxInt64
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		traverse(node.Left)
		nodeVal = append(nodeVal, node.Val)
		traverse(node.Right)
	}
	traverse(root)
	// sort.Ints(nodeVal)
	for i := range nodeVal {
		if i < len(nodeVal)-1 {
			minal := nodeVal[i+1] - nodeVal[i]
			ans = min(ans, minal)
		}
	}
	return ans
}
