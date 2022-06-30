package main

import (
	"fmt"
	"sync"
	"time"
)

func primesAndSleep(n int, sleep time.Duration) []int {
	res := []int{}
	for k := 2; k < n; k++ {
		for i := 2; i < n; i++ {
			if k%i == 0 {
				time.Sleep(sleep)
				if k == i {
					res = append(res, k)
				}
			break
		}
		}
	}
	return res
}

func primesAndSleepGo(n int, sleep time.Duration) []int {
	var wg sync.WaitGroup
	primeChan := make(chan int, n)
	result := []int{}

	for k := 2; k <= n; k++{
		wg.Add(1)
		go func(k int){
			for i := 2; i <= n; i++{
				if k%i == 0 {
					time.Sleep(sleep)
					if k == i{
						primeChan <- k
					}
					wg.Done()
					break
				}
			}
		}(k)
	}

	wg.Wait()
	close(primeChan)

	for num := range primeChan{
		result = append(result, num)
	}

	return result
}

func main() {
	fmt.Println(primesAndSleep(113, 1))
	fmt.Println(primesAndSleepGo(113, 1))
}