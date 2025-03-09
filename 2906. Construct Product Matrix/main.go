package main

func constructProductMatrix(grid [][]int) [][]int {
	pre := make([][]int, len(grid))
	for i := range pre {
		pre[i] = make([]int, len(grid[0]))
	}

	pr := 1
	for i := range grid {
		for j := range grid[i] {
			pre[i][j] = pr
			pr = (pr * (grid[i][j] % 12345)) % 12345
		}
	}

	pr = 1
	for i := len(grid) - 1; i >= 0; i-- {
		for j := len(grid[0]) - 1; j >= 0; j-- {
			pre[i][j] = (pre[i][j] % 12345 * pr) % 12345
			pr = (pr * (grid[i][j] % 12345)) % 12345
		}
	}

	return pre
}
