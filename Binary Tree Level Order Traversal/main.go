package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	ans := levelOrder(&TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}})
	fmt.Println(ans)
}

func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root != nil {
		ans = append(ans, []int{root.Val})
	}
	helper(&ans, root)
	return ans
}
func helper(ans *[][]int, root *TreeNode) {
	l_s := []int{}
	if root == nil || (root.Left == nil && root.Right == nil) {
		return
	}
	if root.Left != nil {
		l_s = append(l_s, root.Left.Val)
	}
	if root.Right != nil {
		l_s = append(l_s, root.Right.Val)
	}
	*ans = append(*ans, l_s)
	helper(ans, root.Left)
	helper(ans, root.Right)
}
