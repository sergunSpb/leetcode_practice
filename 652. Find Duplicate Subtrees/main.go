package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var cache = map[string][]*TreeNode{}
	getHash(root, cache)

	a := []*TreeNode{}
	for _, v := range cache {
		if len(v) > 1 {
			a = append(a, v[0])
		}
	}

	return a
}
func getHash(a *TreeNode, cache map[string][]*TreeNode) string {
	if a == nil {
		return "n"
	}
	r := strconv.Itoa(a.Val)

	if a.Left != nil || a.Right != nil {
		r += "{" + getHash(a.Left, cache) + "|" + getHash(a.Right, cache) + "}"
	}

	if _, ok := cache[r]; !ok {
		cache[r] = []*TreeNode{}
	}

	cache[r] = append(cache[r], a)

	return r
}
