package main

import "fmt"

func main() {
	n := 4
	res := generateParenthesis(n)
	fmt.Println(res)
}

func generateParenthesis(n int) []string {
	ans := []string{}
	o_p, c_p := n, n
	while
	return ans
}

// func generateParenthesis(n int) []string {
// 	ans := []string{}
// 	backtrack("", n, n, &ans)
// 	return ans
// }

// func backtrack(loc_st string, o_p, c_p int, ans *[]string) {
// 	if o_p|c_p == 0 {
// 		*ans = append(*ans, loc_st)
// 	}

// 	if o_p > 0 {
// 		backtrack(loc_st+"(", o_p-1, c_p, ans)
// 	}
// 	if c_p > o_p {
// 		backtrack(loc_st+")", o_p, c_p-1, ans)
// 	}
// }
