package main

import (
	"fmt"
	// "io"
	// "os"
	rev "reverser/reverse_reader"
)

func main() {
	reversedString := rev.NewReverseStringReader("Hello world!")
	fmt.Println(reversedString.Text)
	//io.Copy(os.Stdout, reversedString)
}