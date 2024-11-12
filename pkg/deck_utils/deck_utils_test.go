package deck_utils

import (
	"fmt"
	"reflect"
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

}

func TestFindJocker(t *testing.T) {
	deck := []string{"a", "b", "c", "JA", "d", "e", "JB"}
	ja := FindJocker(deck, "JA")
	jb := FindJocker(deck, "JB")

	if deck[ja] != "JA" {
		t.Error("error: jocker A index is wrong!")
	}

	if deck[jb] != "JB" {
		t.Error("error: jocker B index is wrong!")
	}
}

func TestJockerShift(t *testing.T) {

}

// func TestTripleCut(t *testing.T) {
// 	got := []string{"a", "b", "c", "JA", "d", "e", "JB", "f", "g"}
// 	want := []string{"f", "g", "JA", "d", "e", "JB", "a", "b", "c"}

// 	jockers := []int{3, 6}
// 	got = TripleCut(got, jockers)

// 	for i := range want {
// 		if got[i] == want[i] {
// 			continue
// 		} else {
// 			t.Error("error: got array is not equal to want!")
// 		}
// 	}

// }

func TestTripleCut(t *testing.T) {
	tests := []struct {
		name         string
		inputDeck    []string
		inputIndices []int
		expected     []string
	}{
		{
			name:         "Top of the deck is empty",
			inputDeck:    []string{"JA", "a", "JB", "b"},
			inputIndices: []int{0, 2},
			expected:     []string{"b", "JA", "a", "JB"},
		},
		{
			name:         "Bottom of the deck is empty",
			inputDeck:    []string{"a", "JA", "b", "JB"},
			inputIndices: []int{1, 3},
			expected:     []string{"JA", "b", "JB", "a"},
		},
		{
			name:         "Top of the deck is empty",
			inputDeck:    []string{"JA", "a", "JB", "b"},
			inputIndices: []int{0, 2},
			expected:     []string{"b", "JA", "a", "JB"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TripleCut(tt.inputDeck, tt.inputIndices)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("TripleCut() = %v, want %v", result, tt.expected)
			}
		})
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
