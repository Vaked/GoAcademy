package main

import (
	"flag"
	"log"
	"net/http"
	"sync"
	"time"
)

func pingURL(url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.Printf("Got response for %s: %d\n", url, resp.StatusCode)
	return nil
}

func fetchURLs(urlList []string, concurencyLevel int) chan string {
	urlChan := make(chan string, concurencyLevel)	
	var wg sync.WaitGroup

	for _, url := range urlList {
		wg.Add(1)
		urlChan <- url

		go func(url string) {
			time.Sleep(time.Millisecond * 100)
			pingURL(url)
			<- urlChan
			wg.Done()
		}(url)
	}

	wg.Wait()
	close(urlChan)

	return urlChan
}

func main() {
	var concurentRunners int

	flag.IntVar(&concurentRunners, "c", 2, "URL arguments")
	flag.Parse()
	urlList := flag.Args()

	fetchURLs(urlList, 5)
}