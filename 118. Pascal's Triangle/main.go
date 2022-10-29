package main

import "fmt"

func main() {
	fmt.Printf("%v", generate(5))

}

func generate(numRows int) [][]int {
	var res [][]int
	res = append(res, []int{1})
	for p := 2; p <= numRows; p++ {
		tArr := make([]int, 0, p)
		for i := 0; i < p; i++ {
			if i == 0 || i == p-1 {
				tArr = append(tArr, 1)
			} else {
				tArr = append(tArr, res[p-2][i-1]+res[p-2][i])
			}
		}
		res = append(res, tArr)
	}

	return res
}
