package main

import "fmt"

func main() {
	buildings := [][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}
	fmt.Println(findDiagonalOrder(buildings))
}

func findDiagonalOrder(matrix [][]int) []int {
	maxY := len(matrix)
	if  maxY == 0 {
		return []int{}
	if maxY ==1 {
		return matrix[0]
	}
	maxX := len(matrix[0])
	if maxX == 0 {
		return []int{}
	}
	retval := make([]int, 0, maxX*maxY)
	cur_dir := 0
	x, y := 0, 0
	for {
		retval = append(retval, matrix[y][x])
		if cur_dir == 0 {
			if y == 0 {
				cur_dir = 1
				if x < maxX-1 {
					x++
					continue
				} else if y < maxY-1 {
					y++
					continue
				}
				break
			}
			if x == maxX-1 {
				cur_dir = 1
				if y < maxY-1 {
					y++
					continue
				} else if x < maxX-1 {
					x++
					continue
				}
				break
			}
			x++
			y--
			continue
		}
		if cur_dir == 1 {
			if x == 0 {
				cur_dir = 0
				if y < maxY-1 {
					y++
					continue
				} else if x < maxX-1 {
					x++
					continue
				}
				break
			}
			if y == maxY-1 {
				cur_dir = 0
				if x < maxX-1 {
					x++
					continue
				} else if y < maxY-1 {
					y++
					continue
				}
				break
			}
			x--
			y++
			continue
		}
	}
	return retval
}
