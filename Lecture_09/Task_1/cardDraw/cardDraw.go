package carddraw

import (
	cardgame "cards/cardgame"
)

type dealer interface {
	Deal() (*cardgame.Card, error)
	Done() bool
}

func DrawAllCards(dealer dealer) ([]cardgame.Card, error) {
	var newDeck []cardgame.Card
	
	if dealer.Done(){
		return newDeck, nil
	}

	for {
		card, err := dealer.Deal()

		if card != nil {
			newDeck = append(newDeck, *card)
		} else {
			return newDeck, err
		}
	}
}