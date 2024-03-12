package issubsequence

func isSubsequence(s string, t string) bool {
	ss := []rune(s)
	tt := []rune(t)
	if len(ss) > len(tt) {
		return false
	}
	i, j := 0, 0
	for i < len(ss) && j < len(tt) {
		if ss[i] != tt[j] {
			if j == len(tt)-1 {
				return false
			}
			j++
		} else {
			if i < len(ss)-1 && j == len(tt)-1 {
				return false
			}
			i++
			j++
		}
	}
	return true
}
