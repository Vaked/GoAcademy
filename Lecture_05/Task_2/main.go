package main

import (
	"fmt"
)

type Card struct {
	Value int
	Suit int
}

func checkCardSuit (card Card) bool {
	if card.Suit < 2 || card.Suit > 4 {
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

func sumCardAndSuit(card Card) int {
	return card.Value + card.Suit
}

func maxCard(cards []Card) Card {
	var maxCard = Card {0, 0}
	for idx, card := range cards {
		if compareCards(maxCard, card) > 0 {
			maxCard = cards[idx]
		}
	}
	return maxCard

}


func compareCards(cardOne Card, cardTwo Card) int {
	if checkCard(cardOne) == false || checkCard(cardTwo) == false {
		fmt.Println("Invalid input, card value or card suit is wrong")
		return -2
	}
	firstHandSum := sumCardAndSuit(cardOne)
	secondHandSum := sumCardAndSuit(cardTwo)

	if secondHandSum > firstHandSum {
		fmt.Println("Second hand is bigger")
		return -1
	} else if firstHandSum > secondHandSum {
		fmt.Println("First hand is bigger")
		return 1
	} else {
		fmt.Println("The hands are equal")
		return 0
	}
}



func main() {
	var firstHand = Card {4, 4}
	var secondHand = Card{ 12, 2}

	compareCards(firstHand, secondHand)
}