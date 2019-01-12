/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
给定一个不含重复元素的整数数组。一个以此数组构建的最大二叉树定义如下：

二叉树的根是数组中的最大元素。
左子树是通过数组中最大值左边部分构造出的最大二叉树。
右子树是通过数组中最大值右边部分构造出的最大二叉树。
通过给定的数组构建最大二叉树，并且输出这个树的根节点。
*/

func constructMaximumBinaryTree(nums []int) *TreeNode {
	midx := maxSlice(nums)
	node := &TreeNode{
		Val:   nums[midx],
		Left:  nil,
		Right: nil,
	}

	if midx-1 >= 0 {
		node.Left = constructMaximumBinaryTree(nums[:midx])
	}

	if midx+1 < len(nums) {
		node.Right = constructMaximumBinaryTree(nums[midx+1:])
	}

	return node
}

func maxSlice(ss []int) int {
	if len(ss) == 1 {
		return 0
	}

	mm := ss[0]
	midx := 0
	for idx, v := range ss {
		if mm < v {
			midx = idx
			mm = v
		}
	}

	return midx
}
