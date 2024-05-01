package main

func main() {
	println(uniquePaths(3, 7) == 28)
}

func uniquePaths(m int, n int) int {
	t := make([][]int, m)
	for i := 0; i < m; i++ {
		t[i] = make([]int, n)
	}
	t[m-1][n-1] = 1

	var helper func(i, j int, t [][]int)
	helper = func(i, j int, t [][]int) {
		if t[i][j] != 0 {
			return
		}

		r, l := 0, 0
		if m > i+1 {
			helper(i+1, j, t)
			r = t[i+1][j]
		}

		if n > j+1 {
			helper(i, j+1, t)
			l = t[i][j+1]
		}

		t[i][j] = r + l
	}

	helper(0, 0, t)
	return t[0][0]
}
