package main

func main() {

	println(numIslands([][]byte{{'1', '1'}, {'0', '1'}}))
}

func numIslands(grid [][]byte) int {
	curX, answer := 0, 0
	for curX < len(grid) {
		curY := 0
		for curY < len(grid[0]) {
			if grid[curX][curY] == '1' {
				helper(curX, curY, grid)
				answer++
			}
			curY++
		}
		curX++
	}

	return answer
}

func helper(pX, pY int, grid [][]byte) {
	if pY >= len(grid[0]) || pY < 0 || pX >= len(grid) || pX < 0 {
		return
	}

	if grid[pX][pY] != '1' {
		grid[pX][pY] = '-'
		return
	}
	grid[pX][pY] = '-'

	helper(pX+1, pY, grid)
	helper(pX-1, pY, grid)
	helper(pX, pY+1, grid)
	helper(pX, pY-1, grid)
}
