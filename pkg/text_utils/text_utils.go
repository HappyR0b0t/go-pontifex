package text_utils

import (
	"fmt"
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
	fmt.Println("--- NUMBER TO KEY")
	fmt.Println("NUMBER TO KEY NUMBERED TEXT: ", numberedText)
	fmt.Println("NUMBER TO KEY KEYSTREAM: ", keyStream)
	keyes := []int{}
	m := 26

	for i := range numberedText {
		n := numberedText[i] + keyStream[i]
		fmt.Println((numberedText[i] + keyStream[i]) % m)
		if n%m == 0 {
			keyes = append(keyes, 26)
		} else {
			keyes = append(keyes, n%m)
		}

		// if n < 27 {
		// 	keyes = append(keyes, n)
		// } else {
		// 	keyes = append(keyes, n%m)
		// }
		// keyes = append(keyes, numberedText[i]+(keyStream[i]%m))
	}
	fmt.Println("NUMBER TO KEY KEYES: ", keyes)
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
	fmt.Println("--- KEY TO TEXT:")
	var text string
	fmt.Println("KEY TO TEXT KEYES:", keyes)
	for i := range keyes {
		key := inverseAlphabet[keyes[i]]
		fmt.Println("KEY:", key)
		text += key
		// if i%5 == 0 && i != 1 {
		// 	text += " "
		// }
	}
	return text
}
