package main

func main() {
	println(minFallingPathSum([][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}))
	println(minFallingPathSum([][]int{{-19, 57}, {-40, -5}}))
	println(minFallingPathSum([][]int{{82, 81}, {69, 33}}))

}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func minInSlice(a []int) int {
	t := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] < t {
			t = a[i]
		}
	}

	return t
}

func minFallingPathSum(matrix [][]int) int {
	for i := 1; i < len(matrix); i++ {
		for y := 0; y < len(matrix); y++ {
			matrix[i][y] += minInSlice(matrix[i-1][max(y-1, 0):min(len(matrix), y+2)])
		}
	}

	return minInSlice(matrix[len(matrix)-1])
}
