Create a new version of our homework with the deck of cards.
Move Deck and Card to a different package. Call it cardgame. (you may use a single file in
it called cardgame.go)
Create a second package called carddraw (you may use a single file in it called
carddraw.go)
In the carddraw package, create an unexported interface called dealer. It will specify
the Decks’s Deal method, without being in the same package as Deck. Create a function
called DrawAllCards with the following specification ->

package carddraw

type dealer interface {
    Deal() *cardgame.Card
}

func DrawAllCards(dealer dealer) []cardgame.Card {
    // call the dealer's Draw() method, until you reach a nil Card
}


Use the main function in the main package as the gluing layer between the two
other packages. It will create a concrete implementation of cardgame.Deck
and pass it to carddraw's ' DrawAllCards function. It will then assign the
resulting slice and print it on the screen using fmt.Println()