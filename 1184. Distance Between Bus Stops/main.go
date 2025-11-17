package main

func main() {
	println(distanceBetweenBusStops([]int{10, 3, 4}, 0, 2))
}

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	return min(helper(distance, start, destination), helper(distance, destination, start))
}

func helper(distance []int, start int, destination int) int {
	s := 0
	for i := start; i != destination; i = (i + 1) % len(distance) {
		s += distance[i]
	}
	return s
}
