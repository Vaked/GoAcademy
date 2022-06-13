package main 
 
import (
	"fmt"
	carddraw "cards/cardDraw"
	cardgame "cards/cardgame"
)

func main() {

	deck := &cardgame.Deck{}
	deck.New()
	result, err := carddraw.DrawAllCards(deck)

	fmt.Println()
}