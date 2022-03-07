package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := &TreeNode{10, &TreeNode{8, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}}, &TreeNode{12, &TreeNode{11, nil, nil}, &TreeNode{13, nil, nil}}}

	fmt.Println(lowestCommonAncestor(tree, &TreeNode{10, nil, nil}, &TreeNode{12, nil, nil}))
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if (p.Val >= root.Val && q.Val <= root.Val) || (p.Val <= root.Val && q.Val >= root.Val) {
		return root
	}
	if p.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		return lowestCommonAncestor(root.Right, p, q)
	}
}
