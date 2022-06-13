TASK 1:

Create a new custom type, called Deck (a deck of cards).
type Deck struct {/* decide what to add here */}
Create a factory function New that initializes a new deck. Initializing a new deck, must fill an internal
(for outside users) data structure with 52 cards (of our familiar type Card)
The Deck type must also have two methods: Shuffle and Deal
- Shuffle will randomize the order of the cards inside the deck
- Deal will return the first card from the Deck, removing it from the internal data structure. Please, also
make sure you take care of the case where the deck gets empty. To signal that to the user, you can simply
make the function return a Card pointer.

TASK 2:

Create a new revision of the homework task from last week, where we
created the maxCard function.
Instead of internally using the compareCards function, make the
maxCard function take a second parameter of type CardComparator:

type CardComparator func(cOne Card, cTwo Card) int

func maxCard(cards []Card, comparatorFunc CardComparator) Card {
    // use comparatorFunc here to find the maximum ...
}
