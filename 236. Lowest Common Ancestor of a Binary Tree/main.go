package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := &TreeNode{10, &TreeNode{11, &TreeNode{7, nil, nil}, &TreeNode{25, nil, nil}}, &TreeNode{13, &TreeNode{11, nil, nil}, &TreeNode{2, nil, nil}}}

	fmt.Println(findPathToNode(tree, &TreeNode{13, nil, nil}))
	fmt.Println(findPathToNode(tree, &TreeNode{7, nil, nil}))
	fmt.Println(findPathToNode(tree, &TreeNode{2, nil, nil}))
	fmt.Println(findPathToNode(tree, &TreeNode{10, nil, nil}))
	fmt.Println(findPathToNode(tree, &TreeNode{25, nil, nil}))

	fmt.Println(lowestCommonAncestor(tree, &TreeNode{10, nil, nil}, &TreeNode{13, nil, nil}))
	fmt.Println(lowestCommonAncestor(tree, &TreeNode{25, nil, nil}, &TreeNode{7, nil, nil}))
	fmt.Println(lowestCommonAncestor(tree, &TreeNode{25, nil, nil}, &TreeNode{2, nil, nil}))
}

func findPathToNode(root, n *TreeNode) []*TreeNode {
	if root == nil {
		return []*TreeNode{}
	}
	path := []*TreeNode{root}
	if n.Val == root.Val {
		return path
	}
	if l := findPathToNode(root.Left, n); len(l) > 0 {
		return append(path, l...)
	}
	if r := findPathToNode(root.Right, n); len(r) > 0 {
		return append(path, r...)
	}
	return []*TreeNode{}
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// path1 := []*TreeNode{}
	// path2 := []*TreeNode{}
	// fp, fq := false, false

	// st := []*TreeNode{root}
	// for !(fp && fq) {
	// 	st, el := st[:len(st)-1], st[len(st)-1]
	// 	st = append(st, []*TreeNode{el.Left, el.Right}...)
	// 	if !fp {
	// 		path1 = append(path1, el)
	// 	}
	// 	if !fq {
	// 		path2 = append(path2, el)
	// 	}
	// 	if el.Val == p.Val {
	// 		fp = true
	// 	}
	// 	if el.Val == q.Val {
	// 		fq = true
	// 	}
	// }
	p1 := findPathToNode(root, p)
	p2 := findPathToNode(root, q)
	s := int(math.Min(float64(len(p1)), float64(len(p2))))
	var retval *TreeNode
	for i := 0; i < s; i++ {
		if p1[i].Val != p2[i].Val {
			break
		}
		retval = p1[i]
	}
	return retval
}
