package main

func main() {
	println(canReach("0000000000", 8, 8))
}

func canReach(s string, minJump int, maxJump int) bool {
	// dp approach (O( n * (maxJump - minJump)))
	dp := make([]bool, len(s))
	dp[0] = true

	for i := 1; i < len(s); i++ {
		if s[i] == '1' {
			continue
		}
		if i-minJump < 0 {
			continue
		}
		for j := max(0, i-maxJump); j <= max(0, i-minJump); j++ {
			if dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(dp)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
