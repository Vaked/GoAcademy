package cardGame

import(
	"errors"
)


type Card struct {
	Suit     Suit
	Value    int
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
	var	slice []Card
	for val := 2; val < 15; val++ {
		for suit := 1; suit < 5; suit++ {
			newCard := Card{Suit: suit, Value: val}
			slice = append(slice, newCard)
		}
	}
}

func (deck *Deck) Deal() (*Card, error) {
	lastCard := &deck.Cards[len(deck.Cards)-1]

	if lastCard != nil {
		deck.Cards = deck.Cards[:len(deck.Cards) - 1]
		return lastCard, nil
	} else {
		return nil, errors.New("no cards left in the deck")
	}
}

func (deck *Deck) Done() bool {
	if len(deck.Cards) == 0 {
		return true
	}
	return false
}