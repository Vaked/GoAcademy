package main

import (
	"cards/Task_1/deck"
	"fmt"
)

func main() {
	deck := &deck.Deck{}

	deckSlice := deck.New()
	//deck.Shuffle(&deckSlice)
	fmt.Println()

	card, deckSlice := deck.Deal(deckSlice)
	card2, deckSlice := deck.Deal(deckSlice)
	
	fmt.Println(deckSlice[len(deckSlice) - 1],deckSlice[len(deckSlice) - 2], card, card2)
}