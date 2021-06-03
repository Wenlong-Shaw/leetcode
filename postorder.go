package leetcode

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
/*
590. N 叉树的后序遍历
给定一个 N 叉树，返回其节点值的 后序遍历 。

N 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔（请参见示例）。
进阶：
递归法很简单，你可以使用迭代法完成此题吗?

示例 1：
输入：root = [1,null,3,2,4,null,5,6]
输出：[5,6,3,2,4,1]
*/
type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	ans := []int{}
	var post func(*Node)
	post = func(node *Node) {
		if node == nil {
			return
		}
		for i := range node.Children {
			post(node.Children[i])
		}
		ans = append(ans, node.Val)
	}
	post(root)
	return ans
}

//非递归，迭代解法
func postorder_1(root *Node) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	list := make([]*Node, 1)
	list[0] = root
	for len(list) > 0 {
		node := list[len(list)-1]
		list = list[:len(list)-1]
		if node != nil {
			ans = append(ans, node.Val)
		}
		for i := range node.Children {
			list = append(list, node.Children[i])
		}
	}
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-i-1] = ans[len(ans)-i-1], ans[i]
	}
	return ans
}
