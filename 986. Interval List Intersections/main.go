package main

import "fmt"

func main() {
	fmt.Printf("%+v", intervalIntersection([][]int{{4, 6}, {7, 8}, {10, 17}}, [][]int{{5, 10}}))
}

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	i, j := 0, 0
	var ret [][]int

	for i < len(firstList) && j < len(secondList) {
		if firstList[i][0] <= secondList[j][0] {
			if firstList[i][1] < secondList[j][0] {
				i++
				continue
			}

			if firstList[i][1] >= secondList[j][1] {
				ret = append(ret, []int{secondList[j][0], secondList[j][1]})
				j++
			} else {
				ret = append(ret, []int{secondList[j][0], firstList[i][1]})
				i++
			}
		} else {
			if secondList[j][1] < firstList[i][0] {
				j++
				continue
			}

			if secondList[j][1] >= firstList[i][1] {
				ret = append(ret, []int{firstList[i][0], firstList[i][1]})
				i++
			} else {
				ret = append(ret, []int{firstList[i][0], secondList[j][1]})
				j++
			}
		}
	}

	return ret
}
