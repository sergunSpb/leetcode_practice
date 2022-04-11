func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	c := 0
	curJump := nums[0]
	for _, i := range nums {
		if c > curJump {
			return false
		}
		if curJump-c < i {
			curJump = i
			c = 0
		}
		c++
	}

	return true
}
