package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
Recursion
*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	dep_left := maxDepth(root.Left)
	dep_right := maxDepth(root.Right)

	depth := 0
	if dep_left > dep_right {
		depth = dep_left
	} else {
		depth = dep_right
	}

	return depth + 1
}
