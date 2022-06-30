package cardgame

import (
	"testing"
)

func TestCheckCardSuit(t *testing.T) {
	//Arrange
	var card Card
	card.Suit = 4

	//Act
	result := checkCardSuit(card)
	expectedResult := true

	//Assert
	if result != expectedResult {
		t.Errorf("Expected: %t, got: %t", expectedResult, result)
	}
}

func TestCheckCardValue(t *testing.T) {
	//Arrange
	var card Card
	card.Value = 10

	//Act
	result := checkCardValue(card)
	expectedResult := true

	//Assert
	if result != expectedResult {
		t.Errorf("Expected: %t, got: %t", expectedResult, result)
	}
}

func TestCheckCard(t *testing.T) {
	//Arrange
	var card Card
	card.Value = 10

	//Act
	result := checkCardValue(card)
	expectedResult := true

	//Assert
	if result != expectedResult {
		t.Errorf("Expected: %t, got: %t", expectedResult, result)
	}
}

func TestCheckCardSum(t *testing.T) {
	//Arrange
	var card Card
	card.Suit = 5
	card.Value = 4

	//Act
	result := sumValueAndSuit(card)
	expectedResult := card.Suit + card.Value

	//Assert
	if result != expectedResult {
		t.Errorf("Expected: %d, got: %d", expectedResult, result)
	}
}

func TestCompareCards(t *testing.T) {
	//Arrange
	firstCard := Card{Suit: 2, Value: 12}
	secondCard := Card{Suit: 4, Value: 4}

	//Act
	result := compareCards(firstCard, secondCard)
	expectedResult := 1

	//Assert
	if expectedResult != result {
		t.Errorf("Expected: %d, got: %d", expectedResult, result)
	}
}

func TestMaxCard(t *testing.T) {
	//Arrange
	var firstHand = Card{4, 4}
	var secondHand = Card{12, 2}
	var cards []Card
	cards = append(cards, firstHand)
	cards = append(cards, secondHand)

	//Act
	result := MaxCard(cards)
	expectedResult := secondHand

	//Assert
	if result != expectedResult {
		t.Errorf("Expected: %d, got: %d", expectedResult, result)
	}
}
