package main

import "fmt"

func main() {
	fmt.Print(minFlips("01010111"))
}

func minFlips(s string) int {

	// cache := make( map[string]int )

	t := func(r []rune) int {
		if len(r) == 1 {
			return 0
		}
		// if v , ok := cache[string(r)]; ok {
		//     return v
		// }
		minT2 := 0
		pr := '2'
		for _, v := range r {
			if pr != '2' && pr == v {
				if v == '0' {
					pr = '1'
				} else {
					pr = '0'
				}
				minT2++
			} else {
				pr = v
			}
		}
		// cache[string(r)] = minT2
		return minT2
	}

	m := int(^uint(0) >> 1)
	r := []rune(s)
	r_s := len(r)
	for i := 0; i < r_s-1; i++ {
		c := t(r)
		if c < m {
			m = c
		}
		r = append(r[1:], r[0])
	}

	return m
}
