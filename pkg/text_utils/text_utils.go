package text_utils

import (
	"strings"
)

// func AlphabetGenerator(alphabet string) map[string]int {
// 	alphabetNumbers := make(map[string]int)

// 	for i := 0; i < len(alphabet); i++ {
// 		alphabetNumbers[alphabet[i]] += 1
// 	}
// 	fmt.Println(alphabetNumbers)
// 	return alphabetNumbers
// }

func TextToNumber(text string, alphabet map[string]int) []int {
	text = strings.ToUpper(text)
	text = strings.ReplaceAll(text, " ", "")
	// fmt.Println("TEXT ->", text, "<- TEXT")
	numbers := []int{}
	for i := range text {
		value := alphabet[string(text[i])]
		numbers = append(numbers, value)
	}
	return numbers
}

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

// func mapkey(m map[string]int, value int) (key string, ok bool) {
// 	for k, v := range m {
// 		if v == value {
// 			key = k
// 			ok = true
// 			return
// 		}
// 	}
// 	return
// }

// func KeyToText(keyes []int, inverseAlphabet map[int]string) []string {
// 	text := []string{}
// 	// fmt.Println(keyes)
// 	for i := range keyes {
// 		key, ok := mapkey(inverseAlphabet, keyes[i])
// 		if !ok {
// 			panic("value does not exist in map")
// 		}
// 		text = append(text, key)
// 	}
// 	return text
// }

func KeyToText(keyes []int, inverseAlphabet map[int]string) string {
	var text string
	// fmt.Println(keyes)
	for i := range keyes {
		key := inverseAlphabet[keyes[i]]
		text += key
		if i%5 == 0 {
			text += " "
		}
	}
	return text
}
