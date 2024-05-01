package main

import "strings"

func main() {
	println(reverseWords("    foo     bar   simple   ") == "simple bar foo")
	println(reverseWords("the sky is blue") == "blue is sky the")

}

func reverseWords(s string) string {
	st := -1
	words := make([]string, 0)
	i := 0
	for i < len(s) {
		if s[i] == ' ' {
			if st >= 0 {
				words = append(words, s[st:i])
				st = -1
			}
		} else if st < 0 {
			st = i
		}
		i++
	}

	if st > 0 {
		words = append(words, s[st:])
	}

	var b strings.Builder
	for i := len(words) - 1; i > 0; i-- {
		b.WriteString(words[i])
		b.WriteString(" ")
	}
	b.WriteString(words[0])
	return b.String()
}
