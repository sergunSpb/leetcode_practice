func divisorGame(n int) bool {
	return divisorGameDP(n)
}

func divisorGameDP(n int) bool {
	cache := make([]bool, n+1)
	cache[0] = false
	cache[1] = false

	for i := 2; i <= n; i++ {
		for y := 1; y < i; y++ {
			if i%y == 0 {
				if !cache[i-y] {
					cache[i] = true
					break
				}
			}
		}
	}

	return cache[n]
}

func divisorGameBF(n int) bool {
	for i := 1; i < n; i++ {
		if n%i == 0 && !divisorGameBF(n-i) {
			return true
		}
	}
	return false
}