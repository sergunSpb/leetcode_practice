package main

func main() {
	println(mySqrt(10))
}

func mySqrt(x int) int {
	if x < 2 {
		return 1
	}

	l, r := 1, x

	for l <= r {
		mid := l + (r-l)/2
		res := mid * mid
		if res == x {
			return mid
		}
		if res > x {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r
}
