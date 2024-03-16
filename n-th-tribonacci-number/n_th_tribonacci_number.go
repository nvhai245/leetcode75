package nthtribonaccinumber

func tribonacci(n int) int {
	arr := make([]int, 38)
	arr[0], arr[1], arr[2] = 0, 1, 1
	for i := 3; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2] + arr[i-3]
	}
	return arr[n]
}
