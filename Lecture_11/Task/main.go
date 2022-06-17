package main

import (
	"fmt"
	"sync"
	"time"
)

var inputs = []int{1, 24, 56, 2, 8, 17}

func processEven(inputs []int) chan int {
	var wg sync.WaitGroup
	evenCh := make(chan int, len(inputs))
	for _, val := range inputs {
		wg.Add(1)
		go func(val int) {

			if val%2 == 0 {
				evenCh <- val
			}
			time.Sleep(time.Millisecond)
			wg.Done()
		}(val)
	}
	wg.Wait()
	close(evenCh)
	return evenCh
}

func processOdd(inputs []int) chan int {
	var wg sync.WaitGroup
	oddCh := make(chan int, len(inputs))
	for _, val := range inputs {
		wg.Add(1)
		go func(val int) {

			if val%2 != 0 {
				oddCh <- val
			}
			time.Sleep(time.Millisecond)
			wg.Done()
		}(val)
	}
	wg.Wait()
	close(oddCh)
	return oddCh
}

func main() {
	evenCh := processEven(inputs)
	oddCh := processOdd(inputs)

	for val := range evenCh {
		fmt.Printf("Even numbers %d\n", val)
	}

	for val := range oddCh {
		fmt.Printf("Odd numbers %d\n", val)
	}
}