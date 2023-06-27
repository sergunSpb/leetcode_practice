package main

import "strings"

func main() {
	println(numberToWords(1123456789)) // one billion one hundred 23 millions 4 hundred 56 thousand 7 hundred 89
	println(numberToWords(1))          // one
	println(numberToWords(0))          // one
	println(numberToWords(100))
	println(numberToWords(123)) // one
	println(numberToWords(1000))

}

var m = map[int]string{
	0:  `Zero`,
	1:  `One`,
	2:  `Two`,
	3:  `Three`,
	4:  `Four`,
	5:  `Five`,
	6:  `Six`,
	7:  `Seven`,
	8:  `Eight`,
	9:  `Nine`,
	10: `Ten`,

	11: `Eleven`,
	12: `Twelve`,
	13: `Thirteen`,
	14: `Fourteen`,
	15: `Fifteen`,
	16: `Sixteen`,
	17: `Seventeen`,
	18: `Eighteen`,
	19: `Nineteen`,

	20: `Twenty`,
	30: `Thirty`,
	40: `Forty`,
	50: `Fifty`,
	60: `Sixty`,
	70: `Seventy`,
	80: `Eighty`,
	90: `Ninety`,
}

func numberToWords(num int) string {
	if v, ok := m[num]; ok {
		return v
	}

	var r []string
	ss := []string{"", "Thousand", "Million", "Billion"}
	for i := 0; i < 4; i++ {
		v := num % 1000

		if t := genNumberLessThousand(v); t != "" {
			if ss[i] == "" {
				r = append(r, t)
			} else {
				r = append(r, t+" "+ss[i])
			}
		}
		num = num / 1000

		if num == 0 {
			break
		}
	}

	reverseSlice(r)

	return strings.Join(r, " ")
}

func reverseSlice(i []string) {
	if len(i) < 2 {
		return
	}

	for s := 0; s < len(i)/2; s++ {
		i[s], i[len(i)-s-1] = i[len(i)-s-1], i[s]
	}

}

func genNumberLessThousand(i int) string {
	if i == 0 {
		return ""
	}

	if i >= 100 {
		v := genNumberLessHundred(i % 100)
		if v != "" {
			return m[i/100] + " Hundred " + genNumberLessHundred(i%100)
		}

		return m[i/100] + " Hundred"

	}

	return genNumberLessHundred(i % 100)
}

func genNumberLessHundred(i int) string {
	if i == 0 {
		return ""
	}

	if v, ok := m[i]; ok {
		return v
	}

	v := i - (i % 10)
	res := m[v]
	v = i % 10
	res += " " + m[v]

	return res
}
