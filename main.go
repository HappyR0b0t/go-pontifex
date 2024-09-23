package main

import (
	"fmt"

	"example.com/go-pontifex/pkg/deck_utils"
	"example.com/go-pontifex/pkg/text_utils"
	"example.com/go-pontifex/pkg/utils"
)

var plainText = "Covering topics and trends in large scale system design from the authors of the best selling System Design Interview series."

var alphabet = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"D": 4,
	"E": 5,
	"F": 6,
	"G": 7,
	"H": 8,
	"I": 9,
	"J": 10,
	"K": 11,
	"L": 12,
	"M": 13,
	"N": 14,
	"O": 15,
	"P": 16,
	"Q": 17,
	"R": 18,
	"S": 19,
	"T": 20,
	"U": 21,
	"V": 22,
	"W": 23,
	"X": 24,
	"Y": 25,
	"Z": 26,
}

var inverseAlphabet = map[int]string{
	1:  "A",
	2:  "B",
	3:  "C",
	4:  "D",
	5:  "E",
	6:  "F",
	7:  "G",
	8:  "H",
	9:  "I",
	10: "J",
	11: "K",
	12: "L",
	13: "M",
	14: "N",
	15: "O",
	16: "P",
	17: "Q",
	18: "R",
	19: "S",
	20: "T",
	21: "U",
	22: "V",
	23: "W",
	24: "X",
	25: "Y",
	26: "Z",
}

func main() {

	// deck := deck_utils.DeckGenerator(suit, rank)

	// deckKeyes := deck_utils.DeckShuffle(deck)
	// fmt.Println("INITIAL DECK ->", deckKeyes, "<- INITIAL DECK") // delete

	plainText := utils.ReadText("input_text.txt")
	// fmt.Println("PLAIN TEXT =", plainText)

	cipheredText := CypherText(plainText, alphabet, inverseAlphabet)
	fmt.Println("CIPHERED TEXT =", cipheredText)

	decipheredText := DecypherText(cipheredText, alphabet)
	fmt.Println("DECIPHERED TEXT =", decipheredText)

}

func CypherText(plainText string, alphabet map[string]int, inverseAlphabet map[int]string) string {
	var keyStream = []int{}
	numberedText := text_utils.TextToNumber(plainText, alphabet)
	inputDeck := utils.ReadDeck("input_deck.txt")
	_, keyStream = deck_utils.KeyStream(numberedText, &inputDeck, &keyStream)
	keyes := text_utils.NumberToKey(numberedText, keyStream)
	cipheredText := text_utils.KeyToText(keyes, inverseAlphabet)

	// fmt.Println("FINAL RESULT DECK ->", resultDeck, "<- FINAL RESULT DECK")
	// fmt.Println("CYPHERED TEXT ->", cypheredText, "<- CYPHERED TEXT")

	return cipheredText
}

func DecypherText(cypheredText string, alphabet map[string]int) string {
	var keyStream = []int{}
	numberedText := text_utils.TextToNumber(cypheredText, alphabet)
	inputDeck := utils.ReadDeck("input_deck.txt")
	_, keyStream = deck_utils.KeyStream(numberedText, &inputDeck, &keyStream)
	keyes := text_utils.KeyToNumber(numberedText, keyStream)
	decipheredText := text_utils.KeyToText(keyes, inverseAlphabet)

	// fmt.Println("FINAL RESULT DECK ->", resultDeck, "<- FINAL RESULT DECK")
	// fmt.Println("FINAL TEXT ->", decypheredText, "<- FINAL TEXT")

	return decipheredText
}
