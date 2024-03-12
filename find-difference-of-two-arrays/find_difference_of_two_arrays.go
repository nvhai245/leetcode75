package finddifferenceoftwoarrays

func findDifference(nums1 []int, nums2 []int) [][]int {
	// create hash map
	mapResult := make(map[int]int, len(nums1)+len(nums2))
	for _, num := range nums1 {
		mapResult[num] = 1
	}
	for _, num := range nums2 {
		if _, ok := mapResult[num]; ok && mapResult[num] != 2 {
			mapResult[num] = 0
		} else {
			mapResult[num] = 2
		}
	}

	// create results
	nums1, nums2 = []int{}, []int{}
	for k, v := range mapResult {
		if v == 1 {
			nums1 = append(nums1, k)
		}
		if v == 2 {
			nums2 = append(nums2, k)
		}
	}
	return [][]int{nums1, nums2}
}
