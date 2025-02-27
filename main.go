package main

import (
	"fmt"

	"example.com/go-pontifex/pkg/deck_utils"
	"example.com/go-pontifex/pkg/text_utils"
	"example.com/go-pontifex/pkg/utils"
)

var suit = [4]string{"clubs", "diamonds", "hearts", "spades"}

var rank = [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

func main() {

	// Generate and write input_deck to *.txt file for further usage
	deck := deck_utils.DeckGenerator(suit, rank)
	deckKeys := deck_utils.DeckShuffle(deck)
	utils.WriteGeneratedDeck(deckKeys, "input_deck.txt")

	// Generate and write test_deck to *.txt file
	// testdeck := deck_utils.DeckArrayGenerator(suit, rank)
	// utils.WriteGeneratedDeck(testdeck, "test_input_deck.txt")

	// Reads and prints the plain text to terminal
	plainText := utils.ReadText("input_text.txt")
	fmt.Println("PLAIN TEXT =", plainText)

	// Ciphers plaintext
	cipheredText := CipherText()
	utils.WriteText(cipheredText, "ciphered_text.txt")
	fmt.Println("CIPHERED TEXT =", cipheredText)

	// Decipheres ciphered text
	decipheredText := DecipherText(cipheredText)
	utils.WriteText(decipheredText, "deciphered_text.txt")
	fmt.Println("DECIPHERED TEXT =", decipheredText)

}

// A function to cipher provided text with provided deck
func CipherText() string {
	plainText := utils.ReadText("input_text.txt")
	inputDeck := utils.ReadDeck("input_deck.txt")

	numberedText := text_utils.TextToNumber(plainText)
	var textLength int = len(numberedText)
	_, keyStream := deck_utils.KeyStream(textLength, &inputDeck)
	keys := text_utils.NumberToKey(numberedText, keyStream)
	cipheredText := text_utils.KeyToText(keys)

	return cipheredText
}

// A function to decipher provided text with provided deck
func DecipherText(cipheredText string) string {
	inputDeck := utils.ReadDeck("input_deck.txt")

	numberedText := text_utils.TextToNumber(cipheredText)
	var textLength int = len(numberedText)
	_, keyStream := deck_utils.KeyStream(textLength, &inputDeck)
	keys := text_utils.KeyToNumber(numberedText, keyStream)
	decipheredText := text_utils.KeyToText(keys)

	return decipheredText
}
