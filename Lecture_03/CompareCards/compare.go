package main

import (
	"fmt"
)

func checkCardValue (card1, card2 int) bool {
	if card1 < 2 || card1 > 13 || card2 < 1 || card2 > 13{
		return false
	}
	return true
}

func checkCardSuit (suit1, suit2 int) bool{
	if suit1 < 1 || suit1 > 4 || suit2 < 1 || suit2 > 4 {
		return false
	}
	return true
}

func sumCardAndSuit(card, suit int) int {
	return card + suit
}

func compareCards(cardOneVal int, cardOneSuit int, cardTwoVal int, cardTwoSuit int) int {
	if checkCardValue(cardOneVal, cardTwoVal) && checkCardSuit(cardOneSuit, cardTwoSuit){
		firstHand := sumCardAndSuit(cardOneVal, cardOneSuit)
		secondHand := sumCardAndSuit(cardTwoVal, cardTwoSuit)
		if secondHand > firstHand {
			fmt.Println("Second hand is bigger")
			return -1
		} else if firstHand > secondHand {
			fmt.Println("First hand is bigger")
			return 1
		} else {
			fmt.Println("The hands are equal")
			return 0
		}
	}
	fmt.Println("Invalid input parameters used for the cards and suits")
	return 1
}

func main () {
	fmt.Println(compareCards(42, 4, 4, 4))
}