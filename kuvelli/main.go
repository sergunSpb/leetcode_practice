// Назовем натуральное число "особым", если все его цифры различны.
// Для заданного положительного целого числа n верните количество особых целых чисел,
// принадлежащих интервалу [1, n].
//
// Пример
// Аргумент: n = 20
// Вывод: 19
// Обьяснение: Все числа от 1 до 20, кроме 11, являются особыми.
// Таким образом, существует 19 особых целых чисел.
//
// Ограничения
// 1 <= n <= 2 * 10^9

package main

import (
	"math"
	"strconv"
)

// O(...)

func countSpecialNumbers(n int) int {
	// m:= make(map[int]struct{},10)

	amount := 0

	for i := 1; i <= n; {
		m := make(map[int]int, 10)
		isSpec := true
		v := i
		st := 1
		for v > 0 {
			mod := v % 10
			if _, ok := m[mod]; ok {
				isSpec = false
				prSt := m[mod]
				if prSt == 1 {
					i++
				} else {
					i = (i/int(math.Pow(10, float64(prSt))))*int(math.Pow(10, float64(prSt))) + (mod+1)*int(math.Pow(10, float64(prSt-1)))
				}

				break
			}
			m[mod] = st

			v = v / 10
			st++
		}
		if isSpec {
			i++
			amount++
		}
	}

	return amount
}

func main() {
	type TestCase struct {
		Given    int
		Expected int
	}

	tCases := []TestCase{
		{20, 19},
		{1, 1},
		{10, 10},
		{2000, 1242},
		{2000000000, 5974650},
	}

	for i, tCase := range tCases {
		given := countSpecialNumbers(tCase.Given)
		if given != tCase.Expected {
			println("Test case " + strconv.Itoa(i) + " failed!")
		} else {
			println("Test case " + strconv.Itoa(i) + " passed!")
		}
	}
}
