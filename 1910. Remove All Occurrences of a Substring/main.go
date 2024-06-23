package main

func main() {
	println(removeOccurrences("daabcbaabcbc", "abc"))
}

func removeOccurrencesBrute(s string, part string) string {
	i := 0
	for i < len(s)-len(part)+1 {
		if s[i:i+len(part)] == part {
			return removeOccurrencesBrute(s[0:i]+s[i+len(part):], part)
		}
		i++
	}

	return s
}

// stack
func removeOccurrences(s string, part string) string {
	i := 0
	st := []byte{}

	for i < len(s) {
		st = append(st, s[i])
		if len(st) >= len(part) && string(st[len(st)-len(part):]) == part {
			st = st[:len(st)-len(part)]
		}
		i++
	}

	return string(st)
}
