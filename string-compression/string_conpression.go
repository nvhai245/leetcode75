package stringcompression

import (
	"strconv"
)

func compress(chars []byte) int {
	var currentChar byte
	var repeatBytes []byte
	var repeatCount, currentCharIndex int
	currentCharCountIndex := -1
	for _, c := range chars {
		if c == currentChar {
			repeatCount++
			repeatBytes = repeatNumToBytes(repeatCount)
			currentCharCountIndex = currentCharIndex
			for _, r := range repeatBytes {
				currentCharCountIndex++
				chars[currentCharCountIndex] = r
			}
		} else {
			currentChar = c
			repeatCount = 1
			currentCharCountIndex++
			chars[currentCharCountIndex] = currentChar
			currentCharIndex = currentCharCountIndex
		}
	}
	chars = chars[:currentCharCountIndex+1]
	return len(chars)
}

func repeatNumToBytes(num int) []byte {
	str := strconv.Itoa(num)
	charSlice := []byte(str)
	return charSlice
}
