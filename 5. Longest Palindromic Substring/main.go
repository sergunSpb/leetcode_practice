package main

func main() {
	//bruteforce
}

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}

	r := ""
	str := ""

	for st := 0; st < len(s); st++ {
		for end := st; end < len(s); end++ {
			str = s[st : end+1]
			if check(str) && len(str) > len(r) {
				r = str
			}
		}
	}

	return r
}

func check(s string) bool {
	st, en := 0, len(s)-1
	for st <= en {
		if s[st] != s[en] {
			return false
		}
		st++
		en--
	}

	return true
}
