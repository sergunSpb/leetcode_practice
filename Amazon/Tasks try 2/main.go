package main

func main() {
	println(10 + ((10 - 1) - 2))
}

// func getMinMoves(plates []int32) int32 {
// 	var amount int32
// 	for i := 1; i < len(plates); i++ {
// 		if plates[i] > plates[i-1] {
// 			for j := i; j > 0; j-- {
// 				if plates[j] > plates[j-1] {
// 					plates[j], plates[j-1] = plates[j-1], plates[j]
// 					amount++
// 				} else {
// 					break
// 				}
// 			}
// 		}
// 	}

// 	return amount
// }

func max(a, b int64) int64 {
	if a > b {
		return a
	}

	return b
}

func findMaxProducts(products []int32) int64 {
	if len(products) == 1 {
		return int64(products[0])
	}

	var rMax int64 = -1
	//bruteforce , o(n^2)
outer:
	for i := len(products) - 1; i >= 0; i-- {
		var cSum int64 = 0
		var lastMax int32 = 1000000001
		for j := i; j >= 0; j-- {
			if products[j] < lastMax {
				lastMax = products[j]
			} else {
				//if equal or greater just decrement by 1 ,because 1 is the smallest difference between adjacent piles
				lastMax--
			}
			cSum += int64(lastMax)

			rMax = max(rMax, cSum)
			if lastMax == 0 {
				//stop iterating because 0 is minimum pipe
				continue outer
			}
		}
	}

	return rMax
}
