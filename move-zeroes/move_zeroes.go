package movezeroes

func moveZeroes(nums []int) {
	i, j := 0, 0
	for j < len(nums) && i < len(nums) {
		if nums[j] == 0 {
			j++
			continue
		}
		if nums[i] != 0 {
			i++
			continue
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			j++
		}
	}
}
