package increasingtripletsubsequence

import "math"

func increasingTriplet(nums []int) bool {
	minVal, minVal2 := math.Inf(0), math.Inf(0)
	for _, num := range nums {
		if float64(num) <= minVal {
			minVal = float64(num)
		} else if float64(num) <= minVal2 {
			minVal2 = float64(num)
		} else {
			return true
		}
	}
	return false
}
