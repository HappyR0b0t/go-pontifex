package deck_utils

import (
	"math/rand"
	"strings"
)

var suits = map[string]int{
	"clubs":    0,
	"diamonds": 13,
	"hearts":   26,
	"spades":   39,
}

var rank = map[string]int{
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
	for i := 0; i < len(suit); i++ {
		for j := 0; j < len(rank); j++ {
			deck[suit[i]+"-"+rank[j]] = k
			k++
		}
	}
	deck["JA"] = 0
	deck["JB"] = 0
	return deck
}

func DeckShuffle(deck map[string]int) []string {
	deckKeyes := []string{}
	for k := range deck {
		deckKeyes = append(deckKeyes, k)
	}
	rand.Shuffle(len(deckKeyes), func(i, j int) { deckKeyes[i], deckKeyes[j] = deckKeyes[j], deckKeyes[i] })
	return deckKeyes
}

func MoveJocker(deckKeyes []string, current int, target int) []string {
	for i := len(deckKeyes) - current; i > target; i-- {
		deckKeyes[i-1], deckKeyes[i] = deckKeyes[i], deckKeyes[i-1]
		// fmt.Println("MOVE JOKERS ->", deckKeyes[i])
	}

	return deckKeyes
}

// Finds both Jokers and shifts them accordingly
// TODO Return indexes of both Jokers
func JockerShift(deckKeyes []string) []string {
	// fmt.Println("-----------------JOKERS SHIFT---------------------------")
	// fmt.Println("DECK ->", deckKeyes, " - ", len(deckKeyes))
	shiftOne := 1
	shiftTwo := 2
	a := false
	b := false
	for i := 0; i < len(deckKeyes); i++ {
		if a && b {
			break
		}
		if deckKeyes[i] == "JA" && !a {
			// fmt.Println("IF-1")
			deckKeyes[i], deckKeyes[i+1] = deckKeyes[i+1], deckKeyes[i] // One function could be used. I guess
			a = true
			i = 0
		} else if deckKeyes[len(deckKeyes)-1] == "JA" && !a {
			// fmt.Println("IF-2")
			MoveJocker(deckKeyes, shiftOne, shiftOne)
			a = true
			i = 0
		}
		if deckKeyes[i] == "JB" && !b && a {
			// fmt.Println("IF-3")
			deckKeyes[i], deckKeyes[i+1], deckKeyes[i+2] = deckKeyes[i+1], deckKeyes[i+2], deckKeyes[i]
			b = true
			i += 2
		} else if deckKeyes[len(deckKeyes)-1] == "JB" && !b && a {
			// fmt.Println("IF-4")
			MoveJocker(deckKeyes, shiftOne, shiftTwo)
			b = true
			break
		} else if deckKeyes[len(deckKeyes)-2] == "JB" && !b && a {
			// fmt.Println("IF-4")
			MoveJocker(deckKeyes, shiftTwo, shiftOne)
			b = true
			break
		}
	}
	return deckKeyes
}

// func JockerShift(deckKeyes []string) ([]string, []int) {
// 	fmt.Println("-----------------JOKERS SHIFT---------------------------")
// 	fmt.Println("DECK ->", deckKeyes, " - ", len(deckKeyes))
// 	jokers := []int{0, 0}
// 	shiftOne := 1
// 	shiftTwo := 2
// 	a := false
// 	b := false
// 	for i := range deckKeyes {
// 		if deckKeyes[i] == "JA" && i != len(deckKeyes)-1 && !a {
// 			deckKeyes[i], deckKeyes[i+1] = deckKeyes[i+1], deckKeyes[i]
// 			a = true
// 			i += 1
// 			jokers[0] = i
// 			fmt.Println("JA - ", deckKeyes[i], i)
// 			break
// 		} else if deckKeyes[len(deckKeyes)-1] == "JA" && !a {
// 			MoveJocker(deckKeyes, shiftOne, shiftOne)
// 			a = true
// 			jokers[0] = 1
// 			fmt.Println("JA - ", deckKeyes[i], i)
// 			break
// 		} else {
// 			continue
// 		}
// 	}
// 	for i := range deckKeyes {
// 		if deckKeyes[i] == "JB" && i != len(deckKeyes)-1 && !b {
// 			deckKeyes[i], deckKeyes[i+1], deckKeyes[i+2] = deckKeyes[i+1], deckKeyes[i+2], deckKeyes[i]
// 			b = true
// 			i += 2
// 			jokers[0] = i
// 			fmt.Println("JB - ", deckKeyes[i], i)
// 			break
// 		} else if deckKeyes[len(deckKeyes)-1] == "JB" && !b {
// 			MoveJocker(deckKeyes, shiftOne, shiftTwo)
// 			b = true
// 			jokers[0] = 2
// 			i = 2
// 			fmt.Println("JB - ", deckKeyes[i], i)
// 			break
// 		} else if deckKeyes[len(deckKeyes)-2] == "JB" && !b {
// 			MoveJocker(deckKeyes, shiftTwo, shiftOne)
// 			b = true
// 			jokers[0] = 1
// 			i = 1
// 			fmt.Println("JB - ", deckKeyes[i], i)
// 			break
// 		} else {
// 			continue
// 		}
// 	}
// 	if jokers[0] > jokers[1] {
// 		jokers[0], jokers[1] = jokers[1], jokers[0]
// 	}
// 	return deckKeyes, jokers
// }

func TripleCut(deckKeyes []string) []string {
	// fmt.Println("---> TRIPLE CUT <---")
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

	// top := deckKeyes[:jokers[0]]
	// middle := deckKeyes[jokers[0] : jokers[1]+1]
	// bottom := deckKeyes[jokers[1]+1:]

	// fmt.Println("TOP ->", top, "<- TOP")
	// fmt.Println("MIDDLE ->", middle, "<- MIDDLE")
	// fmt.Println("BOTTOM ->", bottom, "<- BOTTOM")

	deck = append(deck, bottom...)
	deck = append(deck, middle...)
	deck = append(deck, top...)
	// fmt.Println("TRIPLE CUT DECK ->", deck, "<- TRIPLE CUT DECK", len(deck)) // delete
	return deck
}

func CountCut(tripleCutDeck []string) []string {
	// fmt.Println("---> COUNT CUT <---")
	lastIndex := len(tripleCutDeck) - 1
	bottomCard := tripleCutDeck[lastIndex]
	// fmt.Println("BOTTOM CARD ->", bottomCard, "<- BOTTOM CARD") // delete
	number := cardToNumber(tripleCutDeck[lastIndex])
	// fmt.Println("BOTTOM CARD NUMBER ->", number, "<- BOTTOM CARD NUMBER") // delete
	top := tripleCutDeck[0:number]
	bottom := tripleCutDeck[number:lastIndex]
	deck := []string{}
	deck = append(deck, bottom...)
	deck = append(deck, top...)
	deck = append(deck, bottomCard)
	// fmt.Println("COUNT CUT DECK ->", deck, "<- COUNT CUT DECK", len(deck)) // delete

	return deck
}

func cardToNumber(card string) int {
	suitAndRank := strings.Split(card, "-")
	if suitAndRank[0] == "JA" || suitAndRank[0] == "JB" {
		number := 53
		return number
	} else {
		number := suits[suitAndRank[0]] + rank[suitAndRank[1]]
		return number
	}
}

func FindOutput(tripleCutDeck []string) int {
	// fmt.Println("---> FIND OUTPUT <---")
	number := cardToNumber(tripleCutDeck[0])
	// fmt.Println("TOP CARD ->", tripleCutDeck[0], "<- TOP CARD")
	outputNumber := cardToNumber(tripleCutDeck[number])
	// fmt.Println("OUTPUT NUMBER ->", outputNumber, "<- OUTPUT NUMBER")
	return outputNumber
}

func KeyStream(numberedText []int, inputDeck *[]string, keyStream *[]int) ([]string, []int) {
	// keyStream := []int{}
	i := 0
	jokers := []int{0, 0}
	KeyStreamRecusrsive(inputDeck, keyStream, numberedText, i, &jokers)

	return *inputDeck, *keyStream
}

func KeyStreamRecusrsive(inputDeck *[]string, keyStream *[]int, numberedText []int, i int, jokers *[]int) ([]string, []int) {
	// c := []string{}
	key := 0
	if i < len(numberedText) {
		*inputDeck = JockerShift(*inputDeck)
		*inputDeck = TripleCut(*inputDeck)
		*inputDeck = CountCut(*inputDeck)
		key = FindOutput(*inputDeck)
		*keyStream = append(*keyStream, key)
		// fmt.Println("KEYSTREAM --->", keyStream, "<--- KEYSTREAM")
		i += 1
		// fmt.Println("I --->", i, "<--- I")
	} else {
		return *inputDeck, *keyStream
	}
	return KeyStreamRecusrsive(inputDeck, keyStream, numberedText, i, jokers)
}
