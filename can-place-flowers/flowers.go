package canplaceflowers

import "slices"

func canPlaceFlowers(flowerbed []int, n int) bool {
	flowerbed = slices.Concat([]int{0}, flowerbed, []int{0})
	for i := 1; i < len(flowerbed)-1; i++ {
		if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			n--
		}
		if n <= 0 {
			return true
		}
	}
	return false
}
