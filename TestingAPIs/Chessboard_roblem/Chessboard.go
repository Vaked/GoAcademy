package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Grains of rice on a chessboard (doubling every square):\n")

	// world wide rice production in 2012 = 740 million tonnes
	// = 1628000000000 pounds
	// assume 29000 grains in 1 pound of rice

	fmt.Println("---------------------------------------------")
	grains := big.NewInt(1)
	// a chessboard has 64 squares
	for square := 1; square <= 64; square++ {
		// add to total
		fmt.Printf("square %v  grains = %02d\n", square, grains)
		// double grains
		grains.Add(grains, grains)
	}
	fmt.Println("---------------------------------------------")
	fmt.Println("Assume 29000 grains per pound of rice:")
}
