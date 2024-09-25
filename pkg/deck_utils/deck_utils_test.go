package deck_utils

import (
	"testing"
)

func TestDeckGenerator(t *testing.T) {
	a := [4]string{"a", "b", "c", "d"}
	b := [13]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "n"}
	got := DeckGenerator(a, b)
	// fmt.Println(got)
	want := 1
	if got["a-a"] != want {
		t.Error("...")
	}
	if got["JA"] != 53 {
		t.Error("...")
	}
	if got["JB"] != 53 {
		t.Error("...")
	}
}

func TestDeckShuffle(t *testing.T) {}

func TestMoveJocker(t *testing.T) {}

func TestJockerShift(t *testing.T) {}

func TestTripleCut(t *testing.T) {}

func TestCountCut(t *testing.T) {}

func TestCardToNumber(t *testing.T) {}

func TestFindOutputr(t *testing.T) {}

func TestKeyStream(t *testing.T) {}

func TestKeyStreamRecusrsive(t *testing.T) {}
