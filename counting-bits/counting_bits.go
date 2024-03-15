package countingbits

func countBits(n int) []int {
	result := make([]int, n+1)
	result[0] = 0
	currentPow := 1
	for i := 1; i < n+1; i++ {
		if i == currentPow*2 {
			currentPow *= 2
		}
		result[i] = 1 + result[i-currentPow]
	}
	return result
}
