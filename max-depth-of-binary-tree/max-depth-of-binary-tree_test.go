package leetcode

import (
	"testing"
)

func Test(t *testing.T) {
	tree := &TreeNode{
		Left: &TreeNode{
			Left: &TreeNode{},
		},
	}
	t.Log(maxDepth(tree))
}
