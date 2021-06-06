package leetcode

/*
1028. 从先序遍历还原二叉树
我们从二叉树的根节点 root 开始进行深度优先搜索。
在遍历中的每个节点处，我们输出 D 条短划线（其中 D 是该节点的深度），然后输出该节点的值。
（如果节点的深度为 D，则其直接子节点的深度为 D + 1。根节点的深度为 0）。
如果节点只有一个子节点，那么保证该子节点为左子节点。
给出遍历输出 S，还原树并返回其根节点 root。

提示：
原始树中的节点数介于 1 和 1000 之间。
每个节点的值介于 1 和 10 ^ 9 之间。

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/recover-a-tree-from-preorder-traversal/solution/cong-xian-xu-bian-li-huan-yuan-er-cha-shu-by-leetc/
*/
func recoverFromPreorder(traversal string) *TreeNode {
	path, pos := []*TreeNode{}, 0
	for pos < len(traversal) {
		level := 0
		for traversal[pos] == '-' {
			level++
			pos++
		}
		//* 每次收到数字都要先置零value
		value := 0
		//* 计算value的值
		for ; pos < len(traversal) && traversal[pos] >= '0' && traversal[pos] <= '9'; pos++ {
			value = value*10 + int(traversal[pos]-'0')
		}
		node := &TreeNode{Val: value}
		//! 根据当前的level选择父节点
		//TODO: 前序遍历，先左后右
		if level == len(path) {
			if len(path) > 0 {
				path[len(path)-1].Left = node
			}
		} else {
			path = path[:level]
			path[len(path)-1].Right = node
		}
		path = append(path, node)
	}
	return path[0]
}
