package cardgame

import(
	"fmt"
)

type Card struct {
	Value int
	Suit int
}

func checkCardSuit (card Card) bool {
	if card.Suit < 1 || card.Suit > 4 {
		return false
	}
	return true
}

func checkCardValue (card Card) bool {
	if card.Value < 2 || card.Value > 14 {
		return false
	}
	return true
}

func checkCard (card Card) bool {
	if !checkCardSuit(card) || !checkCardValue(card) {
		return false
	}
	return true
}

func sumValueAndSuit(card Card) int {
	return card.Value + card.Suit
}

func MaxCard(cards []Card) Card {
	var maxCard = cards[0]
	for idx, card := range cards {
		if compareCards(maxCard, card) != 0 {
			maxCard = cards[idx]
		}
	}
	return maxCard

}

func compareCards(cardOne Card, cardTwo Card) int {
	if !checkCard(cardOne) || !checkCard(cardTwo) {
		fmt.Println("Invalid input, card value or card suit is wrong")
		return -2
	}
	firstHandSum := sumValueAndSuit(cardOne)
	secondHandSum := sumValueAndSuit(cardTwo)

	if secondHandSum > firstHandSum {
		return -1
	} else if firstHandSum > secondHandSum {
		return 1
	} else {
		return 0
	}
}