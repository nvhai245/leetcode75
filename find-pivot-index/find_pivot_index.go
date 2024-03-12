package findpivotindex

func pivotIndex(nums []int) int {
	sum, leftSum := 0, 0
	for _, num := range nums {
		sum += num
	}
	for i, num := range nums {
		if leftSum*2+num == sum {
			return i
		}
		leftSum += num
	}
	return -1
}
