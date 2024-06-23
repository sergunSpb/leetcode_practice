package main

import "fmt"

func main() {
	println(findStart([]int{1, 5, 6, 7, 8, 9, 10}, 9))
	println(findStart([]int{-10000}, -100))

	fmt.Printf("%v\n", findClosestElements([]int{1, 5, 6, 7, 8, 9, 10, 11, 24, 12}, 4, 10))
	fmt.Printf("%v\n", findClosestElements2([]int{1, 5, 6, 7, 8, 9, 10, 11, 24, 12}, 4, 10))
	fmt.Printf("%v\n", findClosestElements([]int{1, 1, 1, 10, 10, 10}, 1, 9))
	fmt.Printf("%v\n", findClosestElements2([]int{1, 1, 1, 10, 10, 10}, 1, 9))

	println(findStart2([]int{1, 5, 6, 7, 8, 9, 10, 14}, 11))
	println(findStart2([]int{1, 5, 6, 7, 8, 9, 10}, 9))
	println(findStart2([]int{-10000}, -100))
	println(findStart2([]int{-99, -100}, -100))
}

// two pointer + binary search , now extra space
func findClosestElements2(arr []int, k int, x int) []int {
	p := findStart2(arr, x)

	lp, rp := p, p+1

	i := 1
	for i < k {
		i++
		if lp == 0 {
			rp++
			continue
		}
		if rp == len(arr) {
			lp--
			continue
		}

		if x-arr[lp-1] <= arr[rp]-x {
			lp--
			continue
		}
		rp++
	}

	return arr[lp:rp]
}

// two pointer + binary search , with extra space
func findClosestElements(arr []int, k int, x int) []int {
	p := findStart2(arr, x)

	la, ra := []int{}, []int{arr[p]}
	lp, rp := p-1, p+1

	i := 1
	for lp >= 0 && rp <= len(arr)-1 && i < k {
		if lp >= 0 && x-arr[lp] <= arr[rp]-x {
			la = append(la, arr[lp])
			lp--
		} else if rp < len(arr) {
			ra = append(ra, arr[rp])
			rp++
		}
		i++
	}

	res := make([]int, 0, len(arr))

	if i < k {
		if lp < 0 {
			ra = append(ra, arr[rp:rp+k-i]...)
		} else {
			for y := 0; y < (k - i); y++ {
				la = append(la, arr[lp-y])
			}
		}
	}

	for i := len(la) - 1; i >= 0; i-- {
		res = append(res, la[i])
	}

	res = append(res, ra...)

	return res
}

// search closest iterative
func findStart(arr []int, x int) int {
	if x >= arr[len(arr)-1] {
		return len(arr) - 1
	}

	if x <= arr[0] {
		return 0
	}

	for i := 1; i < len(arr); i++ {
		if arr[i] >= x {
			if x-arr[i-1] <= arr[i]-x {
				return i - 1
			}
			return i
		}
	}

	return 0
}

// search closest binary search
func findStart2(arr []int, x int) int {
	st, en := 0, len(arr)

	for st < en {
		p := st + (en-st)/2

		if arr[p] > x {
			en = p
			continue
		}

		st = p + 1
	}

	if st == len(arr) {
		return st - 1
	}

	if st == 0 {
		return 0
	}

	if abs(arr[st-1]-x) <= abs(arr[st]-x) {
		return st - 1
	}

	return st
}

func abs(x int) int {
	if x >= 0 {
		return x
	}

	return -x
}
