Let's create our first real data structure. We will call it MagicList. A MagicList looks like this:

type MagicList struct {
    LastItem *Item
}

Where Item looks like this:

type Item struct {
    Value int
    PrevItem *Item
}

Your goal is to write two functions:
func add(l *MagicList, value int)
This will have to do a couple of things:
— Create a new Item using the int value
— If the list's LastItem is nil, point LastItem to its address
— LastItem becomes your defacto first and only item
— If a LastItem already exists (i.e. not nil)
— take its address and set it to the new Item's PrevItem
— set LastItem to point to the address of the new item

But the more interesting part starts now:
Write another function:
func toSlice(l *MagicList) []int
which, starting from the list's LastItem, traverses it back to the beginning (following PrevItem
until no item has a previous item anymore), gets the int value, and puts it on a regular int slice.
HINT: Think of where you should append to the slice (beginning or end)
Extra bonus points for whoever uses recursion (Google it)

Demonstrate how the program works, by adding 10
integers to the MagicList:
l := &MagicList{}
add(l, 10)
add(l, 22)
add(l, 44)
// etc
fmt.Println(toSlice(l))