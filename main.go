package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"example.com/go-pontifex/pkg/deck_utils"
	"example.com/go-pontifex/pkg/text_utils"
	"example.com/go-pontifex/pkg/utils"
)

var suit = [4]string{"clubs", "diamonds", "hearts", "spades"}

var rank = [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

// A struct for response parsing
type CipherResponse struct {
	Answer string `json:"answer"`
	Deck   string `json:"deck"`
}

// A struct for request parsing
type CipherRequest struct {
	Message string `json:"message"`
	Deck    string `json:"deck"`
}

// A handler for index page
func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	file, err := os.Open("./index.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	http.ServeContent(w, r, fileInfo.Name(), fileInfo.ModTime(), file)
}

// A handler for /cipher page
func cipherHandler(w http.ResponseWriter, r *http.Request) {
	var inputData CipherRequest
	if err := json.NewDecoder(r.Body).Decode(&inputData); err != nil {
		http.Error(w, "Ошибка декодирования JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	var cipherTextAnswer = CipherText(inputData.Message, inputData.Deck)
	// Устанавливаем заголовок ответа как JSON
	w.Header().Set("Content-Type", "application/json")
	// Формируем ответ
	resp := CipherResponse{Answer: cipherTextAnswer, Deck: cipherTextAnswer}

	// Отправляем JSON-ответ
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {

	// // Generate and write input_deck to *.txt file for further usage
	// deck := deck_utils.DeckGenerator(suit, rank)
	// deckKeys := deck_utils.DeckShuffle(deck)
	// utils.WriteGeneratedDeck(deckKeys, "input_deck.txt")

	// // Generate and write test_deck to *.txt file
	// // testdeck := deck_utils.DeckArrayGenerator(suit, rank)
	// // utils.WriteGeneratedDeck(testdeck, "test_input_deck.txt")

	// // Reads and prints the plain text to terminal
	// plainText := utils.ReadText("input_text.txt")
	// fmt.Println("PLAIN TEXT =", plainText)

	// // Ciphers plaintext
	// cipheredText := CipherText()
	// utils.WriteText(cipheredText, "ciphered_text.txt")
	// fmt.Println("CIPHERED TEXT =", cipheredText)

	// // Decipheres ciphered text
	// decipheredText := DecipherText(cipheredText)
	// utils.WriteText(decipheredText, "deciphered_text.txt")
	// fmt.Println("DECIPHERED TEXT =", decipheredText)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/cipher", cipherHandler)

	log.Println("Server is running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}

// A function to cipher provided text with provided deck
func CipherText(plainText string, inputDeck []string) string {
	// plainText := utils.ReadText("input_text.txt")
	// inputDeck := utils.ReadDeck("input_deck.txt")

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
