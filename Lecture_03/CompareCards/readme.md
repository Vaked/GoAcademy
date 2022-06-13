Create a simple program that compares the strengths of two cards and returns one of the
following:
— -1 if the first card has a lower strength than the second one
— 0 if both cards are equal
— 1 if the second card has a greater strength than the first one
A card’s strength consists of two things:
— value (number from 2 to 13). For simplicity, we will assume that J, D, K, and A are also
numbers
— suit (club, diamond, heart, spade) in that order

Use the following function signature:
func compareCards(cardOneVal int, cardOneSuit int, cardTwoVal int, cardTwoSuit int) int {
// ...
}