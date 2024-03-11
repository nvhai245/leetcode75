package reversevowelsofastring

func reverseVowels(s string) string {
	r := []rune(s)
	vowelsMap := map[string]bool{
		"a": true,
		"e": true,
		"u": true,
		"i": true,
		"o": true,
		"A": true,
		"E": true,
		"U": true,
		"I": true,
		"O": true,
	}
	i, j := 0, len(r)-1
	for i < j {
		if _, ok := vowelsMap[string(r[i])]; !ok {
			i++
			continue
		}
		if _, ok := vowelsMap[string(r[j])]; !ok {
			j--
			continue
		}
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}
	return string(r)
}
