package text_utils

import (
	"strings"
)

// Converts text to numbers
func TextToNumber(text string, alphabet map[string]int) []int {
	text = strings.ReplaceAll(strings.ToUpper(text), " ", "")
	numbers := []int{}
	for i := range text {
		value := alphabet[string(text[i])]
		numbers = append(numbers, value)
	}
	return numbers
}

// Converts numbers to key for the keystream
func NumberToKey(numberedText []int, keyStream []int) []int {
	keyes := []int{}
	m := 26
	for i := range numberedText {
		n := numberedText[i] + keyStream[i]
		if n%m == 0 {
			keyes = append(keyes, 26)
		} else {
			keyes = append(keyes, n%m)
		}
	}
	return keyes
}

// Converts keys from the keystream to numbers
func KeyToNumber(numberedText []int, keyStream []int) []int {
	keyes := []int{}
	for i := range numberedText {
		m := 26
		if numberedText[i] < keyStream[i]%m {
			n := (numberedText[i] + m) - keyStream[i]%m
			keyes = append(keyes, n)
		} else {
			n := (numberedText[i]) - keyStream[i]%m
			keyes = append(keyes, n)
		}
	}
	return keyes
}

// Converts keys from the keystream to chars
func KeyToText(keyes []int, inverseAlphabet map[int]string) string {
	var text string
	for i := range keyes {
		key := inverseAlphabet[keyes[i]]
		if i%5 == 0 && i > 4 {
			text += " "
		}
		text += key
	}
	return text
}
