package leetcode

import (
	"fmt"
	"sort"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree2(root *TreeNode) {
	vals := []int{}
	lnodes := []*TreeNode{}
	InOrder(root, &vals, &lnodes)
	sort.Sort(sort.IntSlice(vals[0:]))
	for i, v := range vals {
		lnodes[i].Val = v
	}
}

func InOrder(node *TreeNode, vals *[]int, lnodes *[]*TreeNode) {
	if node == nil {
		return
	}
	InOrder(node.Left, vals, lnodes)
	*vals = append(*vals, node.Val)
	*lnodes = append(*lnodes, node)
	InOrder(node.Right, vals, lnodes)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isPowerOfTwo(n int) bool {
	return (n > 0) && (n&(n-1)) == 0
}

func recoverTree(root *TreeNode) {
	var cur *TreeNode
	var pre *TreeNode
	var parent *TreeNode
	var one *TreeNode
	var two *TreeNode

	var found bool
	cur = root
	for {
		if cur == nil {
			break
		}
		if cur.Left == nil {
			fmt.Println(parent, cur)
			if parent != nil && parent.Val > cur.Val {
				if !found {
					one = parent
					found = true
				}
				two = cur
			}
			parent = cur
			cur = cur.Right
		} else {
			pre = cur.Left
			for {
				if pre.Right != nil && pre.Right != cur {
					pre = pre.Right
				} else {
					break
				}
			}
			if pre.Right == nil {
				pre.Right = cur
				cur = cur.Left
			} else {
				pre.Right = nil
				fmt.Println(parent, cur)
				if parent.Val > cur.Val {
					if !found {
						one = parent
						found = true
					}
					two = cur
				}
				parent = cur
				cur = cur.Right
			}
		}
	}
	if one != nil && two != nil {
		t := one.Val
		one.Val = two.Val
		two.Val = t
	}
}
