package uniquenumberofoccurrences

func uniqueOccurrences(arr []int) bool {
	// build the map
	m := make(map[int]int, len(arr))
	for _, n := range arr {
		m[n]++
	}

	// reverse the map
	mRev := make(map[int]int, len(m))
	for k, v := range m {
		if _, ok := mRev[v]; ok {
			return false
		}
		mRev[v] = k
	}
	return true
}
