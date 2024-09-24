package text_utils

import (
	"strings"
)

// Converts text to numbers
func TextToNumber(text string, alphabet map[string]int) []int {
	text = strings.ToUpper(text)
	text = strings.ReplaceAll(text, " ", "")
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
	for i := range numberedText {
		n := numberedText[i] + keyStream[i]
		m := 26
		if n < 27 {
			keyes = append(keyes, n)
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
			n := (numberedText[i] + 26) - keyStream[i]%m
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
		text += key
		if i%5 == 0 {
			text += " "
		}
	}
	return text
}
