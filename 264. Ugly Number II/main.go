package main

func main() {
	println(nthUglyNumber(10) == 12)
	println(nthUglyNumber(1) == 1)
	println(nthUglyNumber(4) == 4)

}

func min(v1, v2, v3 int) int {
	t := v1
	if v2 < t {
		t = v2
	}

	if v3 < t {
		t = v3
	}

	return t
}

func nthUglyNumber(n int) int {
	res := make([]int, n)
	res[0] = 1

	var p2, p3, p5 int

	cnt := 1
	for cnt < n {
		res[cnt] = min(res[p2]*2, res[p3]*3, res[p5]*5)
		if res[cnt] == res[p2]*2 {
			p2++
		}
		if res[cnt] == res[p3]*3 {
			p3++
		}
		if res[cnt] == res[p5]*5 {
			p5++
		}
		cnt++
	}

	return res[n-1]
}

// func nthUglyNumber(n int) int {
// 	num := 1
// 	am := 1

// 	for am < n {
// 		num++
// 		r := checkPrime(num)
// 		if r {
// 			cache[num] = struct{}{}
// 			am++
// 		}
// 	}

// 	return num
// }

// func checkPrime(n int) bool {
// 	dividers := []int{2, 3, 5}

// 	for n > 1 {
// 		for _, div := range dividers {
// 			if n%div == 0 {
// 				if _, ok := cache[n/div]; ok {
// 					return true
// 				}
// 				return false
// 			}
// 		}
// 		return false
// 	}

// 	return false
// }
