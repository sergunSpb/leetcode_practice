package main

func main() {
	println(findCircleNum([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}))
}

func findCircleNum(isConnected [][]int) int {
	visited := make([]bool, len(isConnected))
	n := 0
	for i := range isConnected {
		if visited[i] {
			continue
		}
		n++
		markVisited(visited, isConnected, i)
	}

	return n
}

func markVisited(visited []bool, isConnected [][]int, i int) {
	visited[i] = true
	for j := range isConnected[i] {
		if visited[j] || i == j || isConnected[i][j] == 0 {
			continue
		}
		markVisited(visited, isConnected, j)
	}
}
