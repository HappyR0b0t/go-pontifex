package text_utils

import (
	"strings"
)

// func NumberToString(nums []int) string {
// 	result := ""
// var c int = 64
// for _, val := range nums {
// 	// fmt.Println(int(val), "<- current val")
// 	result += string(val + c)
// }
// return result
// }

// Converts text to numbers
func TextToNumber(text string) []int {
	text = strings.ReplaceAll(strings.ToUpper(text), " ", "")
	numbers := []int{}
	var c int = 64
	for _, val := range []byte(text) {
		numbers = append(numbers, int(val)-c)
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

// V2 Converts keys from the keystream to chars
func KeyToText(keyes []int) string {
	var text string
	var c int = 64
	for i, val := range keyes {
		// fmt.Println(int(val), "<- current val")
		if i%5 == 0 && i > 4 {
			text += " "
		}
		text += string(byte(val + c))
	}
	return text
}
