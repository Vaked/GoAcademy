package main

import(
	"log"
	"time"
)

func generateThrottled(data string, bufferLimit int, clearInterval time.Duration) <-chan string {
	outChan := make(chan string, bufferLimit)
	go func() {
		for {
			outChan <- data
			if len(outChan) == bufferLimit - 1{
				time.Sleep(clearInterval)
			}
			
		}
	}()
	return outChan
}

func main() {
	out := generateThrottled("foo", 3, time.Second)
	for f := range out {
 	log.Println(f)
	}
}