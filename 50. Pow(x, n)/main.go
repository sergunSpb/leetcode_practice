package main

func main() {
	println(myPow(1, 12312314))
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	neg := false
	if n < 0 {
		n = -n
		neg = true
	}

	if n%2 == 1 {
		t := myPow(x, n/2)
		x = t * t * x
	} else {
		t := myPow(x, n/2)
		x = t * t
	}

	if neg {
		return 1 / x
	}

	return x
}
