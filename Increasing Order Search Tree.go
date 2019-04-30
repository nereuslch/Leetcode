package Leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
	res := &TreeNode{Val:0}
	helper(root,res)
	return res.Right
}

func helper(root,cur *TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil {
		helper(root.Left,cur)
	}

	node := &TreeNode{Val:root.Val}
	for ;cur.Right != nil; cur = cur.Right {}
	cur.Right = node

	if root.Right != nil {
		helper(root.Right,cur)
	}
}
