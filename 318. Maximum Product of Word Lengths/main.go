package main

func main() {
	println(maxProduct([]string{"aaa", "vvv"}))
}

func maxProduct(words []string) int {
	var m int
	masks := make([]int32, len(words))
	for i, w := range words {
		var t int32
		for _, l := range w {
			t |= 1 << (l - 'a')
		}
		masks[i] = t
	}

	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if masks[i]&masks[j] == 0 {
				m = max(m, len(words[i])*len(words[j]))
			}
		}
	}

	return m
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
