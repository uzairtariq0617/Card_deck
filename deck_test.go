package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newdeck()
	if len(d) != 16 {
		t.Errorf("deck of card length is 20,but got %v", len(d))
	}
	if d[0] != "aceoface of spades" {
		t.Errorf("expected firat card of ace of spades,but got %v", d[0])
	}

}
func TESTTOSAVEFILEANDNEWDECKFROMFILE(T *testing.T) {
	os.Remove("_decktesting")
	deck := newdeck()
	deck.savetofile("_decktesting")
	loadedduck := newdeckfromfile("_decktesting")
	if len(loadedduck) != 16 {
		T.Errorf("expected 16 card in decks,got %v", len(loadedduck))
	}
	os.Remove("_decktesting")

}
