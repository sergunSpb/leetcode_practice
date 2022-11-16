package main

import "fmt"

func main() {
	//fmt.Printf("%v\n", matrixBlockSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1)) //[[12,21,16],[27,45,33],[24,39,28]]

	//fmt.Printf("%v\n", matrixBlockSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 2))                                     //[[45,45,45],[45,45,45],[45,45,45]]
	fmt.Printf("%v\n", matrixBlockSum([][]int{{67, 64, 78}, {99, 98, 38}, {82, 46, 46}, {6, 52, 55}, {55, 99, 45}}, 3)) //[[731,731,731],[930,930,930],[930,930,930],[930,930,930],[721,721,721]]

}

func matrixBlockSum(mat [][]int, k int) [][]int {
	sums := make([][]int, 0, len(mat))

	for i := 0; i < len(mat); i++ {
		trSum := 0
		for z := 0; z < min(len(mat[i]), k+1); z++ {
			trSum += mat[i][z]
		}

		tArr := make([]int, 0, len(mat[i]))
		tArr = append(tArr, trSum)

		for j := 1; j < len(mat[i]); j++ {
			if j+k < len(mat[i]) {
				trSum += mat[i][j+k]
			}
			if j-k-1 >= 0 {
				trSum -= mat[i][j-k-1]
			}
			tArr = append(tArr, trSum)
		}
		sums = append(sums, tArr)
	}

	retval := make([][]int, 0, len(mat))

	for i := 0; i < len(mat); i++ {
		tAtt := make([]int, 0, len(mat[i]))
		for j := 0; j < len(mat[i]); j++ {
			tSum := sums[i][j]
			for z := 1; z < k+1; z++ {
				if i-z >= 0 {
					tSum += sums[i-z][j]
				}
				if i+z < len(mat) {
					tSum += sums[i+z][j]
				}
			}
			tAtt = append(tAtt, tSum)
		}
		retval = append(retval, tAtt)

	}

	return retval
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// func sumFromPos(mat [][]int , )
