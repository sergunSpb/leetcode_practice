package main

import "fmt"

func main() {
	tree1 := TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}
	fmt.Print(inorderTraversal(&tree1))
}

//  * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0, 100)
	st := []*TreeNode{}
	cur := root
	for {
		if cur != nil {
			st = append(st, cur)
			cur = cur.Left
		} else {
			if len(st) > 0 {
				var el *TreeNode
				el, st = st[len(st)-1], st[:len(st)-1]
				ans = append(ans, el.Val)
				cur = el.Right
			} else {
				break
			}
		}
	}
	return ans
}
