package main

import "fmt"

func main() {
	n := 9
	dm := make(map[int]interface{})
	for x := 1; x <= n; x++ {
		dm[x] = nil
	}
	s := calc(dm, 2)
	fmt.Println(s)
	fmt.Println(len(s))
}

func calc(n map[int]interface{}, k int) [][]int {
	retcal := make([][]int, 0)

	for digit, _ := range n {
		targetMap := make(map[int]interface{})
		for key, _ := range n {
			targetMap[key] = nil
		}

		delete(targetMap, digit)
		var sl [][]int
		if k > 1 {
			sl = calc(targetMap, k-1)
			for _, x := range sl {
				s := []int{digit}
				s = append(s, x...)
				retcal = append(retcal, s)
			}
		} else {
			s := []int{digit}
			retcal = append(retcal, s)
		}
	}

	return retcal
}
