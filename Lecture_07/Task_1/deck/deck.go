package deck

import (
	"math/rand"
)

type Card struct {
	Suit     Suit
	Value    int
}

type Deck struct {
	cards []Card
}

type Suit = int

const (
	Club Suit = iota + 1
	Diamond
	Heart
	Spade
)

func (deck *Deck) New() []Card{
	var deckSlice []Card
	for val := 2; val < 15; val++ {
		for suit := 1; suit < 5; suit++ {
			newCard := Card{Suit: suit, Value: val}
			deckSlice = append(deckSlice, newCard)
		}
	}
	return deckSlice
}

func (deck *Deck) Shuffle(cards []Card) {
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
}

func (deck *Deck) Deal(cards []Card) (*Card, []Card) {
	if len(cards) != 0 {
		card := cards[len(cards) - 1]
		cards = cards[:len(cards) - 1]
		return &card, cards
	}
	card := cards[len(cards) - 1]
	return &card, cards
}