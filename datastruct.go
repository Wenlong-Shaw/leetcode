package leetcode

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//BinaryTreeConstructor BFS 构建完全二叉树
func BinaryTreeConstructor(nums []int) *TreeNode {
	numsLen := len(nums)
	// nilnode := new(TreeNode)
	// nilnode.Val = -1
	// fmt.Println(nilnode)
	if numsLen < 1 {
		return nil
	}
	var curNode *TreeNode
	nodeArr := make([]*TreeNode, 0)
	for i := range nums {
		curNode = &TreeNode{
			Val: nums[i],
		}
		nodeArr = append(nodeArr, curNode)
	}
	// nodeArr = append(nodeArr, nilnode)
	for i := 1; i <= (numsLen / 2); i++ {
		if 2*i <= numsLen {
			nodeArr[i-1].Left = nodeArr[2*i-1]
		}
		if 2*i+1 <= numsLen {
			nodeArr[i-1].Right = nodeArr[2*i]
		}
	}
	root := nodeArr[0]
	return root
}

func BinaryTreeTraverse(node *TreeNode) {
	defer func() {
		if r := recover(); r != nil {
			// fmt.Println("Panic:", r)
		}
	}()
	if node == nil {
		return
	}
	// 前序遍历
	// visited(node)
	BinaryTreeTraverse(node.Left)
	// 中序遍历
	visited(node)
	BinaryTreeTraverse(node.Right)
	// 后序遍历
	// visited(node)

}

func visited(node *TreeNode) {
	if node.Left == nil {
		fmt.Printf("node value is %v, left Child is %v, right Child is %v.\n\r",
			node.Val, node.Left, node.Right)
	}
	if node.Right == nil {
		fmt.Printf("node value is %v, left Child is %v, right Child is %v.\n\r",
			node.Val, node.Left.Val, node.Right)
	} else {
		fmt.Printf("node value is %v, left Child is %v, right Child is %v.\n\r",
			node.Val, node.Left.Val, node.Right.Val)
	}
}
