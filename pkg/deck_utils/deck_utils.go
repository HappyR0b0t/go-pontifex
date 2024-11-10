package deck_utils

import (
	"fmt"
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

// Moves Jocker to target position V1
func MoveJockerV1(deckKeyes []string, current int, target int) []string {
	for i := current; i > target; i-- {
		deckKeyes[i-1], deckKeyes[i] = deckKeyes[i], deckKeyes[i-1]
	}
	return deckKeyes
}

// Finds both Jokers and shifts them accordingly V1
// TODO Return indexes of both Jokers
// func JockerShiftV1(deckKeyes []string) []string {
// 	shiftOne := 1
// 	shiftTwo := 2
// 	a := false
// 	b := false
// 	for i := 0; i < len(deckKeyes); i++ {
// 		if a && b {
// 			break
// 		}
// 		if deckKeyes[i] == "JA" && !a && i != 53 {
// 			deckKeyes[i], deckKeyes[i+1] = deckKeyes[i+1], deckKeyes[i] // One function could be used. I guess
// 			a = true
// 			i = 0
// 		} else if deckKeyes[len(deckKeyes)-1] == "JA" && !a {
// 			MoveJocker(deckKeyes, shiftOne, shiftOne)
// 			a = true
// 			i = 0
// 		}
// 		if deckKeyes[i] == "JB" && !b && a {
// 			deckKeyes[i], deckKeyes[i+1], deckKeyes[i+2] = deckKeyes[i+1], deckKeyes[i+2], deckKeyes[i]
// 			b = true
// 			i += 2
// 		} else if deckKeyes[len(deckKeyes)-1] == "JB" && !b && a {
// 			MoveJocker(deckKeyes, shiftOne, shiftTwo)
// 			b = true
// 			break
// 		} else if deckKeyes[len(deckKeyes)-2] == "JB" && !b && a {
// 			MoveJocker(deckKeyes, shiftTwo, shiftOne)
// 			b = true
// 			break
// 		}
// 	}
// 	if deckKeyes[53] == "JA" || deckKeyes[53] == "JB" {
// 		fmt.Println("JOCKER SHIFT: JOCKER IS THE LAST CARD!", deckKeyes[53])
// 	}
// 	return deckKeyes
// }

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

// MoveJockerA V2
func MoveJockerA(deck []string, i int) ([]string, int) {
	jockerIndex := 0
	if i == 53 {
		MoveJockerV1(deck, i, 0)
		jockerIndex = 1
	} else {
		deck[i], deck[i+1] = deck[i+1], deck[i]
		jockerIndex = i + 1
	}
	return deck, jockerIndex
}

func MoveJockerB(deck []string, i int) ([]string, int) {
	jockerIndex := 0
	if i == 53 {
		MoveJockerV1(deck, i, 2)
		jockerIndex = 2
	} else if i == 52 {
		MoveJockerV1(deck, i, 1)
		jockerIndex = 1
	} else {
		deck[i], deck[i+1], deck[i+2] = deck[i+1], deck[i+2], deck[i]
		jockerIndex = i + 2
	}
	return deck, jockerIndex
}

// JockerShift V2
func JockerShift(deckKeyes []string) ([]string, []int) {
	// fmt.Println("--- JOCKER SHIFT ---")
	jockers := []int{}
	// fmt.Println("JOCKER SHIFT INIT DECK:", deckKeyes)

	ja := FindJocker(deckKeyes, "JA")
	// fmt.Println("JA INDEX BEFORE MOVE", ja)
	deckKeyes, ja = MoveJockerA(deckKeyes, ja)
	// fmt.Println("JOCKER SHIFT JA DECK:", deckKeyes)
	// fmt.Println("JA INDEX AFTER MOVE", ja)

	jb := FindJocker(deckKeyes, "JB")
	// fmt.Println("JB INDEX BEFORE MOVE", jb)
	deckKeyes, jb = MoveJockerB(deckKeyes, jb)
	// fmt.Println("JOCKER SHIFT JB DECK:", deckKeyes)
	// fmt.Println("JB INDEX AFTER MOVE", jb)

	ja = FindJocker(deckKeyes, "JA")

	jockers = append(jockers, ja, jb)
	if jockers[0] > jockers[1] {
		jockers[0], jockers[1] = jockers[1], jockers[0]
	}
	// fmt.Println("JOCKERS", deckKeyes[jockers[0]], deckKeyes[jockers[1]])
	return deckKeyes, jockers
}

// Performs a triple cut on a deck
func TripleCut(deckKeyes []string, jockers []int) []string {
	// fmt.Println("--- TRIPLE CUT ---")
	// fmt.Println("INITIAL DECK --->", deckKeyes)
	// fmt.Println("JOCKERS", jockers)

	deck := []string{}
	top := deckKeyes[:jockers[0]]
	middle := deckKeyes[jockers[0] : jockers[1]+1]
	bottom := deckKeyes[jockers[1]+1:]

	// fmt.Println("TOP --->", top)
	// fmt.Println("MID --->", middle)
	// fmt.Println("BOTTOM --->", bottom)

	deck = append(deck, bottom...)
	deck = append(deck, middle...)
	deck = append(deck, top...)
	// fmt.Println("RESULT DECK --->", deck, len(deck))
	return deck
}

// Performs a count cut on a deck
func CountCut(tripleCutDeck []string) []string {
	// fmt.Println("--- COUNT CUT ---")

	lastIndex := len(tripleCutDeck) - 1
	value := cardToNumber(tripleCutDeck[lastIndex])

	// fmt.Println("LAST CARD:", tripleCutDeck[lastIndex], "VALUE:", value)
	// fmt.Println("INITIAL DECK --->", tripleCutDeck)

	top := tripleCutDeck[:value]
	middle := tripleCutDeck[value:lastIndex]
	bottom := tripleCutDeck[lastIndex]

	// fmt.Println("TOP --->", top)
	// fmt.Println("MID --->", middle)
	// fmt.Println("BOTTOM --->", bottom)

	countCutDeck := []string{}
	countCutDeck = append(countCutDeck, middle...)
	countCutDeck = append(countCutDeck, top...)
	countCutDeck = append(countCutDeck, bottom)

	// fmt.Println("RESULT DECK --->", countCutDeck, len(countCutDeck))

	if countCutDeck[53] == "JA" || countCutDeck[53] == "JB" {
		fmt.Println("COUNT CUT: JOCKER IS THE LAST CARD!", countCutDeck[53])
	}

	return countCutDeck
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
	if number == 53 {
		fmt.Println("")
		fmt.Println("FIND OUTPUT: OUTPUT CARD IS A JOCKER!", tripleCutDeck[0], number)
		fmt.Println("")
		return 0

	}

	outputNumber := cardToNumber(tripleCutDeck[number+1])
	return outputNumber
}

// Creates a keystream for conversion into chars

// Non-recursive of cipher/decipher process
func KeyStream(textLength int, inputDeck *[]string) ([]string, []int) {
	// i := 0
	// KeyStreamRecusrsive(inputDeck, keyStream, numberedText, i)
	var keyStream = &[]int{}
	fmt.Println("KEYSTREAM GEN: TEXT LEN ==", textLength)

	for i := 0; i < textLength; {
		fmt.Println("KEYSTREAM GEN: I ==", i)
		jockers := &[]int{}
		*inputDeck, *jockers = JockerShift(*inputDeck)
		*inputDeck = TripleCut(*inputDeck, *jockers)
		*inputDeck = CountCut(*inputDeck)
		key := FindOutput(*inputDeck)
		if key == 0 {
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
