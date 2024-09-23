package utils

import (
	"bufio"
	"fmt"
	"os"
)

// Creation of file and writing a deck to this file

func WriteGeneratedDeck(deckKeyes []string) {
	fo, err := os.Create("input_deck.txt")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	for i := range deckKeyes {
		_, err := fo.WriteString(deckKeyes[i] + "\n")
		if err != nil {
			panic(err)
		}
	}
}

// Read from input_deck.txt file
func ReadDeck(deckFile string) []string {
	readFile, err := os.Open(deckFile)

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	input := []string{}

	for fileScanner.Scan() {
		// fmt.Println("INPUT DECK ->", fileScanner.Text(), "<- INPUT DECK") // delete
		input = append(input, fileScanner.Text())
	}
	return input
}

// Read from input_text.txt file
func ReadText(textFile string) string {
	readFile, err := os.Open(textFile)

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var input string

	for fileScanner.Scan() {
		fmt.Println("INPUT TEXT ->", fileScanner.Text(), "<- INPUT TEXT")
		input += fileScanner.Text()
	}
	return input
}
