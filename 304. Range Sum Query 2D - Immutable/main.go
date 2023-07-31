package main

func main() {
	m := Constructor([][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}})
	println(m.SumRegion(2, 1, 4, 3) == 8)
	println(m.SumRegion(1, 2, 2, 4) == 12)
	println(m.SumRegion(2, 2, 2, 2) == 0)
}

type NumMatrix struct {
	sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := make([][]int, 0, len(matrix))

	for y := range matrix {
		line := matrix[y]
		tLine := make([]int, len(line))
		tLine[0] = line[0]
		for i := 1; i < len(line); i++ {
			tLine[i] = tLine[i-1] + line[i]
		}
		if y != 0 {
			// tLine[0] = line[0] + m[y-1][0]
			for i := 0; i < len(line); i++ {
				tLine[i] += m[y-1][i]
			}
		}
		m = append(m, tLine)
	}

	return NumMatrix{m}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	s := this.sums[row2][col2]
	if row1 > 0 {
		s -= this.sums[row1-1][col2]
	}
	if col1 > 0 {
		s -= this.sums[row2][col1-1]
	}

	if row1 > 0 && col1 > 0 {
		s += this.sums[row1-1][col1-1]
	}

	return s
}
