package cardgame

type Card struct {
	Suit     Suit
	Value    int
	PrevCard *Card
}

type Deck struct {
	LastCard *Card
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
			newCard := Card{Suit: suit, Value: val, PrevCard: deck.LastCard}
			deck.LastCard = &newCard
		}
	}
}

func (deck *Deck) Deal() *Card {
	lastCard := deck.LastCard
	if lastCard != nil {
		deck.LastCard = deck.LastCard.PrevCard
	}
	return lastCard
}
