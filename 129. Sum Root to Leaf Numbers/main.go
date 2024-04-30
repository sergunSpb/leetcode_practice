package main

func main() {
	///
}

func sumNumbers(root *TreeNode) int {
	return helper(root, 0)
}

func helper(node *TreeNode, presum *int) int {
	presum = presum*10 + node.Val
	if node.Left == nil && node.Right == nil {
		return presum
	}

	sLeft := 0
	if node.Left != nil {
		sLeft = helper(node.Left, presum)
	}

	sRight := 0
	if node.Right != nil {
		sRight = helper(node.Right, presum)
	}

	return sLeft + sRight
}
