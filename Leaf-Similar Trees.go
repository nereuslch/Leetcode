package Leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var list1,list2 []int
	leafnode(root1,&list1)
	leafnode(root2,&list2)

	if len(list1) != len(list2) {
		return false
	}

	for i := 0;i<len(list1);i ++ {
		if list1[i] != list2[i] {
			return false
		}
	}

	return true

}

func leafnode(root *TreeNode, list *[]int) {
	if root.Left == nil && root.Right == nil {
		*list = append(*list,root.Val)
		return
	}

	if root.Left != nil {
		leafnode(root.Left,list)
	}

	if root.Right != nil {
		leafnode(root.Right,list)
	}
}
