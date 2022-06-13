package main

import (
	carddraw "cards/cardDraw"
	cardgame "cards/cardgame"
	"fmt"
	"log"
)

func main() {
	deck := &cardgame.Deck{}
	deck.New()
	result , err := carddraw.DrawAllCards(deck)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}