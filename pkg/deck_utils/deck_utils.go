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

func DeckArrayGenerator(suit [4]string, rank [13]string) []string {
	deck := []string{}
	for i := range suit {
		for j := range rank {
			deck = append(deck, suit[i]+"-"+rank[j])
		}
	}
	deck = append(deck, "JA")
	deck = append(deck, "JB")
	return deck
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
	for i := current; i > target; i-- {
		deckKeyes[i-1], deckKeyes[i] = deckKeyes[i], deckKeyes[i-1]
	}
	return deckKeyes
}

// Finds a jocker in a deck
func FindJocker(deck []string, jocker string) int {
	jockerIndex := 0
	for i := range deck {
		if deck[i] == jocker {
			jockerIndex = i
			break
		}
	}
	return jockerIndex
}

// Move Jocker A
func MoveJockerA(deck []string, i int) ([]string, int) {
	jockerIndex := 0
	if i == len(deck)-1 {
		MoveJocker(deck, i, 1)
		jockerIndex = 1
	} else {
		MoveJocker(deck, i, i+1)
		jockerIndex = i + 1
	}
	return deck, jockerIndex
}

// Move Jocker B
func MoveJockerB(deck []string, i int) ([]string, int) {
	jockerIndex := 0
	if i == len(deck)-1 {
		MoveJocker(deck, i, 2)
		jockerIndex = 2
	} else if i == len(deck)-2 {
		MoveJocker(deck, i, 1)
		jockerIndex = 1
	} else {
		deck[i], deck[i+1], deck[i+2] = deck[i+1], deck[i+2], deck[i]
		jockerIndex = i + 2
	}
	return deck, jockerIndex
}

// Shifts both jokers accordingly
func JockerShift(deckKeyes []string) ([]string, []int) {
	jockers := []int{}
	ja := FindJocker(deckKeyes, "JA")
	deckKeyes, _ = MoveJockerA(deckKeyes, ja)
	jb := FindJocker(deckKeyes, "JB")
	deckKeyes, jb = MoveJockerB(deckKeyes, jb)
	ja = FindJocker(deckKeyes, "JA")
	jockers = append(jockers, ja, jb)
	if jockers[0] > jockers[1] {
		jockers[0], jockers[1] = jockers[1], jockers[0]
	}
	return deckKeyes, jockers
}

// Performs a triple cut on a deck
func TripleCut(deckKeyes []string, jockers []int) []string {
	deck := []string{}
	top := deckKeyes[:jockers[0]]
	middle := deckKeyes[jockers[0] : jockers[1]+1]
	bottom := deckKeyes[jockers[1]+1:]

	deck = append(deck, bottom...)
	deck = append(deck, middle...)
	deck = append(deck, top...)

	return deck
}

// Performs a count cut on a deck
func CountCut(tripleCutDeck []string, value int) []string {
	lastIndex := len(tripleCutDeck) - 1

	top := tripleCutDeck[:value]
	middle := tripleCutDeck[value:lastIndex]
	bottom := tripleCutDeck[lastIndex]

	countCutDeck := []string{}
	countCutDeck = append(countCutDeck, middle...)
	countCutDeck = append(countCutDeck, top...)
	countCutDeck = append(countCutDeck, bottom)

	return countCutDeck
}

// Converts a card to number
func cardToNumber(card string) int {
	if card == "JA" || card == "JB" {
		return 53
	} else {
		suitAndRank := strings.Split(card, "-")
		return suitsMap[suitAndRank[0]] + rankMap[suitAndRank[1]]
	}
}

// Finds output card in a deck
func FindOutput(tripleCutDeck []string) int {
	return cardToNumber(tripleCutDeck[0])

}

// Creates a keystream for conversion into chars non-recursively

func KeyStream(textLength int, inputDeck *[]string) ([]string, []int) {
	var keyStream = &[]int{}
	for i := 0; i < textLength; {
		jockers := &[]int{}
		lastIndex := len(*inputDeck) - 1

		*inputDeck, *jockers = JockerShift(*inputDeck)
		*inputDeck = TripleCut(*inputDeck, *jockers)
		lastCardValue := cardToNumber((*inputDeck)[lastIndex])
		*inputDeck = CountCut(*inputDeck, lastCardValue)
		key := FindOutput(*inputDeck)
		if key == 53 {
			continue
		}
		*keyStream = append(*keyStream, key)
		i += 1
	}

	return *inputDeck, *keyStream
}

// Recursive function of cipher/decipher process
// func KeyStreamRecusrsive(inputDeck *[]string, keyStream *[]int, numberedText []int, i int) ([]string, []int) {
// 	key := 0
// 	if i < len(numberedText) {
// 		*inputDeck = JockerShift(*inputDeck)
// 		*inputDeck = TripleCut(*inputDeck)
// 		*inputDeck = CountCut(*inputDeck)
// 		if (*inputDeck)[0] == "JA" || (*inputDeck)[0] == "JB" {
// 			fmt.Println("OUTPUT CARD IS A JOCKER!", (*inputDeck)[0])

// 		}
// 		key = FindOutput(*inputDeck)
// 		*keyStream = append(*keyStream, key)
// 		i += 1
// 	} else {
// 		return *inputDeck, *keyStream
// 	}
// 	return KeyStreamRecusrsive(inputDeck, keyStream, numberedText, i)
// }
