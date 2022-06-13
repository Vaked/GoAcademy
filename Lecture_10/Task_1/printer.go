package main
import (
	"fmt"
	"sync"
	"time"
)

type ConcurrentPrinter struct {
	wg sync.WaitGroup
	mu sync.Mutex
}

func (cp *ConcurrentPrinter) printFoo(times int) {
	cp.wg.Add(1)
	go func() {
		defer cp.wg.Done()
		for i := 0; i < times; i++ {
			cp.mu.Lock()
			if i%2 == 0 {
				fmt.Print("foo")
			}
			cp.mu.Unlock()
			time.Sleep(time.Millisecond)
		}
	}()
}

func (cp *ConcurrentPrinter) printBar(times int) {
	cp.wg.Add(1)
	go func() {
		defer cp.wg.Done()
		for i := 0; i < times; i++ {
			cp.mu.Lock()
			if i%2 != 0 {
				fmt.Print("bar")
			}
			cp.mu.Unlock()
			time.Sleep(time.Millisecond)
		}
	}()
}

func main() {
	times := 10
	cp := &ConcurrentPrinter{}
	cp.printFoo(times)
	cp.printBar(times)
	cp.Wait()
}
