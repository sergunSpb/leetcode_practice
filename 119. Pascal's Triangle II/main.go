package main

import "fmt"

func main() {
	fmt.Print(getRow(3))
	fmt.Print(getRow(4))

}

func getRow(rowIndex int) []int {
	prRes := []int{1}
	for i := 1; i <= rowIndex; i++ {
		res := make([]int, 0, i+1)
		for y := 0; y <= i; y++ {
			if y == 0 || y == i {
				res = append(res, 1)
			} else {
				res = append(res, prRes[y]+prRes[y-1])
			}
		}
		prRes = res
	}

	return prRes
}
