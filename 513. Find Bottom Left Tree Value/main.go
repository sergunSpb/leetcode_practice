package main

func main() {
	println(findBottomLeftValue(&TreeNode{0, &TreeNode{-1, nil, nil}, nil}))
	println(findBottomLeftValue(&TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	lastVal := root.Val

	q := []*TreeNode{root}

	for len(q) > 0 {
		curRowLenght := len(q)
		for i := 0; i < curRowLenght; i++ {
			el := q[0]
			q = q[1:]
			if i == 0 {
				lastVal = el.Val
			}

			if el.Left != nil {
				q = append(q, el.Left)
			}
			if el.Right != nil {
				q = append(q, el.Right)
			}
		}
	}

	return lastVal
}

var lastValDFS = 0
var lastLevelDFS = 0

func findBottomLeftValueDFS(root *TreeNode) int {
	lastValDFS = root.Val
	lastLevelDFS = 0
	helper(root, lastLevelDFS)

	return lastValDFS
}

func helper(root *TreeNode, level int) {
	level++
	if lastLevelDFS < level {
		lastLevelDFS, lastValDFS = level, root.Val
	}

	if root.Left != nil {
		helper(root.Left, level)
	}

	if root.Right != nil {
		helper(root.Right, level)
	}
}
