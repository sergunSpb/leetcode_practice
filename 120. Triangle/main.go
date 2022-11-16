package main

func main() {
	println(minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}

	var localMin int
	for i := 1; i < len(triangle); i++ {
		for y := 0; y <= i; y++ {
			if y == 0 {
				triangle[i][y] += triangle[i-1][y]
				localMin = triangle[i][y]
			} else if y == i {
				triangle[i][y] += triangle[i-1][y-1]
			} else {
				triangle[i][y] += min(triangle[i-1][y-1], triangle[i-1][y])
			}
			localMin = min(triangle[i][y], localMin)
		}
	}

	return localMin
}
