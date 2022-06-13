Create a type Square, and a type Circle (add
NewSquare and NewCircle for convenience). Both of
them should also have a method called Area:

func (r *ReceiverType) Area() float64 {}

Also, create an interface called Shape, which expects the same method Area
Create a custom type called Shapes:
type Shapes []Shape
and attach a method to it, called LargestArea

func (s Shapes) LargestArea float64

Demonstrate its use, by appending multiple shapes to an instance of Shapes, and calling LargestArea on it.