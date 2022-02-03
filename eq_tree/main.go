package main

import "fmt"

func main() {
	tree1 := TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	tree2 := TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	fmt.Print(isSameTree(&tree1, &tree2))
}

//  * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil || q != nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
