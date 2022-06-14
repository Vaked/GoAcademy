package main

import(
	"fmt"
	"sync"
)

var inputs = []int{1, 24, 56, 2, 8, 17}
var wg sync.WaitGroup

func processEven(inputs []int) chan int {
	evenCh := make(chan int, len(inputs))

	for _, val := range inputs{
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			if val % 2 == 0 {
				evenCh <- val
			}
		}(val)
	}
	return evenCh
}

// func processOdd(inputs []int) chan int {
// 	oddCh := make(chan int, len(inputs))


// }

func main() {
	evenCh := processEven(inputs)
	//oddCh :=   processOdd(inputs)

//	var valSlice = make([]int, 10)

	defer close(evenCh)
	for val := range evenCh {
		fmt.Println(val)
	}

	//fmt.Println(len(evenCh))

	defer wg.Wait()
	
}