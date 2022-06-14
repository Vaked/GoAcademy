package carddraw

import (
	cardgame "cards/cardgame"
	"fmt"
)

type dealer interface {
	Deal() (*cardgame.Card, error)
	Done() bool
}

func DrawAllCards(dealer dealer) ([]cardgame.Card, error) {
	var newDeck cardgame.Deck
	newDeck.New()
	

	fmt.Println(newDeck)
	// if dealer.Done(){
	// 	return newDeck, nil
	// }

	// for {
	// 	card, err := dealer.Deal()

	// 	if card != nil {
	// 		newDeck = append(newDeck, *card)
	// 	} else {
	// 		return newDeck, err
	// 	}
	// }
}