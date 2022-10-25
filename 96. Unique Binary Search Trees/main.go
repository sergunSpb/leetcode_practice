package main

func main() {
	println(numTrees(3) == 5)
	println(numTrees(1) == 1)
}

func numTrees(n int) int {
	arr := make([]int, n+1)
	arr[0] = 1
	arr[1] = 1
	for t_arr := 2; t_arr <= n; t_arr++ {
		t := 0
		for i := 0; i < t_arr; i++ {
			t += arr[i] * arr[t_arr-1-i]
		}
		arr[t_arr] = t
	}

	return arr[n]
}
