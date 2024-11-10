package deck_utils

import (
	"fmt"
	"testing"

	"example.com/go-pontifex/pkg/utils"
)

var a = [4]string{"clubs", "diamonds", "hearts", "spades"}
var b = [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

func TestDeckGenerator(t *testing.T) {

	got := DeckGenerator(a, b)
	want := 1
	if got["clubs-A"] != want {
		t.Error("Key and value are incorrect")
	}
	if got["JA"] != 53 {
		t.Error("Key and value are incorrect")
	}
	if got["JB"] != 53 {
		t.Error("Key and value are incorrect")
	}
	if len(got) != 54 {
		t.Error("Array length is incorrect")
	}
}

func TestDeckShuffle(t *testing.T) {
	got := DeckGenerator(a, b)
	want := DeckShuffle(got)
	array := []string{}

	array = append(array, want...)
	if len(got) != len(array) {
		t.Error("Length of decks is not equal")
	}
	i := 0
	for k := range got {
		fmt.Println(k, array[i])
		// if k != array[i] {
		// 	t.Error("Decks are identical")
		// }
		i++
	}
}

func TestMoveJocker(t *testing.T) {
	initialDeck := DeckGenerator(a, b)
	// fmt.Println("INITIAL DECK", initialDeck)

	shuffledDeck := DeckShuffle(initialDeck)

	// fmt.Println("SHUFFLED DECK", shuffledDeck)

	movedJockerDeck := MoveJockerV1(shuffledDeck, 1, 1)
	pointer1 := &movedJockerDeck
	// fmt.Println("MOVED JOCKER DECK", pointer1)

	shuffledDeck2 := DeckShuffle(initialDeck)

	flag := 0

	for i := range shuffledDeck2 {
		fmt.Println("SD [i] =", shuffledDeck2[i], "MJD [i] =", (*pointer1)[i])
		if shuffledDeck2[i] != (*pointer1)[i] {
			flag++
			break
		}
	}
	if shuffledDeck2[53] != (*pointer1)[1] {
		flag++
	}
	if flag == 0 {
		fmt.Println("FLAG", flag)
		t.Error("No cards were moved!")
	}
}

func TestJockerShift(t *testing.T) {
	// JA is the last card in the deck. Should be 2nd after shift
	deckOne := utils.ReadDeck("test_input_deck_one.txt")
	shiftedDeck, _ := JockerShift(deckOne)
	if shiftedDeck[1] != "JA" {
		t.Error("casae one: Joker A is shifted into wrong position!")
	}
	// JB is the last card in the deck. Should be 3rd after shift
	deckTwo := utils.ReadDeck("test_input_deck_two.txt")
	shiftedDeck, _ = JockerShift(deckTwo)
	fmt.Println(shiftedDeck[1])
	if shiftedDeck[2] != "JB" {
		t.Error("case two: Joker B is shifted into wrong position!")
	}

	// JB is the second last card in the deck. Should be 2nd after shift
	deckThree := utils.ReadDeck("test_input_deck_three.txt")
	shiftedDeck, _ = JockerShift(deckThree)
	if shiftedDeck[1] != "JB" {
		t.Error("case three: Joker B is shifted into wrong position!")
	}

	deckFour := utils.ReadDeck("test_input_deck_four.txt")
	shiftedDeck, _ = JockerShift(deckFour)
	if shiftedDeck[2] != "JA" && shiftedDeck[3] != "JB" {
		t.Error("case four: Jokers are shifted into wrong positions!")
	}
}

func TestTripleCut(t *testing.T) {
	fmt.Println("--- TRIPLE CUT TEST:")
	deck := utils.ReadDeck("test_input_deck_one.txt")
	jockers := []int{}
	fmt.Println("INITIAL DECK:", deck)

	if len(deck) < 54 {
		t.Error("Deck length is incorrect!")
	}

	deck, jockers = JockerShift(deck)
	fmt.Println("JOKER SHIFT DECK:", deck)

	jOne := 999
	jTwo := 999
	for i := range deck {
		fmt.Println("CURRENT CARD:", deck[i], i)
		if deck[i] == "JA" && jOne == 999 || deck[i] == "JB" && jOne == 999 {
			fmt.Println("---CASE ONE---")
			fmt.Println("TOP JOKER [i]:", deck[i])
			jOne = i
			fmt.Println("JONE =:", jOne)
			i++
		} else if deck[i] == "JA" && jOne != 999 || deck[i] == "JB" && jOne != 999 {
			fmt.Println("---CASE TWO---")
			fmt.Println("BOTTOM JOKER [i]:", deck[i])
			jTwo = i
			fmt.Println("JTWO =:", jTwo)
			break
		}
	}

	fmt.Println("JONE =:", jOne)
	fmt.Println("JTWO =:", jTwo)

	lenTop := jOne
	lenMid := jTwo - jOne
	lenBottom := 53 - jTwo
	indexBottom := lenBottom + lenMid

	fmt.Println("LENGTH ONE:", lenTop)
	fmt.Println("LENGTH TWO:", lenBottom)

	tripleCutDeck := TripleCut(deck, jockers)
	fmt.Println("TRIPLE CUT DECK:", tripleCutDeck)

	fmt.Println("TOP JOKER =", tripleCutDeck[lenBottom])
	if tripleCutDeck[lenBottom] != "JA" && tripleCutDeck[lenBottom] != "JB" {
		t.Error("Top joker index is wrong!")
	}

	fmt.Println("BOTTOM JOKER =", tripleCutDeck[indexBottom])
	if tripleCutDeck[indexBottom] != "JA" && tripleCutDeck[indexBottom] != "JB" {
		t.Error("Bottom joker index is wrong!")
	}

	for i := range deck {
		if deck[i] == tripleCutDeck[i] {
			t.Error("Elements are the same!")
		}
	}
}

func TestCountCut(t *testing.T) {

	initialDeck := utils.ReadDeck("test_input_deck_three.txt")

	value := cardToNumber(initialDeck[53])

	countCutDeck := CountCut(initialDeck)

	lenDeck := len(initialDeck)
	lenTop := value
	lenBottom := 1
	lenMiddle := lenDeck - lenTop - lenBottom

	// Tests if bottom cards are the same
	if initialDeck[53] != countCutDeck[53] {
		t.Error("Last cards are not the same!")
	}
	// Tests if first cards from the middle are the same
	if initialDeck[lenTop] != countCutDeck[0] {
		fmt.Println("TOP CARD :", initialDeck[lenTop], countCutDeck[0])
		t.Error("First card of top is in wrong place!")
	}
	// Tests if first cards from the top are the same
	if initialDeck[0] != countCutDeck[lenMiddle] {
		fmt.Println("TOP CARD OF MIDDLE:", initialDeck[0], countCutDeck[lenMiddle])
		t.Error("First card of middle is in wrong place!")
	}

}

func TestCardToNumber(t *testing.T) {

}

func TestFindOutput(t *testing.T) {}

func TestKeyStream(t *testing.T) {
	numberedText := []int{1, 2, 3}
	textLength := len(numberedText)
	inputDeck := utils.ReadDeck("test_input_deck_five.txt")
	_, keyStream := KeyStream(textLength, &inputDeck)
	fmt.Println(keyStream)
	if len(keyStream) == 0 {
		t.Error("Keystream array length is zero!")
	}
}

func TestKeyStreamRecusrsive(t *testing.T) {}
