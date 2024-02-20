package kidswithcandies

func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))
	max := 0
	for _, v := range candies {
		if v > max {
			max = v
		}
	}
	for i, c := range candies {
		if c+extraCandies >= max {
			result[i] = true
		} else {
			result[i] = false
		}
	}
	return result
}
