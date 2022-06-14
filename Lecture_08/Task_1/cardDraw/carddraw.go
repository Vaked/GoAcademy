package carddraw

import (
	cardgame "cards/cardgame"
	"fmt"
)

type dealer interface {
	Deal() *cardgame.Card
	Done() bool
}

func DrawAllCards(dealer dealer) ([]cardgame.Card, error) {
	var newDeck cardgame.Deck
	newDeck.New()

	var deckSlice []cardgame.Card

	for {
		if !dealer.Done() {
			card := dealer.Deal()

			if card != nil {
				deckSlice = append(deckSlice, cardgame.Card{card.Suit, card.Value})
				fmt.Printf("Print current card :%s , %s\n", card.Suit, card.Value)
			}
		} else {
			deck.Cards, 
		}
	}
	return deckSlice, nil
}