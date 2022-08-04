package main

import "fmt"

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}) == true)
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}) == true)
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}) == false)

}

func wordBreak(s string, wordDict []string) bool {
	var cache map[int]bool = make(map[int]bool)

	var f func(pos int) bool
	f = func(pos int) bool {
		for _, word := range wordDict {
			if pos+len(word) <= len(s) && s[pos:pos+len(word)] == word {
				if _, ok := cache[pos+len(word)]; ok {
					continue
				}
				if pos+len(word) == len(s) {
					return true
				}
				if f(pos + len(word)) {
					return true
				} else {
					cache[pos+len(word)] = false
				}
			}
		}
		return false
	}

	return f(0)
}
