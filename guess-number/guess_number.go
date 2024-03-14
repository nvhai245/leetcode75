package guessnumber

var pick int

func guess(num int) int {
	switch {
	case num > pick:
		return -1
	case num < pick:
		return 1
	default:
		return 0
	}
}

func guessNumber(n int) int {
	return binarySearch(1, n)
}

func binarySearch(a, b int) int {
	if a == b {
		return a
	}
	pivot := (a + b) / 2
	switch guess(pivot) {
	case 1:
		return binarySearch(pivot+1, b)
	case -1:
		return binarySearch(a, pivot-1)
	default:
		return pivot
	}
}
