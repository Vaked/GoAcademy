package carddraw

import (
	cardgame "cards/cardgame"
)

type dealer interface {
	Deal() *cardgame.Card
}

func DrawAllCards(dealer dealer) []cardgame.Card {
	var newDeck []cardgame.Card

	for true {
		card := dealer.Deal()

		if card != nil {
			newDeck = append(newDeck, *card)
		} else {
			break
		}
	}
	return newDeck
}