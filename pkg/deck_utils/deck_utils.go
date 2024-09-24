package deck_utils

import (
	"math/rand"
	"strings"
)

var suitsMap = map[string]int{
	"clubs":    0,
	"diamonds": 13,
	"hearts":   26,
	"spades":   39,
}

var rankMap = map[string]int{
	"A":  1,
	"2":  2,
	"3":  3,
	"4":  4,
	"5":  5,
	"6":  6,
	"7":  7,
	"8":  8,
	"9":  9,
	"10": 10,
	"J":  11,
	"Q":  12,
	"K":  13,
	"JA": 53,
	"JB": 53,
}

func DeckGenerator(suit [4]string, rank [13]string) map[string]int {
	deck := map[string]int{}
	k := 1
	for i := range suit {
		for j := range rank {
			deck[suit[i]+"-"+rank[j]] = k
			k++
		}
	}
	deck["JA"] = 53
	deck["JB"] = 53
	return deck
}

// Shuffles the deck `randomly`
func DeckShuffle(deck map[string]int) []string {
	deckKeyes := []string{}
	for k := range deck {
		deckKeyes = append(deckKeyes, k)
	}
	rand.Shuffle(len(deckKeyes), func(i, j int) { deckKeyes[i], deckKeyes[j] = deckKeyes[j], deckKeyes[i] })
	return deckKeyes
}

// Moves Jocker to target position
func MoveJocker(deckKeyes []string, current int, target int) []string {
	for i := len(deckKeyes) - current; i > target; i-- {
		deckKeyes[i-1], deckKeyes[i] = deckKeyes[i], deckKeyes[i-1]
	}

	return deckKeyes
}

// Finds both Jokers and shifts them accordingly
// TODO Return indexes of both Jokers
func JockerShift(deckKeyes []string) []string {
	shiftOne := 1
	shiftTwo := 2
	a := false
	b := false
	for i := 0; i < len(deckKeyes); i++ {
		if a && b {
			break
		}
		if deckKeyes[i] == "JA" && !a {
			deckKeyes[i], deckKeyes[i+1] = deckKeyes[i+1], deckKeyes[i] // One function could be used. I guess
			a = true
			i = 0
		} else if deckKeyes[len(deckKeyes)-1] == "JA" && !a {
			MoveJocker(deckKeyes, shiftOne, shiftOne)
			a = true
			i = 0
		}
		if deckKeyes[i] == "JB" && !b && a {
			deckKeyes[i], deckKeyes[i+1], deckKeyes[i+2] = deckKeyes[i+1], deckKeyes[i+2], deckKeyes[i]
			b = true
			i += 2
		} else if deckKeyes[len(deckKeyes)-1] == "JB" && !b && a {
			MoveJocker(deckKeyes, shiftOne, shiftTwo)
			b = true
			break
		} else if deckKeyes[len(deckKeyes)-2] == "JB" && !b && a {
			MoveJocker(deckKeyes, shiftTwo, shiftOne)
			b = true
			break
		}
	}
	return deckKeyes
}

// Performs a triple cut on a deck
func TripleCut(deckKeyes []string) []string {
	deck := []string{}
	top := []string{}
	middle := []string{}
	bottom := []string{}
	current := 0
	for i := range deckKeyes {
		if deckKeyes[i] == "JA" {
			current++
		} else if deckKeyes[i] == "JB" {
			middle = append(middle, deckKeyes[i])
			current++
			continue
		}
		if current == 0 {
			top = append(top, deckKeyes[i])
			continue
		} else if current == 1 {
			middle = append(middle, deckKeyes[i])
			continue
		} else if current == 2 {
			bottom = append(bottom, deckKeyes[i])
			continue
		}
	}

	deck = append(deck, bottom...)
	deck = append(deck, middle...)
	deck = append(deck, top...)
	return deck
}

// Performs a count cut on a deck
func CountCut(tripleCutDeck []string) []string {
	lastIndex := len(tripleCutDeck) - 1
	bottomCard := tripleCutDeck[lastIndex]
	number := cardToNumber(tripleCutDeck[lastIndex])
	top := tripleCutDeck[0:number]
	bottom := tripleCutDeck[number:lastIndex]
	deck := []string{}
	deck = append(deck, bottom...)
	deck = append(deck, top...)
	deck = append(deck, bottomCard)

	return deck
}

// Converts a card to number
func cardToNumber(card string) int {
	suitAndRank := strings.Split(card, "-")
	if suitAndRank[0] == "JA" || suitAndRank[0] == "JB" {
		number := 53
		return number
	} else {
		number := suitsMap[suitAndRank[0]] + rankMap[suitAndRank[1]]
		return number
	}
}

// Finds output card in a deck
func FindOutput(tripleCutDeck []string) int {
	number := cardToNumber(tripleCutDeck[0])
	outputNumber := cardToNumber(tripleCutDeck[number])
	return outputNumber
}

// Creates a keystream for conversion into chars
func KeyStream(numberedText []int, inputDeck *[]string, keyStream *[]int) ([]string, []int) {
	i := 0
	jokers := []int{0, 0}
	KeyStreamRecusrsive(inputDeck, keyStream, numberedText, i, &jokers)

	return *inputDeck, *keyStream
}

// Recursive function of cipher/decipher process
func KeyStreamRecusrsive(inputDeck *[]string, keyStream *[]int, numberedText []int, i int, jokers *[]int) ([]string, []int) {
	key := 0
	if i < len(numberedText) {
		*inputDeck = JockerShift(*inputDeck)
		*inputDeck = TripleCut(*inputDeck)
		*inputDeck = CountCut(*inputDeck)
		key = FindOutput(*inputDeck)
		*keyStream = append(*keyStream, key)
		i += 1
	} else {
		return *inputDeck, *keyStream
	}
	return KeyStreamRecusrsive(inputDeck, keyStream, numberedText, i, jokers)
}
