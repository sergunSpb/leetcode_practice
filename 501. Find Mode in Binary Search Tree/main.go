package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	var max = 0
	var m = make(map[int]int, 0)

	var h func(root *TreeNode)
	h = func(root *TreeNode) {
		if root == nil {
			return
		}

		m[root.Val] += 1
		if m[root.Val] > max {
			max = m[root.Val]
		}
		h(root.Left)
		h(root.Right)
		return
	}

	h(root)

	var nodes []int
	for k, v := range m {
		if v == max {
			nodes = append(nodes, k)
		}
	}

	return nodes
}
