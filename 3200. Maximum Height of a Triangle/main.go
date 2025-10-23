package main

func main() {
	println(maxHeightOfTriangle(500, 400))
}

func maxHeightOfTriangle(red int, blue int) int {
	return max(helper(red, blue), helper(blue, red))
}

func helper(red int, blue int) int {
	length := 1
	curRed := false

	for {
		if curRed {
			red -= length
			if red < 0 {
				break
			}
		} else {
			blue -= length
			if blue < 0 {
				break
			}
		}

		curRed = !curRed
		length++
	}

	return length - 1
}
