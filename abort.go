package piscine

func Abort(a, b, c, d, e int) int {
	nums := []int{a, b, c, d, e}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	return nums[2]
}
