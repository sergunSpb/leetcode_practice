package main

import "strconv"

func main() {
	println(numDecodings("226") == 3)
	println(numDecodings("12") == 2)
	println(numDecodings("06") == 0)
	println(numDecodings("27") == 1)

}

func numDecodings(s string) int {
	if string(s[0]) == "0" {
		return 0
	}

	if len(s) < 2 {
		return 1
	}

	decodeMap := make(map[string]struct{}, 26)
	for i := 1; i <= 26; i++ {
		decodeMap[strconv.Itoa(i)] = struct{}{}
	}

	p0, p1 := 0, 1

	if _, ok := decodeMap[string(s[1])]; ok {
		p0++
	}
	if _, ok := decodeMap[string(s[0:2])]; ok {
		p0++
	}
	lenS := len(s)
	for i := 2; i < lenS; i++ {
		t := p1
		p1 = p0
		if _, ok := decodeMap[string(s[i])]; !ok {
			p0 = 0
		}
		if _, ok := decodeMap[string(s[i-1:i+1])]; ok {
			p0 += t
		}
	}

	return p0
}
