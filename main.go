package main

import (
	"fmt"
)

func foo(bar func() string) string {
	return bar()
}

func main() {
	res := foo(func() string {
	return "bar"
})
	fmt.Println(res) // bar
}