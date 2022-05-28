package deck

type Card struct {
	Suit int
	Value int
}

type Deck struct {
	Card *Card
}

func (deck *Deck) New() {
	for val := 2; val < 15; val++ {
		for suit := 1; val < 5; val++ {
			newCard := Card{Suit: suit, Value: val}
			deck.Card = &newCard
		}
	}
} 