Let's communicate possible side effects in our applications from
now and on.
As a first try, copy the last version of your homework with the Deck
once again. This time, make the function Draw communicate an
impossible operation, by returning both a *Card and an error.
Additionally, introduce a new method to both Deck and the dealer
interface, called Done() bool. Return true, if the Deck's internal
pool of cards is empty.

To make things interesting, carddraw.drawAllCards must also
return both []cardgame.Card, and an error. Since we want to
pretend that carddraw and cardgame don't know about it each
other, if an error occurs, first check if the dealer is Done() and if
yes, simply return the drawn cards, and a nil error. Otherwise, do
the opposite.
Modify the the main function to also check for the error, and
terminate the program with log.Fatal(err) if it is present.