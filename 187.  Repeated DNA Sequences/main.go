package main

func findRepeatedDnaSequences(s string) []string {
	if len(s) <= 10 {
		return nil
	}

	e := 9
	seen := make(map[string]int, len(s)/10)
	ans := []string{}

	for e < len(s) {
		t := s[e-9 : e+1]
		if _, ok := seen[t]; ok {
			seen[t]++
		} else {
			seen[t] = 0
		}
		e++
	}

	for k, v := range seen {
		if v == 0 {
			continue
		}
		ans = append(ans, k)
	}

	return ans
}
