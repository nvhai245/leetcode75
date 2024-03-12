package highestaltitude

func largestAltitude(gain []int) int {
	max, prev := 0, 0
	for _, g := range gain {
		prev += g
		if prev > max {
			max = prev
		}
	}
	return max
}
