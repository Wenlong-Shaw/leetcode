package leetcode

/*
1008. 前序遍历构造二叉搜索树
返回与给定前序遍历 preorder 相匹配的二叉搜索树（binary search tree）的根结点。
(回想一下，二叉搜索树是二叉树的一种，其每个节点都满足以下规则，对于 node.left 的任何后代，值总 < node.val，
	而 node.right 的任何后代，值总 > node.val。
	此外，前序遍历首先显示节点 node 的值，然后遍历 node.left，接着遍历 node.right。）
题目保证，对于给定的测试用例，总能找到满足要求的二叉搜索树。



示例：
输入：[8,5,1,7,10,12]
输出：[8,5,10,1,7,null,12]
*/

//TODO: 思路和题105、106同理。即，划分左右子树，再递归求解。
//* [root,[左子树都比root小], [右子树都比root大]]
func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(preorder); i++ {
		if preorder[i] > preorder[0] {
			break
		}
	}
	root.Left = bstFromPreorder(preorder[1:i])
	root.Right = bstFromPreorder(preorder[i:])
	return root
}
