package main 
 
import (
	"fmt"
	carddraw "cards/cardDraw"
	cardgame "cards/cardgame"
)

func main() {

	deck := &cardgame.Deck{}
	deck.New()
	fmt.Println(carddraw.DrawAllCards(deck))
}