package main

import (
	"fmt"
	"cardgame/cardgame"
)

func main() {
	var firstHand = cardgame.Card {Suit:4, Value:4}
	var secondHand = cardgame.Card{Suit:12, Value:2}
	var cards  []cardgame.Card
	cards = append(cards, firstHand)
	cards = append(cards, secondHand)

	fmt.Println(cardgame.MaxCard(cards))
}