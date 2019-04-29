package Leetcode
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

func maxAncestorDiff(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var sum int
	Value(root,&sum)
	return sum
}

func Value(root *TreeNode, sum *int) (int,int) {
	if root.Left == nil && root.Right == nil {
		return root.Val,root.Val
	}

	resmin, resmax := INT_MAX,INT_MIN
	if root.Left != nil {
		lmin,lmax := Value(root.Left,sum)
		resmin = lmin
		resmax = lmax
	}

	if root.Right != nil {
		rmin,rmax := Value(root.Right,sum)
		resmax = max(resmax,rmax)
		resmin = min(resmin,rmin)
	}


	*sum = max(*sum,abs(root.Val-resmin),abs(resmax-root.Val))
	return min(resmin,root.Val),max(resmax,root.Val)
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}

func min(v ...int) int {
	mind := v[0]

	for _,data := range v {
		if data < mind {
			mind = data
		}
	}

	return mind
}

func max(v ...int) int {
	maxd := v[0]

	for _,data := range v {
		if data > maxd {
			maxd = data
		}
	}

	return maxd
}