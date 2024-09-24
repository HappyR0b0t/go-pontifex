package main

import (
	"example.com/go-pontifex/pkg/deck_utils"
	"example.com/go-pontifex/pkg/text_utils"
	"example.com/go-pontifex/pkg/utils"
)

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

var suit = [4]string{"clubs", "diamonds", "hearts", "spades"}

var rank = [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

func main() {

	deck := deck_utils.DeckGenerator(suit, rank)
	deckKeyes := deck_utils.DeckShuffle(deck)
	utils.WriteGeneratedDeck(deckKeyes, "input_deck.txt")

	plainText := utils.ReadText("input_text.txt")

	cipheredText := CipherText(plainText, alphabet, inverseAlphabet)
	utils.WriteText(cipheredText, "ciphered_text.txt")

	decipheredText := DecipherText(cipheredText, alphabet)
	utils.WriteText(decipheredText, "deciphered_text.txt")

}

// A function to cipher provided text with provided deck
func CipherText(plainText string, alphabet map[string]int, inverseAlphabet map[int]string) string {
	var keyStream = []int{}
	numberedText := text_utils.TextToNumber(plainText, alphabet)
	inputDeck := utils.ReadDeck("input_deck.txt")
	_, keyStream = deck_utils.KeyStream(numberedText, &inputDeck, &keyStream)
	keyes := text_utils.NumberToKey(numberedText, keyStream)
	cipheredText := text_utils.KeyToText(keyes, inverseAlphabet)

	return cipheredText
}

// A function to decipher provided text with provided deck
func DecipherText(cipheredText string, alphabet map[string]int) string {
	var keyStream = []int{}
	numberedText := text_utils.TextToNumber(cipheredText, alphabet)
	inputDeck := utils.ReadDeck("input_deck.txt")
	_, keyStream = deck_utils.KeyStream(numberedText, &inputDeck, &keyStream)
	keyes := text_utils.KeyToNumber(numberedText, keyStream)
	decipheredText := text_utils.KeyToText(keyes, inverseAlphabet)

	return decipheredText
}
