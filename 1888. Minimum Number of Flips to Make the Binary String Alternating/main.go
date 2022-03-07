package main

import "fmt"

func main() {
	fmt.Print(minFlips("01010111"), "\n")          //1
	fmt.Print(minFlips("110"), "\n")               //0
	fmt.Print(minFlips("01001001101"), "\n")       //2
	fmt.Print(minFlips("10001100101000000"), "\n") //5

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

	r := []rune(s)
	r_s := len(r)
	if r_s == 1 {
		return 0
	}

	if r_s > 2 {
		i := 0
		for i < r_s-2 {
			if r[i] != r[i+1] {
				i++
			} else {
				r1 := append(r[i+1:r_s], r[0:i+1]...)
				fmt.Print(string(r), "\n")
				fmt.Print(string(r1), "\n")
				if t(r1) < t(r) {
					r = r1
					i = 0
					continue
				}
				break

			}
		}
	}

	return t(r)
}
