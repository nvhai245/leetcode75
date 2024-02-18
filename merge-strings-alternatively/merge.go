package mergestringsalternatively

func mergeAlternatively(word1 string, word2 string) string {
	merged := ""
	if len(word1) == 0 {
		return word2
	}
	for i, r := range word1 {
		if len(word2) < i+1 {
			merged = merged + string(r)
		} else if i == len(word1)-1 && len(word2) > i+1 {
			merged = merged + string(r) + word2[i:]
		} else {
			merged = merged + string(r) + string(word2[i])
		}
	}
	return merged
}
