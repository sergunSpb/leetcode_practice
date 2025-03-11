package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	out := [][]int{[]int{root.Val}}
	dirLR := false

	nl := []*TreeNode{root}
	for {
		nl = nextLevel(nl)
		if len(nl) == 0 {
			break
		}
		r := []int{}
		if !dirLR {
			for i := len(nl) - 1; i >= 0; i-- {
				r = append(r, nl[i].Val)
			}
		} else {
			for i := range nl {
				r = append(r, nl[i].Val)
			}
		}
		dirLR = !dirLR
		out = append(out, r)
	}

	return out
}

func nextLevel(nn []*TreeNode) []*TreeNode {
	r := []*TreeNode{}
	for _, n := range nn {
		if n.Left != nil {
			r = append(r, n.Left)
		}
		if n.Right != nil {
			r = append(r, n.Right)
		}
	}

	return r
}
