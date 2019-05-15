package Leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, L int, R int) int {
	var sum int
	loop(root,L,R,&sum)
	return sum
}

func loop(root *TreeNode,L,R int, sum *int) {
	if root == nil {
		return
	}

	if root.Val >= L && root.Val <= R {
		*sum += root.Val
	}
	if root.Val > L {
		loop(root.Left,L,R,sum)
	}
	if root.Val < R {
		loop(root.Right,L,R,sum)
	}
}