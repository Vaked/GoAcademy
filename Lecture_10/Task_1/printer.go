package main
import (
	"fmt"
	"sync"
	"time"
)

type ConcurrentPrinter struct {
	sync.WaitGroup
	sync.Mutex
}

func (cp *ConcurrentPrinter) printFoo(times int) {
	cp.Add(1)
	go func() {
		defer cp.Done()
		for i := 0; i < times; i++ {
			cp.Lock()
			if i%2 == 0 {
				fmt.Print("foo")
			}
			cp.Unlock()
			time.Sleep(time.Millisecond)
		}
	}()
}

func (cp *ConcurrentPrinter) printBar(times int) {
	cp.Add(1)
	go func() {
		defer cp.Done()
		for i := 0; i < times; i++ {
			cp.Lock()
			if i%2 != 0 {
				fmt.Print("bar")
			}
			cp.Unlock()
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
