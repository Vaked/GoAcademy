package main

import (
	"fmt"
	"sync"
	"time"
)

func primesAndSleep(n int, sleep time.Duration) []int {
	var wg sync.WaitGroup
	res := []int{}

	for k := 2; k < n; k++ {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()
			for i := 2; i < n; i++ {
				wg.Add(1)
				go func(i int){
					defer wg.Done()
					if k%i == 0 {
						time.Sleep(sleep)
						if k == i {
							res = append(res, k)
						}
					}
				}(i)
			}
		}(k)
	}
	wg.Wait()
	return res
}

func main() {
	fmt.Println(primesAndSleep(7, 1))
}