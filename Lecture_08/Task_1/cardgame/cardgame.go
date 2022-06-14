package cardgame

import (
	"log"
	"errors"
)

type Card struct {
	Suit  Suit
	Value int
}

type Deck struct {
	Cards []Card
}

type Suit = int

const (
	Club Suit = iota + 1
	Diamond
	Heart
	Spade
)

func (deck *Deck) New() {
	for val := 2; val < 15; val++ {
		for suit := 1; suit < 5; suit++ {
			newCard := Card{Suit: suit, Value: val}
			deck.Cards = append(deck.Cards, newCard)
		}
	}
}

// 0 1 2 3 4 5 6 7 lastindexof = len-1
// 0 1 2 3 4 5 6   lastindexof = len-2
func (deck *Deck) Deal() (*Card, error) {
	lasCard := &deck.Cards[len(deck.Cards)-1]
	if len(deck.Cards) > 0 {
		deck.Cards = append(deck.Cards, deck.Cards[len(deck.Cards)-2])
		return lasCard, nil
	}
	return nil, errors.New("no more cards in the deck")
}

func (deck *Deck) Done() bool {
	if len(deck.Cards) <= 0{
		return true
	}
	return false
}