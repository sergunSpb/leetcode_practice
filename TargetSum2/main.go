package main

import "fmt"

func main() {
	s := pathSum(&TreeNode{2, &TreeNode{2, nil, nil}, &TreeNode{2, nil, nil}}, 4)
	fmt.Println(s)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	r := [][]int{}
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return [][]int{{targetSum}}
	}
	if root.Left != nil {
		for _, v := range pathSum(root.Left, targetSum-root.Val) {
			r = append(r, append([]int{root.Val}, v...))
		}
	}
	if root.Right != nil {
		for _, v := range pathSum(root.Right, targetSum-root.Val) {
			r = append(r, append([]int{root.Val}, v...))

		}
	}

	return r
}
