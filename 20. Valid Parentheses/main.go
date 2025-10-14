package main

func main() {
	println(isValid("()"))
}

var openPar = map[rune]rune{
	'(': ')',
	'{': '}',
	'[': ']',
}

func isValid(s string) bool {
	var stack = make([]rune, 0, len(s)/2)

	if len(s)%2 != 0 {
		return false
	}

	for _, c := range s {
		if _, ok := openPar[c]; ok {
			stack = append(stack, c)
			continue
		}

		if len(stack) == 0 || len(stack) > len(s)/2 {
			return false
		}

		v := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		if openPar[v] != c {
			return false
		}
	}

	return len(stack) == 0
}
