package gcdofstrings

func gcdOfStrings(str1 string, str2 string) string {
	if len(str1) < len(str2) {
		return gcdOfStrings(str2, str1)
	}
	if str1 == str2 {
		return str1
	}
	if str1[:len(str2)] != str2 {
		return ""
	} else {
		return gcdOfStrings(str1[len(str2):], str2)
	}
}