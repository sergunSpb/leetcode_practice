package main

import "fmt"

func main() {
	ans := combine(4, 2)
	fmt.Println(ans)
	fmt.Println(len(ans))
}

var facts []uint64

func combine(n int, k int) [][]int {

	if n < k || n == 0 || k == 0 {
		return [][]int{}
	}
	length := FactorialMemoization(uint64(n)) / (FactorialMemoization(uint64(k)) * FactorialMemoization(uint64(n-k)))
	ans := make([][]int, length)

	calc(1, n, k, []int{}, &ans)
	return ans
}

func calc(start, end, length int, t_sl []int, ans *[][]int) {
	if length == 0 {
		dst := make([]int, len(t_sl))
		copy(dst, t_sl)
		*ans = append(*ans, dst)
		return
	}
	for i := start; i <= end; i++ {
		t_sl = append(t_sl, i)
		calc(i+1, end, length-1, t_sl, ans)
		t_sl = t_sl[:len(t_sl)-1]
	}
}

func FactorialMemoization(n uint64) (res uint64) {
	if facts[n] != 0 {
		res = facts[n]
		return res
	}

	if n > 0 {
		res = n * FactorialMemoization(n-1)
		return res
	}

	return 1
}
